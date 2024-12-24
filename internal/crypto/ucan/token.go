// Package ucan implements User-Controlled Authorization Network tokens by
// fission:
// https://whitepaper.fission.codes/access-control/ucan/ucan-tokens
//
// From the paper:
// The UCAN format is designed as an authenticated digraph in some larger
// authorization space. The other way to view this is as a function from a set
// of authorizations (“UCAN proofs“) to a subset output (“UCAN capabilities”).
package ucan

import (
	"context"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/crypto"
	mh "github.com/multiformats/go-multihash"
	"github.com/onsonr/sonr/internal/crypto/keys"
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

// Token is a JSON Web Token (JWT) that contains special keys that make the
// token a UCAN
type Token struct {
	// Entire UCAN as a signed JWT string
	Raw      string
	Issuer   keys.DID
	Audience keys.DID
	// the "inputs" to this token, a chain UCAN tokens with broader scopes &
	// deadlines than this token
	Proofs []Proof `json:"prf,omitempty"`
	// the "outputs" of this token, an array of heterogenous resources &
	// capabilities
	Attenuations Attenuations `json:"att,omitempty"`
	// Facts are facts, jack.
	Facts []Fact `json:"fct,omitempty"`
}

// CID calculates the cid of a UCAN using the default prefix
func (t *Token) CID() (cid.Cid, error) {
	pref := cid.Prefix{
		Version:  1,
		Codec:    cid.Raw,
		MhType:   mh.SHA2_256,
		MhLength: -1, // default length
	}

	return t.PrefixCID(pref)
}

// PrefixCID calculates the CID of a token with a supplied prefix
func (t *Token) PrefixCID(pref cid.Prefix) (cid.Cid, error) {
	return pref.Sum([]byte(t.Raw))
}

// Claims is the claims component of a UCAN token. UCAN claims are expressed
// as a standard JWT claims object with additional special fields
type Claims struct {
	*jwt.StandardClaims
	// the "inputs" to this token, a chain UCAN tokens with broader scopes &
	// deadlines than this token
	// Proofs are UCAN chains, leading back to a self-evident origin token
	Proofs []Proof `json:"prf,omitempty"`
	// the "outputs" of this token, an array of heterogenous resources &
	// capabilities
	Attenuations Attenuations `json:"att,omitempty"`
	// Facts are facts, jack.
	Facts []Fact `json:"fct,omitempty"`
}

// Fact is self-evident statement
type Fact struct {
	cidString string
	value     map[string]interface{}
}

// func (fct *Fact) MarshalJSON() (p[])

// func (fct *Fact) UnmarshalJSON(p []byte) error {
// 	var str string
// 	if json.Unmarshal(p, &str); err == nil {
// 	}
// }

// CIDBytesResolver is a small interface for turning a CID into the bytes
// they reference. In practice this may be backed by a network connection that
// can fetch CIDs, eg: IPFS.
type CIDBytesResolver interface {
	ResolveCIDBytes(ctx context.Context, id cid.Cid) ([]byte, error)
}

// Source creates tokens, and provides a verification key for all tokens it
// creates
//
// implementations of Source must conform to the assertion test defined in the
// spec subpackage
type Source interface {
	NewOriginToken(audienceDID string, att Attenuations, fct []Fact, notBefore, expires time.Time) (*Token, error)
	NewAttenuatedToken(parent *Token, audienceDID string, att Attenuations, fct []Fact, notBefore, expires time.Time) (*Token, error)
}

type pkSource struct {
	pk            crypto.PrivKey
	issuerDID     string
	signingMethod jwt.SigningMethod

	verifyKey interface{} // one of: *rsa.PublicKey, *edsa.PublicKey
	signKey   interface{} // one of: *rsa.PrivateKey,
}

// assert pkSource implements tokens at compile time
var _ Source = (*pkSource)(nil)

// NewPrivKeySource creates an authentication interface backed by a single
// private key. Intended for a node running as remote, or providing a public API
func NewPrivKeySource(privKey crypto.PrivKey) (Source, error) {
	rawPrivBytes, err := privKey.Raw()
	if err != nil {
		return nil, fmt.Errorf("getting private key bytes: %w", err)
	}

	var (
		methodStr = ""
		keyType   = privKey.Type()
		signKey   interface{}
		verifyKey interface{}
	)

	switch keyType {
	case crypto.RSA:
		methodStr = "RS256"
		// TODO(b5) - detect if key is encoded as PEM block, here we're assuming it is
		signKey, err = x509.ParsePKCS1PrivateKey(rawPrivBytes)
		if err != nil {
			return nil, err
		}
		rawPubBytes, err := privKey.GetPublic().Raw()
		if err != nil {
			return nil, fmt.Errorf("getting raw public key bytes: %w", err)
		}
		verifyKeyiface, err := x509.ParsePKIXPublicKey(rawPubBytes)
		if err != nil {
			return nil, fmt.Errorf("parsing public key bytes: %w", err)
		}
		var ok bool
		verifyKey, ok = verifyKeyiface.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("public key is not an RSA key. got type: %T", verifyKeyiface)
		}
	case crypto.Ed25519:
		methodStr = "EdDSA"
		signKey = ed25519.PrivateKey(rawPrivBytes)
		rawPubBytes, err := privKey.GetPublic().Raw()
		if err != nil {
			return nil, fmt.Errorf("getting raw public key bytes: %w", err)
		}
		verifyKey = ed25519.PublicKey(rawPubBytes)
	default:
		return nil, fmt.Errorf("unsupported key type for token creation: %q", keyType)
	}

	issuerDID, err := DIDStringFromPublicKey(privKey.GetPublic())
	if err != nil {
		return nil, err
	}

	return &pkSource{
		pk:            privKey,
		signingMethod: jwt.GetSigningMethod(methodStr),
		verifyKey:     verifyKey,
		signKey:       signKey,
		issuerDID:     issuerDID,
	}, nil
}

func (a *pkSource) NewOriginToken(audienceDID string, att Attenuations, fct []Fact, nbf, exp time.Time) (*Token, error) {
	return a.newToken(audienceDID, nil, att, fct, nbf, exp)
}

func (a *pkSource) NewAttenuatedToken(parent *Token, audienceDID string, att Attenuations, fct []Fact, nbf, exp time.Time) (*Token, error) {
	if !parent.Attenuations.Contains(att) {
		return nil, fmt.Errorf("scope of ucan attenuations must be less than it's parent")
	}
	return a.newToken(audienceDID, append(parent.Proofs, Proof(parent.Raw)), att, fct, nbf, exp)
}

// CreateToken returns a new JWT token
func (a *pkSource) newToken(audienceDID string, prf []Proof, att Attenuations, fct []Fact, nbf, exp time.Time) (*Token, error) {
	// create a signer for rsa 256
	t := jwt.New(a.signingMethod)

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
			Issuer:    a.issuerDID,
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

	raw, err := t.SignedString(a.signKey)
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

// DIDPubKeyResolver turns did:key Decentralized IDentifiers into a public key,
// possibly using a network request
type DIDPubKeyResolver interface {
	ResolveDIDKey(ctx context.Context, did string) (keys.DID, error)
}

// DIDStringFromPublicKey creates a did:key identifier string from a public key
func DIDStringFromPublicKey(pub crypto.PubKey) (string, error) {
	id, err := keys.NewDID(pub)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// StringDIDPubKeyResolver implements the DIDPubKeyResolver interface without
// any network backing. Works if the key string given contains the public key
// itself
type StringDIDPubKeyResolver struct{}

// ResolveDIDKey extracts a public key from  a did:key string
func (StringDIDPubKeyResolver) ResolveDIDKey(ctx context.Context, didStr string) (keys.DID, error) {
	return keys.Parse(didStr)
}

// TokenParser parses a raw string into a Token
type TokenParser struct {
	ap   AttenuationConstructorFunc
	cidr CIDBytesResolver
	didr DIDPubKeyResolver
}

// NewTokenParser constructs a token parser
func NewTokenParser(ap AttenuationConstructorFunc, didr DIDPubKeyResolver, cidr CIDBytesResolver) *TokenParser {
	return &TokenParser{
		ap:   ap,
		cidr: cidr,
		didr: didr,
	}
}

// ParseAndVerify will parse, validate and return a token
func (p *TokenParser) ParseAndVerify(ctx context.Context, raw string) (*Token, error) {
	return p.parseAndVerify(ctx, raw, nil)
}

func (p *TokenParser) parseAndVerify(ctx context.Context, raw string, child *Token) (*Token, error) {
	tok, err := jwt.Parse(raw, p.matchVerifyKeyFunc(ctx))
	if err != nil {
		return nil, fmt.Errorf("parsing UCAN: %w", err)
	}

	mc, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("parser fail")
	}

	var iss keys.DID
	// TODO(b5): we're double parsing here b/c the jwt lib we're using doesn't expose
	// an API (that I know of) for storing parsed issuer / audience
	if issStr, ok := mc["iss"].(string); ok {
		iss, err = keys.Parse(issStr)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf(`"iss" key is not in claims`)
	}

	var aud keys.DID
	// TODO(b5): we're double parsing here b/c the jwt lib we're using doesn't expose
	// an API (that I know of) for storing parsed issuer / audience
	if audStr, ok := mc["aud"].(string); ok {
		aud, err = keys.Parse(audStr)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf(`"aud" key is not in claims`)
	}

	var att Attenuations
	if acci, ok := mc[AttKey].([]interface{}); ok {
		for i, a := range acci {
			if mapv, ok := a.(map[string]interface{}); ok {
				a, err := p.ap(mapv)
				if err != nil {
					return nil, err
				}
				att = append(att, a)
			} else {
				return nil, fmt.Errorf(`"att[%d]" is not an object`, i)
			}
		}
	} else {
		return nil, fmt.Errorf(`"att" key is not an array`)
	}

	var prf []Proof
	if prfi, ok := mc[PrfKey].([]interface{}); ok {
		for i, a := range prfi {
			if pStr, ok := a.(string); ok {
				prf = append(prf, Proof(pStr))
			} else {
				return nil, fmt.Errorf(`"prf[%d]" is not a string`, i)
			}
		}
	} else if mc[PrfKey] != nil {
		return nil, fmt.Errorf(`"prf" key is not an array`)
	}

	return &Token{
		Raw:          raw,
		Issuer:       iss,
		Audience:     aud,
		Attenuations: att,
		Proofs:       prf,
	}, nil
}

func (p *TokenParser) matchVerifyKeyFunc(ctx context.Context) func(tok *jwt.Token) (interface{}, error) {
	return func(tok *jwt.Token) (interface{}, error) {
		mc, ok := tok.Claims.(jwt.MapClaims)
		if !ok {
			return nil, fmt.Errorf("parser fail")
		}

		iss, ok := mc["iss"].(string)
		if !ok {
			return nil, fmt.Errorf(`"iss" claims key is required`)
		}

		id, err := p.didr.ResolveDIDKey(ctx, iss)
		if err != nil {
			return nil, err
		}

		return id.VerifyKey()
	}
}
