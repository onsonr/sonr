package token

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
)

// It takes a string, splits it on a space, and returns the second element of the resulting array
func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

// It takes a JWT token and returns the claims from it
func GetClaims(token *jwt.Token) (*jwt.StandardClaims, bool) {
	claims, ok := token.Claims.(*jwt.StandardClaims)

	return claims, ok
}
