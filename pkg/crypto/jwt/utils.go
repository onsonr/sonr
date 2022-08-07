package jwt

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
)

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

func GetClaims(token *jwt.Token) (*jwt.StandardClaims, bool) {
	claims, ok := token.Claims.(*jwt.StandardClaims)

	return claims, ok
}
