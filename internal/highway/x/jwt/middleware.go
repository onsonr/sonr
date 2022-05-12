package jwt

import (
	"errors"
)

type JWTParserMiddleware = func() error

func (j *JWT) BuildJWTParseMiddleware(headers string) JWTParserMiddleware {
	return func() error {
		tokenString, err := ExtractBearerToken(headers)
		if err != nil {
			return errors.New("Unable to parse headers")
		}

		token, err := j.Parse(tokenString)
		if err != nil {
			return err
		}

		_, OK := GetClaims(token)
		if !OK {
			return errors.New("Unable to parse jwt claims")
		}

		return nil
	}
}
