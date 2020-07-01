package utils

import (
	"github.com/go-redis/redis"
	"shici/config"
)

//use
//redis_client := utils_tools.Redisclient()
//str := redis_client.HGet(config.Key, Name).Val()

func Redisclient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	return client
}
