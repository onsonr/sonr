package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonr-io/sonr/internal/blockchain/x/object/types"
	"github.com/spf13/cobra"
)

func CmdListWhatIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-what-is",
		Short: "list all whatIs",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllWhatIsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.WhatIsAll(context.Background(), params)
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

func CmdShowWhatIs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-what-is [did]",
		Short: "shows a whatIs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDid := args[0]

			params := &types.QueryWhatIsRequest{
				Did: argDid,
			}

			res, err := queryClient.WhatIs(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
