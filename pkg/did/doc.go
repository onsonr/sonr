package did

import (
	"encoding/json"
	"fmt"

	"github.com/libp2p/go-libp2p-core/crypto"
	didtypes "github.com/nuts-foundation/go-did"
	"github.com/nuts-foundation/go-did/did"
	"github.com/sonr-io/core/pkg/device"
)

type identityWallet struct {
	didDoc *did.Document
}

func NewIdWallet(pub crypto.PubKey, sname string) (*did.Document, error) {
	// Create DID Document
	didID, err := did.ParseDID(fmt.Sprintf("did:sonr:%s", sname))
	if err != nil {
		return nil, err
	}
	deviceid, err := device.ID()
	if err != nil {
		return nil, err
	}

	didDeviceID, err := did.ParseDID(fmt.Sprintf("did:sonr:%s", deviceid))
	if err != nil {
		return nil, err
	}

	publicKey, err := crypto.PubKeyToStdKey(pub)
	if err != nil {
		return nil, err
	}

	// Empty did document:
	doc := &did.Document{
		Context: []didtypes.URI{did.DIDContextV1URI()},
		ID:      *didID,
		Controller: []did.DID{
			*didDeviceID,
		},
	}

	// Add an assertionMethod
	keyID, err := did.ParseDIDURL(fmt.Sprintf("did:sonr:%s#master", sname))
	if err != nil {
		return nil, err
	}
	verificationMethod, err := did.NewVerificationMethod(*keyID, didtypes.ED25519VerificationKey2018, did.DID{}, publicKey)
	if err != nil {
		return nil, err
	}

	// This adds the method to the VerificationMethod list and stores a reference to the assertion list
	doc.AddAssertionMethod(verificationMethod)
	doc.AddAuthenticationMethod(verificationMethod)

	didJson, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, err
	}
	fmt.Println(string(didJson))

	// Unmarshalling of a json did document:
	parsedDIDDoc := did.Document{}
	err = json.Unmarshal(didJson, &parsedDIDDoc)
	if err != nil {
		return nil, err
	}

	if ok := parsedDIDDoc.IsController(*didDeviceID); !ok {
		return nil, fmt.Errorf("Device ID %s is not a controller of DID Document", *didDeviceID)
	}
	// It can return the key in the convenient lestrrat-go/jwx JWK
	_, err = parsedDIDDoc.AssertionMethod[0].JWK()
	if err != nil {
		return nil, err
	}

	// Or return a native crypto.PublicKey
	_, err = parsedDIDDoc.AssertionMethod[0].PublicKey()
	if err != nil {
		return nil, err
	}

	return &parsedDIDDoc, nil
}
