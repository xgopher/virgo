package oauth

import (
	"app/modules/oauth/database"
	"app/modules/oauth/routes"

	"github.com/gin-gonic/gin"
)

// Register 注册模块服务
func Register(e *gin.Engine) {
	// 自动迁移
	database.Migrate()
	// 加载路由
	routes.RegisterRoutes(e)
}
