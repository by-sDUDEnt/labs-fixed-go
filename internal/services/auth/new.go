package auth

import (
	"context"

	"github.com/google/uuid"

	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/auth/tokens"
)

type Service interface {
	GetToken(ctx context.Context, plaintext string) (models.Token, error)
	UpdateLastVisitedAt(ctx context.Context, plaintext string) error

	Register(ctx context.Context, m *models.RegisterCredentials) (models.TokenResponse, error)
	Login(ctx context.Context, m *models.LoginCredentials) (models.TokenResponse, error)

	//GenerateToken(ctx context.Context, userID uuid.UUID, scope string) (string, error)
	//DeleteToken(ctx context.Context, userID uuid.UUID, scope string) error
	//VerifyToken(ctx context.Context, token string) error
}

type Repo interface {
	models.TXer
	tokens.Repo

	UserByUsername(ctx context.Context, username string) (*models.User, error)
	UserExistsByUsername(ctx context.Context, username string) (bool, error)

	CreateUser(ctx context.Context, user *models.User) error
	CreateCredentials(ctx context.Context, m *models.Credentials) error
	CredentialsByUserID(ctx context.Context, id uuid.UUID) (*models.Credentials, error)
}

type Impl struct {
	repo      Repo
	tokensSrv *tokens.Impl
}

func New(repo Repo) *Impl {
	return &Impl{
		repo:      repo,
		tokensSrv: tokens.New(repo),
	}
}
