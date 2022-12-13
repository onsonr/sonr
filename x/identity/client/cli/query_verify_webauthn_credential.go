package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonr-hq/sonr/x/identity/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVerifyWebauthnCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-webauthn-credential",
		Short: "Query VerifyWebauthnCredential",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryWebauthnRegisterFinishRequest{}

			res, err := queryClient.WebauthnRegistrationFinish(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
