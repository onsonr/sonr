
# Sonr Core

**Sonr Core** is a module for common interfaces used by the following repositories:

- `sonr-io/highway`
- `sonr-io/motor`
- `sonr-io/sonr`

## Description

Sonr is building the most immersive DWeb experience for both Users and Developers alike. We believe the best way to onboard the next billion users is to create a cohesive end-to-end platform that’s composable and interoperable with all existing protocols.

For this we built our Networking layer in [Libp2p](“https://libp2p.io”) and our Layer 1 Blockchain with [Starport](“https://starport.com”). Our network comprises of two separate nodes: [Highway](“https://github.com/sonr-io/highway”) and [Motor](“https://github.com/sonr-io/motor”), which each have a specific use case on the network.

## Getting Started

### Dependencies

- [Golang](https://go.dev)
- [Libp2p](https://libp2p.io)

### Installing

To install the latest version of the Sonr blockchain node's binary, execute the following command on your machine:

``` shell
go get -u https://github.com/sonr-io/core
```

### Configuration

This project is a pseudo-monorepo, meaning it has a single root directory and all of its packages are in subdirectories. The structure is as follows:

``` text
/channel         ->        Real-time Key/Value Store
/common          ->        Core data types and functions.
/device          ->        Node Device management
/docs            ->        Documentation.
/exchange        ->        Data Transfer related Models.
/host            ->        Libp2p Host Configuration
/identity        ->        Identity management models and interfaces
/node            ->        Highway and Motor node builder configuration
/proto           ->        Protobuf Definition Files.
/transmit        ->        Protocol for byte transmission between nodes
/types           ->        Protobuf Compiled Types
  └─ cpp         ->        +   C++ Definition Files
  └─ go          ->        +   Golang Definition Files
  └─ java        ->        +   Java Definition Files
/wallet          ->        Interfaces for managing Universal Wallet
```

## Getting Started

## Decentralized Identifiers (DIDs)

A library to parse and generate W3C [DID Documents](https://www.w3.org/TR/did-core/) and W3C [Verifiable Credentials](https://www.w3.org/TR/vc-data-model/).

### Did usage:

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

### State of the library

Currently, the library is under development. The api can change without notice.
Checkout the issues and PRs to be informed about any development.


## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Authors

Contributors names and contact info

- [Prad Nukala](“https://github.com/prnk28”)

## License

This project facilitated under **Sonr Inc.** is distributed under the **GPLv3 License**. See `LICENSE.md` for more information.

## Acknowledgments

Inspiration, code snippets, etc.

- [Libp2p](https://libp2p.io/)
- [Textile](https://www.textile.io/)
- [Handshake](https://handshake.org/)
