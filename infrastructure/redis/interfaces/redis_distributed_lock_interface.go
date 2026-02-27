package interfaces

import (
	"context"
	"time"
)

type RedisDistributedLockInterface interface {
    Acquire(ctx context.Context, key string, ttl time.Duration) bool
    Release(ctx context.Context, key string) bool
}