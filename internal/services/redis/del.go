package redis

import "context"

func (r Redis) Del(ctx context.Context, key string) error {
	result := r.cli.Del(ctx, key)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

func (r Redis) HDel(ctx context.Context, key string, fields ...string) error {
	result := r.cli.HDel(ctx, key, fields...)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
