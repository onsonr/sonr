package routes

import (
	"github.com/labstack/echo/v4"

	cosmos "github.com/sonrhq/sonr/pkg/handlers/cosmos"
	sonr "github.com/sonrhq/sonr/pkg/handlers/sonr"
)

// RegisterCosmosAPI registers the Cosmos API routes
func RegisterCosmosAPI(e *echo.Echo) {
	e.GET("/balance/:address", cosmos.BankAPI.GetAllBalances)
	e.GET("/balance/:address/:denom", cosmos.BankAPI.GetBalance)
	e.GET("/balance/:address/:denom/spendable", cosmos.BankAPI.GetSpendableBalancesByDenom)
	e.GET("/balance/:address/spendable", cosmos.BankAPI.GetSpendableBalances)
	e.GET("/block", cosmos.CometAPI.GetLatestBlock)
	e.GET("/block/:height", cosmos.CometAPI.GetBlockByHeight)
	e.GET("/constitution", cosmos.GovAPI.GetConstitution)
	e.GET("/delegators/:delegatorAddr", cosmos.StakingAPI.GetDelegatorDelegations)
	e.GET("/delegators/:delegatorAddr/validators", cosmos.StakingAPI.GetDelegatorValidators)
	e.GET("/delegators/:delegatorAddr/validators/:validatorAddr", cosmos.StakingAPI.GetDelegation)
	e.GET("/delegators/:delegatorAddr/validators/:validatorAddr/unbonding", cosmos.StakingAPI.GetUnbondingDelegation)
	e.GET("/delegators/:delegatorAddr/validators/:srcValidatorAddr/redelegate/:dstValidatorAddr", cosmos.StakingAPI.GetRedelegations)
	e.GET("/delegators/:delegatorAddr/unbonding", cosmos.StakingAPI.GetDelegatorUnbondingDelegations)
	e.GET("/health", cosmos.CometAPI.GetNodeInfo)
	e.GET("/history/{height}", cosmos.StakingAPI.GetHistoricalInfo)
	e.GET("/proposals", cosmos.GovAPI.GetProposals)
	e.GET("/proposals/:proposalId", cosmos.GovAPI.GetProposal)
	e.GET("/proposals/:proposalId/deposits", cosmos.GovAPI.GetDeposits)
	e.GET("/proposals/:proposalId/deposits/:depositor", cosmos.GovAPI.GetDeposit)
	e.GET("/proposals/:proposalId/tally", cosmos.GovAPI.GetTally)
	e.GET("/proposals/:proposalId/votes", cosmos.GovAPI.GetVotes)
	e.GET("/proposals/:proposalId/votes/:voter", cosmos.GovAPI.GetVote)
	e.GET("/staking", cosmos.StakingAPI.GetValidators)
	e.GET("/staking/:validatorAddr", cosmos.StakingAPI.GetValidator)
	e.GET("/staking/:validatorAddr/delegations", cosmos.StakingAPI.GetValidatorDelegations)
	e.GET("/supply", cosmos.BankAPI.GetTotalSupply)
	e.GET("/supply/:denom", cosmos.BankAPI.GetSupplyOf)
	e.GET("/syncing", cosmos.CometAPI.GetSyncing)
	e.GET("/validators", cosmos.CometAPI.GetLatestValidatorSet)
	e.GET("/validators/:height", cosmos.CometAPI.GetValidatorSetByHeight)
}

// RegisterSonrAPI registers the Sonr API routes
func RegisterSonrAPI(e *echo.Echo) {
	e.GET("/check/identifier/:id", sonr.AuthAPI.CheckIdentifier)
	e.GET("/service/:origin", sonr.ServiceAPI.QueryOrigin)
	e.GET("/service/:origin/login/:username/start", sonr.ServiceAPI.StartLogin)
	e.POST("/service/:origin/login/:username/finish", sonr.ServiceAPI.FinishLogin)
	e.GET("/service/:origin/register/:username/start", sonr.ServiceAPI.StartRegistration)
	e.POST("/service/:origin/register/:username/finish", sonr.ServiceAPI.FinishRegistration)
	e.GET("/tx/:txHash", cosmos.TxAPI.GetTx)
	e.GET("/tx/block/:height", cosmos.TxAPI.GetBlockWithTxs)
	e.POST("/tx/broadcast", cosmos.TxAPI.BroadcastTx)
	e.POST("/tx/simulate", cosmos.TxAPI.SimulateTx)
}
