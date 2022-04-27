package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdRegisterApplication())
	cmd.AddCommand(CmdRegisterName())
	cmd.AddCommand(CmdAccessName())
	cmd.AddCommand(CmdUpdateName())
	cmd.AddCommand(CmdAccessApplication())
	cmd.AddCommand(CmdUpdateApplication())
	cmd.AddCommand(CmdCreateWhoIs())
	cmd.AddCommand(CmdUpdateWhoIs())
	cmd.AddCommand(CmdDeleteWhoIs())
	// this line is used by starport scaffolding # 1

	return cmd
}
