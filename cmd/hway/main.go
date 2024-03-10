package main

import (
	"github.com/cosmos/cosmos-sdk/version"
)

var (
	// Commit is set by the compiler via -ldflags.
	Commit = "unset"

	// Version is set by the compiler via -ldflags.
	Version = "unset"
)

// init sets the version flags.
func init() {
	version.Name = "Sonr Highway"
	version.AppName = "hway"
	version.Version = Version
	version.Commit = Commit
}

// main is the entry point for the application.
func main() {
	rootCmd.Flags().String("hway-host", "0.0.0.0", "host")
	rootCmd.Flags().Int("hway-port", 8000, "port")
	rootCmd.Flags().String("hway-psql", "postgresql://sonr:sonr@localhost:5432/sonr?sslmode=disable", "psql connection string")
	rootCmd.Flags().String("val-host", "localhost", "validator host")
	rootCmd.Flags().Int("val-rpc", 26657, "validator rpc port")
	rootCmd.Flags().Int("val-grpc", 9090, "validator grpc port")
	rootCmd.AddCommand(versionCmd)
	err := rootCmd.Execute()
	if err != nil {
		rootCmd.PrintErr(err)
	}
}
