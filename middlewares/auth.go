package middlewares

import (
	"strings"
	"net/http"
	"errors"
	
	"app/services"
	"github.com/gin-gonic/gin"
)

var TokenExpired error = errors.New("Token is expired")

// Jwt token auth
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.DefaultQuery("token", "")
		if tokenStr == "" {
			tokenStr = c.Request.Header.Get("Authorization")
			if s := strings.Split(tokenStr, " "); len(s) == 2 {
				tokenStr = s[1]
			}
		}

		j := services.NewJwt()
        // parseToken 解析token包含的信息
        claims, err := j.ParseToken(tokenStr)
        if err != nil {
            if err == TokenExpired {
                c.JSON(http.StatusOK, gin.H{
                    "status": -1,
                    "msg":    "授权已过期",
                })
                c.Abort()
                return
            }
            c.JSON(http.StatusOK, gin.H{
                "status": -1,
                "msg":    err.Error(),
            })
            c.Abort()
            return
        }
        // 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}