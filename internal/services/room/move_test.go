package room

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"go-labs-game-platform/internal/mock_cache"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/cache"
)

func TestImpl_Move(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mock_cache.NewMockService(ctrl)

	impl := New(mockRepo)

	userID1 := uuid.New()
	userID2 := uuid.New()

	room := models.Room{
		ID: uuid.New(),
		PlayerIDs: []uuid.UUID{
			userID1,
			userID2,
		},
		CurrentMovePlayerID: userID1,
		NextMovePlayerID:    userID2,
		Table: [6][7]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
	}

	newRoom := room
	newRoom.CurrentMovePlayerID = userID2
	newRoom.NextMovePlayerID = userID1
	newRoom.Table = [6][7]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0},
	}

	mockRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, key string, value interface{}) error {
			*value.(*models.Room) = room
			return nil
		})
	mockRepo.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mockRepo.EXPECT().Set(gomock.Any(), cache.RoomID(room.ID), &newRoom, gomock.Any()).Return(nil)
	mockRepo.EXPECT().HSet(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	err := impl.Move(context.Background(), userID1, room.ID, ClientMovePacket{
		Packet:   Packet{Type: PacketTypeClientMove},
		Position: 36,
	})
	if err != nil {
		t.Fatal(err)
	}
}
