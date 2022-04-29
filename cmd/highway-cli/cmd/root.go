package highwaycmd

import (
	"context"

	"github.com/sonr-io/sonr/cmd/highway-cli/highwaycmd/bucket"
	"github.com/spf13/cobra"
)

// Execute executes the root command
func Execute() error {
	return bootstrapRootCommand(context.Background()).Execute()
}

func bootstrapRootCommand(ctx context.Context) (rootCmd *cobra.Command) {
	rootCmd = &cobra.Command{
		Use:   "highway",
		Short: "The Sonr Highway CLI tool",
		Long:  `Manage your highway services and blockchain registered types with the Sonr CLI tool.`,

		Run: func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.AddCommand(bootstrapServeCommand(ctx))
	rootCmd.AddCommand(bucket.BootstrapBucketCommand(ctx))

	return
}
