package webauthn

import (
	"bytes"

	"github.com/ucan-wg/go-ucan"
	"lukechampine.com/blake3"
)

// Blake3 is the chosen hash function for Sonr's WebAuthn implementation. This
// is a wrapper around the blake3 hash function. We use this to generate the challenge
// for the WebAuthn session.
//
// Algorithm = blake3(jwkClaims(mpc-share-cid))
//
// 1. Wallet is generated using MPC protocol
// 2. Vault encrypts and stores shares into IPFS
// 3. JWK Created from the Associated UCAN for Wallet CIDs
// 4. Blake3 is used to generate the hash of the JWK claims
//    for the WebAuthn challenge

func HashUCANToken(t *ucan.Token) ([]byte, error) {
	c, err := t.CID()
	if err != nil {
		return nil, err
	}
	bz := bytes.NewBuffer(DeriveKey())
	_, err = bz.Write(c.Bytes())
	if err != nil {
		return nil, err
	}

	hash := blake3.Sum256(bz.Bytes())
	return hash[:], nil
}

func DeriveKey() ([]byte) {
	var subKey []byte
	blake3.DeriveKey(subKey, "sonr.devnet.1", nil)
	return subKey
}
