package room

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go-labs-game-platform/internal/mock_cache"
)

func TestImpl_ListenPackets(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mock_cache.NewMockService(ctrl)

	impl := New(mockRepo)

	mockRepo.EXPECT().Subscribe(gomock.Any(), gomock.Any()).Return(make(chan *redis.Message), nil)

	_, err := impl.ListenPackets(nil, uuid.New(), uuid.New())
	if err != nil {
		t.Error(err)
	}
}
