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
			types.NewBlankDocument(jsonDid)
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDidDocument(
				clientCtx.FromAddress.String(),
				0,
				"",
				types.NewBlankDocument(clientCtx.GetFromAddress().String()),
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

func CmdDeleteDidDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-did-document [did]",
		Short: "Delete a did_document",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexDid := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteDidDocument(
				clientCtx.GetFromAddress().String(),
				indexDid,
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
