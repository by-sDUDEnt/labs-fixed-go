package auth

import (
	"go-labs-game-platform/internal/services/auth"
	"go-labs-game-platform/internal/services/auth/tokens"
)

type Handlers struct {
	tokenSrv tokens.Service
	authSrv  auth.Service
}

func New(authSrv auth.Service) *Handlers {
	return &Handlers{
		authSrv: authSrv,
	}
}
