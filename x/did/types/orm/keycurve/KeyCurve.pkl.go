// Code generated from Pkl module `models`. DO NOT EDIT.
package keycurve

import (
	"encoding"
	"fmt"
)

type KeyCurve string

const (
	P256      KeyCurve = "p256"
	P384      KeyCurve = "p384"
	P521      KeyCurve = "p521"
	X25519    KeyCurve = "x25519"
	X448      KeyCurve = "x448"
	Ed25519   KeyCurve = "ed25519"
	Ed448     KeyCurve = "ed448"
	Secp256k1 KeyCurve = "secp256k1"
	Bls12381  KeyCurve = "bls12381"
	Keccak256 KeyCurve = "keccak256"
)

// String returns the string representation of KeyCurve
func (rcv KeyCurve) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(KeyCurve)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for KeyCurve.
func (rcv *KeyCurve) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "p256":
		*rcv = P256
	case "p384":
		*rcv = P384
	case "p521":
		*rcv = P521
	case "x25519":
		*rcv = X25519
	case "x448":
		*rcv = X448
	case "ed25519":
		*rcv = Ed25519
	case "ed448":
		*rcv = Ed448
	case "secp256k1":
		*rcv = Secp256k1
	case "bls12381":
		*rcv = Bls12381
	case "keccak256":
		*rcv = Keccak256
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid KeyCurve`, str)
	}
	return nil
}
