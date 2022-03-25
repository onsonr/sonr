#### Forked from github.com/nuts-foundation/go-did

# sonr-io/go-did

A library to parse and generate W3C [DID Documents](https://www.w3.org/TR/did-core/) and W3C [Verifiable Credentials](https://www.w3.org/TR/vc-data-model/).

## Did usage:

Creation of a simple DID Document which is its own controller and contains an AssertionMethod.

```go
didID, err := did.ParseDID("did:sonr:123")

// Empty did document:
doc := &did.Document{
    Context:            []did.URI{did.DIDContextV1URI()},
    ID:                 *didID,
}

// Add an assertionMethod
keyID, _ =: did.ParseDIDURL("did:sonr:123#key-1")

keyPair, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
verificationMethod, err := did.NewVerificationMethod(*keyID, did.JsonWebKey2020, did.DID{}, keyPair.Public())

// This adds the method to the VerificationMethod list and stores a reference to the assertion list
doc.AddAssertionMethod(verificationMethod)

didJson, _ := json.MarshalIndent(doc, "", "  ")
fmt.Println(string(didJson))

// Unmarshalling of a json did document:
parsedDIDDoc := did.Document{}
err = json.Unmarshal(didJson, &parsedDIDDoc)

// It can return the key in the convenient lestrrat-go/jwx JWK
parsedDIDDoc.AssertionMethod[0].JWK()

// Or return a native crypto.PublicKey
parsedDIDDoc.AssertionMethod[0].PublicKey()

```

Outputs:

```json
{
  "assertionMethod": ["did:sonr:123#key-1"],
  "@context": "https://www.w3.org/ns/did/v1",
  "controller": "did:sonr:123",
  "id": "did:sonr:123",
  "verificationMethod": [
    {
      "controller": "did:sonr:123",
      "id": "did:sonr:123#key-1",
      "publicKeyJwk": {
        "crv": "P-256",
        "kty": "EC",
        "x": "UANQ8pgvJT33JbrnwMiu1L1JCGQFOEm1ThaNAJcFrWA=",
        "y": "UWm6q5n1iXyeCJLMGDInN40bkkKr8KkoTWDqJBZQXRo="
      },
      "type": "JsonWebKey2020"
    }
  ]
}
```

## Installation

```
go get github.com/nuts-foundation/go-did
```

## State of the library

Currently, the library is under development. The api can change without notice.
Checkout the issues and PRs to be informed about any development.
