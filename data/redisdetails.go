package data

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type FileRepo interface {
	RedisGet(Filename string) string
	RedisSetExp(key string, value string, expiration time.Duration) bool
}

type fileRepo struct {
	rd *redis.Client
}

func NewFileRepo() *fileRepo {
	return &fileRepo{
		rd: RedisClient,
	}
}

func (d *fileRepo) RedisGet(key string) string {
	val, err := d.rd.Get(key).Result()
	if err != nil {
		return ""
	}
	return val
}

func (d *fileRepo) RedisSetExp(key string, value string, expiration time.Duration) bool {

	_, err := d.rd.Set(key, value, expiration).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err == nil
}
