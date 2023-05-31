package tokens

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"go-labs-game-platform/internal/models"
)

type Service interface {
	GetToken(ctx context.Context, plaintext string) (models.Token, error)
	GenerateToken(ctx context.Context, userID uuid.UUID, scope string) (string, error)
	UpdateLastVisitedAt(ctx context.Context, plaintext string) error
	DeleteToken(ctx context.Context, userID uuid.UUID, scope string) error
	VerifyToken(ctx context.Context, token string) error
}

type Repo interface {
	GetToken(ctx context.Context, hash []byte) (models.Token, error)
	GetAllTokens(ctx context.Context) ([]models.Token, error)
	SaveToken(ctx context.Context, token *models.Token) error
	UpdateLastVisitedAt(ctx context.Context, hash []byte, lastVisitedAt time.Time) error
	DeleteToken(ctx context.Context, userID uuid.UUID, scope string) error
}

type Impl struct {
	repo Repo
}

func New(repo Repo) *Impl {
	return &Impl{
		repo: repo,
	}
}

func (s Impl) GetToken(ctx context.Context, plaintext string) (models.Token, error) {
	hash := s.HashToken(plaintext)
	token, err := s.repo.GetToken(ctx, hash)
	if errors.Is(err, models.ErrNotFound) {
		return models.Token{}, models.ErrUnauthorized
	}
	if err != nil {
		return models.Token{}, err
	}

	return token, nil
}

func (s Impl) GenerateToken(ctx context.Context, userID uuid.UUID, scope string) (string, error) {
	if err := s.DeleteToken(ctx, userID, scope); err != nil {
		return "", err
	}

	tokenPlaintext, tokenHash, err := s.CreateToken()
	if err != nil {
		return "", fmt.Errorf("create token: %w", err)
	}

	token := models.Token{
		Hash:          tokenHash,
		Scope:         scope,
		UserID:        userID,
		CreatedAt:     time.Now(),
		LastVisitedAt: time.Now(),
	}

	if err = s.repo.SaveToken(ctx, &token); err != nil {
		return "", fmt.Errorf("save token: %w", err)
	}

	return tokenPlaintext, nil
}

func (s Impl) UpdateLastVisitedAt(ctx context.Context, plaintext string) error {
	hash := s.HashToken(plaintext)
	return s.repo.UpdateLastVisitedAt(ctx, hash, time.Now())
}

func (s Impl) DeleteToken(ctx context.Context, userID uuid.UUID, scope string) error {
	if err := s.repo.DeleteToken(ctx, userID, scope); err != nil {
		return fmt.Errorf("delete token for users %s: %w", userID, err)
	}

	return nil
}

func (s Impl) VerifyToken(ctx context.Context, tokenText string) error {
	token, err := s.GetToken(ctx, tokenText)
	if errors.Is(err, models.ErrNotFound) {
		return models.ErrUnauthorized
	}
	if err != nil {
		return fmt.Errorf("get token: %w", err)
	}
	if !token.IsValid() {
		return models.ErrUnauthorized
	}

	return nil
}
