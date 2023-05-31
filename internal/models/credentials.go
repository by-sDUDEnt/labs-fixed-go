package models

import (
	"time"

	"github.com/google/uuid"
)

type CredentialType string

const (
	CredentialTypePassword CredentialType = "password"
)

type Credentials struct {
	UserID    uuid.UUID      `pg:",pk"`
	Type      CredentialType `pg:",pk"`
	Password  []byte         `pg:",notnull"`
	CreatedAt time.Time      `pg:",notnull,default:now()"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type RegisterCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
