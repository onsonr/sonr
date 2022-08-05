package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/spf13/cobra"
)

func CmdListWhereIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-where-is",
		Short: "list all WhereIs",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllWhereIsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.WhereIsAll(context.Background(), params)
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

func CmdShowWhereIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-where-is [id]",
		Short: "shows a WhereIs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			id := args[0]

			params := &types.QueryGetWhereIsRequest{
				Did: id,
			}

			res, err := queryClient.WhereIs(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
