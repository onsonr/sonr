package didkey

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/ucan-wg/go-ucan"
	"github.com/ucan-wg/go-ucan/didkey"
)

// DIDPubKeyResolver turns did:key Decentralized IDentifiers into a public key,
// possibly using a network request
type DIDPubKeyResolver interface {
	ResolveDIDKey(ctx context.Context, did string) (ID, error)
}

// TokenParser parses a raw string into a Token
type TokenParser struct {
	ap   ucan.AttenuationConstructorFunc
	cidr ucan.CIDBytesResolver
	didr DIDPubKeyResolver
}

// NewTokenParser constructs a token parser
func NewTokenParser(ap ucan.AttenuationConstructorFunc, didr DIDPubKeyResolver, cidr ucan.CIDBytesResolver) *TokenParser {
	return &TokenParser{
		ap:   ap,
		cidr: cidr,
		didr: didr,
	}
}

// ParseAndVerify will parse, validate and return a token
func (p *TokenParser) ParseAndVerify(ctx context.Context, raw string) (*ucan.Token, error) {
	return p.parseAndVerify(ctx, raw, nil)
}

func (p *TokenParser) parseAndVerify(ctx context.Context, raw string, _ *ucan.Token) (*ucan.Token, error) {
	tok, err := jwt.Parse(raw, p.matchVerifyKeyFunc(ctx))
	if err != nil {
		return nil, fmt.Errorf("parsing UCAN: %w", err)
	}

	mc, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("parser fail")
	}

	var iss didkey.ID
	// TODO(b5): we're double parsing here b/c the jwt lib we're using doesn't expose
	// an API (that I know of) for storing parsed issuer / audience
	if issStr, ok := mc["iss"].(string); ok {
		iss, err = didkey.Parse(issStr)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf(`"iss" key is not in claims`)
	}

	var aud didkey.ID
	// TODO(b5): we're double parsing here b/c the jwt lib we're using doesn't expose
	// an API (that I know of) for storing parsed issuer / audience
	if audStr, ok := mc["aud"].(string); ok {
		aud, err = didkey.Parse(audStr)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf(`"aud" key is not in claims`)
	}

	var att ucan.Attenuations
	if acci, ok := mc[ucan.AttKey].([]interface{}); ok {
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

	var prf []ucan.Proof
	if prfi, ok := mc[ucan.PrfKey].([]interface{}); ok {
		for i, a := range prfi {
			if pStr, ok := a.(string); ok {
				prf = append(prf, ucan.Proof(pStr))
			} else {
				return nil, fmt.Errorf(`"prf[%d]" is not a string`, i)
			}
		}
	} else if mc[ucan.PrfKey] != nil {
		return nil, fmt.Errorf(`"prf" key is not an array`)
	}

	return &ucan.Token{
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
