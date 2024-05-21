package kss

import (
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
	"github.com/di-dao/core/x/did/types"
)

// KssI is the interface for the keyshare set
type SetI interface {
	Val() ValI
	Usr() UserI
	PublicKey() *types.PublicKey
}

// KssUserI is the interface for the user keyshare
type UserI interface {
	GetSignFunc(msg []byte) (SignFuncUser, error)
	GetRefreshFunc() (RefreshFuncUser, error)
	PublicKey() *types.PublicKey
}

// KssValI is the interface for the validator keyshare
type ValI interface {
	GetSignFunc(msg []byte) (SignFuncVal, error)
	GetRefreshFunc() (RefreshFuncVal, error)
	PublicKey() *types.PublicKey
}

// RefreshFuncUser is the type for the user refresh function
type RefreshFuncUser = *dklsv1.BobRefresh

// RefreshFuncVal is the type for the validator refresh function
type RefreshFuncVal = *dklsv1.AliceRefresh

// SignFuncVal is the type for the validator sign function
type SignFuncVal = *dklsv1.AliceSign

// SignFuncUser is the type for the user sign function
type SignFuncUser = *dklsv1.BobSign
