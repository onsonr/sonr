package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sonr-io/sonr/x/channel/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateChannel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-channel [name] [description] [owners]",
		Short: "Broadcast message create-channel",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDescription := args[1]

			argTTL, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return err
			}

			argMaxSize, err := strconv.ParseInt(args[4], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateChannel(
				clientCtx.GetFromAddress().String(),
				argName,
				argDescription,
				nil,
				argTTL,
				argMaxSize,
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
