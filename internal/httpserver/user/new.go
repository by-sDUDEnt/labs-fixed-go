package user

import (
	"go-labs-game-platform/internal/services/user"
)

type Handlers struct {
	srv user.Service
}

func New(srv user.Service) *Handlers {
	return &Handlers{srv: srv}
}
