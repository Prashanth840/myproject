package data

import (
	"fmt"
	"time"
)

func RedisGet(key string) string {
	val, err := RedisClient.Get(key).Result()
	if err != nil {
		return ""
	}
	return val
}

func RedisSetExp(key string, value string, expiration time.Duration) bool {

	_, err := RedisClient.Set(key, value, expiration).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err == nil
}
