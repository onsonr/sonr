package jwx

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lestrrat-go/jwx/jwk"
)

/*
	JWK returns the key described by the VerificationMethod as JSON Web Key.
*/
func CreateJWKForEnc(key interface{}) (jwk.Key, error) {
	if key == nil {
		return nil, errors.New("error while creating jwk: public key not provided")
	}

	jwk, err := jwk.New(key)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}

	SetUse(jwk, "enc")
	SetKeyOps(jwk, "encrypt")
	return jwk, nil
}

/*
	JWK returns the key described by the VerificationMethod as JSON Web Key.
*/
func CreateJWKForSig(key interface{}) (jwk.Key, error) {
	if key == nil {
		return nil, errors.New("error while creating jwk: public key not provided")
	}

	jwk, err := jwk.New(key)
	if err != nil {
		return nil, err
	}

	err = SetUse(jwk, "sig")
	if err != nil {
		return nil, err
	}
	err = SetKeyOps(jwk, "sign")

	if err != nil {
		return nil, err
	}

	return jwk, nil
}

func Marshall(jwk *jwk.Key) ([]byte, error) {
	keyAsJSON, err := json.Marshal(jwk)
	if err != nil {
		return nil, err
	}

	return keyAsJSON, nil
}

func Unmarshall(key []byte) (*map[string]interface{}, error) {
	keyAsMap := map[string]interface{}{}
	json.Unmarshal(key, &keyAsMap)

	return &keyAsMap, nil
}

func SetKid(key jwk.Key, value string) error {
	return key.Set("kid", value)
}

func SetUse(key jwk.Key, value string) error {
	return key.Set("use", value)
}

func SetKeyOps(key jwk.Key, value string) error {
	return key.Set("key_ops", value)
}
