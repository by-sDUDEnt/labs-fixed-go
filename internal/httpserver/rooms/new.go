package rooms

import (
	"go-labs-game-platform/internal/services/room"
)

type Handlers struct {
	srv room.Service
}

func New(srv room.Service) *Handlers {
	return &Handlers{srv: srv}
}
