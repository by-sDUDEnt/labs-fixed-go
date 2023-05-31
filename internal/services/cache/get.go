package cache

import (
	"context"
	"encoding/json"
)

func (r Impl) Get(ctx context.Context, key string, dest any) error {
	result := r.cli.Get(ctx, key)
	if result.Err() != nil {
		return result.Err()
	}

	if err := json.Unmarshal([]byte(result.Val()), dest); err != nil {
		return err
	}

	r.cli.HGetAll(ctx, key)

	return nil
}

func (r Impl) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	result := r.cli.HGetAll(ctx, key)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return result.Val(), nil
}
