package room

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/redis"
)

func (i Impl) GetByID(ctx context.Context, roomID uuid.UUID) (*models.Room, error) {
	var room models.Room

	if err := i.redis.Get(ctx, redis.RoomID(roomID), &room); err != nil {
		return nil, fmt.Errorf("get room from redis: %w", err)
	}

	return &room, nil
}
