package rc

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type contextKey string

const redisKey = contextKey("redis")

func InjectRedis(ctx context.Context, rdb *redis.Client) context.Context {
	log.Println("🔥 InjectRedis was called")
	return context.WithValue(ctx, redisKey, rdb)
}

func GetRedis(ctx context.Context) *redis.Client {
	val := ctx.Value(redisKey)
	if rdb, ok := val.(*redis.Client); ok {
		log.Println("✅ Retrieved existing Redis client from context")
		return rdb
	}
	return nil
}
