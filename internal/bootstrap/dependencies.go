package bootstrap

import (
	"go-labs-game-platform/internal/services/auth"
	"go-labs-game-platform/internal/services/cache"
	"go-labs-game-platform/internal/services/room"
	"go-labs-game-platform/internal/services/user"
)

type Dependencies struct {
	CacheCli cache.Service
	AuthSrv  auth.Service
	UserSrv  user.Service
	RoomSrv  room.Service
}

func NewDependencies(redisCli cache.Service, authSrv auth.Service,
	userSrv user.Service, roomSrv room.Service) *Dependencies {
	return &Dependencies{
		CacheCli: redisCli,
		AuthSrv:  authSrv,
		UserSrv:  userSrv,
		RoomSrv:  roomSrv,
	}
}
