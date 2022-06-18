package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/registry module sentinel errors
var (
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrInsufficientFunds    = sdkerrors.Register(ModuleName, 1502, "insufficient funds")

	// Alias
	ErrAliasNotFound    = sdkerrors.Register(ModuleName, 1301, "alias not found")
	ErrAliasExists      = sdkerrors.Register(ModuleName, 1302, "alias already exists")
	ErrAliasUnavailable = sdkerrors.Register(ModuleName, 1303, "invalid alias")

	// Controller
	ErrControllerNotFound = sdkerrors.Register(ModuleName, 1401, "controller not found")
	ErrControllerExists   = sdkerrors.Register(ModuleName, 1402, "controller already exists")

	// Extensions
	ErrExtensionNotFound    = sdkerrors.Register(ModuleName, 1601, "extension not found")
	ErrExtensionExists      = sdkerrors.Register(ModuleName, 1602, "extension already exists")
	ErrExtensionUnavailable = sdkerrors.Register(ModuleName, 1603, "invalid extension")

	// DID
	ErrDidNotFound        = sdkerrors.Register(ModuleName, 1701, "did not found")
	ErrDidDeactivated     = sdkerrors.Register(ModuleName, 1702, "did deactivated")
	ErrDidDocumentInvalid = sdkerrors.Register(ModuleName, 1703, "did document invalid")

	// Query Account Creation
	ErrAccountExists    = sdkerrors.Register(ModuleName, 1801, "cannot create whois, account already exists")
	ErrInvalidBech32    = sdkerrors.Register(ModuleName, 1802, "invalid bech32 provided to creator")
	ErrFailedNewAccount = sdkerrors.Register(ModuleName, 1803, "failed to create new account")
)
