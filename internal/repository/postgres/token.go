package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"

	"go-labs-game-platform/internal/models"
)

func (r *Repo) GetToken(ctx context.Context, hash []byte) (models.Token, error) {
	var token models.Token
	err := r.getQuery(ctx, &token).
		Where("hash = ?", hash).First()

	return token, r.err(err)
}

func (r *Repo) GetAllTokens(ctx context.Context) ([]models.Token, error) {
	var tokens []models.Token
	err := r.getQuery(ctx, &tokens).
		Order("hash").
		Select()
	return tokens, r.err(err)
}

func (r *Repo) SaveToken(ctx context.Context, token *models.Token) error {
	_, err := r.getQuery(ctx, token).
		Insert()
	return r.err(err)
}

func (r *Repo) UpdateLastVisitedAt(ctx context.Context, hash []byte, lastVisitedAt time.Time) error {
	res, err := r.getQuery(ctx, &models.Token{}).
		Set("last_visited_at = ?", lastVisitedAt).
		Where("hash = ?", hash).
		Update()
	if err != nil {
		return r.err(err)
	}

	if res.RowsAffected() != 1 {
		return models.ErrNotFound
	}

	return nil
}

func (r *Repo) DeleteToken(ctx context.Context, userID uuid.UUID, scope string) error {
	_, err := r.getQuery(ctx, &models.Token{}).
		Where("user_id = ?", userID).
		Where("scope = ?", scope).Delete()
	return r.err(err)
}
