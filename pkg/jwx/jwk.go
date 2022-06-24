package jwx

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

/*
	JWK returns the key described by the VerificationMethod as JSON Web Key.
*/
func (x *jwxImpl) CreateEncJWK() (jwk.Key, error) {
	if x.key == nil {
		return nil, errors.New("error while creating jwk: public key not provided")
	}

	jwk, err := jwk.FromRaw(x.key)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}

	setUse(jwk, "enc")
	setKeyOps(jwk, "encrypt")

	x.jwk = jwk

	return jwk, nil
}

/*
	JWK returns the key described by the VerificationMethod as JSON Web Key.
*/
func (x *jwxImpl) CreateSignJWK() (jwk.Key, error) {
	if x.key == nil {
		return nil, errors.New("error while creating jwk: public key not provided")
	}

	jwk, err := jwk.FromRaw(x.key)
	if err != nil {
		return nil, err
	}

	err = setUse(jwk, "sig")
	if err != nil {
		return nil, err
	}
	err = setKeyOps(jwk, "sign")

	if err != nil {
		return nil, err
	}

	x.jwk = jwk

	return jwk, nil
}

func (x *jwxImpl) MarshallJSON() ([]byte, error) {
	keyAsJSON, err := json.Marshal(x.jwk)
	if err != nil {
		return nil, err
	}

	return keyAsJSON, nil
}

func (x *jwxImpl) UnmarshallJSON(key []byte) (*map[string]interface{}, error) {
	keyAsMap := map[string]interface{}{}
	json.Unmarshal(key, &keyAsMap)

	return &keyAsMap, nil
}

func setKid(key jwk.Key, value string) error {
	return key.Set("kid", value)
}

func setUse(key jwk.Key, value string) error {
	return key.Set("use", value)
}

func setKeyOps(key jwk.Key, value string) error {
	return key.Set("key_ops", value)
}
