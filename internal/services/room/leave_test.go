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

func TestImpl_Leave(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mock_cache.NewMockService(ctrl)

	impl := New(mockRepo)

	roomID := uuid.New()
	userID := uuid.New()

	mockRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, key string, value interface{}) error {
			// remove room: prefix
			parsedRoomID := uuid.MustParse(key[5:])
			*value.(*models.Room) = models.Room{
				ID:        parsedRoomID,
				PlayerIDs: []uuid.UUID{userID},
			}
			return nil
		})
	mockRepo.EXPECT().Del(gomock.Any(), cache.RoomID(roomID)).Return(nil)
	mockRepo.EXPECT().HDel(gomock.Any(), cache.RoomList, roomID.String()).Return(nil)

	err := impl.Leave(context.Background(), userID, roomID)
	if err != nil {
		t.Fatal(err)
	}
}
