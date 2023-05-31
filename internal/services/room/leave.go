package room

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/redis"
)

func (i Impl) Leave(ctx context.Context, userID, roomID uuid.UUID) error {
	var room models.Room

	if err := i.redis.Get(ctx, redis.RoomID(roomID), &room); err != nil {
		return fmt.Errorf("get room from redis: %w", err)
	}

	if !room.RemovePlayer(userID) {
		return ErrRoomIsFull
	}

	if room.IsEmpty() {
		if err := i.redis.Del(ctx, redis.RoomID(room.ID)); err != nil {
			return err
		}

		if err := i.redis.HDel(ctx, redis.RoomList, room.ID.String()); err != nil {
			return err
		}

		return nil
	}

	if err := i.saveRoom(ctx, &room); err != nil {
		return err
	}

	return nil
}
