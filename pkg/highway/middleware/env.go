package middleware

import (
	"crypto/rand"
	"os"

	"github.com/go-webauthn/webauthn/protocol"
)

const (
	// Standard ports for the sonr grpc and rpc api endpoints.
	SonrGrpcPort = "0.0.0.0:9090"
	SonrRpcPort  = "0.0.0.0:26657"

	// CurrentChainID is the current chain ID.
	CurrentChainID = "sonrdevnet-1"

	// ChallengeLength - Length of bytes to generate for a challenge
	ChallengeLength = 32
)

func GrpcEndpoint() string {
	if env := os.Getenv("ENVIRONMENT"); env != "prod" {
		return SonrGrpcPort
	}
	return SonrGrpcPort
}

func RpcEndpoint() string {
	if env := os.Getenv("ENVIRONMENT"); env != "prod" {
		return SonrRpcPort
	}
	return SonrRpcPort
}

func GetHomeDir() string {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		homeDir = os.Getenv("USERPROFILE") // windows
	}
	return homeDir
}

func ValidatorAddress() (string, bool) {
	if address := os.Getenv("SONR_VALIDATOR_ADDRESS"); address != "" {
		return address, true
	}
	return "", false
}

// createChallenge creates a new challenge that should be signed and returned by the authenticator. The spec recommends
// using at least 16 bytes with 100 bits of entropy. We use 32 bytes.
func createChallenge() (challenge protocol.URLEncodedBase64, err error) {
	challenge = make([]byte, ChallengeLength)

	if _, err = rand.Read(challenge); err != nil {
		return nil, err
	}
	return challenge, nil
}
