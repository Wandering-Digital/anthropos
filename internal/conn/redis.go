package conn

import (
	"context"
	"log"

	"github.com/Wandering-Digital/anthropos/internal/config"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func ConnectRedis() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis().Address,
		Password: config.Redis().Password,
		DB:       config.Redis().DB,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Println(err)
		return err
	}

	redisClient = rdb
	log.Println("Redis connected successfully!")
	return nil
}

func Redis() *redis.Client {
	return redisClient
}
