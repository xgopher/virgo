package services

import (
	"app/services/jwt"
)

// UserPayload 载荷
type UserPayload struct {
	ID string `json:"id"`
	jwt.Payload
}
