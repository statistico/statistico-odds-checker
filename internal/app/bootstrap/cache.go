package bootstrap

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/statistico/statistico-odds-checker/internal/app/cache"
	"strconv"
)

func (c Container) Cache() cache.Store {
	config := c.Config.RedisConfig

	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)

	db, err := strconv.Atoi(config.Database)

	if err != nil {
		db = 0
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       db,
		PoolSize: 50,
	})
	
	return cache.NewRedisStore(client)
}
