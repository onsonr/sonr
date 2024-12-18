package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/internal/context"
	"lukechampine.com/blake3"
)

func Spawn(c echo.Context) (CreatePasskeyParams, error) {
	cc := c.(*GatewayContext)
	block := fmt.Sprintf("%d", CurrentBlock(c))
	handle := GetHandle(c)
	origin := GetOrigin(c)
	challenge := GetSessionChallenge(c)
	sid := GetSessionID(c)
	nonce, err := calcNonce(sid)
	if err != nil {
		return defaultCreatePasskeyParams(), err
	}
	encl, err := mpc.GenEnclave(nonce)
	if err != nil {
		return defaultCreatePasskeyParams(), err
	}
	cc.stagedEnclaves[sid] = encl
	context.WriteCookie(c, context.SonrAddress, encl.Address())
	return CreatePasskeyParams{
		Address:       encl.Address(),
		Handle:        handle,
		Name:          origin,
		Challenge:     challenge,
		CreationBlock: block,
	}, nil
}

func Claim() (CreatePasskeyParams, error) {
	return CreatePasskeyParams{}, nil
}

// Uses blake3 to hash the sessionID to generate a nonce of length 12 bytes
func calcNonce(sessionID string) ([]byte, error) {
	hash := blake3.New(32, nil)
	_, err := hash.Write([]byte(sessionID))
	if err != nil {
		return nil, err
	}
	// Read the hash into a byte slice
	nonce := make([]byte, 12)
	_, err = hash.Write(nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

// ╭───────────────────────────────────────────────────────────╮
// │            Create Passkey (/register/passkey)             │
// ╰───────────────────────────────────────────────────────────╯

// defaultCreatePasskeyParams returns a default CreatePasskeyParams
func defaultCreatePasskeyParams() CreatePasskeyParams {
	return CreatePasskeyParams{
		Address:       "",
		Handle:        "",
		Name:          "",
		Challenge:     "",
		CreationBlock: "",
	}
}

// CreatePasskeyParams represents the parameters for creating a passkey
type CreatePasskeyParams struct {
	Address       string
	Handle        string
	Name          string
	Challenge     string
	CreationBlock string
}
