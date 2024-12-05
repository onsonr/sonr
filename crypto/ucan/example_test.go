package ucan_test

import (
	"context"
	"fmt"
	"time"

	"github.com/ucan-wg/go-ucan"
)

func Example() {
	source, err := ucan.NewPrivKeySource(keyOne)
	panicIfError(err)

	audienceDID, err := ucan.DIDStringFromPublicKey(keyOne.GetPublic())
	panicIfError(err)

	caps := ucan.NewNestedCapabilities("SUPER_USER", "OVERWRITE", "SOFT_DELETE", "REVISE", "CREATE")
	att := ucan.Attenuations{
		{caps.Cap("SUPER_USER"), ucan.NewStringLengthResource("api", "*")},
		{caps.Cap("SUPER_USER"), ucan.NewStringLengthResource("dataset", "b5:world_bank_population:*")},
	}
	zero := time.Time{}

	// create a root UCAN
	origin, err := source.NewOriginToken(audienceDID, att, nil, zero, zero)
	panicIfError(err)

	id, err := origin.CID()
	panicIfError(err)

	fmt.Printf("cid of root UCAN: %s\n", id.String())

	att = ucan.Attenuations{
		{caps.Cap("SUPER_USER"), ucan.NewStringLengthResource("dataset", "third:resource")},
	}

	if _, err = source.NewAttenuatedToken(origin, audienceDID, att, nil, zero, zero); err != nil {
		fmt.Println(err)
	}

	att = ucan.Attenuations{
		{caps.Cap("OVERWRITE"), ucan.NewStringLengthResource("dataset", "b5:world_bank_population:*")},
	}

	derivedToken, err := source.NewAttenuatedToken(origin, audienceDID, att, nil, zero, zero)
	panicIfError(err)

	id, err = derivedToken.CID()
	panicIfError(err)

	fmt.Printf("cid of derived UCAN: %s\n", id.String())

	p := exampleParser()
	tok, err := p.ParseAndVerify(context.Background(), origin.Raw)
	panicIfError(err)

	fmt.Printf("issuer DID key type: %s\n", tok.Issuer.Type().String())

	// Output:
	// cid of root UCAN: bafkreihl4b2ncrijeutlkppykgspz6wm3q2o4wiej6njl6tj7k2xa3zcue
	// scope of ucan attenuations must be less than it's parent
	// cid of derived UCAN: bafkreifhpoxctmbmvocdevfbmio6cpzltwauesyyjycipnylocoykwghzu
	// issuer DID key type: RSA
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func exampleParser() *ucan.TokenParser {
	caps := ucan.NewNestedCapabilities("SUPER_USER", "OVERWRITE", "SOFT_DELETE", "REVISE", "CREATE")

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
