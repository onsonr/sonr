package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/marstr/guid"

	did "github.com/sonr-io/sonr/x/identity/types"
)

type JWT struct {
	options JWTOptions
}

func DefaultNew() *JWT {
	return &JWT{
		options: options.DefaultTestConfig(),
	}
}

/*
Creates jwt from parsed with passed did as issuer
*/
func (j *JWT) Generate(doc *did.DidDocument) (string, error) {
	if doc == nil {
		return "", errors.New("highway/jwt Document cannot be nil")
	}
	creatorDID := doc.GetID()
	time := time.Now().Unix()
	exp := time + j.options.ttl // expiers in one hour after issue
	// Create a new token object, specifying signing method and the claims
	// Will use current timespant at time of execution for token issue time.
	token := jwt.NewWithClaims(j.options.singingMethod, jwt.StandardClaims{
		IssuedAt:  time,
		ExpiresAt: exp,
		Issuer:    creatorDID,
		Id:        guid.NewGUID().String(),
		Subject:   "",
		NotBefore: time,
	})

	if token.Header["typ"] != "JWT" {
		return "", errors.New("token type is not JWT")
	}
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(j.options.secret)

	return tokenString, err
}

/*
Parses a jwt token if possible and returns as Token type.
*/
func (j *JWT) Parse(token string) (*jwt.Token, error) {
	parsed, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.options.secret, nil
	})

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, jwt.ErrECDSAVerification
		} else {
			return nil, jwt.ErrECDSAVerification
		}
	}

	return parsed, nil
}
