package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chanxuehong/wechat/oauth2"
	mpoauth2 "github.com/chanxuehong/wechat/open/oauth2"
	"github.com/gin-gonic/gin"
)

var appID string
var appSecret string
var oauth2Endpoint oauth2.Endpoint
var oauth2RedirectURI string

const (
	oauth2Scope = "snsapi_userinfo" // 填上自己的参数
)

func init() {
	appID = os.Getenv("WECHAT_APP_ID")         // 填上自己的参数
	appSecret = os.Getenv("WECHAT_APP_SECRET") // 填上自己的参数
	oauth2RedirectURI = os.Getenv("WECHAT_OAUTH2_REDIRECT_URI")

	oauth2Endpoint = mpoauth2.NewEndpoint(appID, appSecret)
}

// OauthWechatController ...
type OauthWechatController struct {
}

// NewOauthWechatController ...
func NewOauthWechatController() *OauthWechatController {
	return &OauthWechatController{}
}

// AuthCode 获取扫码二维码
func (i *OauthWechatController) AuthCode(c *gin.Context) {
	fmt.Println(appID)
	AuthCodeURL := mpoauth2.AuthCodeURL(appID, oauth2RedirectURI+"?redirect_uri=http://localhost:8180/api/v1/oauth/xxx", "snsapi_login", "abcdef")
	fmt.Println(AuthCodeURL)

	c.Redirect(http.StatusMovedPermanently, AuthCodeURL)
	c.Abort()
}

// GetUserInfo 第二步：通过code获取access_token
func (i *OauthWechatController) GetUserInfo(c *gin.Context) {

	code := c.Query("code")

	// AuthCodeURL := mpoauth2.AuthCodeURL(appID, oauth2RedirectURI, "snsapi_login", "abcdef")
	// fmt.Println(AuthCodeURL)
	fmt.Println(code)
	// c.Redirect(http.StatusMovedPermanently, AuthCodeURL)
	// c.Abort()

	oauth2Client := oauth2.Client{
		Endpoint: oauth2Endpoint,
	}

	token, _ := oauth2Client.ExchangeToken(code)

	userinfo, _ := mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)

	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   userinfo,
	})
}
