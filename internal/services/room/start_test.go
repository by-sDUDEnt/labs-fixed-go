package room

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"go-labs-game-platform/internal/mock_cache"
	"go-labs-game-platform/internal/models"
)

func TestImpl_StartGame(t *testing.T) {
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
	}

	mockRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, key string, value interface{}) error {
			*value.(*models.Room) = room
			return nil
		})
	mockRepo.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockRepo.EXPECT().HSet(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockRepo.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	err := impl.StartGame(context.Background(), room.ID)
	if err != nil {
		t.Fatal(err)
	}
}
