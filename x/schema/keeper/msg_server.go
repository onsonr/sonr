package keeper

import (
	"os"

	"github.com/sonr-io/sonr/pkg/protocol"
	"github.com/sonr-io/sonr/pkg/store"
	"github.com/sonr-io/sonr/x/schema/types"
)

var IPFSShellURL = os.Getenv("IPFS_API")

type msgServer struct {
	Keeper

	ipfs protocol.IPFS
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	memStore := store.NewMemoryStore()
	return &msgServer{
		Keeper: keeper,
		ipfs:   protocol.NewIPFSShell(IPFSShellURL, memStore.Datastore()),
	}
}

var _ types.MsgServer = msgServer{}
