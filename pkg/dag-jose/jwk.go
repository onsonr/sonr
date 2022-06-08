package dag_jose

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
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}

	err = SetUse(jwk, "sig")

	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}

	return jwk, nil
}

/*
	Convert to JSON and back to fix encoding of key material to make sure
	an unmarshalled and newly created VerificationMethod are equal on object level.
	The format of PublicKeyJwk in verificationMethod is a map[string]interface{}.
	We can't use the Key.AsMap since the values of the map will all be internal jwk lib structs.
	After unmarshalling all the fields will be map[string]string.
*/
func CreateUnmarshalledKey(jwk *jwk.Key) (*map[string]interface{}, error) {
	keyAsJSON, err := json.Marshal(jwk)
	if err != nil {
		return nil, err
	}
	keyAsMap := map[string]interface{}{}
	json.Unmarshal(keyAsJSON, &keyAsMap)

	return &keyAsMap, nil
}

func SetKid(key jwk.Key, value string) error {
	return key.Set("kid", value)
}

func SetUse(key jwk.Key, value string) error {
	return key.Set("use", value)
}
