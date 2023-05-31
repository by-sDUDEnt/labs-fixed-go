package user

import (
	"context"

	"github.com/google/uuid"

	"go-labs-game-platform/internal/models"
)

func (s *Impl) ByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	if id == uuid.Nil {
		return nil, models.ErrNotFound
	}

	return s.repo.UserByID(ctx, id)
}
