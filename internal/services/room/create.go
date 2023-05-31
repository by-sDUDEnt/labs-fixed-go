package room

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/cache"
)

func (i Impl) Create(ctx context.Context, userID uuid.UUID) (*models.Room, error) {
	room := &models.Room{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		GameType:  models.GameTypeFourInARow,
		PlayerIDs: []uuid.UUID{userID},
	}

	if err := i.saveRoom(ctx, room); err != nil {
		return nil, err
	}

	return room, nil
}

func (i Impl) saveRoom(ctx context.Context, room *models.Room) error {
	if err := i.cache.Set(ctx, cache.RoomID(room.ID), room, 0); err != nil {
		return err
	}

	if err := i.cache.HSet(ctx, cache.RoomList, room.ID.String(), room); err != nil {
		return err
	}

	return nil
}
