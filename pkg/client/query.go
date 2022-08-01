package client

import (
	"context"

	rt "github.com/sonr-io/sonr/x/registry/types"
	st "github.com/sonr-io/sonr/x/schema/types"
	"google.golang.org/grpc"
)

// func (c *Client) QueryAccount(address string) (*at.BaseAccount, error) {
// 	// Create a connection to the gRPC server.
// 	grpcConn, err := grpc.Dial(
// 		c.GetRPCAddress(),   // Or your gRPC server address.
// 		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
// 	)
// 	defer grpcConn.Close()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// We then call the QueryAccount method on this client.
// 	res, err := at.NewQueryClient(grpcConn).Account(context.Background(), &at.QueryAccountRequest{Address: address})
// 	if err != nil {
// 		return nil, err
// 	}

// 	acc := &at.BaseAccount{}
// 	err = acc.Unmarshal(res.GetAccount().GetValue())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return acc, nil
// }

func (c *Client) QueryWhoIs(did string) (*rt.WhoIs, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()
	if err != nil {
		return nil, err
	}

	// Create a new request.
	req := &rt.QueryWhoIsRequest{Did: did}
	// We then call the QueryWhoIs method on this client.
	res, err := rt.NewQueryClient(grpcConn).WhoIs(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.GetWhoIs(), nil
}

func (c *Client) QueryWhoIsByAlias(alias string) (*rt.WhoIs, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
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

func (c *Client) QueryWhoIsByController(controller string) (*rt.WhoIs, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
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

func (c *Client) QueryWhatIsByController(creator string, did string) (*st.WhatIs, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()
	if err != nil {
		return nil, err
	}
	qc := st.NewQueryClient(grpcConn)

	// We then call the QueryWhoIs method on this client.
	res, err := qc.WhatIs(context.Background(), &st.QueryWhatIsRequest{
		Creator: creator,
		Did:     creator,
	})

	if err != nil {
		return nil, err
	}

	return res.WhatIs, nil
}
