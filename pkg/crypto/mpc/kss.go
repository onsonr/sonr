package mpc

import (
	"fmt"
	"time"

	"github.com/ucan-wg/go-ucan"
)

type KeyshareSource interface {
	ucan.Source
}

type keyshareSource struct {
	userShare Share
	valShare  Share
}

func KeyshareSetFromArray(arr []Share) (KeyshareSource, error) {
	if len(arr) != 2 {
		return nil, fmt.Errorf("invalid keyshare array length")
	}
	return keyshareSource{
		userShare: arr[0],
		valShare:  arr[1],
	}, nil
}

func (k keyshareSource) NewOriginToken(audienceDID string, att ucan.Attenuations, fct []ucan.Fact, notBefore, expires time.Time) (*ucan.Token, error) {
	// Create a new token with the user's keyshare
	token := ucan.NewToken()
	token.Issuer = k.userShare.GetPublicKey()
	token.Audience = []byte(audienceDID)
	token.Attenuations = att
	token.Facts = fct
	token.NotBefore = notBefore
	token.Expiration = expires

	// Sign the token using MPC
	signFunc, err := k.userShare.SignFunc(token.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to create sign function: %w", err)
	}

	valSignFunc, err := k.valShare.SignFunc(token.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to create validator sign function: %w", err)
	}

	sig, err := RunSignProtocol(valSignFunc, signFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to run sign protocol: %w", err)
	}

	sigBytes, err := SerializeSignature(sig)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize signature: %w", err)
	}

	token.Signature = sigBytes
	return token, nil
}

func (k keyshareSource) NewAttenuatedToken(parent *ucan.Token, audienceDID string, att ucan.Attenuations, fct []ucan.Fact, notBefore, expires time.Time) (*ucan.Token, error) {
	// Create attenuated token
	token := ucan.NewToken()
	token.Issuer = k.userShare.GetPublicKey()
	token.Audience = []byte(audienceDID)
	token.Attenuations = att
	token.Facts = fct
	token.NotBefore = notBefore
	token.Expiration = expires
	token.Proof = parent

	// Sign the token using MPC
	signFunc, err := k.userShare.SignFunc(token.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to create sign function: %w", err)
	}

	valSignFunc, err := k.valShare.SignFunc(token.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to create validator sign function: %w", err)
	}

	sig, err := RunSignProtocol(valSignFunc, signFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to run sign protocol: %w", err)
	}

	sigBytes, err := SerializeSignature(sig)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize signature: %w", err)
	}

	token.Signature = sigBytes
	return token, nil
}
