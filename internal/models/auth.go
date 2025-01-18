package models

import (
	"encoding/json"
	"fmt"
)

type AccessRole string

const (
	AdminRole    AccessRole = "ADMIN"
	EmployeeRole AccessRole = "EMPLOYEE"
)

type AuthenticationRequest struct {
	Username     string `json:"username" validate:"required"`
	PasswordHash string `json:"password" validate:"required,sha256"`
}

type AuthenticationResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type AuthorizationRequest struct {
	Token string `json:"token" validate:"required"`
}

type AuthorizationResponse struct {
	Success bool   `json:"success"`
	Role    string `json:"role"`
	UserID  int    `json:"user_id"`
}

type Credentials struct {
	AuthenticationRequest
	UserID int `json:"user_id"`
}

func (a AuthenticationResponse) ToJSON() []byte {
	data, _ := json.Marshal(a)
	return data
}

func NewAuthenticationRequestFromJSON(data []byte) (*AuthenticationRequest, error) {
	authReq := AuthenticationRequest{}
	err := json.Unmarshal(data, &authReq)
	if err != nil {
		return nil, err
	}
	return &authReq, nil
}

var FailedAuthenticationResponse = AuthenticationResponse{Success: false}

func NewAuthorizationRequestFromJSON(data []byte) (*AuthorizationRequest, error) {
	authReq := AuthorizationRequest{}
	err := json.Unmarshal(data, &authReq)
	if err != nil {
		return nil, err
	}
	return &authReq, nil
}

var FailedAuthorizationResponse = AuthorizationResponse{Success: false}

type ErrInvalidCredentials struct {
	message string
}

func (e ErrInvalidCredentials) Error() string {
	return fmt.Sprintf("Invalid credentials: %s", e.message)
}

type ErrAuthenticationFailed struct {
	message string
}

func (e ErrAuthenticationFailed) Error() string {
	return fmt.Sprintf("Authentication failed: %s", e.message)
}
