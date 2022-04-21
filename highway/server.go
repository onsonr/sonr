package highway

import (
	context "context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/duo-labs/webauthn.io/session"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/gorilla/mux"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/kataras/golog"
	"github.com/patrickmn/go-cache"
	"github.com/sonr-io/core/channel"
	"github.com/sonr-io/core/device"
	"github.com/sonr-io/core/highway/client"
	"github.com/sonr-io/core/highway/config"
	hn "github.com/sonr-io/core/host"
	"github.com/sonr-io/core/host/ipfs"
	v1 "go.buf.build/grpc/go/sonr-io/core/highway/v1"
	"google.golang.org/grpc"
)

// Error Definitions
var (
	logger                 = golog.Default.Child("node/highway")
	ErrEmptyQueue          = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery        = errors.New("No SName or PeerID provided.")
	ErrMissingParam        = errors.New("Paramater is missing.")
	ErrProtocolsNotSet     = errors.New("Node Protocol has not been initialized.")
	ErrMethodUnimplemented = errors.New("Method is not implemented.")
)

// HighwayServer is the RPC Service for the Custodian Node.
type HighwayServer struct {
	v1.HighwayServer
	node hn.HostImpl

	cosmos *client.Cosmos

	// Properties
	ctx          context.Context
	listener     net.Listener
	grpc         *grpc.Server
	router       *mux.Router
	ipfsProtocol *ipfs.IPFSProtocol

	// Configuration
	auth         *webauthn.WebAuthn
	cache        *cache.Cache
	sessionStore *session.Store

	ipfs iface.CoreAPI

	// List of Entries
	channels map[string]channel.Channel
}

// NewHighwayServer creates a new Highway service stub for the node.
func NewHighway(ctx context.Context, opts ...config.Option) (*HighwayServer, error) {
	c := config.DefaultConfig()
	for _, opt := range opts {
		opt(c)
	}

	node, err := hn.NewHost(ctx, device.Role_HIGHWAY, c)
	if err != nil {
		return nil, err
	}

	// Get the Listener for the Host
	lst, err := node.Listener()
	if err != nil {
		return nil, err
	}

	// Create a new Cosmos Client for Sonr Blockchain
	cosmos, err := client.NewCosmos(ctx, c)
	if err != nil {
		return nil, err
	}

	// Create a WebAuthn instance
	web, err := webauthn.New(c.WebauthnConfig())
	if err != nil {
		return nil, err
	}

	// Create a new Session Store
	sessionStore, err := session.NewStore()
	if err != nil {
		return nil, err
	}

	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	cche := cache.New(5*time.Minute, 10*time.Minute)

	// Create the RPC Service
	stub := &HighwayServer{
		cosmos: cosmos,
		node:   node,
		cache:  cche,
		ctx:    ctx,
		grpc:   grpc.NewServer(),

		listener:     lst,
		auth:         web,
		sessionStore: sessionStore,
	}

	stub.ipfsProtocol, err = ipfs.New(ctx, stub.node)
	if err != nil {
		return nil, err
	}

	// TODO Implement P2P Protocols for Sonr Network
	// Set Discovery Protocol
	// stub.DiscoverProtocol, err = discover.New(ctx, node, stub)
	// if err != nil {
	// 	logger.Errorf("%s - Failed to start DiscoveryProtocol", err)
	// 	return nil, err
	// }

	// Set Transmit Protocol
	// stub.ExchangeProtocol, err = exchange.New(ctx, node, stub)
	// if err != nil {
	// 	logger.Errorf("%s - Failed to start TransmitProtocol", err)
	// 	return nil, err
	// }

	// Register RPC Service
	v1.RegisterHighwayServer(stub.grpc, stub)
	return stub, nil
}

// Serve starts the RPC Service.
func (s *HighwayServer) Serve() {
	logger.Infof("Starting RPC Server on %s", s.listener.Addr().String())
	go s.serveCtxListener(s.ctx, s.listener)
}

// Serve serves the RPC Service on the given port.
func (s *HighwayServer) serveCtxListener(ctx context.Context, listener net.Listener) {
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	if err := s.grpc.Serve(listener); err != nil {
		logger.Errorf("%s - Failed to start HTTP server", err)
	}
	s.node.Persist()
}

// from: https://github.com/duo-labs/webauthn.io/blob/3f03b482d21476f6b9fb82b2bf1458ff61a61d41/server/response.go#L15
func JsonResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}
