# Crypto Package

| **Important**: Kryptology package Original development by Coinbase - now maintained internally at Sonr.

[![Go Reference](https://pkg.go.dev/badge/github.com/sonrhq/sonr/crypto.svg)](https://pkg.go.dev/github.com/sonrhq/sonr/crypto)

Coinbase's advanced cryptography library used for Sonr's internal cryptographic operations.

## Usage

Use the latest version of this library:

```$xslt
go get github.com/sonrhq/sonr/crypto
```

Pin a specific release of this library:

```$xslt
go get github.com/sonrhq/sonr/crypto@v0.8.6
```

## Components

The following is the list of primitives and protocols that are implemented in this repository.

### Curves

The curve abstraction code can be found at [crypto/core/curves/curve.go](core/curves/curve.go)

The curves that implement this abstraction are as follows.

- [BLS12377](core/curves/bls12377_curve.go)
- [BLS12381](core/curves/bls12381_curve.go)
- [Ed25519](core/curves/ed25519_curve.go)
- [Secp256k1](core/curves/k256_curve.go)
- [P256](core/curves/p256_curve.go)
- [Pallas](core/curves/pallas_curve.go)

### Protocols

The generic protocol interface [crypto/core/protocol/protocol.go](core/protocol/protocol.go).
This abstraction is currently only used in DKLs18 implementation.

- [Cryptographic Accumulators](accumulator)
- [Bulletproof](bulletproof)
- Oblivious Transfer
  - [Verifiable Simplest OT](ot/base/simplest)
  - [KOS OT Extension](ot/extension/kos)
- Threshold ECDSA Signature
  - [DKLs18 - DKG and Signing](tecdsa/dkls/v1)
  - GG20: The authors of GG20 have stated that the protocol is obsolete and should not be used. See [https://eprint.iacr.org/2020/540.pdf](https://eprint.iacr.org/2020/540.pdf).
    - [GG20 - DKG](dkg/gennaro)
    - [GG20 - Signing](tecdsa/gg20)
- Threshold Schnorr Signature
  - [FROST threshold signature - DKG](dkg/frost)
  - [FROST threshold signature - Signing](ted25519/frost)
- [Paillier encryption system](paillier)
- Secret Sharing Schemes
  - [Shamir's secret sharing scheme](sharing/shamir.go)
  - [Pedersen](sharing/pedersen.go)
  - [Feldman](sharing/feldman.go)
- [Verifiable encryption](verenc)
- [ZKP Schnorr](zkp/schnorr)

## Contributing

- [Versioning](https://blog.golang.org/publishing-go-modules): `vMajor.Minor.Patch`
  - Major revision indicates breaking API change or significant new features
  - Minor revision indicates no API breaking changes and may include significant new features or documentation
  - Patch indicates no API breaking changes and may include only fixes

## [References](docs/)

- [[GG20] _One Round Threshold ECDSA with Identifiable Abort._](https://eprint.iacr.org/2020/540.pdf)
- [[specV5] _One Round Threshold ECDSA for Coinbase._](docs/Coinbase_Pseudocode_v5.pdf)
- [[EL20] _Eliding RSA Group Membership Checks._](docs/rsa-membership.pdf) [src](https://www.overleaf.com/project/5f9c3b0624a9a600012037a3)
- [[P99] _Public-Key Cryptosystems Based on Composite Degree Residuosity Classes._](http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.112.4035&rep=rep1&type=pdf)
