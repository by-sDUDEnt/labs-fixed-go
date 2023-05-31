package room

import (
	"github.com/google/uuid"
)

const (
	PacketTypeServerStart = iota
	PacketTypeClientStart
	PacketTypeClientMove
	PacketTypeServerMove
	PacketTypeEnd
)

type (
	PacketType int

	Packet struct {
		Type PacketType `json:"type"`
	}

	ClientStartPacket struct {
		Packet
	}

	ServerStartPacket struct {
		Packet

		Players           []uuid.UUID `json:"players"`
		FirstMovePlayerID uuid.UUID   `json:"first_move_player_id"`
	}

	ClientMovePacket struct {
		Packet

		Position uint8 `json:"position"`
	}

	ServerMovePacket struct {
		Packet

		PlayerID uuid.UUID `json:"player_id"`
		Position uint8     `json:"position"`
	}

	EndPacket struct {
		Packet

		WinPlayerID uuid.UUID `json:"win_player_id"`
	}
)
