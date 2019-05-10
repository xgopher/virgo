package jwt

import (
	"fmt"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestEncode(t *testing.T) {

	payload := Payload{
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "com.github.virgo",
		},
	}

	token, err := Encode(payload)

	fmt.Printf("%v %v \n", token, err)
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDAwLCJpc3MiOiJjb20uZ2l0aHViLnZpcmdvIn0.LcbozwqmF6YpqrZ4vxM7L02xDr9jmytZanlRkmxJpLg <nil>
}

func TestDecode(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDAwLCJpc3MiOiJjb20uZ2l0aHViLnZpcmdvIn0.LcbozwqmF6YpqrZ4vxM7L02xDr9jmytZanlRkmxJpLg"

	// sample token is expired.  override time so it parses as valid
	at(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if claims, ok := token.Claims.(*Payload); ok && token.Valid {
			fmt.Printf("%v %v \n", claims.StandardClaims.ExpiresAt, claims.StandardClaims.Issuer)
		} else {
			fmt.Println(err)
		}
	})
}

// Override time value for tests.  Restore default value after.
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
