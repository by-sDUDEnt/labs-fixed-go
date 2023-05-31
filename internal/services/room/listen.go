package room

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go-labs-game-platform/internal/services/cache"
)

func (i Impl) ListenPackets(ctx context.Context, roomID uuid.UUID, userID uuid.UUID) (<-chan *redis.Message, error) {
	fmt.Println("listen packets", cache.RoomUserChannelID(roomID, userID))
	subscribe, err := i.cache.Subscribe(ctx, cache.RoomUserChannelID(roomID, userID))
	if err != nil {
		return nil, err
	}

	return subscribe, nil
}
