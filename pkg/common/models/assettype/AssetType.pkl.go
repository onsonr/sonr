// Code generated from Pkl module `sonr.orm.Models`. DO NOT EDIT.
package assettype

import (
	"encoding"
	"fmt"
)

type AssetType string

const (
	Native  AssetType = "native"
	Wrapped AssetType = "wrapped"
	Staking AssetType = "staking"
	Pool    AssetType = "pool"
	Ibc     AssetType = "ibc"
	Cw20    AssetType = "cw20"
)

// String returns the string representation of AssetType
func (rcv AssetType) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(AssetType)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for AssetType.
func (rcv *AssetType) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "native":
		*rcv = Native
	case "wrapped":
		*rcv = Wrapped
	case "staking":
		*rcv = Staking
	case "pool":
		*rcv = Pool
	case "ibc":
		*rcv = Ibc
	case "cw20":
		*rcv = Cw20
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid AssetType`, str)
	}
	return nil
}
