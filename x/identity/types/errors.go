package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/identity module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1001, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")
)

// x/identity module sentinel errors
var (
	ErrMpc                   = sdkerrors.Register(ModuleName, 1100, "internal mpc error")
	ErrMpcTimeout            = sdkerrors.Register(ModuleName, 1101, "mpc timeout")
	ErrMpcNotReady           = sdkerrors.Register(ModuleName, 1102, "mpc not ready")
	ErrWebauthnCred          = sdkerrors.Register(ModuleName, 1200, "webauthn credential error")
	ErrWebauthnCredNotFound  = sdkerrors.Register(ModuleName, 1201, "webauthn credential not found")
	ErrWebauthnCredCollision = sdkerrors.Register(ModuleName, 1202, "webauthn credential already exists")
	ErrWebauthnCredAssign    = sdkerrors.Register(ModuleName, 1203, "webauthn credential assignment error")
	ErrWebauthnCredVerify    = sdkerrors.Register(ModuleName, 1204, "webauthn credential verification error")
	ErrUnauthorized          = sdkerrors.Register(ModuleName, 1300, "unauthorized")
	ErrInvalidDid            = sdkerrors.Register(ModuleName, 1400, "invalid did")
	ErrDidCollision          = sdkerrors.Register(ModuleName, 2100, "did already exists")
	ErrDidNotFound           = sdkerrors.Register(ModuleName, 2101, "did not found")
	ErrServiceCollision      = sdkerrors.Register(ModuleName, 3100, "service already exists")
	ErrServiceNotFound       = sdkerrors.Register(ModuleName, 3101, "service not found")
	ErrAliasCollision        = sdkerrors.Register(ModuleName, 4100, "alias already exists")
	ErrAliasNotFound         = sdkerrors.Register(ModuleName, 4101, "alias not found")
	ErrWalletAccountCreation = sdkerrors.Register(ModuleName, 5100, "wallet account creation error")
	ErrWalletAccountNotFound = sdkerrors.Register(ModuleName, 5101, "wallet account not found")
)
