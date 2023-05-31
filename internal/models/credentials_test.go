package models

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCredentials(t *testing.T) {
	credentials := Credentials{
		UserID:    uuid.New(),
		Type:      CredentialTypePassword,
		Password:  []byte("password123"),
		CreatedAt: time.Now(),
	}

	// Test for UserID
	if credentials.UserID == uuid.Nil {
		t.Error("UserID should not be nil")
	}

	// Test for Type
	if credentials.Type != CredentialTypePassword {
		t.Error("Type should be CredentialTypePassword")
	}

	// Test for Password
	expectedPassword := []byte("password123")
	if string(credentials.Password) != string(expectedPassword) {
		t.Errorf("Expected password: %s, got: %s", string(expectedPassword), string(credentials.Password))
	}

	// Test for CreatedAt
	if credentials.CreatedAt.IsZero() {
		t.Error("CreatedAt should not be zero")
	}
}

func TestTokenResponse(t *testing.T) {
	token := "abc123"
	tokenResponse := TokenResponse{Token: token}

	// Test for Token
	if tokenResponse.Token != token {
		t.Errorf("Expected token: %s, got: %s", token, tokenResponse.Token)
	}
}

func TestRegisterCredentials(t *testing.T) {
	username := "john_doe"
	password := "pass123"
	registerCredentials := RegisterCredentials{Username: username, Password: password}

	// Test for Username
	if registerCredentials.Username != username {
		t.Errorf("Expected username: %s, got: %s", username, registerCredentials.Username)
	}

	// Test for Password
	if registerCredentials.Password != password {
		t.Errorf("Expected password: %s, got: %s", password, registerCredentials.Password)
	}
}

func TestLoginCredentials(t *testing.T) {
	username := "john_doe"
	password := "pass123"
	loginCredentials := LoginCredentials{Username: username, Password: password}

	// Test for Username
	if loginCredentials.Username != username {
		t.Errorf("Expected username: %s, got: %s", username, loginCredentials.Username)
	}

	// Test for Password
	if loginCredentials.Password != password {
		t.Errorf("Expected password: %s, got: %s", password, loginCredentials.Password)
	}
}
