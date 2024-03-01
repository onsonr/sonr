package shared

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JwtControllerClaims is a custom claims type for the JWT middleware
type JwtControllerClaims struct {
	Handle  string `json:"handle"`
	Address string `json:"address"`
	Origin  string `json:"origin"`
	jwt.RegisteredClaims
}

func NewJWTControllerToken(handle, address, origin string) *jwt.Token {
	claims := &JwtControllerClaims{
		Handle:  handle,
		Address: address,
		Origin:  origin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token
}
