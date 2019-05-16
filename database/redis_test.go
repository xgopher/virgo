package database

import (
	// _ "app/config" // 加载 .env 配置文件
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestRedis(t *testing.T) {
	InitRedis()

	conn := Redis

	_, err := conn.Do("SET", "name", "hello world")
	if err != nil {
		fmt.Println("redis set error:", err)
	}

	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Got name: %s \n", name)
	}
}
