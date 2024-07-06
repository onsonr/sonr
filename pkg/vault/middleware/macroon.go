package middleware

import (
	"time"

	"gopkg.in/macaroon.v2"
)

// TODO: Replace session authentication with macroons.
func GenerateSessionMacaroon(userId string, audience string) (*macaroon.Macaroon, error) {
	m, err := macaroon.New([]byte("secret-key"), []byte(userId), "sonr-oidc", macaroon.LatestVersion)
	if err != nil {
		return nil, err
	}

	err = m.AddFirstPartyCaveat([]byte("exp=" + time.Now().Add(1*time.Hour).Format(time.RFC3339)))
	if err != nil {
		return nil, err
	}

	err = m.AddFirstPartyCaveat([]byte("aud=" + audience))
	if err != nil {
		return nil, err
	}

	return m, nil
}

func ValidateSessionMacaroon(m *macaroon.Macaroon) (bool, error) {
	return false, nil
}
