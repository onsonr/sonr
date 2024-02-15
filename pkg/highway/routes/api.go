package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/highway/handlers"
)

// RegisterAPIRoutes registers the API routes
func RegisterAPIRoutes(e *echo.Echo) {
	registerBankRoutes(e)
	registerCometRoutes(e)
	registerGovRoutes(e)
	registerServiceRoutes(e)
	registerStakingRoutes(e)
}

// registerBankRoutes registers the bank routes
func registerBankRoutes(e *echo.Echo) {
	e.GET("/balance/:address", handlers.BankAPI.GetAllBalances)
	e.GET("/balance/:address/spendable", handlers.BankAPI.GetSpendableBalances)
	e.GET("/balance/:address/:denom", handlers.BankAPI.GetBalance)
	e.GET("/balance/:address/:denom/spendable", handlers.BankAPI.GetSpendableBalancesByDenom)
	e.GET("/supply", handlers.BankAPI.GetTotalSupply)
	e.GET("/supply/:denom", handlers.BankAPI.GetSupplyOf)
}

// registerCometRoutes registers the comet routes
func registerCometRoutes(e *echo.Echo) {
	e.GET("/block", handlers.CometAPI.GetLatestBlock)
	e.GET("/block/:height", handlers.CometAPI.GetBlockByHeight)
	e.GET("/health", handlers.CometAPI.GetNodeInfo)
	e.GET("/syncing", handlers.CometAPI.GetSyncing)
	e.GET("/validators", handlers.CometAPI.GetLatestValidatorSet)
	e.GET("/validators/:height", handlers.CometAPI.GetValidatorSetByHeight)
}

// registerGovRoutes registers the gov routes
func registerGovRoutes(e *echo.Echo) {
	e.GET("/constitution", handlers.GovAPI.GetConstitution)
	e.GET("/proposals", handlers.GovAPI.GetProposals)
	e.GET("/proposals/:proposalId", handlers.GovAPI.GetProposal)
	e.GET("/proposals/:proposalId/deposits", handlers.GovAPI.GetDeposits)
	e.GET("/proposals/:proposalId/deposits/:depositor", handlers.GovAPI.GetDeposit)
	e.GET("/proposals/:proposalId/tally", handlers.GovAPI.GetTally)
	e.GET("/proposals/:proposalId/votes", handlers.GovAPI.GetVotes)
	e.GET("/proposals/:proposalId/votes/:voter", handlers.GovAPI.GetVote)
}

// registerServiceRoutes registers the service routes
func registerServiceRoutes(e *echo.Echo) {
	e.GET("/service/:origin", handlers.ServiceAPI.QueryOrigin)
	e.GET("/service/:origin/login/:username/start", handlers.ServiceAPI.StartLogin)
	e.POST("/service/:origin/login/:username/finish", handlers.ServiceAPI.FinishLogin)
	e.GET("/service/:origin/register/:username/start", handlers.ServiceAPI.StartRegistration)
	e.POST("/service/:origin/register/:username/finish", handlers.ServiceAPI.FinishRegistration)
}

// registerStakingRoutes registers the staking routes
func registerStakingRoutes(e *echo.Echo) {
	e.GET("/delegators/:delegatorAddr", handlers.StakingAPI.GetDelegatorDelegations)
	e.GET("/delegators/:delegatorAddr/unbonding", handlers.StakingAPI.GetDelegatorUnbondingDelegations)
	e.GET("/delegators/:delegatorAddr/validators", handlers.StakingAPI.GetDelegatorValidators)
	e.GET("/delegators/:delegatorAddr/validators/:validatorAddr", handlers.StakingAPI.GetDelegation)
	e.GET("/delegators/:delegatorAddr/validators/:validatorAddr/unbonding", handlers.StakingAPI.GetUnbondingDelegation)
	e.GET("/delegators/:delegatorAddr/validators/:srcValidatorAddr/redelegate/:dstValidatorAddr", handlers.StakingAPI.GetRedelegations)
	e.GET("/history/{height}", handlers.StakingAPI.GetHistoricalInfo)
	e.GET("/staking", handlers.StakingAPI.GetValidators)
	e.GET("/staking/:validatorAddr", handlers.StakingAPI.GetValidator)
	e.GET("/staking/:validatorAddr/delegations", handlers.StakingAPI.GetValidatorDelegations)
}
