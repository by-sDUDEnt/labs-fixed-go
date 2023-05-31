package room

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/cache"
)

func (i Impl) GetByID(ctx context.Context, roomID uuid.UUID) (*models.Room, error) {
	var room models.Room

	if err := i.cache.Get(ctx, cache.RoomID(roomID), &room); err != nil {
		return nil, fmt.Errorf("get room from cache: %w", err)
	}

	return &room, nil
}

func (i Impl) GetList(ctx context.Context) (models.RoomList, error) {
	all, err := i.cache.HGetAll(ctx, cache.RoomList)
	if err != nil {
		return nil, fmt.Errorf("get room list from cache: %w", err)
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
