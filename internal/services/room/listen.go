package room

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	redis2 "github.com/redis/go-redis/v9"
	"go-labs-game-platform/internal/services/redis"
)

func (i Impl) ListenPackets(ctx context.Context, roomID uuid.UUID, userID uuid.UUID) (<-chan *redis2.Message, error) {
	fmt.Println("listen packets", redis.RoomUserChannelID(roomID, userID))
	subscribe, err := i.redis.Subscribe(ctx, redis.RoomUserChannelID(roomID, userID))
	if err != nil {
		return nil, err
	}

	return subscribe, nil
}
