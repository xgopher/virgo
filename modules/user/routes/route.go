package routes

import (
	"app/modules/user/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes ...
func RegisterRoutes(router *gin.Engine) {
	// userController := controllers.UserController{} // 正常使用方法
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
