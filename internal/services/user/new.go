package user

import (
	"context"

	"github.com/google/uuid"

	"go-labs-game-platform/internal/models"
)

type Service interface {
	ByID(ctx context.Context, uuid uuid.UUID) (*models.User, error)
}

type Repo interface {
	UserByID(ctx context.Context, uuid uuid.UUID) (*models.User, error)
}

type Impl struct {
	repo Repo
}

func New(repo Repo) *Impl {
	return &Impl{
		repo: repo,
	}
}
