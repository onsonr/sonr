package server

import (
	context "context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kataras/golog"
	"github.com/sonr-io/core/channel"
	t "go.buf.build/grpc/go/sonr-io/core/types/v1"

	"github.com/sonr-io/core/config"
	hn "github.com/sonr-io/core/host"
	"github.com/sonr-io/core/host/discover"
	"github.com/sonr-io/core/host/exchange"
	v1 "go.buf.build/grpc/go/sonr-io/core/highway/v1"

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

// HighwayStub is the RPC Service for the Custodian Node.
type HighwayStub struct {
	v1.HighwayServiceServer
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

// startHighwayStub creates a new Highway service stub for the node.
func StartHighwayRPCServer(ctx context.Context, n hn.HostImpl, loc *t.Location, lst net.Listener) (*HighwayStub, error) {
	// create an instance of cosmosclient
	cosmos, err := cosmosclient.New(ctx)
	if err != nil {
		return nil, err
	}

	// Create the RPC Service
	stub := &HighwayStub{
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
func (s *HighwayStub) Serve(ctx context.Context, listener net.Listener) {
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
func (s *HighwayStub) AccessName(ctx context.Context, req *v1.AccessNameRequest) (*v1.AccessNameResponse, error) {
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
func (s *HighwayStub) RegisterName(ctx context.Context, req *v1.RegisterNameRequest) (*v1.RegisterNameResponse, error) {
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
func (s *HighwayStub) UpdateName(ctx context.Context, req *v1.UpdateNameRequest) (*v1.UpdateNameResponse, error) {
	return nil, ErrMethodUnimplemented
}

// AccessService accesses a service.
func (s *HighwayStub) AccessService(ctx context.Context, req *v1.AccessServiceRequest) (*v1.AccessServiceResponse, error) {
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
func (s *HighwayStub) RegisterService(ctx context.Context, req *v1.RegisterServiceRequest) (*v1.RegisterServiceResponse, error) {
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
func (s *HighwayStub) UpdateService(ctx context.Context, req *v1.UpdateServiceRequest) (*v1.UpdateServiceResponse, error) {
	return nil, ErrMethodUnimplemented
}

// CreateChannel creates a new channel.
func (s *HighwayStub) CreateChannel(ctx context.Context, req *v1.CreateChannelRequest) (*v1.CreateChannelResponse, error) {
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
func (s *HighwayStub) ReadChannel(ctx context.Context, req *v1.ReadChannelRequest) (*v1.ReadChannelResponse, error) {
	// Find channel by DID
	ch, ok := s.channels[req.GetDid()]
	if !ok {
		return nil, ErrInvalidQuery
	}

	// Read the channel
	peers := ch.Read()
	logger.Debugf("Read %d peers from channel %s", len(peers), peers)
	return &v1.ReadChannelResponse{
		// Peers: peers,
	}, nil
}

// UpdateChannel updates a channel.
func (s *HighwayStub) UpdateChannel(ctx context.Context, req *v1.UpdateChannelRequest) (*v1.UpdateChannelResponse, error) {
	return nil, ErrMethodUnimplemented
}

// DeleteChannel deletes a channel.
func (s *HighwayStub) DeleteChannel(ctx context.Context, req *v1.DeleteChannelRequest) (*v1.DeleteChannelResponse, error) {
	return nil, ErrMethodUnimplemented
}

// ListenChannel listens to a channel.
func (s *HighwayStub) ListenChannel(req *v1.ListenChannelRequest, stream v1.HighwayService_ListenChannelServer) error {
	// Find channel by DID
	ch, ok := s.channels[req.GetDid()]
	if !ok {
		return ErrInvalidQuery
	}

	// Listen to the channel
	chListen, err := ch.Listen()
	if err != nil {
		return err
	}

	// Listen to the channel
	for {
		select {
		case msg := <-chListen:
			// Send peer to client
			if err := stream.Send(&v1.ListenChannelResponse{
				Message: msg.GetData(),
			}); err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}

// CreateBucket creates a new bucket.
func (s *HighwayStub) CreateBucket(ctx context.Context, req *v1.CreateBucketRequest) (*v1.CreateBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}

// ReadBucket reads a bucket.
func (s *HighwayStub) ReadBucket(ctx context.Context, req *v1.ReadBucketRequest) (*v1.ReadBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}

// UpdateBucket updates a bucket.
func (s *HighwayStub) UpdateBucket(ctx context.Context, req *v1.UpdateBucketRequest) (*v1.UpdateBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}

// DeleteBucket deletes a bucket.
func (s *HighwayStub) DeleteBucket(ctx context.Context, req *v1.DeleteBucketRequest) (*v1.DeleteBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}

// ListenBucket listens to a bucket.
func (s *HighwayStub) ListenBucket(req *v1.ListenBucketRequest, stream v1.HighwayService_ListenBucketServer) error {
	return nil
}

// CreateObject creates a new object.
func (s *HighwayStub) CreateObject(ctx context.Context, req *v1.CreateObjectRequest) (*v1.CreateObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}

// ReadObject reads an object.
func (s *HighwayStub) ReadObject(ctx context.Context, req *v1.ReadObjectRequest) (*v1.ReadObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}

// UpdateObject updates an object.
func (s *HighwayStub) UpdateObject(ctx context.Context, req *v1.UpdateObjectRequest) (*v1.UpdateObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}

// DeleteObject deletes an object.
func (s *HighwayStub) DeleteObject(ctx context.Context, req *v1.DeleteObjectRequest) (*v1.DeleteObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}

// UploadBlob uploads a blob.
func (s *HighwayStub) UploadBlob(req *v1.UploadBlobRequest, stream v1.HighwayService_UploadBlobServer) error {
	// hash, err := s.ipfs.Upload(req.Path)
	// if err != nil {
	// 	return err
	// }
	logger.Debug("Uploaded blob to IPFS", "hash")
	return nil
}

// DownloadBlob downloads a blob.
func (s *HighwayStub) DownloadBlob(req *v1.DownloadBlobRequest, stream v1.HighwayService_DownloadBlobServer) error {
	// path, err := s.ipfs.Download(req.GetDid())
	// if err != nil {
	// 	return err
	// }
	logger.Debug("Downloaded blob from IPFS", "path")
	return nil
}

// SyncBlob synchronizes a blob with remote version.
func (s *HighwayStub) SyncBlob(req *v1.SyncBlobRequest, stream v1.HighwayService_SyncBlobServer) error {
	return nil
}

// DeleteBlob deletes a blob.
func (s *HighwayStub) DeleteBlob(ctx context.Context, req *v1.DeleteBlobRequest) (*v1.DeleteBlobResponse, error) {
	return nil, ErrMethodUnimplemented
}

// ParseDid parses a DID.
func (s *HighwayStub) ParseDid(ctx context.Context, req *v1.ParseDidRequest) (*v1.ParseDidResponse, error) {
	//d, err := s.node.ParseDid(req.GetDid())
	// if err != nil {
	// 	return nil, err
	// }
	return &v1.ParseDidResponse{
		Did: "",
	}, nil
}

// ResolveDid resolves a DID.
func (s *HighwayStub) ResolveDid(ctx context.Context, req *v1.ResolveDidRequest) (*v1.ResolveDidResponse, error) {
	return &v1.ResolveDidResponse{
		DidDocument: "",
	}, nil
}
