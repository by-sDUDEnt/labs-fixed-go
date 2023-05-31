package room

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

func (i Impl) HandlePacket(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, msg []byte) error {
	var packet Packet

	err := json.Unmarshal(msg, &packet)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	switch packet.Type {
	case PacketTypeClientStart:
		var clientStartPacket ClientStartPacket

		if err = json.Unmarshal(msg, &clientStartPacket); err != nil {
			return fmt.Errorf("json.Unmarshal: %w", err)
		}

		if err = i.startGame(ctx, userID, roomID); err != nil {
			return fmt.Errorf("roomSrv.startGame: %w", err)
		}

		fmt.Println("start")
	case PacketTypeClientMove:
		var clientMovePacket ClientMovePacket

		if err = json.Unmarshal(msg, &clientMovePacket); err != nil {
			return fmt.Errorf("json.Unmarshal: %w", err)
		}

		if err = i.move(ctx, userID, roomID, clientMovePacket); err != nil {
			return fmt.Errorf("roomSrv.move: %w", err)
		}
	}

	return nil
}
