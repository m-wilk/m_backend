package usecase

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type logout struct {
	RedisClient *redis.Client
}

func NewLogout(redis *redis.Client) logout {
	return logout{RedisClient: redis}
}

func (l *logout) Base(userId string) error {
	l.RedisClient.Del(context.Background(), userId)
	return nil
}
