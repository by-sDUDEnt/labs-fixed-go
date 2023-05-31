package postgres

import (
	"context"

	"github.com/google/uuid"

	"go-labs-game-platform/internal/models"
)

func (r *Repo) CreateCredentials(ctx context.Context, m *models.Credentials) error {
	_, err := r.getQuery(ctx, m).Insert()
	return r.err(err)
}

func (r *Repo) CredentialsByUserID(ctx context.Context, userID uuid.UUID) (*models.Credentials, error) {
	var credentials models.Credentials
	err := r.getQuery(ctx, &credentials).
		Where("user_id = ?", userID).
		Select()
	if err != nil {
		return nil, r.err(err)
	}

	return &credentials, nil
}
