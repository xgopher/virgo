package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "app/config" // 加载 .env 配置文件

	"github.com/chanxuehong/rand"
	"github.com/chanxuehong/session"
	"github.com/chanxuehong/sid"
	"github.com/chanxuehong/wechat/oauth2"
	mpoauth2 "github.com/chanxuehong/wechat/open/oauth2"
	"github.com/gin-gonic/gin"
)

var appID string
var appSecret string
var oauth2Endpoint oauth2.Endpoint
var oauth2RedirectURI string
var sessionStorage = session.New(20*60, 60*60)

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
	sid := sid.New()
	state := string(rand.NewHex())

	if err := sessionStorage.Add(sid, state); err != nil {
		// io.WriteString(w, err.Error())
		c.JSON(404, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	c.SetCookie("sid", sid, 3600, "/", "", false, true)

	AuthCodeURL := mpoauth2.AuthCodeURL(appID, oauth2RedirectURI, "snsapi_login", state)

	// 302 临时跳转 - 重定向到微信扫码登录页面
	c.Redirect(http.StatusFound, AuthCodeURL)
	c.Abort()
}

// GetUserInfo 第二步：通过code获取access_token
func (i *OauthWechatController) GetUserInfo(c *gin.Context) {

	cookie, err := c.Cookie("sid")

	if err != nil {

		c.JSON(404, gin.H{"error": "111"})
		log.Println(err)
		return
	}

	session, err := sessionStorage.Get(cookie)
	if err != nil {
		c.JSON(404, gin.H{"error": "xxx"})
		log.Println(err)
		return
	}

	savedState := session.(string) // 一般是要序列化的, 这里保存在内存所以可以这么做

	code := c.Query("code")
	if code == "" {
		log.Println("用户禁止授权")
		return
	}

	queryState := c.Query("state")
	if queryState == "" {
		log.Println("state 参数为空")
		return
	}

	if savedState != queryState {
		str := fmt.Sprintf("state 不匹配, session 中的为 %q, url 传递过来的是 %q", savedState, queryState)
		c.JSON(404, str)
		log.Println(str)
		return
	}

	oauth2Client := oauth2.Client{
		Endpoint: oauth2Endpoint,
	}

	// 通过 code 换取网页授权 access_token
	token, err := oauth2Client.ExchangeToken(code)

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	log.Printf("token: %+v\r\n", token)
	// access_token 调取微信用户信息
	// 返回unionid、openid 等信息
	userinfo, err := mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)

	// 根据 unionid 获取本地 user 信息 (如：关联 users 表)
	// ... todo...

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   userinfo,
	})
}
