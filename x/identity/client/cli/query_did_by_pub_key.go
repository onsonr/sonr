package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDidByPubKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "did-by-pub-key",
		Short: "Query DidByPubKey",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDidByPubKeyRequest{}

			res, err := queryClient.DidByPubKey(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
