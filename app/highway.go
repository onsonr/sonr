// ---
// Highway Implementation
//
// Highway is a service node that provides a high bandwidth, low latency, and enhanced experience for the Sonr network.
// Highway nodes are incentivized to provide a high quality of service to the network, and are rewarded for doing so.
//
// Modules: Vault, Comm, DID Utility
// Interface: IPFS host
// Transports: TCP, UDP, QUIC, HTTP, WebTransport, WebRTC, WebSockets
// ---

package app

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/pkg/node"
	"github.com/sonr-hq/sonr/pkg/node/config"
	"github.com/sonr-hq/sonr/pkg/vault"
)

// `HighwayProtocol` is a struct that contains a libp2p host, a wallet, a context, a client context, a
// serve mux, a vault service, and an IPFS service.
// @property Node - The libp2p host
// @property Wallet - The wallet is the account that the node will use to sign transactions.
// @property ctx - The context of the node.
// @property clientCtx - This is the context for the client. It is used to create a client for the
// vault service.
// @property serveMux - This is the gRPC server mux. It's used to register the gRPC services.
// @property vs - The VaultService is the service that manages the vault. It is responsible for
// creating the vault, adding and removing members, and managing the vault's state.
// @property ipfs - The IPFS node
type HighwayProtocol struct {
	// Node is the libp2p host
	node config.IPFSNode

	// Initialization properties
	ctx   context.Context
	vs    *vault.VaultService
	cache *gocache.Cache
}

// It creates a new IPFS node and returns it
func StartHighwayProtocol() *HighwayProtocol {
	ctx := context.Background()
	return &HighwayProtocol{
		ctx:   ctx,
		cache: gocache.New(time.Minute*2, time.Minute*10),
	}
}

// It's registering the gRPC gateway routes.
func (h *HighwayProtocol) RegisterGRPCGatewayRoutes(cctx client.Context, server *runtime.ServeMux) error {
	node, err := node.NewIPFS(h.ctx, config.WithClientContext(cctx, true))
	if err != nil {
		return err
	}
	h.node = node
	vs, err := vault.NewService(cctx, server, h.node, h.cache)
	if err != nil {
		return err
	}
	h.vs = vs
	return nil
}
