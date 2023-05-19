package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonrhq/core/x/service/types"
	"github.com/spf13/cobra"
)

func CmdListServiceRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-service-record",
		Short: "list all ServiceRecord",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.ListServiceRecordsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ListServiceRecords(context.Background(), params)
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

func CmdShowServiceRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-service-record [Id]",
		Short: "shows a ServiceRecord",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argId := args[0]

			params := &types.QueryServiceRecordRequest{
				Origin: argId,
			}

			res, err := queryClient.ServiceRecord(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
