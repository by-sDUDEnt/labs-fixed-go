package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"go-labs-game-platform/internal/config"
)

type Redis struct {
	cli *redis.Client
}

func New() (Redis, error) {
	cfg := config.Get().Redis
	return Redis{
		redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Password: cfg.Password,
		}),
	}, nil
}
