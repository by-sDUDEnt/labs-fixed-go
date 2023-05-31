package auth

import (
	"context"

	"go-labs-game-platform/internal/models"
)

func (srv Impl) GetToken(ctx context.Context, plaintext string) (models.Token, error) {
	return srv.tokensSrv.GetToken(ctx, plaintext)
}

func (srv Impl) UpdateLastVisitedAt(ctx context.Context, plaintext string) error {
	return srv.tokensSrv.UpdateLastVisitedAt(ctx, plaintext)
}
