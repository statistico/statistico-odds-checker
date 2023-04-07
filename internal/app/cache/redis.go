package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Store interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}

type redisStore struct {
	client *redis.Client
}

func (r *redisStore) Set(ctx context.Context, key, value string) error {
	err := r.client.Set(ctx, key, value, 12*time.Hour).Err()

	if err != nil {
		return err
	}

	fmt.Printf("Value %s set in cache %s", key, value)

	return nil
}

func (r *redisStore) Get(ctx context.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	fmt.Printf("Fetchd from Redis cache: %s", value)

	return value, nil
}

func NewRedisStore(c *redis.Client) Store {
	return &redisStore{client: c}
}
