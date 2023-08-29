package config

import "github.com/spf13/cobra"

var (
	// FlagGrpcAddress is the grpc address of the application.
	FlagGrpcAddress = Flag("grpc.address")

	// FlagAPIAddress is the api address of the application.
	FlagAPIAddress = Flag("api.address")

	// FlagMinimumGasPrices is the minimum gas prices of the application.
	FlagMinimumGasPrices = Flag("minimum-gas-prices")

	// FlagSeeds is the seeds of the application.
	FlagSeeds = Flag("p2p.seeds")

	// FlagPersistentPeers is the persistent peers of the application.
	FlagPersistentPeers = Flag("p2p.persistent_peers")

	// FlagPrivatePeerIds is the private peer ids of the application.
	FlagPrivatePeerIds = Flag("p2p.private_peer_ids")

	// FlagHomeDir is the home directory of the application.
	FlagHomeDir = Flag("home")
)

// Flag is a type alias for string.
type Flag string

// String returns the string representation of the flag.
func (f Flag) String() string {
	return string(f)
}

// AppendFlags appends the flags to the given command.
func AppendFlags(cmd *cobra.Command) {
    cmd.Flags().String(FlagGrpcAddress.String(), "0.0.0.0:26657", "grpc address of the application")
    cmd.Flags().String(FlagAPIAddress.String(), "0.0.0.0:1317", "api address of the application")
    cmd.Flags().String(FlagMinimumGasPrices.String(), "0.00usnr", "minimum gas prices of the application")
    cmd.Flags().String(FlagSeeds.String(), "", "seeds of the application")
    cmd.Flags().String(FlagPersistentPeers.String(), "", "persistent peers of the application")
    cmd.Flags().String(FlagPrivatePeerIds.String(), "", "private peer ids of the application")
    cmd.Flags().String(FlagHomeDir.String(), "$HOME/.sonr", "home directory of the application")
}

// Flags is a struct to contain the values of the flags.
type Flags struct {
    // GrpcAddress is the grpc address of the application.
    GrpcAddress string

    // APIAddress is the api address of the application.
    APIAddress string

    // MinimumGasPrices is the minimum gas prices of the application.
    MinimumGasPrices string

    // Seeds is the seeds of the application.
    Seeds string

    // PersistentPeers is the persistent peers of the application.
    PersistentPeers string

    // PrivatePeerIds is the private peer ids of the application.
    PrivatePeerIds string

    // HomeDir is the home directory of the application.
    HomeDir string
}

// GetFlags returns the flags of the given command.
func GetFlags(cmd *cobra.Command) Flags {
    return Flags{
        GrpcAddress: cmd.Flag(FlagGrpcAddress.String()).Value.String(),
        APIAddress: cmd.Flag(FlagAPIAddress.String()).Value.String(),
        MinimumGasPrices: cmd.Flag(FlagMinimumGasPrices.String()).Value.String(),
        Seeds: cmd.Flag(FlagSeeds.String()).Value.String(),
        PersistentPeers: cmd.Flag(FlagPersistentPeers.String()).Value.String(),
        PrivatePeerIds: cmd.Flag(FlagPrivatePeerIds.String()).Value.String(),
        HomeDir: cmd.Flag(FlagHomeDir.String()).Value.String(),
    }
}
