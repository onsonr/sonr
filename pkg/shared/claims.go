package shared

import "github.com/golang-jwt/jwt/v5"

// JwtControllerClaims is a custom claims type for the JWT middleware
type JwtControllerClaims struct {
	Handle  string `json:"handle"`
	Address string `json:"address"`
	Origin  string `json:"origin"`
	jwt.RegisteredClaims
}

// JwtSessionClaims is a custom claims type for the JWT middleware
type JwtSessionClaims struct {
	Handle string `json:"handle"`
	jwt.RegisteredClaims
}
