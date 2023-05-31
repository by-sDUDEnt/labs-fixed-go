//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"go-labs-game-platform/internal/services/redis"
	"go-labs-game-platform/internal/services/room"

	"go-labs-game-platform/internal/repository/postgres"
	"go-labs-game-platform/internal/services/auth"
	"go-labs-game-platform/internal/services/user"
)

func Up() (*Dependencies, error) {
	wire.Build(
		wire.Bind(new(auth.Service), new(*auth.Impl)),
		wire.Bind(new(user.Service), new(*user.Impl)),
		wire.Bind(new(room.Service), new(*room.Impl)),

		wire.Bind(new(auth.Repo), new(*postgres.Repo)),
		wire.Bind(new(user.Repo), new(*postgres.Repo)),
		wire.Bind(new(room.Repo), new(*postgres.Repo)),

		redis.New,
		postgres.New,
		auth.New,
		user.New,
		room.New,
		NewDependencies,
	)
	return &Dependencies{}, nil
}
