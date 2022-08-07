package mpc

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rsa"
	"encoding/asn1"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/fxamacker/cbor/v2"
)

//Error represents an error in a WebAuthn relying party operation
type Error struct {
	err     string
	wrapped error
}

//Error implements the error interface
func (e Error) Error() string {
	return e.err
}

//Unwrap allows for error unwrapping
func (e Error) Unwrap() error {
	return e.wrapped
}

//Wrap returns a new error which contains the provided error wrapped with this
//error
func (e Error) Wrap(err error) Error {
	n := e
	n.wrapped = err
	return n
}

//Is establishes equality for error types
func (e Error) Is(target error) bool {
	return e.Error() == target.Error()
}

//NewError returns a new Error with a custom message
func NewError(fmStr string, els ...interface{}) Error {
	return Error{
		err: fmt.Sprintf(fmStr, els...),
	}
}

//Categorical top-level errors
var (
	ErrDecodeAttestedCredentialData = Error{err: "error decoding attested credential data"}
	ErrDecodeAuthenticatorData      = Error{err: "error decoding authenticator data"}
	ErrDecodeCOSEKey                = Error{err: "error decoding raw public key"}
	ErrECDAANotSupported            = Error{err: "ECDAA not supported"}
	ErrEncodeAttestedCredentialData = Error{err: "error encoding attested credential data"}
	ErrEncodeAuthenticatorData      = Error{err: "error encoding authenticator data"}
	ErrGenerateChallenge            = Error{err: "error generating challenge"}
	ErrMarshalAttestationObject     = Error{err: "error marshaling attestation object"}
	ErrOption                       = Error{err: "option error"}
	ErrNotImplemented               = Error{err: "not implemented"}
	ErrUnmarshalAttestationObject   = Error{err: "error unmarshaling attestation object"}
	ErrVerifyAttestation            = Error{err: "error verifying attestation"}
	ErrVerifyAuthentication         = Error{err: "error verifying authentication"}
	ErrVerifyClientExtensionOutput  = Error{err: "error verifying client extension output"}
	ErrVerifyRegistration           = Error{err: "error verifying registration"}
	ErrVerifySignature              = Error{err: "error verifying signature"}
)

//COSEKey represents a key decoded from COSE format.
type COSEKey struct {
	Kty       int             `cbor:"1,keyasint,omitempty"`
	Kid       []byte          `cbor:"2,keyasint,omitempty"`
	Alg       int             `cbor:"3,keyasint,omitempty"`
	KeyOpts   int             `cbor:"4,keyasint,omitempty"`
	IV        []byte          `cbor:"5,keyasint,omitempty"`
	CrvOrNOrK cbor.RawMessage `cbor:"-1,keyasint,omitempty"` // K for symmetric keys, Crv for elliptic curve keys, N for RSA modulus
	XOrE      cbor.RawMessage `cbor:"-2,keyasint,omitempty"` // X for curve x-coordinate, E for RSA public exponent
	Y         cbor.RawMessage `cbor:"-3,keyasint,omitempty"` // Y for curve y-cooridate
	D         []byte          `cbor:"-4,keyasint,omitempty"`
}

//COSEAlgorithmIdentifier is a number identifying a cryptographic algorithm
type COSEAlgorithmIdentifier int

//enum values for COSEAlgorithmIdentifier type
const (
	AlgorithmRS1   COSEAlgorithmIdentifier = -65535
	AlgorithmRS512 COSEAlgorithmIdentifier = -259
	AlgorithmRS384 COSEAlgorithmIdentifier = -258
	AlgorithmRS256 COSEAlgorithmIdentifier = -257
	AlgorithmPS512 COSEAlgorithmIdentifier = -39
	AlgorithmPS384 COSEAlgorithmIdentifier = -38
	AlgorithmPS256 COSEAlgorithmIdentifier = -37
	AlgorithmES512 COSEAlgorithmIdentifier = -36
	AlgorithmES384 COSEAlgorithmIdentifier = -35
	AlgorithmEdDSA COSEAlgorithmIdentifier = -8
	AlgorithmES256 COSEAlgorithmIdentifier = -7
)

//COSEEllipticCurve is a number identifying an elliptic curve
type COSEEllipticCurve int

//enum values for COSEEllipticCurve type
const (
	CurveP256 COSEEllipticCurve = 1
	CurveP384 COSEEllipticCurve = 2
	CurveP521 COSEEllipticCurve = 3
)

//COSEKeyType is a number identifying a key type
type COSEKeyType int

//enum values for COSEKeyType type
const (
	KeyTypeOKP COSEKeyType = 1
	KeyTypeEC2 COSEKeyType = 2
	KeyTypeRSA COSEKeyType = 3
)

//VerifySignature verifies a signature using a provided COSEKey, message, and
//signature
func VerifySignature(rawKey cbor.RawMessage, message, sig []byte) error {
	coseKey := COSEKey{}
	err := cbor.Unmarshal(rawKey, &coseKey)
	if err != nil {
		return ErrVerifySignature.Wrap(ErrDecodeCOSEKey.Wrap(err))
	}

	publicKey, err := DecodePublicKey(&coseKey)
	if err != nil {
		return ErrVerifySignature.Wrap(err)
	}

	switch COSEAlgorithmIdentifier(coseKey.Alg) {
	case AlgorithmES256,
		AlgorithmES384,
		AlgorithmES512:
		pk, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return ErrVerifySignature.Wrap(NewError("Invalid public key type for ECDSA algorithm"))
		}
		switch COSEAlgorithmIdentifier(coseKey.Alg) {
		case AlgorithmES256:
			return verifyECDSASignature(pk, crypto.SHA256, message, sig)
		case AlgorithmES384:
			return verifyECDSASignature(pk, crypto.SHA384, message, sig)
		case AlgorithmES512:
			return verifyECDSASignature(pk, crypto.SHA512, message, sig)
		}

	case AlgorithmRS1,
		AlgorithmRS512,
		AlgorithmRS384,
		AlgorithmRS256,
		AlgorithmPS512,
		AlgorithmPS384,
		AlgorithmPS256:
		pk, ok := publicKey.(*rsa.PublicKey)
		if !ok {
			return ErrVerifySignature.Wrap(NewError("Invalid public key type for RSA algorithm"))
		}

		switch COSEAlgorithmIdentifier(coseKey.Alg) {
		case AlgorithmRS1:
			return verifyRSAPKCS1v15Signature(pk, crypto.SHA1, message, sig)
		case AlgorithmRS256:
			return verifyRSAPKCS1v15Signature(pk, crypto.SHA256, message, sig)
		case AlgorithmRS384:
			return verifyRSAPKCS1v15Signature(pk, crypto.SHA384, message, sig)
		case AlgorithmRS512:
			return verifyRSAPKCS1v15Signature(pk, crypto.SHA512, message, sig)
		case AlgorithmPS256:
			return verifyRSAPSSSignature(pk, crypto.SHA256, message, sig)
		case AlgorithmPS384:
			return verifyRSAPSSSignature(pk, crypto.SHA384, message, sig)
		case AlgorithmPS512:
			return verifyRSAPSSSignature(pk, crypto.SHA512, message, sig)
		}

	case AlgorithmEdDSA:
		pk, ok := publicKey.(ed25519.PublicKey)
		if !ok {
			return ErrVerifySignature.Wrap(NewError("Invalid public key type for EdDSA algorithm"))
		}
		if ed25519.Verify(pk, message, sig) {
			return nil
		}
		return ErrVerifySignature.Wrap(NewError("EdDSA signature verification failed"))
	}
	return ErrVerifySignature.Wrap(NewError("COSE algorithm ID %d not supported", coseKey.Alg))
}

//DecodePublicKey parses a crypto.PublicKey from a COSEKey
func DecodePublicKey(coseKey *COSEKey) (crypto.PublicKey, error) {
	var publicKey crypto.PublicKey

	switch COSEKeyType(coseKey.Kty) {
	case KeyTypeOKP:
		k, err := decodeEd25519PublicKey(coseKey)
		if err != nil {
			return nil, ErrDecodeCOSEKey.Wrap(err)
		}
		publicKey = k
	case KeyTypeEC2:
		k, err := decodeECDSAPublicKey(coseKey)
		if err != nil {
			return nil, ErrDecodeCOSEKey.Wrap(err)
		}
		publicKey = k
	case KeyTypeRSA:
		k, err := decodeRSAPublicKey(coseKey)
		if err != nil {
			return nil, ErrDecodeCOSEKey.Wrap(err)
		}
		publicKey = k
	default:
		return nil, ErrDecodeCOSEKey.Wrap(NewError("COSE key type %d not supported", coseKey.Kty))
	}

	return publicKey, nil
}

func decodeECDSAPublicKey(coseKey *COSEKey) (*ecdsa.PublicKey, error) {
	var curve elliptic.Curve
	var curveID int
	if err := cbor.Unmarshal(coseKey.CrvOrNOrK, &curveID); err != nil {
		return nil, NewError("Error decoding elliptic curve ID").Wrap(err)
	}

	switch COSEEllipticCurve(curveID) {
	case CurveP256:
		curve = elliptic.P256()
	case CurveP384:
		curve = elliptic.P384()
	case CurveP521:
		curve = elliptic.P521()
	default:
		return nil, NewError("COSE elliptic curve %d not supported", curveID)
	}

	var xBytes, yBytes []byte
	if err := cbor.Unmarshal(coseKey.XOrE, &xBytes); err != nil {
		return nil, NewError("Error decoding elliptic X parameter").Wrap(err)
	}
	if err := cbor.Unmarshal(coseKey.Y, &yBytes); err != nil {
		return nil, NewError("Error decoding elliptic Y parameter").Wrap(err)
	}

	x, y := big.NewInt(0), big.NewInt(0)
	x = x.SetBytes(xBytes)
	y = y.SetBytes(yBytes)

	return &ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}, nil
}

func verifyECDSASignature(pubKey *ecdsa.PublicKey, hash crypto.Hash, message, sig []byte) error {
	type ECDSASignature struct {
		R, S *big.Int
	}

	ecdsaSig := ECDSASignature{}
	_, err := asn1.Unmarshal(sig, &ecdsaSig)
	if err != nil {
		return ErrVerifySignature.Wrap(NewError("Unable to parse ECDSA signature").Wrap(err))
	}

	hasher := hash.New()
	hasher.Write(message)

	if ecdsa.Verify(pubKey, hasher.Sum(nil), ecdsaSig.R, ecdsaSig.S) {
		return nil
	}
	return ErrVerifySignature.Wrap(NewError("ECDSA signature verification failed"))
}

func decodeRSAPublicKey(coseKey *COSEKey) (*rsa.PublicKey, error) {
	var nBytes, eBytes []byte
	if err := cbor.Unmarshal(coseKey.CrvOrNOrK, &nBytes); err != nil {
		return nil, NewError("Error unmarshaling RSA modulus").Wrap(err)
	}
	if err := cbor.Unmarshal(coseKey.XOrE, &eBytes); err != nil {
		return nil, NewError("Error unmarshaling RSA exponent").Wrap(err)
	}

	n := big.NewInt(0)
	var e int32
	n = n.SetBytes(nBytes)
	if err := binary.Read(bytes.NewBuffer(eBytes), binary.BigEndian, &e); err != nil {
		return nil, NewError("Error decoding RSA exponent").Wrap(err)
	}

	return &rsa.PublicKey{
		N: n,
		E: int(e),
	}, nil
}

func verifyRSAPKCS1v15Signature(pubKey *rsa.PublicKey, hash crypto.Hash, message, sig []byte) error {
	hasher := hash.New()
	hasher.Write(message)

	err := rsa.VerifyPKCS1v15(pubKey, hash, hasher.Sum(nil), sig)
	if err != nil {
		return ErrVerifySignature.Wrap(NewError("RSA signature verification failed").Wrap(err))
	}
	return nil
}

func verifyRSAPSSSignature(pubKey *rsa.PublicKey, hash crypto.Hash, message, sig []byte) error {
	hasher := hash.New()
	hasher.Write(message)

	err := rsa.VerifyPSS(pubKey, hash, hasher.Sum(nil), sig, nil)
	if err != nil {
		return ErrVerifySignature.Wrap(NewError("RSA signature verification failed").Wrap(err))
	}
	return nil
}

func decodeEd25519PublicKey(coseKey *COSEKey) (ed25519.PublicKey, error) {
	var kBytes []byte
	if err := cbor.Unmarshal(coseKey.XOrE, &kBytes); err != nil {
		return nil, NewError("Error unmarshaling Ed25519 public key").Wrap(err)
	}

	k := ed25519.PublicKey(kBytes)
	return k, nil
}
