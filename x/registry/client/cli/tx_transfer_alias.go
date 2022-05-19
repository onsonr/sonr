package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sonr-io/sonr/x/registry/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdTransferAlias() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-name-alias [did] [alias] [recipient]",
		Short: "Broadcast message TransferNameAlias",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDid := args[0]
			argAlias := args[1]
			argRecipient := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferAlias(
				clientCtx.GetFromAddress().String(),
				argDid,
				argAlias,
				argRecipient,
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
