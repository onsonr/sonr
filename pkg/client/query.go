package client

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/query"
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

func (c *Client) QueryWhatIsByCreator(creator string, pagination *query.PageRequest) ([]*st.WhatIs, error) {
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
		Creator:    creator,
		Pagination: pagination,
	})

	if err != nil {
		return nil, err
	}

	return res.WhatIs, nil
}

func (c *Client) QueryWhatIsByDid(did string) (*st.WhatIs, error) {
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
	res, err := qc.WhatIsByDid(context.Background(), &st.QueryWhatIsByDidRequest{
		Did: did,
	})

	if err != nil {
		return nil, err
	}

	return res.WhatIs, nil
}

func (c *Client) QueryBucket(uuid string) (*bt.BucketConfig, error) {
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

	// We then call the QueryWhoIs method on this client.
	res, err := qc.Bucket(context.Background(), &bt.QueryGetBucketRequest{
		Uuid: uuid,
	})

	if err != nil {
		return nil, err
	}

	return &res.Bucket, nil
}


func (c *Client) QueryBucketsByCreator(creator string) ([]bt.BucketConfig, error) {
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

	// We then call the QueryWhoIs method on this client.
	res, err := qc.BucketByCreator(context.Background(), &bt.QueryGetBucketByCreatorRequest{
		Creator: creator,
	})

	if err != nil {
		return nil, err
	}

	return res.Buckets, nil
}


func (c *Client) QueryBucketsByName(name string) ([]bt.BucketConfig, error) {
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

	// We then call the QueryWhoIs method on this client.
	res, err := qc.BucketByName(context.Background(), &bt.QueryGetBucketByNameRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return res.Buckets, nil
}


func (c *Client) QueryAllBuckets() ([]bt.BucketConfig, error) {
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

	// We then call the QueryWhoIs method on this client.
	res, err := qc.BucketsAll(context.Background(), &bt.QueryAllBucketsRequest{

	})

	if err != nil {
		return nil, err
	}

	return res.Buckets, nil
}
