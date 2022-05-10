package bucket

import (
	"context"
	"fmt"

	"github.com/sonr-io/sonr/cmd/highway-cli/prompt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	bt "go.buf.build/grpc/go/sonr-io/blockchain/bucket"
)

/**
 * WIP
 */

func bootstrapCreateBucketCommand(ctx context.Context) (createBucketCmd *cobra.Command) {

	createBucketCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a bucket with the given label",

		Run: func(cmd *cobra.Command, args []string) {
			if l := len(args); l < 1 {
				cobra.CheckErr(fmt.Errorf("expected 1 argument but found %d", l))
			}

			var (
				label       = prompt.ForString("Enter bucket label: ", prompt.NonEmpty)
				creator     = prompt.ForString("Enter creator: ", prompt.NonEmpty)
				description = prompt.ForString("Enter description: ", prompt.NonEmpty)
				kind        = prompt.ForString("Enter bucket kind: ", prompt.NonEmpty)
			)

			fmt.Printf("%s<>%s<>%s<>%s\n", label, creator, description, kind)

			rpcClient := newRpcClient(fmt.Sprintf("127.0.0.1:%d", viper.GetInt("PORT")))
			var reply interface{} // bt.MsgCreateBucketResponse
			err := rpcClient.Call(args[0], &bt.MsgCreateBucket{
				Label:       label,
				Creator:     creator,
				Description: description,
				Kind:        kind,
			}, &reply)
			cobra.CheckErr(err)

			fmt.Println("request finished")
			fmt.Printf("%v\n", reply)
			// fmt.Printf("code: %d\n", reply.GetCode())
			// fmt.Printf(" did: %s\n", reply.GetWhichIs().GetDid())
		},
	}

	return
}
