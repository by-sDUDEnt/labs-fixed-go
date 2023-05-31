package room

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/cache"
)

func (i Impl) Join(ctx context.Context, userID, roomID uuid.UUID) (*models.Room, error) {
	var room models.Room

	if err := i.cache.Get(ctx, cache.RoomID(roomID), &room); err != nil {
		return nil, fmt.Errorf("get room from cache: %w", err)
	}

	if room.IsFull() {
		return nil, ErrRoomIsFull
	}

	if room.ContainsPlayer(userID) {
		return nil, ErrRoomAlreadyContainsPlayer
	}

	room.AddPlayer(userID)

	if err := i.saveRoom(ctx, &room); err != nil {
		return nil, err
	}

	return &room, nil
}
