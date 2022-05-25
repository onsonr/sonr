package jwt

import (
	"errors"
)

type JWTParserMiddleware = func() error

func (j *JWT) BuildJWTParseMiddleware(headers string) JWTParserMiddleware {
	return func() error {
		token, err := j.Parse(headers)
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
