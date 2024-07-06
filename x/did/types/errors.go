package types

import sdkerrors "cosmossdk.io/errors"

var (
	ErrInvalidGenesisState     = sdkerrors.Register(ModuleName, 100, "invalid genesis state")
	ErrInvalidETHAddressFormat = sdkerrors.Register(ModuleName, 200, "invalid ETH address format")
	ErrInvalidBTCAddressFormat = sdkerrors.Register(ModuleName, 201, "invalid BTC address format")
	ErrInvalidIDXAddressFormat = sdkerrors.Register(ModuleName, 202, "invalid IDX address format")
	ErrInvalidEmailFormat      = sdkerrors.Register(ModuleName, 203, "invalid email format")
	ErrInvalidPhoneFormat      = sdkerrors.Register(ModuleName, 204, "invalid phone format")
	ErrMinimumAssertions       = sdkerrors.Register(ModuleName, 300, "at least one assertion is required for account initialization")
	ErrInvalidControllers      = sdkerrors.Register(ModuleName, 301, "no more than one controller can be used for account initialization")
	ErrMaximumAuthenticators   = sdkerrors.Register(ModuleName, 302, "more authenticators provided than the total accepted count")
)
