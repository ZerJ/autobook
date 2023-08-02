package booking

import (
	"github.com/go-redis/redis"
)

// 定义一个全局变量
var Redisclient *redis.Client

func InitRedis() (err error) {
	Redisclient = redis.NewClient(&redis.Options{
		Addr:     "192.168.58.200:6379", // 指定
		Password: "",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	_, err = Redisclient.Ping().Result()
	return
}
