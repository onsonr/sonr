// Code generated from Pkl module `format`. DO NOT EDIT.
package didmethod

import (
	"encoding"
	"fmt"
)

type DIDMethod string

const (
	Ipfs     DIDMethod = "ipfs"
	Sonr     DIDMethod = "sonr"
	Bitcoin  DIDMethod = "bitcoin"
	Ethereum DIDMethod = "ethereum"
	Ibc      DIDMethod = "ibc"
	Webauthn DIDMethod = "webauthn"
	Dwn      DIDMethod = "dwn"
	Service  DIDMethod = "service"
)

// String returns the string representation of DIDMethod
func (rcv DIDMethod) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(DIDMethod)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for DIDMethod.
func (rcv *DIDMethod) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "ipfs":
		*rcv = Ipfs
	case "sonr":
		*rcv = Sonr
	case "bitcoin":
		*rcv = Bitcoin
	case "ethereum":
		*rcv = Ethereum
	case "ibc":
		*rcv = Ibc
	case "webauthn":
		*rcv = Webauthn
	case "dwn":
		*rcv = Dwn
	case "service":
		*rcv = Service
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid DIDMethod`, str)
	}
	return nil
}
