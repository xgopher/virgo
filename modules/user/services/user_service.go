package services

import (
	"os"
	"time"
	"app/services/jwt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"app/modules/user/models"
)

// UserPayload 载荷
type UserPayload struct {
	ID int `json:"id"`
	jwt.Payload
}

// 从模型生成token
func GetTokenFromUser(user models.User) (string, error) {
	
	claims := UserPayload{
		user.ID,
		jwt.Payload{
			jwtgo.StandardClaims{
				NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
				ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
				Issuer:    os.Getenv("JWT_ISSUER"),         //签名的发行者
			},
		},
	}

	token, err := jwt.Encode(claims)

	if err != nil {
		return "", err
	}

	return token, nil
}
