package configuration

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/donnyirianto/go-be-fiber/exception"
	"github.com/redis/go-redis/v9"
)

func NewRedis(config Config) (*redis.Client, error) {
	requiredParams := []string{"REDIS_HOST", "REDIS_PORT", "REDIS_PASSWORD", "REDIS_POOL_MAX_SIZE", "REDIS_POOL_MIN_IDLE_SIZE"}
	for _, param := range requiredParams {
		if config.GetString(param) == "" {
			err := errors.New("Missing required configuration parameter: " + param)
			return nil, err
		}
	}

	host := config.GetString("REDIS_HOST")
	port := config.GetString("REDIS_PORT")
	password := config.GetString("REDIS_PASSWORD")
	maxPoolSize := config.GetInt("REDIS_POOL_MAX_SIZE")
	minIdlePoolSize := config.GetInt("REDIS_POOL_MIN_IDLE_SIZE")

	redisStore := redis.NewClient(&redis.Options{
		Addr:         host + ":" + port,
		Password:     password,
		PoolSize:     maxPoolSize,
		MinIdleConns: minIdlePoolSize,
	})
	return redisStore, nil
}

func SetCache[T any](cacheManager *redis.Client, ctx context.Context, prefix string, key string, executeData func(context.Context, string) (T, error)) *T {
	var data []byte
	var object T
	if err := cacheManager.Get(ctx, prefix+"_"+key).Scan(&data); err == nil {
		err := json.Unmarshal(data, &object)
		if err != nil {
			// Log the error
			exception.PanicLogging(err)
		}
		return &object
	}

	value, err := executeData(ctx, key)
	if err != nil {
		// Log the error
		exception.PanicLogging(err)
		return nil
	}

	cacheValue, err := json.Marshal(value)
	if err != nil {
		// Log the error
		exception.PanicLogging(err)
		return nil
	}

	if err := cacheManager.Set(ctx, prefix+"_"+key, cacheValue, -1).Err(); err != nil {
		// Log the error
		exception.PanicLogging(err)
		return nil
	}
	return &value
}
