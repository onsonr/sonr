package e2e

import (
	"context"
	"fmt"
	"testing"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/strangelove-ventures/poa"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	numPOAVals = 2
)

func TestPOA(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	// setup base chain
	chains := interchaintest.CreateChainWithConfig(t, numPOAVals, NumberFullNodes, Name, ChainImage.Version, DefaultChainConfig)
	chain := chains[0].(*cosmos.CosmosChain)

	enableBlockDB := false
	ctx, _, _, _ := interchaintest.BuildInitialChain(t, chains, enableBlockDB)

	// setup accounts
	acc0, err := interchaintest.GetAndFundTestUserWithMnemonic(ctx, "acc0", AccMnemonic, GenesisFundsAmount, chain)
	require.NoError(t, err)
	acc1, err := interchaintest.GetAndFundTestUserWithMnemonic(ctx, "acc1", Acc1Mnemonic, GenesisFundsAmount, chain)
	require.NoError(t, err)

	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), GenesisFundsAmount, chain)
	incorrectUser := users[0]

	// get validator operator addresses
	vals, err := chain.StakingQueryValidators(ctx, stakingtypes.Bonded.String())
	require.NoError(t, err)
	require.Equal(t, len(vals), numPOAVals)

	validators := make([]string, len(vals))
	for i, v := range vals {
		validators[i] = v.OperatorAddress
	}

	// === Test Cases ===
	testStakingDisabled(t, ctx, chain, validators, acc0, acc1)
	testPowerErrors(t, ctx, chain, validators, incorrectUser, acc0)
	testPending(t, ctx, chain, acc0)
	testRemoveValidator(t, ctx, chain, validators, acc0)
}

func testRemoveValidator(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, validators []string, acc0 ibc.Wallet) {
	t.Log("\n===== TEST REMOVE VALIDATOR =====")
	powerOne := int64(9_000_000_000_000)
	powerTwo := int64(2_500_000)

	res, _ := POASetPower(t, ctx, chain, acc0, validators[0], powerOne, "--unsafe")
	fmt.Printf("%+v", res)
	res, _ = POASetPower(t, ctx, chain, acc0, validators[1], powerTwo, "--unsafe")
	fmt.Printf("%+v", res)

	// decode res.TxHash into a TxResponse
	txRes, err := chain.GetTransaction(res.TxHash)
	require.NoError(t, err)
	fmt.Printf("%+v", txRes)

	if err := testutil.WaitForBlocks(ctx, 2, chain); err != nil {
		t.Fatal(err)
	}

	vals, err := chain.StakingQueryValidators(ctx, stakingtypes.Bonded.String())
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("%d", powerOne), vals[0].Tokens.String())
	require.Equal(t, fmt.Sprintf("%d", powerTwo), vals[1].Tokens.String())

	// validate the validators both have a conesnsus-power of /1_000_000
	p1 := GetPOAConsensusPower(t, ctx, chain, vals[0].OperatorAddress)
	require.EqualValues(t, powerOne/1_000_000, p1) // = 9000000
	p2 := GetPOAConsensusPower(t, ctx, chain, vals[1].OperatorAddress)
	require.EqualValues(t, powerTwo/1_000_000, p2) // = 2

	// remove the 2nd validator (lower power)
	POARemove(t, ctx, chain, acc0, validators[1])

	// allow the poa.BeginBlocker to update new status
	if err := testutil.WaitForBlocks(ctx, 5, chain); err != nil {
		t.Fatal(err)
	}

	vals, err = chain.StakingQueryValidators(ctx, stakingtypes.Bonded.String())
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("%d", powerOne), vals[0].Tokens.String())
	require.Equal(t, 1, len(vals))

	vals, err = chain.StakingQueryValidators(ctx, stakingtypes.Unbonded.String())
	require.NoError(t, err)
	require.Equal(t, "0", vals[0].Tokens.String())
	require.Equal(t, 1, len(vals))
	p2 = GetPOAConsensusPower(t, ctx, chain, vals[0].OperatorAddress)
	require.EqualValues(t, 0, p2)
}

func testStakingDisabled(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, validators []string, acc0, acc1 ibc.Wallet) {
	t.Log("\n===== TEST STAKING DISABLED =====")

	err := chain.GetNode().StakingDelegate(ctx, acc0.KeyName(), validators[0], "1stake")
	require.Error(t, err)
	require.Contains(t, err.Error(), poa.ErrStakingActionNotAllowed.Error())

	granter := acc1
	grantee := acc0

	// Grant grantee (acc0) the ability to delegate from granter (acc1)
	res, err := chain.GetNode().AuthzGrant(ctx, granter, grantee.FormattedAddress(), "generic", "--msg-type", "/cosmos.staking.v1beta1.MsgDelegate")
	require.NoError(t, err)
	require.EqualValues(t, res.Code, 0)

	// Generate nested message
	nested := []string{"tx", "staking", "delegate", validators[0], "1stake"}
	nestedCmd := TxCommandBuilder(ctx, chain, nested, granter.FormattedAddress())

	// Execute nested message via a wrapped Exec
	_, err = chain.GetNode().AuthzExec(ctx, grantee, nestedCmd)
	require.Error(t, err)
	require.Contains(t, err.Error(), poa.ErrStakingActionNotAllowed.Error())
}

func testPending(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, acc0 ibc.Wallet) {
	t.Log("\n===== TEST PENDING =====")

	res, _ := POACreatePendingValidator(t, ctx, chain, acc0, "pl3Q8OQwtC7G2dSqRqsUrO5VZul7l40I+MKUcejqRsg=", "testval", "0.10", "0.25", "0.05")
	require.EqualValues(t, 0, res.Code)

	require.NoError(t, testutil.WaitForBlocks(ctx, 2, chain))

	pv := GetPOAPending(t, ctx, chain)
	require.Equal(t, 1, len(pv))
	require.Equal(t, "0", pv[0].Tokens.String())
	require.Equal(t, "1", pv[0].MinSelfDelegation.String())

	res, _ = POARemovePending(t, ctx, chain, acc0, pv[0].OperatorAddress)
	require.EqualValues(t, 0, res.Code)

	// validate it was removed
	pv = GetPOAPending(t, ctx, chain)
	require.Equal(t, 0, len(pv))
}

func testPowerErrors(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, validators []string, incorrectUser ibc.Wallet, admin ibc.Wallet) {
	t.Log("\n===== TEST POWER ERRORS =====")
	var res sdk.TxResponse
	var err error

	t.Run("fail: set-power message from a non authorized user", func(t *testing.T) {
		res, _ = POASetPower(t, ctx, chain, incorrectUser, validators[1], 1_000_000)
		res, err := chain.GetTransaction(res.TxHash)
		require.NoError(t, err)
		require.Contains(t, res.RawLog, poa.ErrNotAnAuthority.Error())
	})

	t.Run("fail: set-power message below minimum power requirement (self bond)", func(t *testing.T) {
		res, err = POASetPower(t, ctx, chain, admin, validators[0], 1)
		require.Error(t, err) // cli validate error
		require.Contains(t, err.Error(), poa.ErrPowerBelowMinimum.Error())
	})

	t.Run("fail: set-power message above 30%% without unsafe flag", func(t *testing.T) {
		res, _ = POASetPower(t, ctx, chain, admin, validators[0], 9_000_000_000_000_000)
		res, err := chain.GetTransaction(res.TxHash)
		require.NoError(t, err)
		require.Contains(t, res.RawLog, poa.ErrUnsafePower.Error())
	})
}
