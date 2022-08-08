package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"webapp.io/settings"
)

// 声明一个全局 rdb 变量
var rdb *redis.Client

// Init It creates a new Redis client and returns an error if it can't connect to the Redis server
func Init(conf *settings.RedisConf) (err error) {
	redisAddr := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	// 创建一个 链接 client 对象
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: conf.Password,
		DB:       conf.DB,
		PoolSize: conf.PoolSize, // 连接池大小
	})
	_, err = rdb.Ping().Result()
	return
}

// Close `Close()` closes the connection to the Redis server
func Close() {
	_ = rdb.Close()
}
