package types

import "github.com/di-dao/core/crypto/tecdsa/dklsv1"

// KssI is the interface for the keyshare set
type KssI interface {
	Val() KssValI
	Usr() KssUserI
	PublicKey() *PublicKey
}

// KssUserI is the interface for the user keyshare
type KssUserI interface {
	GetSignFunc(msg []byte) (SignFuncUser, error)
	GetRefreshFunc() (RefreshFuncUser, error)
	PublicKey() *PublicKey
}

// KssValI is the interface for the validator keyshare
type KssValI interface {
	GetSignFunc(msg []byte) (SignFuncVal, error)
	GetRefreshFunc() (RefreshFuncVal, error)
	PublicKey() *PublicKey
}

// RefreshFuncUser is the type for the user refresh function
type RefreshFuncUser = *dklsv1.BobRefresh

// RefreshFuncVal is the type for the validator refresh function
type RefreshFuncVal = *dklsv1.AliceRefresh

// SignFuncVal is the type for the validator sign function
type SignFuncVal = *dklsv1.AliceSign

// SignFuncUser is the type for the user sign function
type SignFuncUser = *dklsv1.BobSign
