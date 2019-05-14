package routes

import (
	"app/modules/oauth/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes ...
func RegisterRoutes(router *gin.Engine) {
	oauthController := controllers.NewOauthWechatController()

	router.GET("/api/v1/oauth/authCode", oauthController.AuthCode)
	router.GET("/api/v1/oauth/getUserInfo", oauthController.GetUserInfo)
}
