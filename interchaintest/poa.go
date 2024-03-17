package e2e

import (
	"context"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/poa"
	"github.com/stretchr/testify/require"
)

func POASetPower(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, user ibc.Wallet, valoper string, power int64, flags ...string) (sdk.TxResponse, error) {
	cmd := TxCommandBuilder(ctx, chain, []string{"tx", "poa", "set-power", valoper, fmt.Sprintf("%d", power)}, user.KeyName(), flags...)
	return ExecuteTransaction(ctx, chain, cmd)
}

func POARemove(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, user ibc.Wallet, valoper string) (sdk.TxResponse, error) {
	cmd := TxCommandBuilder(ctx, chain, []string{"tx", "poa", "remove", valoper}, user.KeyName())
	return ExecuteTransaction(ctx, chain, cmd)
}

func POARemovePending(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, user ibc.Wallet, valoper string) (sdk.TxResponse, error) {
	cmd := TxCommandBuilder(ctx, chain, []string{"tx", "poa", "remove-pending", valoper}, user.KeyName())
	return ExecuteTransaction(ctx, chain, cmd)
}

func POACreatePendingValidator(
	t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, user ibc.Wallet,
	ed25519PubKey string, moniker string, commissionRate string, commissionMaxRate string, commissionMaxChangeRate string,
) (sdk.TxResponse, error) {
	file := "validator_file.json"

	content := fmt.Sprintf(`{
		"pubkey": {"@type":"/cosmos.crypto.ed25519.PubKey","key":"%s"},
		"amount": "0%s",
		"moniker": "%s",
		"identity": "",
		"website": "https://website.com",
		"security": "security@cosmos.xyz",
		"details": "description",
		"commission-rate": "%s",
		"commission-max-rate": "%s",
		"commission-max-change-rate": "%s",
		"min-self-delegation": "1"

	}`, ed25519PubKey, chain.Config().Denom, moniker, commissionRate, commissionMaxRate, commissionMaxChangeRate)

	err := chain.GetNode().WriteFile(ctx, []byte(content), file)
	require.NoError(t, err)

	filePath := fmt.Sprintf("%s/%s", chain.GetNode().HomeDir(), file)
	cmd := TxCommandBuilder(ctx, chain, []string{"tx", "poa", "create-validator", filePath}, user.KeyName())
	return ExecuteTransaction(ctx, chain, cmd)
}

func SubmitGovernanceProposalForValidatorChanges(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, user ibc.Wallet, validator string, power uint64, unsafe bool) string {
	govAddr := "cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn"

	powerMsg := []cosmos.ProtoMessage{
		&poa.MsgSetPower{
			Sender:           govAddr,
			ValidatorAddress: validator,
			Power:            power,
			Unsafe:           unsafe,
		},
	}

	title := fmt.Sprintf("Update" + validator + "Power")
	desc := fmt.Sprintf("Updating power for validator %s to %d", validator, power)

	proposal, err := chain.BuildProposal(powerMsg, title, desc, desc, fmt.Sprintf(`50%s`, chain.Config().Denom), user.FormattedAddress(), false)
	require.NoError(t, err, "error building proposal")

	fmt.Printf("proposal: %+v\n", proposal)

	txProp, err := chain.SubmitProposal(ctx, user.KeyName(), proposal)
	t.Log("txProp", txProp)
	require.NoError(t, err, "error submitting proposal")

	return txProp.ProposalID
}

func GetPOAParams(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain) poa.Params {
	var res poa.ParamsResponse
	ExecuteQuery(ctx, chain, []string{"query", "poa", "params"}, &res)
	return res.Params
}

func GetPOAConsensusPower(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, valoperAddr string) int64 {
	res, err := poa.NewQueryClient(chain.GetNode().GrpcConn).ConsensusPower(ctx, &poa.QueryConsensusPowerRequest{ValidatorAddress: valoperAddr})
	require.NoError(t, err)
	return res.ConsensusPower
}

func GetPOAPending(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain) []poa.Validator {
	res, err := poa.NewQueryClient(chain.GetNode().GrpcConn).PendingValidators(ctx, &poa.QueryPendingValidatorsRequest{})
	require.NoError(t, err)

	return res.Pending
}

func POAUpdateParams(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, user ibc.Wallet, admins []string, gracefulExit bool) (sdk.TxResponse, error) {
	// admin1,admin2,admin3
	adminList := ""
	for _, admin := range admins {
		adminList += admin + ","
	}
	adminList = adminList[:len(adminList)-1]

	gracefulParam := "true"
	if !gracefulExit {
		gracefulParam = "false"
	}

	cmd := TxCommandBuilder(ctx, chain, []string{"tx", "poa", "update-params", adminList, gracefulParam}, user.KeyName())
	return ExecuteTransaction(ctx, chain, cmd)
}

func POAUpdateStakingParams(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, user ibc.Wallet, sp stakingtypes.Params) (sdk.TxResponse, error) {
	command := []string{"tx", "poa", "update-staking-params",
		sp.UnbondingTime.String(),
		fmt.Sprintf("%d", sp.MaxValidators),
		fmt.Sprintf("%d", sp.MaxEntries),
		fmt.Sprintf("%d", sp.HistoricalEntries),
		sp.BondDenom,
		fmt.Sprintf("%d", sp.MinCommissionRate),
	}

	cmd := TxCommandBuilder(ctx, chain, command, user.KeyName())
	return ExecuteTransaction(ctx, chain, cmd)
}
