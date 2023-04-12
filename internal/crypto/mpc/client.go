package mpc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CustomMPC is an interface that defines the methods needed for MPC-based transactions.
type CustomMPC interface {
	CreateAndSignTx(from, to sdk.AccAddress, amount sdk.Coin, memo string, accountNumber, sequence uint64) (sdk.Tx, error)
	BroadcastTx(tx sdk.Tx) error
}

// MPCClient is an implementation of the CustomMPC interface.
type MPCClient struct {
	// ... other fields
}
