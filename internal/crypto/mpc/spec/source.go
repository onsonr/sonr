package spec

import (
	"context"
	"fmt"
	"time"

	"github.com/onsonr/sonr/internal/crypto/keys"
	"github.com/onsonr/sonr/internal/crypto/ucan"
	"lukechampine.com/blake3"
)

type KeyshareSource interface {
	ucan.Source

	Address() string
	Issuer() string
	ChainCode() ([]byte, error)
	OriginToken() (*Token, error)
	SignData(data []byte) ([]byte, error)
	VerifyData(data []byte, sig []byte) (bool, error)
	UCANParser() *ucan.TokenParser
}

//	func NewSource(ks mpc.KeyEnclave) (KeyshareSource, error) {
//		iss, addr, err := getIssuerDID(val.PublicKey())
//		if err != nil {
//			return nil, err
//		}
//
//		return ucanKeyshare{
//			issuerDID: iss,
//			addr:      addr,
//		}, nil
//	}
//
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

// DefaultOriginToken returns a default token with the keyshare's issuer as the audience
func (k ucanKeyshare) OriginToken() (*Token, error) {
	// att := ucan.NewSmartAccount(k.addr)
	zero := time.Time{}
	// return k.NewOriginToken(k.issuerDID, att, nil, zero, zero)
	return k.newToken(k.issuerDID, nil, nil, nil, zero, zero)
}

func (k ucanKeyshare) SignData(data []byte) ([]byte, error) {
	// // Create signing functions
	// signFunc, err := k.userShare.SignFunc(data)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create sign function: %w", err)
	// }
	//
	// valSignFunc, err := k.valShare.SignFunc(data)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create validator sign function: %w", err)
	// }

	// Run the signing protocol
	// return mpc.ExecuteSigning(valSignFunc, signFunc)
	return nil, nil
}

func (k ucanKeyshare) VerifyData(data []byte, sig []byte) (bool, error) {
	return false, nil
	// return k.valShare.PublicKey().Verify(data, sig)
}

// TokenParser returns a token parser that can be used to parse tokens
func (k ucanKeyshare) UCANParser() *ucan.TokenParser {
	caps := ucan.AccountPermissions.GetCapabilities()
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
			}
		}

		return ucan.Attenuation{
			Rsc: rsc,
			Cap: caps.Cap(cap),
		}, nil
	}

	store := ucan.NewMemTokenStore()
	return ucan.NewTokenParser(ac, customDIDPubKeyResolver{}, store.(ucan.CIDBytesResolver))
}

// customDIDPubKeyResolver implements the DIDPubKeyResolver interface without
// any network backing. Works if the key string given contains the public key
// itself
type customDIDPubKeyResolver struct{}

// ResolveDIDKey extracts a public key from  a did:key string
func (customDIDPubKeyResolver) ResolveDIDKey(ctx context.Context, didStr string) (keys.DID, error) {
	return keys.Parse(didStr)
}
