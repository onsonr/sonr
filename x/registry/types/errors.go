package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/registry module sentinel errors
var (
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")

	// Controller
	ErrControllerNotFound = sdkerrors.Register(ModuleName, 1401, "controller not found")
	ErrControllerExists   = sdkerrors.Register(ModuleName, 1402, "controller already exists")

	// Alias
	ErrAliasNotFound    = sdkerrors.Register(ModuleName, 1403, "alias not found")
	ErrAliasExists      = sdkerrors.Register(ModuleName, 1404, "alias already exists")
	ErrAliasUnavailable = sdkerrors.Register(ModuleName, 1405, "invalid alias")

	// Extensions
	ErrExtensionNotFound    = sdkerrors.Register(ModuleName, 1406, "extension not found")
	ErrExtensionExists      = sdkerrors.Register(ModuleName, 1407, "extension already exists")
	ErrExtensionUnavailable = sdkerrors.Register(ModuleName, 1408, "invalid extension")

	// DID
	ErrDidNotFound        = sdkerrors.Register(ModuleName, 1409, "did not found")
	ErrDidDeactivated     = sdkerrors.Register(ModuleName, 1410, "did deactivated")
	ErrDidDocumentInvalid = sdkerrors.Register(ModuleName, 1411, "did document invalid")
)
