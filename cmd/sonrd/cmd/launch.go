package cmd

import (
	"github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/sonrhq/core/config"
	"github.com/spf13/cobra"
)

// LaunchCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func LaunchCmd(mbm module.BasicManager, txEncCfg client.TxEncodingConfig, genBalIterator types.GenesisBalancesIterator, defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "launch",
		Short: "Launch sonr node with all files needed for Tendermint and the respective application.",
		Long:  `Initialize validators's and node's configuration files.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := baseInitFunc(mbm, cmd, args)
            if err != nil {
                return err
            }
            return baseGentxFunc(mbm, txEncCfg, genBalIterator, cmd, args)
		},
	}
	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().BoolP(FlagOverwrite, "o", false, "overwrite the genesis.json file")
	cmd.Flags().Bool(FlagRecover, false, "provide seed phrase to recover existing key instead of creating")
	cmd.Flags().String(flags.FlagChainID, config.ChainID(), "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().Int64(flags.FlagInitHeight, 1, "specify the initial block height at genesis")
	return cmd
}

func setupSeeds(cmd *cobra.Command) {
}

func setupPersistentPeers(cmd *cobra.Command) {
}

func setupPrivValidator(cmd *cobra.Command) {
}


