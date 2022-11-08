package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	//"github.com/mr-tron/base58/base58"
	"github.com/sonr-io/sonr/pkg/did"

	"github.com/sonr-io/sonr/x/registry/types"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

const (
	FlagPubKey = "pubkey"
)

func FlagSetPublicKey() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagPubKey, "", "The creator's public key")
	return fs
}

func CmdCreateWhoIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-who-is [didJson]",
		Short: "Create a new WhoIs",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			fs := cmd.Flags()
			pubkeyStr, err := fs.GetString(FlagPubKey)
			if err != nil {
				return err
			}

			pub := &secp256k1.PubKey{
				Key: []byte(pubkeyStr),
			}
			whoIsType := types.WhoIsType(1)
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			rawCreator := clientCtx.GetFromAddress().String()
			didUrl, err := did.ParseDID(fmt.Sprintf("did:snr:%s", rawCreator))
			if err != nil {
				return err
			}
			doc, err := did.NewDocument(didUrl.String())

			if err != nil {
				return err
			}

			docBz, err := doc.MarshalJSON()
			if err != nil {
				return err
			}
			msg := types.NewMsgCreateWhoIs(clientCtx.GetFromAddress().String(), pub, docBz, whoIsType)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FlagSetPublicKey())

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdUpdateWhoIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-who-is [did] [didJson]",
		Short: "Update a WhoIs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateWhoIs(clientCtx.GetFromAddress().String(), []byte(args[0]))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdDeactivateWhoIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-who-is [did]",
		Short: "Delete a WhoIs by Did",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeactivateWhoIs(clientCtx.GetFromAddress().String(), args[0])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
