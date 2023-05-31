package room

import (
	"context"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/cache"
)

type Service interface {
	GetByID(ctx context.Context, roomID uuid.UUID) (*models.Room, error)
	GetList(ctx context.Context) (models.RoomList, error)
	Create(ctx context.Context, userID uuid.UUID) (*models.Room, error)
	Join(ctx context.Context, userID, roomID uuid.UUID) (*models.Room, error)
	Leave(ctx context.Context, userID, roomID uuid.UUID) error

	HandlePacket(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, msg []byte) error
	ListenPackets(ctx context.Context, roomID uuid.UUID, userID uuid.UUID) (<-chan *redis.Message, error)
}

type Impl struct {
	cache cache.Service
}

func New(cache cache.Service) *Impl {
	return &Impl{
		cache: cache,
	}
}
