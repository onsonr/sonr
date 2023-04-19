package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonrhq/core/x/domain/types"
	"github.com/spf13/cobra"
)

func CmdListTLDRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-tld-record",
		Short: "list all TLDRecord",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTLDRecordRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TLDRecordAll(context.Background(), params)
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

func CmdShowTLDRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-tld-record [index]",
		Short: "shows a TLDRecord",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetTLDRecordRequest{
				Index: argIndex,
			}

			res, err := queryClient.TLDRecord(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
