package cli

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/spf13/cobra"
)

func CmdCreateBucket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-bucket",
		Short: "Create a new Bucket",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if len(args[0]) < 1 {
				return errors.New("label must be defined")
			}

			// parse strictly
			// visiblityConv, err := strconv.ParseUint(args[2], 10, 32)
			// visilbity := types.BucketConfigVisibility(visiblityConv)

			msg := types.NewMsgDefineBucket(clientCtx.GetFromAddress().String(), args[0])

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
