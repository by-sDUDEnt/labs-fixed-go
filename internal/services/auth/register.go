package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"go-labs-game-platform/internal/models"
)

func (srv Impl) Register(ctx context.Context, m *models.RegisterCredentials) (models.TokenResponse, error) {
	exists, err := srv.repo.UserExistsByUsername(ctx, m.Username)
	if err != nil {
		return models.TokenResponse{}, fmt.Errorf("check user exists: %w", err)
	}

	if exists {
		return models.TokenResponse{}, models.ErrUserAlreadyExists
	}

	user := models.User{
		ID:        uuid.New(),
		Username:  m.Username,
		CreatedAt: time.Now(),
	}

	ctx, err = srv.repo.Begin(ctx)
	if err != nil {
		return models.TokenResponse{}, fmt.Errorf("begin transaction: %w", err)
	}

	if err = srv.repo.CreateUser(ctx, &user); err != nil {
		return models.TokenResponse{}, fmt.Errorf("create user: %w", err)
	}

	password, err := srv.tokensSrv.HashPassword(m.Password)
	if err != nil {
		return models.TokenResponse{}, fmt.Errorf("hash password: %w", err)
	}

	if err = srv.repo.CreateCredentials(ctx, &models.Credentials{
		UserID:    user.ID,
		Type:      models.CredentialTypePassword,
		Password:  password,
		CreatedAt: time.Now(),
	}); err != nil {
		return models.TokenResponse{}, fmt.Errorf("create credentials: %w", err)
	}

	token, err := srv.tokensSrv.GenerateToken(ctx, user.ID, models.ScopeSessionUser)
	if err != nil {
		return models.TokenResponse{}, fmt.Errorf("generate token: %w", err)
	}

	if err = srv.repo.Commit(ctx); err != nil {
		return models.TokenResponse{}, fmt.Errorf("commit transaction: %w", err)
	}

	return models.TokenResponse{Token: token}, nil
}
