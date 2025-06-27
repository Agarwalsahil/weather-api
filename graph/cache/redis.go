package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func InitRedis(ctx context.Context) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	fmt.Printf("Redis connection successful: %s", ping)
}
