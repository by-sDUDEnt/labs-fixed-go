package cache

import (
	"context"
	"encoding/json"
	"time"
)

func (r Impl) Set(ctx context.Context, key string, val any, exp time.Duration) error {
	jsonData, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return r.cli.Set(ctx, key, jsonData, exp).Err()
}

func (r Impl) HSet(ctx context.Context, key string, id, val any) error {
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
