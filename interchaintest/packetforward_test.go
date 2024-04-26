package e2e

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"cosmossdk.io/math"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	interchaintestrelayer "github.com/strangelove-ventures/interchaintest/v8/relayer"
	"github.com/strangelove-ventures/interchaintest/v8/testreporter"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PacketMetadata struct {
	Forward *ForwardMetadata `json:"forward"`
}

type ForwardMetadata struct {
	Receiver       string        `json:"receiver"`
	Port           string        `json:"port"`
	Channel        string        `json:"channel"`
	Timeout        time.Duration `json:"timeout"`
	Retries        *uint8        `json:"retries,omitempty"`
	Next           *string       `json:"next,omitempty"`
	RefundSequence *uint64       `json:"refund_sequence,omitempty"`
}

func TestPacketForwardMiddleware(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		ctx                             = context.Background()
		client, network                 = interchaintest.DockerSetup(t)
		rep                             = testreporter.NewNopReporter()
		eRep                            = rep.RelayerExecReporter(t)
		chainID_A, chainID_B, chainID_C = "chain-a", "chain-b", "chain-c"
		chainA, chainB, chainC          *cosmos.CosmosChain
	)

	// base config which all networks will use as defaults.
	baseCfg := DefaultChainConfig

	// Set specific chain ids for each so they are their own unique networks
	baseCfg.ChainID = chainID_A
	configA := baseCfg

	baseCfg.ChainID = chainID_B
	configB := baseCfg

	baseCfg.ChainID = chainID_C
	configC := baseCfg

	// Create chain factory with multiple individual networks.
	numVals := 1
	numFullNodes := 0

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:          configA.Name,
			ChainConfig:   configA,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
		{
			Name:          configA.Name,
			ChainConfig:   configB,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
		{
			Name:          configA.Name,
			ChainConfig:   configC,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
	})

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	chainA, chainB, chainC = chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain), chains[2].(*cosmos.CosmosChain)

	r := interchaintest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		interchaintestrelayer.CustomDockerImage(RelayerRepo, RelayerVersion, "100:1000"),
		interchaintestrelayer.StartupFlags("--processor", "events", "--block-history", "100"),
	).Build(t, client, network)

	const pathAB = "ab"
	const pathBC = "bc"

	ic := interchaintest.NewInterchain().
		AddChain(chainA).
		AddChain(chainB).
		AddChain(chainC).
		AddRelayer(r, "relayer").
		AddLink(interchaintest.InterchainLink{
			Chain1:  chainA,
			Chain2:  chainB,
			Relayer: r,
			Path:    pathAB,
		}).
		AddLink(interchaintest.InterchainLink{
			Chain1:  chainB,
			Chain2:  chainC,
			Relayer: r,
			Path:    pathBC,
		})

	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),

		SkipPathCreation: false,
	}))
	t.Cleanup(func() {
		_ = ic.Close()
	})

	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), GenesisFundsAmount, chainA, chainB, chainC)

	abChan, err := ibc.GetTransferChannel(ctx, r, eRep, chainID_A, chainID_B)
	require.NoError(t, err)

	baChan := abChan.Counterparty

	cbChan, err := ibc.GetTransferChannel(ctx, r, eRep, chainID_C, chainID_B)
	require.NoError(t, err)

	bcChan := cbChan.Counterparty

	// Start the relayer on all paths
	err = r.StartRelayer(ctx, eRep, pathAB, pathBC)
	require.NoError(t, err)

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				t.Logf("an error occurred while stopping the relayer: %s", err)
			}
		},
	)

	// Get original account balances
	userA, userB, userC := users[0], users[1], users[2]

	var transferAmount math.Int = math.NewInt(100_000)

	// Compose the prefixed denoms and ibc denom for asserting balances
	firstHopDenom := transfertypes.GetPrefixedDenom(baChan.PortID, baChan.ChannelID, chainA.Config().Denom)
	secondHopDenom := transfertypes.GetPrefixedDenom(cbChan.PortID, cbChan.ChannelID, firstHopDenom)

	firstHopDenomTrace := transfertypes.ParseDenomTrace(firstHopDenom)
	secondHopDenomTrace := transfertypes.ParseDenomTrace(secondHopDenom)

	firstHopIBCDenom := firstHopDenomTrace.IBCDenom()
	secondHopIBCDenom := secondHopDenomTrace.IBCDenom()

	firstHopEscrowAccount := sdk.MustBech32ifyAddressBytes(chainA.Config().Bech32Prefix, transfertypes.GetEscrowAddress(abChan.PortID, abChan.ChannelID))
	secondHopEscrowAccount := sdk.MustBech32ifyAddressBytes(chainB.Config().Bech32Prefix, transfertypes.GetEscrowAddress(bcChan.PortID, bcChan.ChannelID))

	t.Run("multi-hop a->b->c", func(t *testing.T) {
		// Send packet from Chain A->Chain B->Chain C

		transfer := ibc.WalletAmount{
			Address: userB.FormattedAddress(),
			Denom:   chainA.Config().Denom,
			Amount:  transferAmount,
		}

		firstHopMetadata := &PacketMetadata{
			Forward: &ForwardMetadata{
				Receiver: userC.FormattedAddress(),
				Channel:  bcChan.ChannelID,
				Port:     bcChan.PortID,
			},
		}

		memo, err := json.Marshal(firstHopMetadata)
		require.NoError(t, err)

		chainAHeight, err := chainA.Height(ctx)
		require.NoError(t, err)

		transferTx, err := chainA.SendIBCTransfer(ctx, abChan.ChannelID, userA.KeyName(), transfer, ibc.TransferOptions{Memo: string(memo)})
		require.NoError(t, err)
		_, err = testutil.PollForAck(ctx, chainA, chainAHeight, chainAHeight+30, transferTx.Packet)
		require.NoError(t, err)
		err = testutil.WaitForBlocks(ctx, 1, chainA)
		require.NoError(t, err)

		chainABalance, err := chainA.GetBalance(ctx, userA.FormattedAddress(), chainA.Config().Denom)
		require.NoError(t, err)

		chainBBalance, err := chainB.GetBalance(ctx, userB.FormattedAddress(), firstHopIBCDenom)
		require.NoError(t, err)

		chainCBalance, err := chainC.GetBalance(ctx, userC.FormattedAddress(), secondHopIBCDenom)
		require.NoError(t, err)

		require.Equal(t, GenesisFundsAmount.Sub(transferAmount).Int64(), chainABalance.Int64())
		require.Equal(t, int64(0), chainBBalance.Int64())
		require.Equal(t, int64(100000), chainCBalance.Int64())

		firstHopEscrowBalance, err := chainA.GetBalance(ctx, firstHopEscrowAccount, chainA.Config().Denom)
		require.NoError(t, err)

		secondHopEscrowBalance, err := chainB.GetBalance(ctx, secondHopEscrowAccount, firstHopIBCDenom)
		require.NoError(t, err)

		require.Equal(t, transferAmount.Int64(), firstHopEscrowBalance.Int64())
		require.Equal(t, transferAmount.Int64(), secondHopEscrowBalance.Int64())
	})
}
