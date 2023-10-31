package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonrhq/sonr/x/identity/types"
	"github.com/spf13/cobra"
)

func CmdListControllerAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-controller-account",
		Short: "list all controllerAccount",
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

			params := &types.QueryAllControllerAccountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ControllerAccountAll(cmd.Context(), params)
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

func CmdShowControllerAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-controller-account [id]",
		Short: "shows a controllerAccount",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			addr := args[0]

			params := &types.QueryGetControllerAccountRequest{
				Address: addr,
			}

			res, err := queryClient.ControllerAccount(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
