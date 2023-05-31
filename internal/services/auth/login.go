package auth

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"go-labs-game-platform/internal/models"
)

func (srv Impl) Login(ctx context.Context, m *models.LoginCredentials) (models.TokenResponse, error) {
	user, err := srv.repo.UserByUsername(ctx, m.Username)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return models.TokenResponse{}, models.ErrUserNotFound
		}

		return models.TokenResponse{}, fmt.Errorf("get user by username: %w", err)
	}

	credentials, err := srv.repo.CredentialsByUserID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return models.TokenResponse{}, models.ErrUserNotFound
		}

		return models.TokenResponse{}, fmt.Errorf("get credentials by user id: %w", err)
	}

	ok, err := srv.tokensSrv.PasswordMatches(m.Password, credentials.Password)
	if err != nil || !ok {
		if !ok || errors.Is(err, models.ErrNotFound) {
			return models.TokenResponse{}, models.ErrUserNotFound
		}

		return models.TokenResponse{}, fmt.Errorf("password matches: %w", err)
	}

	token, err := srv.tokensSrv.GenerateToken(ctx, user.ID, models.ScopeSessionUser)
	if err != nil {
		return models.TokenResponse{}, err
	}

	return models.TokenResponse{Token: token}, nil
}
