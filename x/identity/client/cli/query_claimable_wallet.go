package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/spf13/cobra"
)

func CmdListClaimableWallet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-claimable-wallet",
		Short: "list all ClaimableWallet",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllClaimableWalletRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ClaimableWalletAll(context.Background(), params)
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

func CmdShowClaimableWallet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-claimable-wallet [id]",
		Short: "shows a ClaimableWallet",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetClaimableWalletRequest{
				Id: id,
			}

			res, err := queryClient.ClaimableWallet(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
