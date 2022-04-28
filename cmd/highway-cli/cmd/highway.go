package cmd

import (
	"context"
	"fmt"

	"github.com/sonr-io/sonr/pkg/highway"
	"github.com/spf13/cobra"
)

func BootstrapBlobCommand(ctx context.Context, highway *highway.HighwayServer) *cobra.Command {
	// serveCmd represents the serve command
	highwayBlobCmd := &cobra.Command{
		Use:   "blob",
		Short: "Upload/Download/Delete Blobs stored on IPFS",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("serve called")
		},
	}

	return highwayBlobCmd
}

func BootstrapChannelCommand(ctx context.Context, highway *highway.HighwayServer) *cobra.Command {
	// serveCmd represents the serve command
	highwayChannelCmd := &cobra.Command{
		Use:   "channel",
		Short: "Manage Channels on the Sonr Blockchain.",
	}

	highwayChannelCmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	}

	return highwayChannelCmd
}

func BootstrapObjectCommand(ctx context.Context, highway *highway.HighwayServer) *cobra.Command {
	// createHwyCmd represents the serve command
	highwayObjectCmd := &cobra.Command{
		Use:   "object",
		Short: "Manage Objects on the Sonr Blockchain.",
	}

	highwayObjectCmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	}

	highwayObjectCmd.AddCommand(&cobra.Command{
		Use:   "object",
		Short: "Manage Objects on the Sonr Blockchain.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("serve called")
		},
	})

	return highwayObjectCmd
}

func BootstrapHighwayCommand(ctx context.Context, highway *highway.HighwayServer) *cobra.Command {
	// HighwayCmd represents the deploy command
	highwayCmd := &cobra.Command{
		Use:   "highway",
		Short: "Manage your Highway node on the Sonr Testnet and Local Dev Enviorment",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		SuggestFor: []string{"highway", "h"},
		Aliases:    []string{"h"},
	}

	highwayCmd.Run = func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	}

	highwayCmd.PreRun = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			return
		}
	}

	return highwayCmd
}
