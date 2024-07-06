// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
package types

import (
	"bytes"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	_ authtypes.AccountI                 = (*EthAccount)(nil)
	_ EthAccountI                        = (*EthAccount)(nil)
	_ authtypes.GenesisAccount           = (*EthAccount)(nil)
	_ codectypes.UnpackInterfacesMessage = (*EthAccount)(nil)
)

var emptyCodeHash = crypto.Keccak256(nil)

const (
	// AccountTypeEOA defines the type for externally owned accounts (EOAs)
	AccountTypeEOA = int8(iota + 1)
	// AccountTypeContract defines the type for contract accounts
	AccountTypeContract
)

// EthAccountI represents the interface of an EVM compatible account
type EthAccountI interface {
	authtypes.AccountI
	// EthAddress returns the ethereum Address representation of the AccAddress
	EthAddress() common.Address
	// CodeHash is the keccak256 hash of the contract code (if any)
	GetCodeHash() common.Hash
	// SetCodeHash sets the code hash to the account fields
	SetCodeHash(code common.Hash) error
	// Type returns the type of Ethereum Account (EOA or Contract)
	Type() int8
}

// ----------------------------------------------------------------------------
// Main Evmos account
// ----------------------------------------------------------------------------

// ProtoAccount defines the prototype function for BaseAccount used for an
// AccountKeeper.
func ProtoAccount() authtypes.AccountI {
	return &EthAccount{
		BaseAccount: &authtypes.BaseAccount{},
		CodeHash:    common.BytesToHash(emptyCodeHash).String(),
	}
}

// GetAccountNumber returns account number.
func (acc EthAccount) GetAccountNumber() uint64 {
	return acc.BaseAccount.AccountNumber
}

// GetAddress returns account address.
func (acc EthAccount) GetAddress() sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(acc.BaseAccount.Address)
	return addr
}

// GetBaseAccount returns base account.
func (acc EthAccount) GetBaseAccount() *authtypes.BaseAccount {
	return acc.BaseAccount
}

// GetPubKey returns the PubKey
func (acc EthAccount) GetPubKey() cryptotypes.PubKey {
	return acc.GetBaseAccount().GetPubKey()
}

// GetSequence returns the sequence
func (acc EthAccount) GetSequence() uint64 {
	return acc.BaseAccount.Sequence
}

// EthAddress returns the account address ethereum format.
func (acc EthAccount) EthAddress() common.Address {
	return common.BytesToAddress(acc.GetAddress().Bytes())
}

// GetCodeHash returns the account code hash in byte format
func (acc EthAccount) GetCodeHash() common.Hash {
	return common.HexToHash(acc.CodeHash)
}

// SetAccountNumber sets the account number
func (acc *EthAccount) SetAccountNumber(accNum uint64) error {
	acc.BaseAccount.AccountNumber = accNum
	return nil
}

// SetAddress sets the address
func (acc *EthAccount) SetAddress(addr sdk.AccAddress) error {
	acc.BaseAccount.Address = addr.String()
	return nil
}

// SetCodeHash sets the account code hash to the EthAccount fields
func (acc *EthAccount) SetCodeHash(codeHash common.Hash) error {
	acc.CodeHash = codeHash.Hex()
	return nil
}

// SetPubKey sets the pubkey
func (acc *EthAccount) SetPubKey(pubkey cryptotypes.PubKey) error {
	acc.BaseAccount.PubKey = codectypes.UnsafePackAny(pubkey)
	return nil
}

// SetSequence sets the sequence
func (acc *EthAccount) SetSequence(seq uint64) error {
	acc.BaseAccount.Sequence = seq
	return nil
}

// Type returns the type of Ethereum Account (EOA or Contract)
func (acc EthAccount) Type() int8 {
	if bytes.Equal(emptyCodeHash, common.HexToHash(acc.CodeHash).Bytes()) {
		return AccountTypeEOA
	}
	return AccountTypeContract
}

// Stringreturns the string representation of the EthAccount
func (acc EthAccount) String() string {
	return acc.EthAddress().String()
}

// Validate checks if the Evmos account fields are valid
func (acc EthAccount) Validate() error {
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (acc EthAccount) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return nil
}
