package services

import (
	"context"
	"time"

	"event-backend/infrastructure/redis/interfaces"

	"github.com/redis/go-redis/v9"
)

type RedisCacheService struct {
    client *redis.Client
}

func NewRedisCacheService(client *redis.Client) interfaces.RedisCacheInterface {
    return &RedisCacheService{client: client}
}

func (r *RedisCacheService) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
    return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisCacheService) Get(ctx context.Context, key string) (string, error) {
    return r.client.Get(ctx, key).Result()
}

func (r *RedisCacheService) Delete(ctx context.Context, key string) error {
    return r.client.Del(ctx, key).Err()
}

func (r *RedisCacheService) Exists(ctx context.Context, key string) (bool, error) {
    result, err := r.client.Exists(ctx, key).Result()
    if err != nil {
        return false, err
    }
    return result > 0, nil
}