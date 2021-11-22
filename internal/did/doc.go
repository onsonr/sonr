package did

import (
	"fmt"
	"regexp"
	"strings"
)

// DID is parsed according to the generic syntax: https://w3c.github.io/did-core/#generic-did-syntax
type DID struct {
	Scheme           string // Scheme is always "did"
	Method           string // Method is the specific DID methods
	MethodSpecificID string // MethodSpecificID is the unique ID computed or assigned by the DID method
}

// String returns a string representation of this DID.
func (d *DID) String() string {
	return fmt.Sprintf("%s:%s:%s", d.Scheme, d.Method, d.MethodSpecificID)
}

// Parse parses the string according to the generic DID syntax.
// See https://w3c.github.io/did-core/#generic-did-syntax.
func Parse(did string) (*DID, error) {
	// I could not find a good ABNF parser :(
	const idchar = `a-zA-Z0-9-_\.`
	regex := fmt.Sprintf(`^did:[a-z0-9]+:(:+|[:%s]+)*[%%:%s]+[^:]$`, idchar, idchar)

	r, err := regexp.Compile(regex)
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex=%s (this should not have happened!). %w", regex, err)
	}

	if !r.MatchString(did) {
		return nil, fmt.Errorf(
			"invalid did: %s. Make sure it conforms to the DID syntax: https://w3c.github.io/did-core/#did-syntax", did)
	}

	parts := strings.SplitN(did, ":", 3)

	return &DID{
		Scheme:           "did",
		Method:           parts[1],
		MethodSpecificID: parts[2],
	}, nil
}



// func NewDoc(pub crypto.PubKey, sname string) (*did.Document, error) {
// 	// Create DID Document
// 	didID, err := did.ParseDID(fmt.Sprintf("did:sonr:%s", sname))
// 	if err != nil {
// 		return nil, err
// 	}
// 	deviceid, err := device.ID()
// 	if err != nil {
// 		return nil, err
// 	}

// 	didDeviceID, err := did.ParseDID(fmt.Sprintf("did:sonr:%s", deviceid))
// 	if err != nil {
// 		return nil, err
// 	}

// 	publicKey, err := crypto.PubKeyToStdKey(pub)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Empty did document:
// 	doc := &did.Document{
// 		Context: []didtypes.URI{did.DIDContextV1URI()},
// 		ID:      *didID,
// 		Controller: []did.DID{
// 			*didDeviceID,
// 		},
// 	}

// 	// Add an assertionMethod
// 	keyID, err := did.ParseDIDURL(fmt.Sprintf("did:sonr:%s#master", sname))
// 	if err != nil {
// 		return nil, err
// 	}
// 	verificationMethod, err := did.NewVerificationMethod(*keyID, didtypes.ED25519VerificationKey2018, did.DID{}, publicKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// This adds the method to the VerificationMethod list and stores a reference to the assertion list
// 	doc.AddAssertionMethod(verificationMethod)
// 	doc.AddAuthenticationMethod(verificationMethod)

// 	didJson, err := json.MarshalIndent(doc, "", "  ")
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println(string(didJson))

// 	// Unmarshalling of a json did document:
// 	parsedDIDDoc := did.Document{}
// 	err = json.Unmarshal(didJson, &parsedDIDDoc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if ok := parsedDIDDoc.IsController(*didDeviceID); !ok {
// 		return nil, fmt.Errorf("Device ID %s is not a controller of DID Document", *didDeviceID)
// 	}
// 	// It can return the key in the convenient lestrrat-go/jwx JWK
// 	_, err = parsedDIDDoc.AssertionMethod[0].JWK()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Or return a native crypto.PublicKey
// 	_, err = parsedDIDDoc.AssertionMethod[0].PublicKey()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &parsedDIDDoc, nil
// }
