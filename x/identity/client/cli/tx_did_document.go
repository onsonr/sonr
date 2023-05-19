package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/spf13/cobra"
)

func CmdCreateDidDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-did",
		Short: "Create a new did_document",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			jsonDid := args[0]
			types.NewSonrIdentity(jsonDid)
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDidDocument(
				clientCtx.FromAddress.String(),
				0,
				"",
				types.NewSonrIdentity(clientCtx.GetFromAddress().String()),
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

func CmdUpdateDidDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-did-document [did] [context] [controller] [verification-method] [authentication] [assertion-method] [capibility-invocation] [capability-delegation] [key-agreement] [service] [also-known-as]",
		Short: "Update a did_document",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDidDocument(
				clientCtx.GetFromAddress().String(),
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
