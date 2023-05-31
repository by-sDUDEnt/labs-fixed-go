package room

import (
	"context"
	"encoding/json"
	"fmt"

	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/redis"
)

func (i Impl) GetList(ctx context.Context) (models.RoomList, error) {
	all, err := i.redis.HGetAll(ctx, redis.RoomList)
	if err != nil {
		return nil, fmt.Errorf("get room list from redis: %w", err)
	}

	rooms := make(models.RoomList, 0, len(all))

	for _, room := range all {
		var r models.Room

		if err = json.Unmarshal([]byte(room), &r); err != nil {
			return nil, err
		}

		rooms = append(rooms, r)
	}

	return rooms, nil
}
