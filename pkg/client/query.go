package client

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	rt "github.com/sonr-io/sonr/x/registry/types"
	"google.golang.org/grpc"
)

func CheckBalance(address string) (types.Coins, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		SONR_RPC_ADDR_PUBLIC, // Or your gRPC server address.
		grpc.WithInsecure(),  // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()
	if err != nil {
		return nil, err
	}
	resp, err := banktypes.NewQueryClient(grpcConn).AllBalances(context.Background(), &banktypes.QueryAllBalancesRequest{
		Address: address, // "snr155huqeqlm4lh5vvgychum0sa2xw70654m3kucq7vufjkx89hzvuqx0jmqc",
	})
	if err != nil {
		return nil, err
	}
	return resp.GetBalances(), nil
}

func QueryWhoIs(did string) (*rt.WhoIs, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		SONR_RPC_ADDR_PUBLIC, // Or your gRPC server address.
		grpc.WithInsecure(),  // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()
	if err != nil {
		return nil, err
	}
	res, err := rt.NewQueryClient(grpcConn).WhoIs(context.Background(), &rt.QueryWhoIsRequest{Did: did})
	if err != nil {
		return nil, err
	}
	return res.GetWhoIs(), nil
}

func QueryWhoIsByAlias(alias string) (*rt.WhoIs, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		SONR_RPC_ADDR_PUBLIC, // Or your gRPC server address.
		grpc.WithInsecure(),  // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()
	if err != nil {
		return nil, err
	}
	qc := rt.NewQueryClient(grpcConn)
	// We then call the QueryWhoIs method on this client.
	res, err := qc.WhoIsAlias(context.Background(), &rt.QueryWhoIsAliasRequest{Alias: alias})
	if err != nil {
		return nil, err
	}
	return res.GetWhoIs(), nil
}

func QueryWhoIsByController(controller string) (*rt.WhoIs, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		SONR_RPC_ADDR_PUBLIC, // Or your gRPC server address.
		grpc.WithInsecure(),  // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()
	if err != nil {
		return nil, err
	}
	qc := rt.NewQueryClient(grpcConn)
	// We then call the QueryWhoIs method on this client.
	res, err := qc.WhoIsController(context.Background(), &rt.QueryWhoIsControllerRequest{Controller: controller})
	if err != nil {
		return nil, err
	}
	return res.GetWhoIs(), nil
}
