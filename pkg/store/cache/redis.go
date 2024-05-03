package cache

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(dsn string) (*redis.Client, error) {
	opts, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	return client, nil
}
