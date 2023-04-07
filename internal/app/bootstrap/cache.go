package bootstrap

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/statistico/statistico-odds-checker/internal/app/cache"
	"os"
	"strconv"
	"time"
)

func (c Container) Cache() cache.Store {
	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	db, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))

	if err != nil {
		db = 0
	}

	d, _ := time.ParseDuration("10s")

	client := redis.NewClient(&redis.Options{
		Addr:        addr,
		DB:          db,
		DialTimeout: d,
	})

	return cache.NewRedisStore(client)
}
