package redis

import "github.com/google/uuid"

const (
	RoomList = "roomList"
)

func RoomID(id uuid.UUID) string {
	return "room:" + id.String()
}

func RoomUserChannelID(roomID, userID uuid.UUID) string {
	return "room:" + roomID.String() + ":user:" + userID.String()
}
