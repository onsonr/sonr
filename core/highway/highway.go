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

package highway

import (
	"context"
	"log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/ipfs"
)

type HighwayNode struct {
	// Node is the libp2p host
	Node   *ipfs.IPFS
	Wallet common.Wallet

	// Properties
	ctx       context.Context
	clientCtx client.Context
	serveMux  *runtime.ServeMux
	vs        *VaultService
}

func NewHighwayNode() *HighwayNode {
	ctx := context.Background()
	node, err := ipfs.New(ctx)
	if err != nil {
		log.Println("Failed to create IPFS node:", err)
	}
	return &HighwayNode{
		ctx:  ctx,
		Node: node,
	}
}

func (h *HighwayNode) RegisterGRPCGatewayRoutes(cctx client.Context, server *runtime.ServeMux) error {
	h.serveMux = server
	vs, err := NewVaultService(h.ctx, server)
	if err != nil {
		return err
	}
	h.vs = vs
	return nil
}
