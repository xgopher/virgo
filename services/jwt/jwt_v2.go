package jwt

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWT 密钥
var jwtSecret []byte

func init() {
	jwtSecret = []byte("123456")
}

// Payload 载体
type Payload struct {
	jwt.StandardClaims
}

// Encode 用 HS256 创建 token 签名
func Encode(claims jwt.Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

// Decode 解析验证 token
func Decode(tokenString string) (jwt.Claims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	}

	// 分析错误原因
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

	// 解析失败，返回空载体 + 错误
	return nil, err

}
