package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonr-io/sonr/x/registry/types"
	"github.com/spf13/cobra"
)

func CmdListWhoIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-who-is",
		Short: "list all whoIs",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllWhoIsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.WhoIsAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowWhoIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-who-is [did]",
		Short: "shows a whoIs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDid := args[0]

			params := &types.QueryWhoIsRequest{
				Did: argDid,
			}

			res, err := queryClient.WhoIs(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
