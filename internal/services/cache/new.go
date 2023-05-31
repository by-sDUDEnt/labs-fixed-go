package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go-labs-game-platform/internal/config"
)

type Service interface {
	Set(ctx context.Context, key string, val any, exp time.Duration) error
	HSet(ctx context.Context, key string, id, val any) error
	Get(ctx context.Context, key string, dest any) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	Del(ctx context.Context, key string) error
	HDel(ctx context.Context, key string, fields ...string) error
	Subscribe(ctx context.Context, channel string) (<-chan *redis.Message, error)
	Publish(ctx context.Context, channel string, message any) error
}

type Impl struct {
	cli *redis.Client
}

func New() (*Impl, error) {
	cfg := config.Get().Redis
	return &Impl{
		redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Password: cfg.Password,
		}),
	}, nil
}
