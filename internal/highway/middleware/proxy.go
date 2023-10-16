package middleware

import (
	"context"
	"fmt"

	domaintypes "github.com/sonr-io/core/x/domain/types"
	identitytypes "github.com/sonr-io/core/x/identity/types"
	servicetypes "github.com/sonr-io/core/x/service/types"
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

// GetEmailRecordCreator returns the DIDDocument of a given DID or Alias
func GetEmailRecordCreator(ctx context.Context, index string) (string, error) {
	addr := fmt.Sprintf("%s:%d", viper.GetString("node.grpc.host"), viper.GetInt("node.grpc.port"))
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return "", err
	}
	defer grpcConn.Close()
	qc := domaintypes.NewQueryClient(grpcConn)
	resp, err := qc.UsernameRecord(ctx, &domaintypes.QueryGetUsernameRecordsRequest{Index: index})
	if err != nil {
		return "", err
	}
	return resp.UsernameRecords.Address, nil
}

// CheckAliasAvailable returns the DIDDocument of a given DID or Alias
func CheckAliasAvailable(ctx context.Context, alias string) (bool, error) {
	addr, err := GetEmailRecordCreator(ctx, alias)
	if err != nil {
		return false, err
	}
	return addr == "", nil
}

// CheckAliasUnavailable returns the DIDDocument of a given DID or Alias
func CheckAliasUnavailable(ctx context.Context, alias string) (bool, error) {
	addr, err := GetEmailRecordCreator(ctx, alias)
	if err != nil {
		return false, err
	}
	return addr != "", nil
}

// GetServiceRecord returns the DIDDocument of a given DID or Alias
func GetServiceRecord(ctx context.Context, origin string) (*servicetypes.ServiceRecord, error) {
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
	qc := servicetypes.NewQueryClient(grpcConn)
	resp, err := qc.ServiceRecord(ctx, &servicetypes.QueryGetServiceRecordRequest{Origin: origin})
	if err != nil {
		return nil, err
	}
	return &resp.ServiceRecord, nil
}
