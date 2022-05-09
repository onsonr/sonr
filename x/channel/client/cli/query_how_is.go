package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonr-io/sonr/x/channel/types"
	"github.com/spf13/cobra"
)

func CmdListHowIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-how-is",
		Short: "list all howIs",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllHowIsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.HowIsAll(context.Background(), params)
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

func CmdShowHowIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-how-is [did]",
		Short: "shows a howIs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDid := args[0]

			params := &types.QueryHowIsRequest{
				Did: argDid,
			}

			res, err := queryClient.HowIs(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
