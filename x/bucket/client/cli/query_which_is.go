package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/spf13/cobra"
)

func CmdListWhichIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-which-is",
		Short: "list all whichIs",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllWhichIsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.WhichIsAll(context.Background(), params)
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

func CmdShowWhichIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-which-is [did]",
		Short: "shows a whichIs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDid := args[0]

			params := &types.QueryWhichIsRequest{
				Did: argDid,
			}

			res, err := queryClient.WhichIs(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
