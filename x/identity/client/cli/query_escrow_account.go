package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonr-io/core/x/identity/types"
	"github.com/spf13/cobra"
)

func CmdListEscrowAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-escrow-account",
		Short: "list all escrowAccount",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllEscrowAccountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.EscrowAccountAll(cmd.Context(), params)
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

func CmdShowEscrowAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-escrow-account [id]",
		Short: "shows a escrowAccount",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			addr := args[0]
			params := &types.QueryGetEscrowAccountRequest{
				Address: addr,
			}

			res, err := queryClient.EscrowAccount(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
