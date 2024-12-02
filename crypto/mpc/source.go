package mpc

import (
	"context"
	"fmt"
	"time"

	"github.com/onsonr/sonr/crypto/didkey"
	"github.com/onsonr/sonr/x/dwn/types/attns"
	"github.com/ucan-wg/go-ucan"
	"lukechampine.com/blake3"
)

type KeyshareSource interface {
	ucan.Source

	Address() string
	Issuer() string
	ChainCode() ([]byte, error)
	OriginToken() (*Token, error)
	PublicKey() PublicKey
	SignData(data []byte) ([]byte, error)
	VerifyData(data []byte, sig []byte) (bool, error)
	UCANParser() *didkey.TokenParser
}

func NewSource(ks Keyset) (KeyshareSource, error) {
	val := ks.Val()
	user := ks.User()
	iss, addr, err := ComputeIssuerDID(val.GetPublicKey())
	if err != nil {
		return nil, err
	}
	return ucanKeyshare{
		userShare: user,
		valShare:  val,
		addr:      addr,
		issuerDID: iss,
	}, nil
}

// Address returns the address of the keyshare
func (k ucanKeyshare) Address() string {
	return k.addr
}

// Issuer returns the DID of the issuer of the keyshare
func (k ucanKeyshare) Issuer() string {
	return k.issuerDID
}

// ChainCode returns the chain code of the keyshare
func (k ucanKeyshare) ChainCode() ([]byte, error) {
	sig, err := k.SignData([]byte(k.addr))
	if err != nil {
		return nil, err
	}
	hash := blake3.Sum256(sig)
	// Return the first 32 bytes of the hash
	return hash[:32], nil
}

// PublicKey returns the public key of the keyshare
func (k ucanKeyshare) PublicKey() PublicKey {
	return createPublicKey(k.valShare.PublicKey, "secp256k1")
}

// DefaultOriginToken returns a default token with the keyshare's issuer as the audience
func (k ucanKeyshare) OriginToken() (*Token, error) {
	att := attns.CreateSmartAccountAttenuations(k.addr)
	zero := time.Time{}
	return k.NewOriginToken(k.issuerDID, att, nil, zero, zero)
}

func (k ucanKeyshare) SignData(data []byte) ([]byte, error) {
	// Create signing functions
	signFunc, err := k.userShare.SignFunc(data)
	if err != nil {
		return nil, fmt.Errorf("failed to create sign function: %w", err)
	}

	valSignFunc, err := k.valShare.SignFunc(data)
	if err != nil {
		return nil, fmt.Errorf("failed to create validator sign function: %w", err)
	}

	// Run the signing protocol
	sig, err := RunSignProtocol(valSignFunc, signFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to run sign protocol: %w", err)
	}
	return SerializeSignature(sig)
}

func (k ucanKeyshare) VerifyData(data []byte, sig []byte) (bool, error) {
	return VerifySignature(k.userShare.PublicKey, data, sig)
}

// TokenParser returns a token parser that can be used to parse tokens
func (k ucanKeyshare) UCANParser() *didkey.TokenParser {
	caps := attns.AttentuationSmartAccount.GetCapabilities()
	ac := func(m map[string]interface{}) (ucan.Attenuation, error) {
		var (
			cap string
			rsc ucan.Resource
		)
		for key, vali := range m {
			val, ok := vali.(string)
			if !ok {
				return ucan.Attenuation{}, fmt.Errorf(`expected attenuation value to be a string`)
			}

			if key == ucan.CapKey {
				cap = val
			} else {
				rsc = ucan.NewStringLengthResource(key, val)
			}
		}

		return ucan.Attenuation{
			Rsc: rsc,
			Cap: caps.Cap(cap),
		}, nil
	}

	store := ucan.NewMemTokenStore()
	return didkey.NewTokenParser(ac, customDIDPubKeyResolver{}, store.(ucan.CIDBytesResolver))
}

// customDIDPubKeyResolver implements the DIDPubKeyResolver interface without
// any network backing. Works if the key string given contains the public key
// itself
type customDIDPubKeyResolver struct{}

// ResolveDIDKey extracts a public key from  a did:key string
func (customDIDPubKeyResolver) ResolveDIDKey(ctx context.Context, didStr string) (didkey.ID, error) {
	return didkey.Parse(didStr)
}
