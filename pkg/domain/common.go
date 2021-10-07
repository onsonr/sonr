package domain

import (
	"errors"
	"os"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/host"
)

var (
	logger              = golog.Child("domain")
	ErrMissingAPIKey    = errors.New("Missing Namebase Handshake Key in env")
	ErrMissingAPISecret = errors.New("Missing Namebase Handshake Secret in env")
	ErrParameters       = errors.New("Failed to create new DomainProtocol, invalid parameters")
)

// fetchApiKeys fetches the Textile Api/Secrect keys from the environment
func fetchApiKeys() (string, string, error) {
	// Get API Key
	key, ok := os.LookupEnv("HANDSHAKE_KEY")
	if !ok {
		return "", "", ErrMissingAPIKey
	}

	// Get API Secret
	secret, ok := os.LookupEnv("HANDSHAKE_SECRET")
	if !ok {
		return "", "", ErrMissingAPISecret
	}
	return key, secret, nil
}

func checkParams(host *host.SNRHost) error {
	if host == nil {
		logger.Error("Host provided is nil", ErrParameters)
		return ErrParameters
	}
	return host.IsReady()
}
