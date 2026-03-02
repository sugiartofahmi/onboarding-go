package services

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"event-backend/infrastructure/redis/interfaces"

	"github.com/redis/go-redis/v9"
)

const luaRelease = `
if redis.call('GET', KEYS[1]) == ARGV[1] then
    return redis.call('DEL', KEYS[1])
else
    return 0
end`

type RedisDistributedLockService struct {
	client     *redis.Client
	instanceId string
}

func NewRedisDistributedLockService(client *redis.Client) interfaces.RedisDistributedLockInterface {
	hostname, _ := os.Hostname()
	instanceId := fmt.Sprintf("%s-%d-%d", hostname, os.Getpid(), rand.Int63())
	return &RedisDistributedLockService{client: client, instanceId: instanceId}
}

func (r *RedisDistributedLockService) Acquire(ctx context.Context, key string, ttl time.Duration) bool {
	acquired, err := r.client.SetNX(ctx, key, r.instanceId, ttl).Result()
	if err != nil {
		log.Printf("failed to acquire lock for key %q: %v", key, err)
		return false
	}
	return acquired
}

func (r *RedisDistributedLockService) Release(ctx context.Context, key string) bool {
	result, err := r.client.Eval(ctx, luaRelease, []string{key}, r.instanceId).Result()
	if err != nil {
		log.Printf("failed to release lock for key %q: %v", key, err)
		return false
	}
	return result.(int64) == 1
}
