package token

import (
	"crypto"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ucan-wg/go-ucan"
)

// ErrInvalidToken indicates an access token is invalid
var ErrInvalidToken = errors.New("invalid access token")

const (
	// UCANVersion is the current version of the UCAN spec
	UCANVersion = "0.7.0"
	// UCANVersionKey is the key used in version headers for the UCAN spec
	UCANVersionKey = "ucv"
	// PrfKey denotes "Proofs" in a UCAN. Stored in JWT Claims
	PrfKey = "prf"
	// FctKey denotes "Facts" in a UCAN. Stored in JWT Claims
	FctKey = "fct"
	// AttKey denotes "Attenuations" in a UCAN. Stored in JWT Claims
	AttKey = "att"
	// CapKey indicates a resource Capability. Used in an attenuation
	CapKey = "cap"
)

// NewUnsignedToken It creates a new token, encodes it, and returns it
func NewUnsignedUCAN(p crypto.PublicKey, audienceDID string, prf []ucan.Proof, att ucan.Attenuations, fct []ucan.Fact, nbf, exp time.Time) (string, error) {
	t := jwt.New(jwt.SigningMethodES256)

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
	t.Claims = &ucan.Claims{
		StandardClaims: &jwt.StandardClaims{
			Issuer:    p.(crypto.Signer).Public().(crypto.PublicKey).(crypto.SignerOpts).HashFunc().String(),
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
	return t.SigningString()
}
