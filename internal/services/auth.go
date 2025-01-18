package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"hss/internal/models"
	"hss/internal/repositories"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/square/go-jose/v3"
	"github.com/square/go-jose/v3/jwt"
)

type AuthService struct {
	authRepository *repositories.AuthRepository
	encrypter      jose.Encrypter
	signer         jose.Signer
}

func decodeKey(encodedKey string) ([]byte, error) {
	if encodedKey == "" {
		return nil, fmt.Errorf("key must not be empty")
	}

	key, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		return nil, fmt.Errorf("invalid key format: %w", err)
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("key must be 32 bytes long")
	}

	return key, nil
}

func NewAuthService(authRepository *repositories.AuthRepository) (*AuthService, error) {
	env := os.Getenv("JWT_ENCRYPTION_KEY")
	log.Print(env)

	encryptionKey, err := decodeKey(os.Getenv("JWT_ENCRYPTION_KEY"))
	if err != nil {
		log.Fatalf("failed to fetch and decode encryption key: %v", err)
	}

	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.DIRECT, Key: encryptionKey}, nil)
	if err != nil {
		log.Fatalf("failed to instantiate encrypter: %v", err)
	}

	signatureKey, err := decodeKey(os.Getenv("JWT_SIGNATURE_KEY"))
	if err != nil {
		log.Fatalf("failed to fetch and decode signature key: %v", err)
	}

	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: signatureKey}, nil)
	if err != nil {
		log.Fatalf("failed to instantiate signer: %v", err)
	}

	return &AuthService{authRepository: authRepository, encrypter: encrypter, signer: signer}, nil
}

func (s *AuthService) ValidateCredentials(ctx context.Context, authReq *models.AuthenticationRequest) (models.AuthenticationResponse, error) {
	if strings.Contains(authReq.Username, "/") {
		return s.validateAddressCredentials(ctx, authReq)
	} else {
		return s.validateCompanyCredentials(ctx, authReq)
	}
}

func (s *AuthService) validateAddressCredentials(ctx context.Context, authReq *models.AuthenticationRequest) (models.AuthenticationResponse, error) {
	usernameParts := strings.Split(authReq.Username, "/")

	if len(usernameParts) != 2 {
		return models.FailedAuthenticationResponse, models.ErrInvalidCredentials{}
	}

	companyUsername := usernameParts[0]
	addressUsername := usernameParts[1]

	companyExists, err := s.authRepository.CheckCompanyByUsername(ctx, companyUsername)
	if err != nil {
		return models.FailedAuthenticationResponse, models.ErrInvalidCredentials{}
	}
	if !companyExists {
		return models.FailedAuthenticationResponse, fmt.Errorf("company %s does not exist", companyUsername)
	}

	addressID, err := s.authRepository.ValidateAddressCredentials(ctx, companyUsername, addressUsername, authReq.PasswordHash)
	if err != nil {
		return models.FailedAuthenticationResponse, models.ErrAuthenticationFailed{}
	}
	if addressID < 0 {
		return models.FailedAuthenticationResponse, models.ErrInvalidCredentials{}
	}

	token, err := s.generateToken(addressID, models.EmployeeRole)
	if err != nil {
		return models.FailedAuthenticationResponse, err
	}

	return models.AuthenticationResponse{Success: true, Token: token}, nil
}

func (s *AuthService) validateCompanyCredentials(ctx context.Context, authReq *models.AuthenticationRequest) (models.AuthenticationResponse, error) {
	companyID, err := s.authRepository.ValidateCompanyCredentials(ctx, authReq.Username, authReq.PasswordHash)
	if err != nil {
		return models.FailedAuthenticationResponse, models.ErrInvalidCredentials{}
	}

	if companyID < 0 {
		return models.FailedAuthenticationResponse, models.ErrInvalidCredentials{}
	}

	token, err := s.generateToken(companyID, models.AdminRole)
	if err != nil {
		return models.FailedAuthenticationResponse, err
	}

	return models.AuthenticationResponse{Success: true, Token: token}, nil
}

type CustomClaims struct {
	Role models.AccessRole `json:"role"`
	jwt.Claims
}

func (s *AuthService) generateToken(userID int, accessRole models.AccessRole) (string, error) {

	customClaims := CustomClaims{
		Role: accessRole,
		Claims: jwt.Claims{
			Subject:   strconv.Itoa(userID),
			Issuer:    "hss",
			Expiry:    jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token, err := jwt.SignedAndEncrypted(s.signer, s.encrypter).Claims(customClaims).CompactSerialize()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, authReq *models.AuthorizationRequest) (models.AuthorizationResponse, error) {
	parsedJWT, err := jwt.ParseSignedAndEncrypted(authReq.Token)
	if err != nil {
		return models.FailedAuthorizationResponse, err
	}

	encryptionKey, err := decodeKey(os.Getenv("JWT_ENCRYPTION_KEY"))
	if err != nil {
		log.Fatalf("failed to fetch and decode encryption key: %v", err)
	}

	decryptedJWT, err := parsedJWT.Decrypt(encryptionKey)
	if err != nil {
		return models.FailedAuthorizationResponse, err
	}

	signatureKey, err := decodeKey(os.Getenv("JWT_SIGNATURE_KEY"))
	if err != nil {
		log.Fatalf("failed to fetch and decode signature key: %v", err)
	}

	claims := CustomClaims{}
	if err := decryptedJWT.Claims(signatureKey, &claims); err != nil {
		return models.FailedAuthorizationResponse, err
	}

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return models.FailedAuthorizationResponse, err
	}

	return models.AuthorizationResponse{Success: true, Role: string(claims.Role), UserID: userID}, nil
}
