// Code generated from Pkl module `common.types.ORM`. DO NOT EDIT.
package keytype

import (
	"encoding"
	"fmt"
)

type KeyType string

const (
	Octet     KeyType = "octet"
	Elliptic  KeyType = "elliptic"
	Rsa       KeyType = "rsa"
	Symmetric KeyType = "symmetric"
	Hmac      KeyType = "hmac"
	Mpc       KeyType = "mpc"
	Zk        KeyType = "zk"
	Webauthn  KeyType = "webauthn"
	Bip32     KeyType = "bip32"
)

// String returns the string representation of KeyType
func (rcv KeyType) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(KeyType)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for KeyType.
func (rcv *KeyType) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "octet":
		*rcv = Octet
	case "elliptic":
		*rcv = Elliptic
	case "rsa":
		*rcv = Rsa
	case "symmetric":
		*rcv = Symmetric
	case "hmac":
		*rcv = Hmac
	case "mpc":
		*rcv = Mpc
	case "zk":
		*rcv = Zk
	case "webauthn":
		*rcv = Webauthn
	case "bip32":
		*rcv = Bip32
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid KeyType`, str)
	}
	return nil
}
