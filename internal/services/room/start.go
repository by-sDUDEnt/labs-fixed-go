package room

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/cache"
)

func (i Impl) StartGame(ctx context.Context, roomID uuid.UUID) error {
	var room models.Room

	if err := i.cache.Get(ctx, cache.RoomID(roomID), &room); err != nil {
		return fmt.Errorf("cache.Get: %w", err)
	}

	if !room.CanBeStarted() {
		return fmt.Errorf("room is not ready to start")
	}

	room.Status = models.GameStatusInProgress

	firstMovePlayerID := room.PlayerIDs[0]

	startPacket := ServerStartPacket{
		Packet: Packet{
			Type: PacketTypeServerStart,
		},
		Players:           room.PlayerIDs,
		FirstMovePlayerID: firstMovePlayerID,
	}

	if err := i.Broadcast(ctx, &room, startPacket); err != nil {
		return err
	}

	room.CurrentMovePlayerID = firstMovePlayerID
	room.NextMovePlayerID = room.PlayerIDs[1]

	if err := i.saveRoom(ctx, &room); err != nil {
		return err
	}

	return nil
}
