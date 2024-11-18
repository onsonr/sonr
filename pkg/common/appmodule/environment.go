package appmodule

import (
	"cosmossdk.io/core/event"
	"cosmossdk.io/core/gas"
	"cosmossdk.io/core/header"
	"cosmossdk.io/core/store"

	"github.com/onsonr/sonr/pkg/common/branch"
	"github.com/onsonr/sonr/pkg/common/log"
	"github.com/onsonr/sonr/pkg/common/router"
	"github.com/onsonr/sonr/pkg/common/transaction"
)

// Environment is used to get all services to their respective module.
// Contract: All fields of environment are always populated by runtime.
type Environment struct {
	Logger log.Logger

	BranchService      branch.Service
	EventService       event.Service
	GasService         gas.Service
	HeaderService      header.Service
	QueryRouterService router.Service
	MsgRouterService   router.Service
	TransactionService transaction.Service

	KVStoreService  store.KVStoreService
	MemStoreService store.MemoryStoreService
}
