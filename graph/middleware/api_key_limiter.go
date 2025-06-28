package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func RateLimiter(rdb *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")
			if apiKey == "" {
				http.Error(w, "Missing API Key", http.StatusUnauthorized)
				return
			}

			key := fmt.Sprintf("ratelimit:%s", strings.ToLower(apiKey))
			ctx := r.Context()
			count, err := rdb.Get(ctx, key).Int()

			if err == redis.Nil {
				_ = rdb.Set(ctx, key, 1, time.Minute).Err()
			} else if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			} else if count > 5 {
				log.Println("Request count exceeded the count limit")
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			} else {
				_ = rdb.Incr(ctx, key).Err()
			}

			next.ServeHTTP(w, r)
		})
	}
}
