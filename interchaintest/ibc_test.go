package e2e

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	interchaintestrelayer "github.com/strangelove-ventures/interchaintest/v8/relayer"
	"github.com/strangelove-ventures/interchaintest/v8/testreporter"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

const (
	ibcPath = "ibc-path"
)

func TestIBC(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	cfgA := DefaultChainConfig
	cfgA.ChainID = cfgA.ChainID + "-1"

	cfgB := DefaultChainConfig
	cfgB.ChainID = cfgB.ChainID + "-2"

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t, zaptest.Level(zapcore.DebugLevel)), []*interchaintest.ChainSpec{
		{
			Name:          DefaultChainConfig.Name,
			Version:       ChainImage.Version,
			ChainName:     cfgA.ChainID,
			NumValidators: &NumberVals,
			NumFullNodes:  &NumberFullNodes,
			ChainConfig:   cfgA,
		},
		{
			Name:          DefaultChainConfig.Name,
			Version:       ChainImage.Version,
			ChainName:     cfgB.ChainID,
			NumValidators: &NumberVals,
			NumFullNodes:  &NumberFullNodes,
			ChainConfig:   cfgB,
		},
	})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chainA, chainB := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

	// Relayer Factory
	client, network := interchaintest.DockerSetup(t)
	rf := interchaintest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t, zaptest.Level(zapcore.DebugLevel)),
		interchaintestrelayer.CustomDockerImage(RelayerRepo, RelayerVersion, "100:1000"),
		interchaintestrelayer.StartupFlags("--processor", "events", "--block-history", "100"),
	)

	r := rf.Build(t, client, network)

	ic := interchaintest.NewInterchain().
		AddChain(chainA).
		AddChain(chainB).
		AddRelayer(r, "relayer").
		AddLink(interchaintest.InterchainLink{
			Chain1:  chainA,
			Chain2:  chainB,
			Relayer: r,
			Path:    ibcPath,
		})

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	// Build interchain
	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:         t.Name(),
		Client:           client,
		NetworkID:        network,
		SkipPathCreation: false,
	}))

	// Create and Fund User Wallets
	fundAmount := math.NewInt(10_000_000)
	users := interchaintest.GetAndFundTestUsers(t, ctx, "default", fundAmount, chainA, chainB)
	userA := users[0]
	userB := users[1]

	userAInitial, err := chainA.GetBalance(ctx, userA.FormattedAddress(), chainA.Config().Denom)
	require.NoError(t, err)
	require.True(t, userAInitial.Equal(fundAmount))

	// Get Channel ID
	aInfo, err := r.GetChannels(ctx, eRep, chainA.Config().ChainID)
	require.NoError(t, err)
	aChannelID := aInfo[0].ChannelID

	bInfo, err := r.GetChannels(ctx, eRep, chainB.Config().ChainID)
	require.NoError(t, err)
	bChannelID := bInfo[0].ChannelID

	// Send Transaction
	amountToSend := math.NewInt(1_000_000)
	dstAddress := userB.FormattedAddress()
	transfer := ibc.WalletAmount{
		Address: dstAddress,
		Denom:   chainA.Config().Denom,
		Amount:  amountToSend,
	}

	_, err = chainA.SendIBCTransfer(ctx, aChannelID, userA.KeyName(), transfer, ibc.TransferOptions{})
	require.NoError(t, err)

	// relay MsgRecvPacket to chainB, then MsgAcknowledgement back to chainA
	require.NoError(t, r.Flush(ctx, eRep, ibcPath, aChannelID))

	// test source wallet has decreased funds
	expectedBal := userAInitial.Sub(amountToSend)
	aNewBal, err := chainA.GetBalance(ctx, userA.FormattedAddress(), chainA.Config().Denom)
	require.NoError(t, err)
	require.True(t, aNewBal.Equal(expectedBal))

	// Trace IBC Denom
	srcDenomTrace := transfertypes.ParseDenomTrace(transfertypes.GetPrefixedDenom("transfer", bChannelID, chainA.Config().Denom))
	dstIbcDenom := srcDenomTrace.IBCDenom()

	// Test destination wallet has increased funds
	bNewBal, err := chainB.GetBalance(ctx, userB.FormattedAddress(), dstIbcDenom)
	require.NoError(t, err)
	require.True(t, bNewBal.Equal(amountToSend))
}
