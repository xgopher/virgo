package controllers

import (
	"app/database"
	"app/modules/oauth/models"

	"github.com/gin-gonic/gin"
)

// OauthWechatController ...
type OauthWechatController struct {
}

func NewOauthWechatController() *OauthWechatController {
	return &OauthWechatController{}
}

// Index ...
func (i *OauthWechatController) Index(c *gin.Context) {
	// Connection to the database
	db := database.DB

	var oauths []models.OauthWechat
	// SELECT * FROM oauths
	db.Find(&oauths)

	// Display JSON result
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "success!",
		"data":   oauths,
	})
}
