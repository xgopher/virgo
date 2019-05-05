package main

import (
	"app/config"
	"app/database"
	"app/middlewares"
	"app/modules/user"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载 .env 配置文件
	config.Loadenv()
	// 初始化DB链接池
	database.InitPool()
	// 服务停止时清理数据库链接
	defer database.DB.Close()

	e := gin.Default()
	// 调用跨域中间件
	e.Use(middlewares.Cors())

	// 注册模块
	user.Register(e)
	// Listen and Server in 0.0.0.0:8080
	port := os.Getenv("HTTP_PORT")
	e.Run(":" + port)
}
