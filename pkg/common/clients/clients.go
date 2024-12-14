package clients

import (
	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	nodev1beta1 "github.com/cosmos/cosmos-sdk/client/grpc/node"
	didv1 "github.com/onsonr/sonr/api/did/v1"
	dwnv1 "github.com/onsonr/sonr/api/dwn/v1"
	svcv1 "github.com/onsonr/sonr/api/svc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func conn(addr string) (*grpc.ClientConn, error) {
	grpcConn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return grpcConn, nil
}

func BankQuery(addr string) (bankv1beta1.QueryClient, error) {
	conn, err := conn(addr)
	if err != nil {
		return nil, err
	}
	return bankv1beta1.NewQueryClient(conn), nil
}

func DIDQuery(addr string) (didv1.QueryClient, error) {
	conn, err := conn(addr)
	if err != nil {
		return nil, err
	}
	return didv1.NewQueryClient(conn), nil
}

func DWNQuery(addr string) (dwnv1.QueryClient, error) {
	conn, err := conn(addr)
	if err != nil {
		return nil, err
	}
	return dwnv1.NewQueryClient(conn), nil
}

func NodeQuery(addr string) (nodev1beta1.ServiceClient, error) {
	conn, err := conn(addr)
	if err != nil {
		return nil, err
	}
	return nodev1beta1.NewServiceClient(conn), nil
}

func SVCQuery(addr string) (svcv1.QueryClient, error) {
	conn, err := conn(addr)
	if err != nil {
		return nil, err
	}
	return svcv1.NewQueryClient(conn), nil
}
