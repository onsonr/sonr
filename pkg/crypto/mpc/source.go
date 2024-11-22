package mpc

import (
	"fmt"
	"time"

	"github.com/ucan-wg/go-ucan"
)

type KeyshareSource interface {
	ucan.Source

	Address() string
	Issuer() string
	DefaultOriginToken() (*Token, error)
	PublicKey() []byte
	TokenParser() *ucan.TokenParser
	SignData(data []byte) ([]byte, error)
	VerifyData(data []byte, sig []byte) (bool, error)
}

func createKeyshareSource(val *ValKeyshare, user *UserKeyshare) (KeyshareSource, error) {
	iss, addr, err := ComputeIssuerDID(val.GetPublicKey())
	if err != nil {
		return nil, err
	}
	return keyshareSource{
		userShare: user,
		valShare:  val,
		addr:      addr,
		issuerDID: iss,
	}, nil
}

// Address returns the address of the keyshare
func (k keyshareSource) Address() string {
	return k.addr
}

// Issuer returns the DID of the issuer of the keyshare
func (k keyshareSource) Issuer() string {
	return k.issuerDID
}

// PublicKey returns the public key of the keyshare
func (k keyshareSource) PublicKey() []byte {
	return k.valShare.PublicKey
}

// DefaultOriginToken returns a default token with the keyshare's issuer as the audience
func (k keyshareSource) DefaultOriginToken() (*Token, error) {
	caps := NewSmartAccountCapabilities()
	att := CreateSmartAccountAttenuations(caps, k.addr)
	zero := time.Time{}
	return k.NewOriginToken(k.issuerDID, att, nil, zero, zero)
}

// TokenParser returns a token parser that can be used to parse tokens
func (k keyshareSource) TokenParser() *ucan.TokenParser {
	caps := NewSmartAccountCapabilities()
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
	return ucan.NewTokenParser(ac, ucan.StringDIDPubKeyResolver{}, store.(ucan.CIDBytesResolver))
}

func (k keyshareSource) SignData(data []byte) ([]byte, error) {

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

func (k keyshareSource) VerifyData(data []byte, sig []byte) (bool, error) {
	return VerifySignature(k.userShare.PublicKey, data, sig)
}
