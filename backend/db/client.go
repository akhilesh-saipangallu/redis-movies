package db

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var REDIS_CLIENT *redis.Client

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetRedisClient() *redis.Client {
	if REDIS_CLIENT != nil {
		return REDIS_CLIENT
	}

	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	return redis.NewClient(&redis.Options{
		Addr:     redisAddr, // change this
		Password: "",        // no password set
		DB:       0,         // use default DB
		Protocol: 2,
	})
}
