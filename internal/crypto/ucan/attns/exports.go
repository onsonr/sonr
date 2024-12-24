// Package attns implements the UCAN resource and capability types
package attns

import (
	"github.com/onsonr/sonr/internal/crypto/ucan"
	"github.com/onsonr/sonr/internal/crypto/ucan/attns/capaccount"
	"github.com/onsonr/sonr/internal/crypto/ucan/attns/capinterchain"
	"github.com/onsonr/sonr/internal/crypto/ucan/attns/capvault"
	"github.com/onsonr/sonr/internal/crypto/ucan/attns/resaccount"
	"github.com/onsonr/sonr/internal/crypto/ucan/attns/resinterchain"
	"github.com/onsonr/sonr/internal/crypto/ucan/attns/resvault"
)

// Capability hierarchy for sonr network
// -------------------------------------
// VAULT (DWN)
//
//	└─ CRUD/ASSET
//	└─ CRUD/AUTHZGRANT
//	└─ CRUD/PROFILE
//	└─ CRUD/RECORD
//	└─ USE/RECOVERY
//	└─ USE/SYNC
//	└─ USE/SIGNER
//
// ACCOUNT (DID)
//
//	└─ EXEC/BROADCAST
//	└─ EXEC/QUERY
//	└─ EXEC/SIMULATE
//	└─ EXEC/VOTE
//	└─ EXEC/DELEGATE
//	└─ EXEC/INVOKE
//	└─ EXEC/SEND
//
// INTERCHAIN
//
//	└─ TRANSFER/SWAP
//	└─ TRANSFER/SEND
//	└─ TRANSFER/ATOMIC
//	└─ TRANSFER/BATCH
//	└─ TRANSFER/P2P
//	└─ TRANSFER/SEND

type Capability string

const (
	CapExecBroadcast = capaccount.ExecBroadcast
	CapExecQuery     = capaccount.ExecQuery
	CapExecSimulate  = capaccount.ExecSimulate
	CapExecVote      = capaccount.ExecVote
	CapExecDelegate  = capaccount.ExecDelegate
	CapExecInvoke    = capaccount.ExecInvoke
	CapExecSend      = capaccount.ExecSend

	CapTransferSwap   = capinterchain.TransferSwap
	CapTransferSend   = capinterchain.TransferSend
	CapTransferAtomic = capinterchain.TransferAtomic
	CapTransferBatch  = capinterchain.TransferBatch
	CapTransferP2P    = capinterchain.TransferP2p

	CapCrudAsset      = capvault.CrudAsset
	CapCrudAuthzgrant = capvault.CrudAuthzgrant
	CapCrudProfile    = capvault.CrudProfile
	CapCrudRecord     = capvault.CrudRecord
	CapUseRecovery    = capvault.UseRecovery
	CapUseSync        = capvault.UseSync
	CapUseSigner      = capvault.UseSigner
)

type NewCapFunc func(string) ucan.Capability

type BuildResourceFunc func(string, string) ucan.Resource

func CreateArray(attns ...ucan.Attenuation) ucan.Attenuations {
	return ucan.Attenuations(attns)
}

func New(cap ucan.Capability, rsc ucan.Resource) ucan.Attenuation {
	return ucan.Attenuation{
		Cap: cap,
		Rsc: rsc,
	}
}

// NewAccountCap creates a new account capability
func NewAccountCap(ty capaccount.CapAccount) ucan.Capability {
	return capaccount.NewCap(ty)
}

// NewInterchainCap creates a new interchain capability
func NewInterchainCap(ty capinterchain.CapInterchain) ucan.Capability {
	return capinterchain.NewCap(ty)
}

// NewVaultCap creates a new vault capability
func NewVaultCap(ty capvault.CapVault) ucan.Capability {
	return capvault.NewCap(ty)
}

// BuildAccountResource creates a new account resource
func BuildAccountResource(ty resaccount.ResAccount, value string) ucan.Resource {
	return resaccount.Build(ty, value)
}

// BuildInterchainResource creates a new interchain resource
func BuildInterchainResource(ty resinterchain.ResInterchain, value string) ucan.Resource {
	return resinterchain.Build(ty, value)
}

// BuildVaultResource creates a new vault resource
func BuildVaultResource(ty resvault.ResVault, value string) ucan.Resource {
	return resvault.Build(ty, value)
}
