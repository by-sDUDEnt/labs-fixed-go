package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (r Impl) Subscribe(ctx context.Context, channel string) (<-chan *redis.Message, error) {
	sub := r.cli.Subscribe(ctx, channel)

	return sub.Channel(), nil
}

func (r Impl) Publish(ctx context.Context, channel string, message any) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	result := r.cli.Publish(ctx, channel, messageJSON)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
