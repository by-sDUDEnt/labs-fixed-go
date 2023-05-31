package room

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/redis"
)

func (i Impl) startGame(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error {
	var room models.Room

	if err := i.redis.Get(ctx, redis.RoomID(roomID), &room); err != nil {
		return fmt.Errorf("redis.Get: %w", err)
	}

	if !room.CanBeStarted() {
		return fmt.Errorf("room is not ready to start")
	}

	room.Status = models.GameStatusInProgress

	//firstMovePlayerID := room.PlayerIDs[rand.Intn(1)] TODO revert
	firstMovePlayerID := room.PlayerIDs[0]
	fmt.Println("publish packet", redis.RoomUserChannelID(roomID, userID))

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
	room.NextMovePlayerID = room.PlayerIDs[1] // TODO change to oposite player

	if err := i.saveRoom(ctx, &room); err != nil {
		return err
	}

	return nil
}
