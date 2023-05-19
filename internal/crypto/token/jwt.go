package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/marstr/guid"

	did "github.com/sonrhq/core/x/identity/types"
)

// `JWT` is a struct that has a field called `options` of type `JWTOptions`.
// @property {JWTOptions} options - This is a struct that contains the options for the JWT.
type JWT struct {
	options JWTOptions
}

// `DefaultNew()` returns a new instance of `JWT` with the default options
func DefaultNew() *JWT {
	return &JWT{
		options: options.DefaultTestConfig(),
	}
}

// Creating a new token with the claims of the document.
func (j *JWT) Generate(doc *did.Identity) (string, error) {
	if doc == nil {
		return "", errors.New("highway/jwt Document cannot be nil")
	}
	creatorDID := doc.GetId()
	time := time.Now().Unix()
	exp := time + j.options.ttl
	// expires in one hour after issue
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

// Parses a jwt token if possible and returns as Token type.
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
