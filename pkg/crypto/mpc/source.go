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
	TokenParser() *ucan.TokenParser
}

func KeyshareSourceFromArray(arr []Share) (KeyshareSource, error) {
	if len(arr) != 2 {
		return nil, fmt.Errorf("invalid keyshare array length")
	}
	iss, addr, err := ComputeIssuerDID(arr[0].GetPublicKey())
	if err != nil {
		return nil, err
	}
	return keyshareSource{
		userShare: arr[0],
		valShare:  arr[1],
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
