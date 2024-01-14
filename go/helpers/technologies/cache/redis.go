package cache

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func NewRedisDB() (*redis.Client, error) {
	// Read environment variables
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	// Create Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
		DB:       0, // Use default database
	})

	// Ping Redis to check the connection
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
