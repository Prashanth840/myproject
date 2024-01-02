package services

import (
	"myproject/global"
	"strconv"
	"strings"
	"time"
)

func Filehandler(Filename string, Filesize int64) (string, string) {

	size := strconv.Itoa(int(Filesize))
	Filename = strings.ReplaceAll(Filename, " ", "")
	global.RedisSetExp(Filename, size, time.Duration(time.Second*120))
	result := global.RedisGet(Filename)

	return Filename, result
}
