package main

import (
	_ "app/config" // 加载 .env 配置文件
	"app/database"
	"app/middlewares"
	"app/modules/oauth"
	"app/modules/user"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	dubugMode, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))

	if dubugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化DB链接池
	database.InitDb()
	// 服务停止时清理数据库链接
	defer database.DB.Close()

	database.InitRedis()
	// 服务停止时清理 redis 链接
	defer database.Redis.Close()

	e := gin.Default()
	// 调用跨域中间件
	e.Use(middlewares.Cors())

	// 注册模块
	user.Register(e)
	oauth.Register(e)
	// Listen and Server in 0.0.0.0:8080
	port := os.Getenv("HTTP_PORT")
	e.Run(":" + port)
}
