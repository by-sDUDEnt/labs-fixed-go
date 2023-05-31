package room

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"go-labs-game-platform/internal/mock_cache"
)

func TestImpl_Create(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mock_cache.NewMockService(ctrl)

	impl := New(mockRepo)

	userID := uuid.New()

	mockRepo.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockRepo.EXPECT().HSet(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	create, err := impl.Create(context.Background(), userID)
	if err != nil {
		t.Fatal(err)
	}

	if create.ID == uuid.Nil {
		t.Fatal("room id is nil")
	}
}
