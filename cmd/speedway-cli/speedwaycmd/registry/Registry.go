package MotorRegistry

import (
	"context"

	"github.com/spf13/cobra"
)

func BootstrapRegistryCommand(ctx context.Context) (registryCmd *cobra.Command) {
	registryCmd = &cobra.Command{
		Use:   "speedway-registry",
		Short: "Provides commands for managing registries on the Sonr network",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	registryCmd.AddCommand(bootstrapCreateAccountCommand(ctx))
	registryCmd.AddCommand(bootstrapLoginCommand(ctx))
	return
}
