package models

import (
	"time"

	"github.com/google/uuid"

	"go-labs-game-platform/internal/config"
)

const (
	ScopeSessionUser = "session_user"
)

type Token struct {
	Hash          []byte `pg:",pk"`
	UserID        uuid.UUID
	Scope         string
	CreatedAt     time.Time
	LastVisitedAt time.Time
}

func (t *Token) IsValid() bool {
	now := time.Now()

	switch t.Scope {
	case ScopeSessionUser:
		absoluteTimeout := config.Get().Security.UserAbsoluteTimeout
		idleTimeout := config.Get().Security.UserIdleTimeout
		return now.Before(t.CreatedAt.Add(absoluteTimeout)) && now.Before(t.LastVisitedAt.Add(idleTimeout))
	}

	return false
}

type LoginResponse struct {
	ID          uuid.UUID `json:"id"`
	Token       string    `json:"token"`
	AgreedTerms *bool     `json:"agreed_terms,omitempty"`
}
