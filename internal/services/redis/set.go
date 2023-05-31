package redis

import (
	"context"
	"encoding/json"
	"time"
)

func (r Redis) Set(ctx context.Context, key string, val any, exp time.Duration) error {
	jsonData, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return r.cli.Set(ctx, key, jsonData, exp).Err()
}

func (r Redis) HSet(ctx context.Context, key string, id, val any) error {
	valJSON, err := json.Marshal(val)
	if err != nil {
		return err
	}

	result := r.cli.HSet(ctx, key, id, valJSON)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
