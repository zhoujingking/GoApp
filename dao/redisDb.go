package dao

import (
	"github.com/go-redis/redis"
)

var (
	redisClient *redis.Client
)

func GetRedisClient() *redis.Client {
	return redisClient
}

func InitRedisConnection(address string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	redisClient = client
	return nil
}
