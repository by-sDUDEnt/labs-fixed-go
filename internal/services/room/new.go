package room

import (
	"context"

	"github.com/google/uuid"
	redis2 "github.com/redis/go-redis/v9"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/redis"
)

type Service interface {
	GetByID(ctx context.Context, roomID uuid.UUID) (*models.Room, error)
	GetList(ctx context.Context) (models.RoomList, error)
	Create(ctx context.Context, userID uuid.UUID) (*models.Room, error)
	Join(ctx context.Context, userID, roomID uuid.UUID) (*models.Room, error)
	Leave(ctx context.Context, userID, roomID uuid.UUID) error

	HandlePacket(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, msg []byte) error
	ListenPackets(ctx context.Context, roomID uuid.UUID, userID uuid.UUID) (<-chan *redis2.Message, error)
}

type Repo interface {
	//UserByID(ctx context.Context, uuid uuid.UUID) (*models.User, error)
}

type Impl struct {
	repo  Repo
	redis redis.Redis
}

func New(repo Repo, redis redis.Redis) *Impl {
	return &Impl{
		repo:  repo,
		redis: redis,
	}
}
