package redis

import "github.com/redis/go-redis/v9"

type Cache struct {
	RedisClient *redis.Client
}

func NewRedisClient(dsn string) (*redis.Client, error) {
	//opt, err := redis.ParseURL(dsn)
	//if err != nil {
	//	return nil, err
	//}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhostttttt",
		Password: "",
		DB:       0,
	})
	return rdb, nil
}
