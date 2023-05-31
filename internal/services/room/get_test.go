package room

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"go-labs-game-platform/internal/mock_cache"
	"go-labs-game-platform/internal/models"
)

func TestImpl_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mock_cache.NewMockService(ctrl)

	impl := New(mockRepo)

	roomID := uuid.New()

	room := models.Room{
		ID: roomID,
	}

	mockRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, key string, value interface{}) error {
			*value.(*models.Room) = room
			return nil
		})

	get, err := impl.GetByID(context.Background(), roomID)
	if err != nil {
		t.Fatal(err)
	}

	if get.ID == uuid.Nil {
		t.Fatal("room id is nil")
	}
}
func TestImpl_GetByID_failed(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mock_cache.NewMockService(ctrl)

	impl := New(mockRepo)

	room := models.Room{
		ID: uuid.New(),
	}

	mockRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, key string, value interface{}) error {
			// remove room: prefix
			roomID := uuid.MustParse(key[5:])
			*value.(*models.Room) = models.Room{ID: roomID}
			return nil
		})

	get, err := impl.GetByID(context.Background(), uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	if get.ID == room.ID {
		t.Fatal("room id must not be equal")
	}
}
