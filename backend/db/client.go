package db

import "github.com/redis/go-redis/v9"

var REDIS_CLIENT *redis.Client

func GetRedisClient() *redis.Client {
	if REDIS_CLIENT != nil {
		return REDIS_CLIENT
	}

	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // change this
		Password: "",               // no password set
		DB:       0,                // use default DB
		// UnstableResp3: true,
		Protocol: 2,
	})
}
