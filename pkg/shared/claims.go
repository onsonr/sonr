package shared

import "github.com/golang-jwt/jwt/v5"

// JwtControllerClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtControllerClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}
