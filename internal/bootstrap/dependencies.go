package bootstrap

import (
	"go-labs-game-platform/internal/services/auth"
	"go-labs-game-platform/internal/services/redis"
	"go-labs-game-platform/internal/services/room"
	"go-labs-game-platform/internal/services/user"
)

type Dependencies struct {
	RedisCli redis.Redis

	AuthSrv auth.Service
	UserSrv user.Service
	RoomSrv room.Service
}

func NewDependencies(redisCli redis.Redis, authSrv auth.Service,
	userSrv user.Service, roomSrv room.Service) *Dependencies {
	return &Dependencies{
		RedisCli: redisCli,
		AuthSrv:  authSrv,
		UserSrv:  userSrv,
		RoomSrv:  roomSrv,
	}
}
