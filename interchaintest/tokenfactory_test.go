package e2e

import (
	"testing"

	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/stretchr/testify/require"
)

func TestTokenFactory(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	// setup base chain
	chains := interchaintest.CreateChainWithConfig(t, NumberVals, NumberFullNodes, Name, ChainImage.Version, DefaultChainConfig)
	chain := chains[0].(*cosmos.CosmosChain)
	ctx, ic, _, _ := interchaintest.BuildInitialChain(t, chains, false)

	users := interchaintest.GetAndFundTestUsers(t, ctx, "default", GenesisFundsAmount, chain, chain)
	user := users[0]
	user2 := users[1]

	uaddr := user.FormattedAddress()
	uaddr2 := user2.FormattedAddress()

	node := chain.GetNode()

	tfDenom, _, err := node.TokenFactoryCreateDenom(ctx, user, "ictestdenom", 2_500_00)
	t.Log("TF Denom: ", tfDenom)
	require.NoError(t, err)

	t.Log("Mint TF Denom to user")
	node.TokenFactoryMintDenom(ctx, user.FormattedAddress(), tfDenom, 100)
	if balance, err := chain.GetBalance(ctx, uaddr, tfDenom); err != nil {
		t.Fatal(err)
	} else if balance.Int64() != 100 {
		t.Fatal("balance not 100")
	}

	t.Log("Mint TF Denom to another user")
	node.TokenFactoryMintDenomTo(ctx, user.FormattedAddress(), tfDenom, 70, user2.FormattedAddress())
	if balance, err := chain.GetBalance(ctx, uaddr2, tfDenom); err != nil {
		t.Fatal(err)
	} else if balance.Int64() != 70 {
		t.Fatal("balance not 70")
	}

	t.Log("Change admin to uaddr2")
	_, err = node.TokenFactoryChangeAdmin(ctx, user.KeyName(), tfDenom, uaddr2)
	require.NoError(t, err)

	// ensure the admin is the contract
	res, err := chain.TokenFactoryQueryAdmin(ctx, tfDenom)
	require.NoError(t, err)
	require.EqualValues(t, res.AuthorityMetadata.Admin, uaddr2, "admin not uaddr2. Did not properly transfer.")

	t.Cleanup(func() {
		_ = ic.Close()
	})

}
