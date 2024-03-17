package e2e

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/stretchr/testify/require"
)

type GetCountResponse struct {
	// {"data":{"count":0}}
	Data *GetCountObj `json:"data"`
}

type GetCountObj struct {
	Count int64 `json:"count"`
}

func TestCosmWasmIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	// Base setup
	chains := interchaintest.CreateChainWithConfig(t, 1, 0, Name, ChainImage.Version, DefaultChainConfig)
	ctx, ic, _, _ := interchaintest.BuildInitialChain(t, chains, false)

	require.NotNil(t, ic)
	require.NotNil(t, ctx)

	chain := chains[0].(*cosmos.CosmosChain)

	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), GenesisFundsAmount, chain)
	user := users[0]

	StdExecute(t, ctx, chain, user)

	t.Cleanup(func() {
		_ = ic.Close()
	})
}

func StdExecute(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, user ibc.Wallet) (contractAddr string) {
	_, contractAddr = SetupContract(t, ctx, chain, user.KeyName(), "contracts/cw_template.wasm", `{"count":0}`)
	chain.ExecuteContract(ctx, user.KeyName(), contractAddr, `{"increment":{}}`, "--fees", "10000"+chain.Config().Denom)

	var res GetCountResponse
	err := SmartQueryString(t, ctx, chain, contractAddr, `{"get_count":{}}`, &res)
	require.NoError(t, err)

	require.Equal(t, int64(1), res.Data.Count)

	return contractAddr
}

func SmartQueryString(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, contractAddr, queryMsg string, res interface{}) error {
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(queryMsg), &jsonMap); err != nil {
		t.Fatal(err)
	}
	err := chain.QueryContract(ctx, contractAddr, jsonMap, &res)
	return err
}

func SetupContract(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, keyname string, fileLoc string, message string, extraFlags ...string) (codeId, contract string) {
	codeId, err := chain.StoreContract(ctx, keyname, fileLoc)
	if err != nil {
		t.Fatal(err)
	}

	needsNoAdminFlag := true
	for _, flag := range extraFlags {
		if flag == "--admin" {
			needsNoAdminFlag = false
		}
	}

	contractAddr, err := chain.InstantiateContract(ctx, keyname, codeId, message, needsNoAdminFlag, extraFlags...)
	if err != nil {
		t.Fatal(err)
	}

	return codeId, contractAddr
}
