package serviceproxy

import (
	"context"
	"fmt"

	servicetypes "github.com/sonrhq/core/x/service/types"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

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
