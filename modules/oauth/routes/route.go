package routes

import (
	"app/modules/oauth/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes ...
func RegisterRoutes(router *gin.Engine) {
	oauthController := controllers.NewOauthWechatController()

	router.GET("/api/v1/oauth/login", oauthController.Login)
	router.GET("/api/v1/oauth/callback", oauthController.Callback)
}
