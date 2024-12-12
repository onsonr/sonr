package main

import (
	"context"
	"log"
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
			ctx := context.Background()
			evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
			if err != nil {
				return err
			}
			defer evaluator.Close()

			appPath := formatConfigPath(cmd, "app.toml")
			configPath := formatConfigPath(cmd, "config.toml")

			// Create app.toml
			if err := createAppToml(evaluator, appPath); err != nil {
				cmd.PrintErrf("Failed to create app.toml: %v\n", err)
				return err
			}
			cmd.Printf("Successfully created %s\n", appPath)

			// Create config.toml
			if err := createConfigToml(evaluator, configPath); err != nil {
				cmd.PrintErrf("Failed to create config.toml: %v\n", err)
				return err
			}
			cmd.Printf("Successfully created %s\n", configPath)

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
	log.Printf("res: %s", res)
	return writeConfigFile(path, res)
}

func createConfigToml(evaluator pkl.Evaluator, path string) error {
	configSource := pkl.UriSource("https://pkl.sh/sonr.chain/0.0.2/Config.pkl")
	res, err := evaluator.EvaluateOutputText(context.Background(), configSource)
	if err != nil {
		return err
	}
	log.Printf("res: %s", res)
	return writeConfigFile(path, res)
}

func formatConfigPath(cmd *cobra.Command, fileName string) string {
	configDir := cmd.Flag("config-dir").Value.String()
	// Expand home directory if needed
	if configDir[:2] == "~/" {
		home, err := os.UserHomeDir()
		if err == nil {
			configDir = filepath.Join(home, configDir[2:])
		}
	}
	return filepath.Join(configDir, fileName)
}

func writeConfigFile(path string, content string) error {
	// Create the directory path if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	// Check if file already exists
	if _, err := os.Stat(path); err == nil {
		// File exists, create backup
		backupPath := path + ".backup"
		if err := os.Rename(path, backupPath); err != nil {
			return err
		}
	}

	// Write the new config file
	return os.WriteFile(path, []byte(content), 0o644)
}
