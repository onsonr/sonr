package jwt

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/host"
)

type JWT struct {
	options JWTOptions
}

func New(ctx context.Context, hn host.SonrHost) JWT {
	options = JWTOptions{}
	options.secret = []byte(hn.Config().Secret)
	options.singingMethod = hn.Config().SigningMethod
	options.ttl = hn.Config().Expiration

	return JWT{
		options: options,
	}
}

func DefaultNew() JWT {
	return JWT{
		options: options.DefaultTestConfig(),
	}
}

/*

 */
func (j *JWT) Generate(doc *did.Document) (string, error) {
	if doc == nil {
		return "", errors.New("highway/jwt Document cannot be nil")
	}

	time := time.Now().Unix()
	exp := time + j.options.ttl // expiers in one hour after issue
	// Create a new token object, specifying signing method and the claims
	// Will use current timespant at time of execution for token issue time.
	token := jwt.NewWithClaims(j.options.singingMethod, jwt.StandardClaims{
		IssuedAt:  time,
		ExpiresAt: exp,
		Issuer:    doc.ID.DID.String(),
	})

	if token.Header["typ"] != "JWT" {
		return "", errors.New("token type is not JWT")
	}
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(j.options.secret)

	return tokenString, err
}

/*

 */
func (j *JWT) Parse(token string) (*jwt.Token, error) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return j.options.secret, nil
	})

	if parsed.Valid {
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, jwt.ErrECDSAVerification
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, jwt.ErrECDSAVerification
		} else {
			return nil, jwt.ErrECDSAVerification
		}
	}

	return parsed, nil
}
