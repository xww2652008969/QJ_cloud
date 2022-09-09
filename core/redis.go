package core

import (
	"gitgub.com/xww2652008969/QJ-cloud/global"
	"github.com/go-redis/redis/v8"
	"os"
)

func InitRedis() {
	global.QjRedis = redis.NewClient(&redis.Options{
		Addr:     "localhost:49155",
		Password: "redispw",
		DB:       0,
	})
	_, err := global.QjRedis.Ping(global.QjRedis.Context()).Result()
	if err != nil {
		global.QjLog.Print(err)
		os.Exit(-1)
	} else {
		global.QjLog.Print("Redis启动成功")
	}
}
