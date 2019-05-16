package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chanxuehong/wechat/oauth2"
	mpoauth2 "github.com/chanxuehong/wechat/open/oauth2"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"

	_ "app/config" // 加载 .env 配置文件
	"app/database"
	"app/modules/oauth/models"
	"app/modules/oauth/services"
)

var appID string
var appSecret string
var oauth2Endpoint oauth2.Endpoint
var oauth2RedirectURI string

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

// Login 扫码登陆, 发起微信授权登录请求
// 第一步：请求CODE
func (i *OauthWechatController) Login(c *gin.Context) {

	state := services.RandString(32)

	conn := database.Redis

	var err error

	_, err = conn.Do("SET", state, 1)
	if err != nil {
		fmt.Println("redis set error:", err)
		return
	}

	// 有效期 5 分钟
	_, err = conn.Do("expire", state, 300)

	if err != nil {
		fmt.Println("set expire error: ", err)
		return
	}

	AuthCodeURL := mpoauth2.AuthCodeURL(appID, oauth2RedirectURI, "snsapi_login", state)

	// 302 临时跳转 - 重定向到微信扫码登录页面
	c.Redirect(http.StatusFound, AuthCodeURL)
	c.Abort()
}

// Callback 微信授权回调
// 第二步：通过code交换 access_token, 接着 access_token 获取微信用户信息
func (i *OauthWechatController) Callback(c *gin.Context) {

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

	if !checkState(queryState) {
		str := fmt.Sprintf("state 不匹配, url 传递过来的是 %q", queryState)
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

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	// 保存 or 更新微信用户信息
	db := database.DB
	// 序列化 - 用户特权信息
	jsonPrivilege, err := json.Marshal(userinfo.Privilege)

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	oauthWechat := models.OauthWechat{
		OpenID:     userinfo.OpenId,  // 字段有区别
		UnionID:    userinfo.UnionId, // 字段有区别
		Nickname:   userinfo.Nickname,
		HeadImgURL: userinfo.HeadImageURL, // 字段有区别
		Sex:        userinfo.Sex,
		Province:   userinfo.Province,
		City:       userinfo.City,
		Country:    userinfo.Country,
		Privilege:  string(jsonPrivilege), // 字段有区别
	}

	var count int

	db.Model(&models.OauthWechat{}).Where("unionid = ?", userinfo.UnionId).Count(&count)

	// 下面这一段，又长又臭，可以考虑封装在 repositories 层
	// FirstOrCreate, FirstOrInit 这2个方法效率太低，丢弃!!!
	// Save 方法，需要 id 主键, 放弃!!!
	if count > 0 {
		// 更新
		db.Model(&models.OauthWechat{}).Where("unionid = ?", userinfo.UnionId).UpdateColumns(models.OauthWechat{
			Nickname:   userinfo.Nickname,
			HeadImgURL: userinfo.HeadImageURL, // 字段有区别
			Sex:        userinfo.Sex,
			Province:   userinfo.Province,
			City:       userinfo.City,
			Country:    userinfo.Country,
			Privilege:  string(jsonPrivilege), // 字段有区别
		})
	} else {
		// 创建 -- 多了一次 SELECT xx from 表名 WHERE ID=xx; 有点多余
		db.Create(&oauthWechat)
	}

	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   userinfo,
	})
}

// 验证 state 正确性
func checkState(state string) bool {

	conn := database.Redis
	savedState, err := redis.String(conn.Do("GET", state))

	if err != nil {
		fmt.Println("redis get error:", err)
		return false
	}

	if "1" == savedState {
		return true
	}

	str := fmt.Sprintf("state 不匹配, url 传递过来的是 %q", state)
	log.Println(str)
	return false

}
