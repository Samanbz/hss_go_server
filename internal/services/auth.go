package services

import (
	"context"
	"fmt"
	"hss/internal/models"
	"hss/internal/repositories"
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

func NewAuthService(authRepository *repositories.AuthRepository) (*AuthService, error) {
	encryptionKey := os.Getenv("JWT_ENCRYPTION_KEY")
	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.DIRECT, Key: encryptionKey}, nil)
	if err != nil {
		return nil, err
	}

	signatureKey := os.Getenv("JWT_SIGNATURE_KEY")
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: signatureKey}, nil)
	if err != nil {
		return nil, err
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

	encryptionKey := os.Getenv("JWT_ENCRYPTION_KEY")
	decryptedJWT, err := parsedJWT.Decrypt(encryptionKey)
	if err != nil {
		return models.FailedAuthorizationResponse, err
	}

	signatureKey := os.Getenv("JWT_SIGNATURE_KEY")
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
