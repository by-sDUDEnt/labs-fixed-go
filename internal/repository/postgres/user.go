package postgres

import (
	"context"

	"github.com/google/uuid"

	"go-labs-game-platform/internal/models"
)

func (r *Repo) UserExistsByUsername(ctx context.Context, username string) (bool, error) {
	exists, err := r.getQuery(ctx, (*models.User)(nil)).
		Where("username = ?", username).
		Exists()
	if err != nil {
		return false, r.err(err)
	}

	return exists, nil
}

func (r *Repo) UserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.getQuery(ctx, &user).
		Where("id = ?", id).
		Select()
	if err != nil {
		return nil, r.err(err)
	}

	return &user, nil
}

func (r *Repo) UserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.getQuery(ctx, &user).
		Where("username = ?", username).
		Select()
	if err != nil {
		return nil, r.err(err)
	}

	return &user, nil
}

func (r *Repo) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.getQuery(ctx, user).Insert()
	if err != nil {
		return r.err(err)
	}

	return nil
}
