package client

import (
	"context"
	"errors"

	bt "github.com/sonr-io/sonr/x/bucket/types"
	rt "github.com/sonr-io/sonr/x/registry/types"
	st "github.com/sonr-io/sonr/x/schema/types"
	"google.golang.org/grpc"
)

func (c *Client) QueryWhoIs(did string) (*rt.WhoIs, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

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
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

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
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	qc := rt.NewQueryClient(grpcConn)
	// We then call the QueryWhoIs method on this client.
	res, err := qc.WhoIsController(context.Background(), &rt.QueryWhoIsControllerRequest{Controller: controller})
	if err != nil {
		return nil, err
	}
	return res.GetWhoIs(), nil
}

func (c *Client) QueryWhatIs(creator string, did string) (*st.WhatIs, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	qc := st.NewQueryClient(grpcConn)
	// We then call the QueryWhoIs method on this client.
	res, err := qc.WhatIs(context.Background(), &st.QueryWhatIsRequest{
		Creator: creator,
		Did:     did,
	})

	if err != nil {
		return nil, err
	}

	return res.WhatIs, nil
}

func (c *Client) QueryWhatIsByCreator(creator string) ([]*st.WhatIs, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	qc := st.NewQueryClient(grpcConn)

	// We then call the QueryWhoIs method on this client.
	res, err := qc.WhatIsByCreator(context.Background(), &st.QueryWhatIsCreatorRequest{
		Creator: creator,
	})

	if err != nil {
		return nil, err
	}

	return res.WhatIs, nil
}

func (c *Client) QueryWhereIs(did string, address string) (*bt.WhereIs, error) {
	if did == "" {
		return nil, errors.New("did invalid for Get WhereIs by Creator request")
	}
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	qc := bt.NewQueryClient(grpcConn)
	resp, err := qc.WhereIs(context.Background(), &bt.QueryGetWhereIsRequest{
		Creator: address,
		Did:     did,
	})

	if err != nil {
		return nil, err
	}

	return &resp.WhereIs, nil
}

func (c *Client) QueryWhereIsByCreator(address string) (*bt.QueryGetWhereIsByCreatorResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	qc := bt.NewQueryClient(grpcConn)
	res, err := qc.WhereIsByCreator(context.Background(), &bt.QueryGetWhereIsByCreatorRequest{
		Creator:    address,
		Pagination: nil,
	})

	if err != nil {
		return nil, err
	}
	return res, nil
}
