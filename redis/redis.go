package redis

import (
	"github.com/go-redis/redis"
	"studentScoreManagement/config"
	"time"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// InitRedis 初始化连接
func InitRedis() {
	addr := config.GetRedis().Host + ":" + config.GetRedis().Port
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.GetRedis().Password,
	})
	if _, err := rdb.Ping().Result(); err != nil {
		panic("redis连接失败 " + err.Error())
	}
}

func Set(key, value string, timeout time.Duration) error {
	return rdb.Set(key, value, timeout).Err()
}

func Get(key string) (string, error) {
	return rdb.Get(key).Result()
}
