package rc

import (
	"context"
	"log"
	"net/http"

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

func WithRedis(rdb *redis.Client) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx := InjectRedis(r.Context(), rdb)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
