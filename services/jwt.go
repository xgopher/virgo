package services

import (
	"errors"
	"time"
	"os"
	"app/modules/user/models"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}


// 一些常量
var (
    TokenExpired     error  = errors.New("Token is expired")
    TokenNotValidYet error  = errors.New("Token not active yet")
    TokenMalformed   error  = errors.New("Token format error")
    TokenInvalid     error  = errors.New("Couldn't handle this token:")
    SignKey          string = os.Getenv("JWT_SIGN")
)


// 载荷
type customClaims struct {
    ID    string `json:"id"`
    jwt.StandardClaims
}

// 新建一个jwt实例
func NewJwt() *JWT {
    return &JWT{
        []byte(getSignKey()),
    }
}

// 获取signKey
func getSignKey() string {
    return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
    SignKey = key
    return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims customClaims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*customClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
        return j.SigningKey, nil
    })
    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorMalformed != 0 {
                return nil, TokenMalformed
            } else if ve.Errors&jwt.ValidationErrorExpired != 0 {
                // Token is expired
                return nil, TokenExpired
            } else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
                return nil, TokenNotValidYet
            } else {
                return nil, TokenInvalid
            }
        }
    }
    if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
    jwt.TimeFunc = func() time.Time {
        return time.Unix(0, 0)
    }
    token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
        return j.SigningKey, nil
    })
    if err != nil {
        return "", err
    }
    if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
        jwt.TimeFunc = time.Now
        claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
        return j.CreateToken(*claims)
    }
    return "", TokenInvalid
}


// 生成令牌
func (j *JWT) GetTokenFromUser(user models.User) (string, error) {

    claims := customClaims{
        string(user.ID),
        jwt.StandardClaims{
            NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
            ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
            Issuer:    os.Getenv("JWT_ISSUER"),        //签名的发行者
        },
    }

    token, err := j.CreateToken(claims)

    if err != nil {
        return "", err
    }
   
    return token, nil
}