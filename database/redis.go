package database

import (
	_ "app/config" // 加载 .env 配置文件
	"os"

	"github.com/gomodule/redigo/redis"
)

// Redis ...
var Redis redis.Conn

// InitRedis ...
func InitRedis() {
	network := os.Getenv("REDIS_NETWORK")
	address := os.Getenv("REDIS_ADDRESS")

	Redis, err = redis.Dial(network, address)
	if err != nil {
		panic(err)
	}
}
