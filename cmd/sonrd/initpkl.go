package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/apple/pkl-go/pkl"
	"github.com/spf13/cobra"
)

var configDir string

func newPklInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init-pkl",
		Short: "Initialize the Sonrd configuration using PKL",
		RunE: func(cmd *cobra.Command, args []string) error {
			evaluator, err := pkl.NewEvaluator(context.Background(), pkl.PreconfiguredOptions)
			if err != nil {
				return err
			}
			err = createAppToml(evaluator, formatConfigPath(cmd, "app.toml"))
			if err != nil {
				return err
			}
			err = createConfigToml(evaluator, formatConfigPath(cmd, "config.toml"))
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&configDir, "config-dir", "~/.sonr/config", "Path to where pkl files should be output")
	return cmd
}

func createAppToml(evaluator pkl.Evaluator, path string) error {
	appSource := pkl.UriSource("https://pkl.sh/sonr.chain/0.0.2/App.pkl")
	res, err := evaluator.EvaluateOutputText(context.Background(), appSource)
	if err != nil {
		return err
	}
	return writeConfigFile(path, res)
}

func createConfigToml(evaluator pkl.Evaluator, path string) error {
	configSource := pkl.UriSource("https://pkl.sh/sonr.conf/0.0.2/Config.pkl")
	res, err := evaluator.EvaluateOutputText(context.Background(), configSource)
	if err != nil {
		return err
	}
	return writeConfigFile(path, res)
}

func formatConfigPath(cmd *cobra.Command, fileName string) string {
	configDir := cmd.Flag("config-dir").Value.String()
	return filepath.Join(configDir, fileName)
}

func writeConfigFile(path string, content string) error {
	if err := os.MkdirAll(path, 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0o644)
}
