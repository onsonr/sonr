package domainproxy

import (
	"context"
	"fmt"

	domaintypes "github.com/sonr-io/core/x/domain/types"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

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
