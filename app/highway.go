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
	"log"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/pkg/node/ipfs"
	"github.com/sonr-hq/sonr/pkg/vault"
)

// `HighwayNode` is a struct that contains a libp2p host, a wallet, a context, a client context, a
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
type HighwayNode struct {
	// Node is the libp2p host
	IPFS ipfs.IPFS

	// Properties
	ctx       context.Context
	clientCtx client.Context
	serveMux  *runtime.ServeMux
	vs        *vault.VaultService
	cache     *gocache.Cache
}

// It creates a new IPFS node and returns it
func NewHighwayNode() *HighwayNode {
	ctx := context.Background()
	node, err := ipfs.NewLocal(ctx)
	if err != nil {
		log.Println("Failed to create IPFS node:", err)
	}
	return &HighwayNode{
		ctx:   ctx,
		IPFS:  node,
		cache: gocache.New(time.Minute*2, time.Minute*10),
	}
}

// It's registering the gRPC gateway routes.
func (h *HighwayNode) RegisterGRPCGatewayRoutes(cctx client.Context, server *runtime.ServeMux) error {
	h.serveMux = server
	vs, err := vault.NewVaultService(h.ctx, server, h.IPFS, h.cache)
	if err != nil {
		return err
	}
	h.vs = vs
	return nil
}
