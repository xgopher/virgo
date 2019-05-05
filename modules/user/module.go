package user

import (
	"app/database"
	"app/modules/user/controllers"
	"app/modules/user/models"

	"github.com/gin-gonic/gin"
)

// Register 注册模块服务
func Register(e *gin.Engine) {
	// 自动迁移
	autoMigrate()
	// 加载路由
	loadRoute(e)
}

func loadRoute(router *gin.Engine) {
	userController := controllers.NewUserController()
	v1 := router.Group("api/v1")
	{
		v1.GET("/users", userController.Index)
		v1.POST("/users", userController.Store)
		v1.GET("/users/:id", userController.Show)
		v1.PUT("/users/:id", userController.Update)
		v1.DELETE("/users/:id", userController.Destroy)
	}
}

// autoMigrate 自动迁移
func autoMigrate() {
	// Creating the table
	if !database.DB.HasTable(&models.User{}) {
		database.DB.CreateTable(&models.User{})
		database.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.User{})
	}
}
