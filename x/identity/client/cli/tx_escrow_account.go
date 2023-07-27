package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/spf13/cobra"
)

func CmdCreateEscrowAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-escrow-account [address] [public-key] [lockup-usd-balance]",
		Short: "Create a new escrowAccount",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argPublicKey := args[1]
			argLockupUsdBalance := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateEscrowAccount(clientCtx.GetFromAddress().String(), argAddress, argPublicKey, argLockupUsdBalance)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateEscrowAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-escrow-account [id] [address] [public-key] [lockup-usd-balance]",
		Short: "Update a escrowAccount",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argAddress := args[1]

			argPublicKey := args[2]

			argLockupUsdBalance := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateEscrowAccount(clientCtx.GetFromAddress().String(), id, argAddress, argPublicKey, argLockupUsdBalance)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteEscrowAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-escrow-account [id]",
		Short: "Delete a escrowAccount by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
	addr := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteEscrowAccount(clientCtx.GetFromAddress().String(), addr)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
