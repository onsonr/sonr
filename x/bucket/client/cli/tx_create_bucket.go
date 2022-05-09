package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateBucket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-bucket [label] [description] [kind]",
		Short: "Broadcast message create-bucket",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argLabel := args[0]
			argDescription := args[1]
			argKind := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBucket(
				clientCtx.GetFromAddress().String(),
				argLabel,
				argDescription,
				argKind,
				nil,
				nil,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
