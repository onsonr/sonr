package cmd

import (
	"os/exec"

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
            // Update with flags
            err = updateAppToml(config.GetFlags(cmd))
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
    config.AppendFlags(cmd)
	return cmd
}

func updateAppToml(flags config.Flags) error {
    if flags.GrpcAddress != "" {
        cmd := exec.Command("sed", "-i", "s/grpc.address = .*/grpc.address = \""+flags.GrpcAddress+"\"/", flags.UsrHomeDir+"/.sonr/config/app.toml")
        err := cmd.Run()
        if err != nil {
            return err
        }
    }
    if flags.APIAddress != "" {
        cmd := exec.Command("sed", "-i", "s/api.address = .*/api.address = \""+flags.APIAddress+"\"/", flags.UsrHomeDir+"/.sonr/config/app.toml")
        err := cmd.Run()
        if err != nil {
            return err
        }
    }
    if flags.MinimumGasPrices != "" {
        cmd := exec.Command("sed", "-i", "s/minimum-gas-prices = .*/minimum-gas-prices = \""+flags.MinimumGasPrices+"\"/", flags.UsrHomeDir+"/.sonr/config/app.toml")
        err := cmd.Run()
        if err != nil {
            return err
        }
    }
    if flags.Seeds != "" {
        cmd := exec.Command("sed", "-i", "s/p2p.seeds = .*/p2p.seeds = \""+flags.Seeds+"\"/", flags.UsrHomeDir+"/.sonr/config/config.toml")
        err := cmd.Run()
        if err != nil {
            return err
        }
    }
    if flags.PersistentPeers != "" {
        cmd := exec.Command("sed", "-i", "s/p2p.persistent_peers = .*/p2p.persistent_peers = \""+flags.PersistentPeers+"\"/", flags.UsrHomeDir+"/.sonr/config/config.toml")
        err := cmd.Run()
        if err != nil {
            return err
        }
    }
    if flags.PrivatePeerIds != "" {
        cmd := exec.Command("sed", "-i", "s/p2p.private_peer_ids = .*/p2p.private_peer_ids = \""+flags.PrivatePeerIds+"\"/", flags.UsrHomeDir+"/.sonr/config/config.toml")
        err := cmd.Run()
        if err != nil {
            return err
        }
    }
    return nil
}
