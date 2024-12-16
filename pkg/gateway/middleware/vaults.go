package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/internal/context"
	"github.com/onsonr/sonr/pkg/gateway/types"
	"lukechampine.com/blake3"
)

func Spawn(c echo.Context) (types.CreatePasskeyParams, error) {
	cc := c.(*GatewayContext)
	block := fmt.Sprintf("%d", CurrentBlock(c))
	handle := GetHandle(c)
	origin := GetOrigin(c)
	challenge := GetSessionChallenge(c)
	sid := GetSessionID(c)
	nonce, err := calcNonce(sid)
	if err != nil {
		return types.DefaultCreatePasskeyParams(), err
	}
	encl, err := mpc.GenEnclave(nonce)
	if err != nil {
		return types.DefaultCreatePasskeyParams(), err
	}
	cc.stagedEnclaves[sid] = encl
	context.WriteCookie(c, context.SonrAddress, encl.Address())
	return types.CreatePasskeyParams{
		Address:       encl.Address(),
		Handle:        handle,
		Name:          origin,
		Challenge:     challenge,
		CreationBlock: block,
	}, nil
}

func Claim() (types.CreatePasskeyParams, error) {
	return types.CreatePasskeyParams{}, nil
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
