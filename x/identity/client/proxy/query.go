package identityproxy

import (
	"context"
	"fmt"

	identitytypes "github.com/sonrhq/core/x/identity/types"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// GetControllerAccount returns the DIDDocument of a given DID or Alias
func GetControllerAccount(ctx context.Context, address string) (*identitytypes.ControllerAccount, error) {
	addr := fmt.Sprintf("%s:%d", viper.GetString("node.grpc.host"), viper.GetInt("node.grpc.port"))
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()
	qc := identitytypes.NewQueryClient(grpcConn)
	resp, err := qc.ControllerAccount(ctx, &identitytypes.QueryGetControllerAccountRequest{Address: address})
	if err != nil {
		return nil, err
	}
	return &resp.ControllerAccount, nil
}

// GetDIDByAlias returns the DIDDocument of a given DID or Alias
func GetDIDByAlias(ctx context.Context, alias string) (*identitytypes.DIDDocument, error) {
	addr := fmt.Sprintf("%s:%d", viper.GetString("node.grpc.host"), viper.GetInt("node.grpc.port"))
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()
	qc := identitytypes.NewQueryClient(grpcConn)
	resp, err := qc.DidByAlsoKnownAs(ctx, &identitytypes.QueryDidByAlsoKnownAsRequest{Alias: alias})
	if err != nil {
		return nil, err
	}
	return &resp.DidDocument, nil
}
