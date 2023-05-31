package room

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/cache"
)

func (i Impl) Leave(ctx context.Context, userID, roomID uuid.UUID) error {
	var room models.Room

	if err := i.cache.Get(ctx, cache.RoomID(roomID), &room); err != nil {
		return fmt.Errorf("get room from cache: %w", err)
	}

	if !room.RemovePlayer(userID) {
		return ErrRoomDoesNotContainPlayer
	}

	if room.IsEmpty() {
		if err := i.cache.Del(ctx, cache.RoomID(room.ID)); err != nil {
			return err
		}

		if err := i.cache.HDel(ctx, cache.RoomList, room.ID.String()); err != nil {
			return err
		}

		return nil
	}

	if err := i.saveRoom(ctx, &room); err != nil {
		return err
	}

	return nil
}
