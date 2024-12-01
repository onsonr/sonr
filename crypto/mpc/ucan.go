package mpc

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ucan-wg/go-ucan"
)

type (
	Token  = ucan.Token
	Claims = ucan.Claims
	Proof  = ucan.Proof

	Attenuations = ucan.Attenuations
	Fact         = ucan.Fact
)

var (
	UCANVersion    = ucan.UCANVersion
	UCANVersionKey = ucan.UCANVersionKey
	PrfKey         = ucan.PrfKey
	FctKey         = ucan.FctKey
	AttKey         = ucan.AttKey
	CapKey         = ucan.CapKey
)

type ucanKeyshare struct {
	userShare *UserKeyshare
	valShare  *ValKeyshare

	addr      string
	issuerDID string
}

func (k ucanKeyshare) NewOriginToken(audienceDID string, att Attenuations, fct []Fact, notBefore, expires time.Time) (*ucan.Token, error) {
	return k.newToken(audienceDID, nil, att, fct, notBefore, expires)
}

func (k ucanKeyshare) NewAttenuatedToken(parent *Token, audienceDID string, att ucan.Attenuations, fct []ucan.Fact, nbf, exp time.Time) (*Token, error) {
	if !parent.Attenuations.Contains(att) {
		return nil, fmt.Errorf("scope of ucan attenuations must be less than it's parent")
	}
	return k.newToken(audienceDID, append(parent.Proofs, Proof(parent.Raw)), att, fct, nbf, exp)
}

func (k ucanKeyshare) newToken(audienceDID string, prf []Proof, att Attenuations, fct []Fact, nbf, exp time.Time) (*ucan.Token, error) {
	t := jwt.New(NewJWTSigningMethod("MPC256", k))

	// if _, err := did.Parse(audienceDID); err != nil {
	// 	return nil, fmt.Errorf("invalid audience DID: %w", err)
	// }

	t.Header[UCANVersionKey] = UCANVersion

	var (
		nbfUnix int64
		expUnix int64
	)

	if !nbf.IsZero() {
		nbfUnix = nbf.Unix()
	}
	if !exp.IsZero() {
		expUnix = exp.Unix()
	}

	// set our claims
	t.Claims = &Claims{
		StandardClaims: &jwt.StandardClaims{
			Issuer:    k.issuerDID,
			Audience:  audienceDID,
			NotBefore: nbfUnix,
			// set the expire time
			// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
			ExpiresAt: expUnix,
		},
		Attenuations: att,
		Facts:        fct,
		Proofs:       prf,
	}

	raw, err := t.SignedString(nil)
	if err != nil {
		return nil, err
	}

	return &Token{
		Raw:          raw,
		Attenuations: att,
		Facts:        fct,
		Proofs:       prf,
	}, nil
}
