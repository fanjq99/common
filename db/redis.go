package db

import (
	"time"

	"github.com/go-redis/redis"
)

func GetRedisClient(addr, password string, database int) (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		PoolSize:     4,
		IdleTimeout:  5 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MinIdleConns: 1,
		DB:           database,
	})

	return c, c.Ping().Err()
}
