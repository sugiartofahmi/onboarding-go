package interfaces

import (
	"context"
	"time"
)

type RedisCacheInterface interface {
    Set(ctx context.Context, key string, value any, ttl time.Duration) error
    Get(ctx context.Context, key string) (string, error)
    Delete(ctx context.Context, key string) error
    Exists(ctx context.Context, key string) (bool, error)
}