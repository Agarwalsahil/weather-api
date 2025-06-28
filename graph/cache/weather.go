package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"weather-api/graph/model"

	"github.com/redis/go-redis/v9"
)

func GetWeatherFromCache(ctx context.Context, rdb *redis.Client, key string) (*model.WeatherResponse, error) {
	val, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		log.Println("Key not found")
		return nil, nil
	} else if err != nil {
		log.Printf("error getting data from cache: %v", err)
		return nil, err
	}
	var weather model.WeatherResponse
	err = json.Unmarshal([]byte(val), &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}

func SetWeatherInCache(ctx context.Context, rdb *redis.Client, key string, data *model.WeatherResponse) error {
	bytes, err := json.Marshal(data)

	if err != nil {
		log.Printf("error in marshalling the weather data: %v", err)
		return err
	}

	return rdb.Set(ctx, key, bytes, 3*time.Hour).Err()
}
