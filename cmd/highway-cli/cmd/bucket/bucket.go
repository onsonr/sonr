package bucket

import (
	"context"
	"net/rpc"

	"github.com/spf13/cobra"
)

func BootstrapBucketCommand(ctx context.Context) (bucketCmd *cobra.Command) {

	bucketCmd = &cobra.Command{
		Use:   "bucket",
		Short: "Provides commands for managing buckets on the Sonr network",

		Run: func(cmd *cobra.Command, args []string) {},
	}

	bucketCmd.AddCommand(bootstrapCreateBucketCommand(ctx))

	return
}

func newRpcClient(address string) *rpc.Client {
	client, err := rpc.Dial("tcp", address)
	cobra.CheckErr(err)
	return client
}
