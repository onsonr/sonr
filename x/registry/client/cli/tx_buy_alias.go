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

func CmdBuyAlias() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-app-alias [did] [name] ",
		Short: "Broadcast message BuyAppAlias",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDid := args[0]
			argName := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyAlias(
				clientCtx.GetFromAddress().String(),
				argDid,
				argName,
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
