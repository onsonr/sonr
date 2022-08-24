package bucket

import (
	"context"
	"errors"

	bt "github.com/sonr-io/sonr/x/bucket/types"
	"github.com/sonr-io/sonr/x/bucket/types"
	"google.golang.org/grpc"
)

func (b *bucketImpl) ResolveBuckets(address string) error {
	// Create a connection to the gRPC server.
	conn, err := grpc.Dial(
		b.rpcEndpoint,         // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer conn.Close()
	if err != nil {
		return err
	}
	client := bt.NewQueryClient(conn)
	if b.whereIs == nil {
		return errors.New("top level bucket not provided")
	}

	for _, bi := range b.whereIs.Content {
		if bi.Type == types.ResourceIdentifier_DID {
			resp, err := client.WhereIs(context.Background(), &types.QueryGetWhereIsRequest{
				Creator: address,
				Did:     bi.Uri,
			})

			if err != nil {
				return err
			}
			b.contentCache[bi.Uri] = &BucketContent{
				Item:        New(b.address, &resp.WhereIs, b.shell, b.rpcEndpoint),
				Id:          bi.Uri,
				ContentType: types.ResourceIdentifier_DID,
			}
		}
	}

	return nil
}

func (b *bucketImpl) ResolveContent() error {
	if b.whereIs == nil {
		return errors.New("top level bucket not provided")
	}

	for _, bi := range b.whereIs.Content {
		if bi.Type == types.ResourceIdentifier_CID {
			var dag map[string]interface{}
			err := b.shell.DagGet(bi.Uri, &dag)

			if err != nil {
				return err
			}
			b.contentCache[bi.Uri] = &BucketContent{
				Item:        dag,
				Id:          bi.Uri,
				ContentType: types.ResourceIdentifier_DID,
			}
		}
	}

	return nil
}
