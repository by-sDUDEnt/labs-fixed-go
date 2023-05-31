package room

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/cache"
)

func (i Impl) Move(ctx context.Context, userID uuid.UUID, id uuid.UUID, move ClientMovePacket) error {
	var room models.Room

	if err := i.cache.Get(ctx, cache.RoomID(id), &room); err != nil {
		return err
	}

	if !room.ContainsPlayer(userID) {
		return ErrRoomDoesNotContainPlayer
	}

	if room.CurrentMovePlayerID != userID {
		return ErrNotYourTurn
	}

	room.CurrentMovePlayerID = room.NextMovePlayerID
	room.NextMovePlayerID = userID

	row := move.Position / 7
	col := move.Position % 7

	if room.Table[row][col] != 0 {
		return ErrInvalidMove
	}

	room.Table[row][col] = room.IndexOfPlayer(userID) + 1

	var packet any = ServerMovePacket{
		Packet: Packet{
			Type: PacketTypeServerMove,
		},
		PlayerID: userID,
		Position: move.Position,
	}

	if err := i.Broadcast(ctx, &room, packet); err != nil {
		return fmt.Errorf("broadcast: %w", err)
	}

	if room.HasWinner() {
		packet = EndPacket{
			Packet: Packet{
				Type: PacketTypeEnd,
			},
			WinPlayerID: userID,
		}
		if err := i.Broadcast(ctx, &room, packet); err != nil {
			return fmt.Errorf("broadcast: %w", err)
		}

		room.Status = models.GameStatusFinished
	}

	if err := i.saveRoom(ctx, &room); err != nil { // TODO move to saveRoom added by me
		return fmt.Errorf("saveRoom: %w", err)
	}

	return nil
}

func (i Impl) Broadcast(ctx context.Context, room *models.Room, packet any) error {
	for _, playerID := range room.PlayerIDs {
		if err := i.cache.Publish(ctx, cache.RoomUserChannelID(room.ID, playerID), packet); err != nil {
			return fmt.Errorf("broadcast: %w", err)
		}
	}

	return nil
}
