package context

import (
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/pkg/common"
	"lukechampine.com/blake3"
)

func (cc *GatewayContext) Spawn(handle, origin string) (*CreatePasskeyParams, error) {
	challenge := GetAuthChallenge(cc)
	sid := GetSessionID(cc)
	nonce, err := calcNonce(sid)
	if err != nil {
		return nil, err
	}
	encl, err := mpc.GenEnclave(nonce)
	if err != nil {
		return nil, err
	}
	cc.stagedEnclaves[sid] = encl
	common.WriteCookie(cc, common.SonrAddress, encl.Address())
	return &CreatePasskeyParams{
		Address:       encl.Address(),
		Handle:        handle,
		Name:          origin,
		Challenge:     challenge,
		CreationBlock: cc.StatusBlock(),
	}, nil
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
