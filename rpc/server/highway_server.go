package server

import (
	context "context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kataras/golog"
	bt "github.com/sonr-io/blockchain/x/bucket/types"
	ct "github.com/sonr-io/blockchain/x/channel/types"
	ot "github.com/sonr-io/blockchain/x/object/types"
	rt "github.com/sonr-io/blockchain/x/registry/types"
	"github.com/sonr-io/core/channel"
	"github.com/sonr-io/core/config"
	hn "github.com/sonr-io/core/node"
	"github.com/sonr-io/core/node/discover"
	"github.com/sonr-io/core/node/exchange"
	v1 "go.buf.build/grpc/go/sonr-io/core/highway/v1"
	t "go.buf.build/grpc/go/sonr-io/core/types/v1"

	"github.com/tendermint/starport/starport/pkg/cosmosclient"
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
	config.CallbackImpl
	node   hn.HostImpl
	cosmos cosmosclient.Client

	// Properties
	ctx context.Context
	mux *runtime.ServeMux
	*discover.DiscoverProtocol
	*exchange.ExchangeProtocol

	// Configuration
	// ipfs *storage.IPFSService

	// List of Entries
	channels map[string]channel.Channel
}

// NewHighwayServer creates a new Highway service stub for the node.
func NewHighwayServer(ctx context.Context, n hn.HostImpl, loc *t.Location, lst net.Listener) (*HighwayServer, error) {
	// create an instance of cosmosclient
	cosmos, err := cosmosclient.New(ctx)
	if err != nil {
		return nil, err
	}

	// Create the RPC Service
	stub := &HighwayServer{
		node:   n,
		ctx:    ctx,
		mux:    runtime.NewServeMux(),
		cosmos: cosmos,
	}

	// // Set IPFS Service
	// stub.ipfs, err = storage.Init()
	// if err != nil {
	// 	return nil, err
	// }

	// Set Discovery Protocol
	stub.DiscoverProtocol, err = discover.New(ctx, n, stub, discover.WithLocation(loc))
	if err != nil {
		logger.Errorf("%s - Failed to start DiscoveryProtocol", err)
		return nil, err
	}

	// Set Transmit Protocol
	stub.ExchangeProtocol, err = exchange.New(ctx, n, stub)
	if err != nil {
		logger.Errorf("%s - Failed to start TransmitProtocol", err)
		return nil, err
	}

	// Register the RPC Service
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// if err = highway.RegisterHighwayServiceServer(stub.mux, "localhost:8080", opts); err != nil {
	// 	logger.Errorf("%s - Failed to register RPC Service", err)
	// 	return nil, err
	// }
	go stub.Serve(ctx, lst)
	return stub, nil
}

// Serve serves the RPC Service on the given port.
func (s *HighwayServer) Serve(ctx context.Context, listener net.Listener) {
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	if err := http.ListenAndServe(":8081", s.mux); err != nil {
		logger.Errorf("%s - Failed to start HTTP server", err)
	}

	for {
		// Stop Serving if context is done
		select {
		case <-ctx.Done():
			s.DiscoverProtocol.Close()
			return
		}
	}
}

// AccessName accesses a name.
func (s *HighwayServer) AccessName(ctx context.Context, req *v1.MsgAccessName) (*v1.MsgAccessNameResponse, error) {
	// // instantiate a query client for your `blog` blockchain
	// queryClient := registry.NewQueryClient(s.cosmos.Context)

	// // query the blockchain using the client's `didAll` method to get all dids
	// // store all dids in queryResp
	// queryResp, err := queryClient.Dids(context.Background(), &types.QueryDidsRequest{})
	// if err != nil {
	// 	return nil, err
	// }

	// print response from querying all the dids
	fmt.Print("\n\nAll Dids:\n\n")
	// fmt.Println(queryResp)
	return nil, ErrMethodUnimplemented
}

// RegisterName registers a name.
func (s *HighwayServer) RegisterName(ctx context.Context, req *rt.MsgRegisterName) (*rt.MsgRegisterNameResponse, error) {
	// // account `alice` was initialized during `starport chain serve`
	// accountName := "alice"

	// // get account from the keyring by account name and return a bech32 address
	// address, err := s.cosmos.Address(accountName)
	// if err != nil {
	// 	return nil, err
	// }

	// // define a message to create a did
	// msg := &types.MsgCreateDidDocument{
	// 	Creator: address.String(),
	// }

	// // broadcast a transaction from account `alice` with the message to create a did
	// // store response in txResp
	// txResp, err := s.cosmos.BroadcastTx(accountName, msg)
	// if err != nil {
	// 	return nil, err
	// }

	// print response from broadcasting a transaction
	fmt.Print("MsgCreateDidDocument:\n\n")
	// fmt.Println(txResp)
	return nil, ErrMethodUnimplemented
}

// UpdateName updates a name.
func (s *HighwayServer) UpdateName(ctx context.Context, req *rt.MsgUpdateName) (*rt.MsgUpdateNameResponse, error) {
	return nil, ErrMethodUnimplemented
}

// AccessService accesses a service.
func (s *HighwayServer) AccessService(ctx context.Context, req *rt.MsgAccessService) (*rt.MsgAccessServiceResponse, error) {
	// // instantiate a query client for your `blog` blockchain
	// queryClient := types.NewQueryClient(s.cosmos.Context)

	// // query the blockchain using the client's `didAll` method to get all dids
	// // store all dids in queryResp
	// queryResp, err := queryClient.Dids(context.Background(), &types.QueryDidsRequest{})
	// if err != nil {
	// 	return nil, err
	// }

	// print response from querying all the dids
	fmt.Print("\n\nAll Dids:\n\n")
	// fmt.Println(queryResp)
	return nil, ErrMethodUnimplemented
}

// RegisterService registers a service.
func (s *HighwayServer) RegisterService(ctx context.Context, req *rt.MsgRegisterService) (*rt.MsgRegisterServiceResponse, error) {
	// // account `alice` was initialized during `starport chain serve`
	// accountName := "alice"

	// // get account from the keyring by account name and return a bech32 address
	// address, err := s.cosmos.Address(accountName)
	// if err != nil {
	// 	return nil, err
	// }

	// // define a message to create a did
	// msg := &types.MsgCreateDidDocument{
	// 	Creator: address.String(),
	// }

	// // broadcast a transaction from account `alice` with the message to create a did
	// // store response in txResp
	// txResp, err := s.cosmos.BroadcastTx(accountName, msg)
	// if err != nil {
	// 	return nil, err
	// }

	// print response from broadcasting a transaction
	fmt.Print("MsgCreateDidDocument:\n\n")
	// fmt.Println(txResp)
	return nil, ErrMethodUnimplemented
}

// UpdateService updates a service.
func (s *HighwayServer) UpdateService(ctx context.Context, req *rt.MsgUpdateService) (*rt.MsgUpdateServiceResponse, error) {
	return nil, ErrMethodUnimplemented
}

// CreateChannel creates a new channel.
func (s *HighwayServer) CreateChannel(ctx context.Context, req *ct.MsgCreateChannel) (*ct.MsgCreateChannelResponse, error) {
	// Create the Channel
	ch, err := channel.New(ctx, s.node, req.Name)
	if err != nil {
		return nil, err
	}

	// Add to the list of Channels
	s.channels[req.Name] = ch
	return nil, ErrMethodUnimplemented
}

// ReadChannel reads a channel.
func (s *HighwayServer) ReadChannel(ctx context.Context, req *ct.MsgReadChannel) (*ct.MsgReadChannelResponse, error) {
	// Find channel by DID
	ch, ok := s.channels[req.GetDid()]
	if !ok {
		return nil, ErrInvalidQuery
	}

	// Read the channel
	peers := ch.Read()
	logger.Debugf("Read %d peers from channel %s", len(peers), peers)
	return &ct.MsgReadChannelResponse{
		// Peers: peers,
	}, nil
}

// UpdateChannel updates a channel.
func (s *HighwayServer) UpdateChannel(ctx context.Context, req *ct.MsgUpdateChannel) (*ct.MsgUpdateChannelResponse, error) {
	return nil, ErrMethodUnimplemented
}

// DeleteChannel deletes a channel.
func (s *HighwayServer) DeleteChannel(ctx context.Context, req *ct.MsgDeleteChannel) (*ct.MsgDeleteChannelResponse, error) {
	return nil, ErrMethodUnimplemented
}

// // ListenChannel listens to a channel.
// func (s *HighwayServer) ListenChannel(req *ct.ListenChannelRequest, stream v1.HighwayService_ListenChannelServer) error {
// 	// Find channel by DID
// 	ch, ok := s.channels[req.GetDid()]
// 	if !ok {
// 		return ErrInvalidQuery
// 	}

// 	// Listen to the channel
// 	chListen, err := ch.Listen()
// 	if err != nil {
// 		return err
// 	}

// 	// Listen to the channel
// 	for {
// 		select {
// 		case msg := <-chListen:
// 			// Send peer to client
// 			if err := stream.Send(&v1.ListenChannelResponse{
// 				Message: msg.GetData(),
// 			}); err != nil {
// 				return err
// 			}
// 		case <-stream.Context().Done():
// 			return nil
// 		}
// 	}
// }

// CreateBucket creates a new bucket.
func (s *HighwayServer) CreateBucket(ctx context.Context, req *bt.MsgCreateBucket) (*bt.MsgCreateBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}

// ReadBucket reads a bucket.
func (s *HighwayServer) ReadBucket(ctx context.Context, req *bt.MsgReadBucket) (*bt.MsgReadBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}

// UpdateBucket updates a bucket.
func (s *HighwayServer) UpdateBucket(ctx context.Context, req *bt.MsgUpdateBucket) (*bt.MsgUpdateBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}

// DeleteBucket deletes a bucket.
func (s *HighwayServer) DeleteBucket(ctx context.Context, req *bt.MsgDeleteBucket) (*bt.MsgDeleteBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}

// CreateObject creates a new object.
func (s *HighwayServer) CreateObject(ctx context.Context, req *ot.MsgCreateObject) (*ot.MsgCreateObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}

// ReadObject reads an object.
func (s *HighwayServer) ReadObject(ctx context.Context, req *ot.MsgReadObject) (*ot.MsgReadObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}

// UpdateObject updates an object.
func (s *HighwayServer) UpdateObject(ctx context.Context, req *ot.MsgUpdateObject) (*ot.MsgUpdateObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}

// DeleteObject deletes an object.
func (s *HighwayServer) DeleteObject(ctx context.Context, req *ot.MsgDeleteObject) (*ot.MsgDeleteObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}

// // UploadBlob uploads a blob.
// func (s *HighwayServer) UploadBlob(req *v1.UploadBlobRequest, stream v1.HighwayService_UploadBlobServer) error {
// 	// hash, err := s.ipfs.Upload(req.Path)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	logger.Debug("Uploaded blob to IPFS", "hash")
// 	return nil
// }

// // DownloadBlob downloads a blob.
// func (s *HighwayServer) DownloadBlob(req *v1.DownloadBlobRequest, stream v1.HighwayService_DownloadBlobServer) error {
// 	// path, err := s.ipfs.Download(req.GetDid())
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	logger.Debug("Downloaded blob from IPFS", "path")
// 	return nil
// }

// // SyncBlob synchronizes a blob with remote version.
// func (s *HighwayServer) SyncBlob(req *v1.SyncBlobRequest, stream v1.HighwayService_SyncBlobServer) error {
// 	return nil
// }

// // DeleteBlob deletes a blob.
// func (s *HighwayServer) DeleteBlob(ctx context.Context, req *v1.DeleteBlobRequest) (*v1.DeleteBlobResponse, error) {
// 	return nil, ErrMethodUnimplemented
// }

// // ParseDid parses a DID.
// func (s *HighwayServer) ParseDid(ctx context.Context, req *v1.ParseDidRequest) (*v1.ParseDidResponse, error) {
// 	//d, err := s.node.ParseDid(req.GetDid())
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	return &v1.ParseDidResponse{
// 		Did: "",
// 	}, nil
// }

// // ResolveDid resolves a DID.
// func (s *HighwayServer) ResolveDid(ctx context.Context, req *v1.ResolveDidRequest) (*v1.ResolveDidResponse, error) {
// 	return &v1.ResolveDidResponse{
// 		DidDocument: "",
// 	}, nil
// }
