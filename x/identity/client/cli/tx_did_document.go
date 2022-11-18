package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sonr-io/sonr/x/identity/types"
	"github.com/spf13/cobra"
)

func CmdCreateDidDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-did",
		Short: "Create a new did_document",
		Args:  cobra.ExactArgs(11),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			jsonDid := args[0]
			types.NewDocumentFromJson([]byte(jsonDid))
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDidDocument(
				clientCtx.GetFromAddress().String(),
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
		Args:  cobra.ExactArgs(11),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexDid := args[0]

			// Get value arguments
			argContext := args[1]
			argController := args[2]
			argVerificationMethod := args[3]
			argAuthentication := args[4]
			argAssertionMethod := args[5]
			argCapibilityInvocation := args[6]
			argCapabilityDelegation := args[7]
			argKeyAgreement := args[8]
			argService := args[9]
			argAlsoKnownAs := args[10]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDidDocument(
				clientCtx.GetFromAddress().String(),
				indexDid,
				argContext,
				argController,
				argVerificationMethod,
				argAuthentication,
				argAssertionMethod,
				argCapibilityInvocation,
				argCapabilityDelegation,
				argKeyAgreement,
				argService,
				argAlsoKnownAs,
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
