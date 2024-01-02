package global

import (
	"fmt"
	"myproject/data"
	"time"
)

func RedisGet(key string) string {
	val, err := data.RedisClient.Get(key).Result()
	if err != nil {
		return ""
	}
	return val
}

func RedisSetExp(key string, value string, expiration time.Duration) bool {

	_, err := data.RedisClient.Set(key, value, expiration).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err == nil
}
