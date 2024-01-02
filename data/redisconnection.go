package data

import (
	"fmt"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client
var Redishost = "127.0.0.1"
var Redisport = "6379"

func InitRedisDB() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     Redishost + ":" + Redisport,
		Password: "",
		DB:       0,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
	}

}
