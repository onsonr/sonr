package routes

import (
	"github.com/labstack/echo/v4"

	handlers "github.com/sonrhq/sonr/pkg/handlers"
	"github.com/sonrhq/sonr/pkg/middleware/common"
)

// RegisterCosmosAPI registers the Cosmos API routes
func RegisterCosmosAPI(e *echo.Echo) {
	e.GET("/balance/:address", handlers.BankAPI.GetAllBalances)
	e.GET("/balance/:address/:denom", handlers.BankAPI.GetBalance)
	e.GET("/balance/:address/:denom/spendable", handlers.BankAPI.GetSpendableBalancesByDenom)
	e.GET("/balance/:address/spendable", handlers.BankAPI.GetSpendableBalances)
	e.GET("/block", handlers.CometAPI.GetLatestBlock)
	e.GET("/block/:height", handlers.CometAPI.GetBlockByHeight)
	e.GET("/constitution", handlers.GovAPI.GetConstitution)
	e.GET("/delegators/:delegatorAddr", handlers.StakingAPI.GetDelegatorDelegations)
	e.GET("/delegators/:delegatorAddr/validators", handlers.StakingAPI.GetDelegatorValidators)
	e.GET("/delegators/:delegatorAddr/validators/:validatorAddr", handlers.StakingAPI.GetDelegation)
	e.GET("/delegators/:delegatorAddr/validators/:validatorAddr/unbonding", handlers.StakingAPI.GetUnbondingDelegation)
	e.GET("/delegators/:delegatorAddr/validators/:srcValidatorAddr/redelegate/:dstValidatorAddr", handlers.StakingAPI.GetRedelegations)
	e.GET("/delegators/:delegatorAddr/unbonding", handlers.StakingAPI.GetDelegatorUnbondingDelegations)
	e.GET("/health", handlers.CometAPI.GetNodeInfo)
	e.GET("/history/{height}", handlers.StakingAPI.GetHistoricalInfo)
	e.GET("/proposals", handlers.GovAPI.GetProposals)
	e.GET("/proposals/:proposalId", handlers.GovAPI.GetProposal)
	e.GET("/proposals/:proposalId/deposits", handlers.GovAPI.GetDeposits)
	e.GET("/proposals/:proposalId/deposits/:depositor", handlers.GovAPI.GetDeposit)
	e.GET("/proposals/:proposalId/tally", handlers.GovAPI.GetTally)
	e.GET("/proposals/:proposalId/votes", handlers.GovAPI.GetVotes)
	e.GET("/proposals/:proposalId/votes/:voter", handlers.GovAPI.GetVote)
	e.GET("/staking", handlers.StakingAPI.GetValidators)
	e.GET("/staking/:validatorAddr", handlers.StakingAPI.GetValidator)
	e.GET("/staking/:validatorAddr/delegations", handlers.StakingAPI.GetValidatorDelegations)
	e.GET("/supply", handlers.BankAPI.GetTotalSupply)
	e.GET("/supply/:denom", handlers.BankAPI.GetSupplyOf)
	e.GET("/syncing", handlers.CometAPI.GetSyncing)
	e.GET("/validators", handlers.CometAPI.GetLatestValidatorSet)
	e.GET("/validators/:height", handlers.CometAPI.GetValidatorSetByHeight)
}

// RegisterSonrAPI registers the Sonr API routes
func RegisterSonrAPI(e *echo.Echo) {
	e.GET("/service/:origin", handlers.ServiceAPI.QueryOrigin)
	e.GET("/service/:origin/email/:email/confirm", handlers.ServiceAPI.VerifyEmail)
	e.POST("/service/:origin/email/:email/verify", handlers.ServiceAPI.SendVerificationEmail)
	e.GET("/service/:origin/login/:username/start", handlers.ServiceAPI.StartLogin)
	e.POST("/service/:origin/login/:username/finish", handlers.ServiceAPI.FinishLogin)
	e.GET("/service/:origin/register/:username/start", handlers.ServiceAPI.StartRegistration)
	e.POST("/service/:origin/register/:username/finish", handlers.ServiceAPI.FinishRegistration)
	e.GET("/tx/:hash", handlers.TxAPI.GetTx)
	e.GET("/tx/block/:height", handlers.TxAPI.GetBlockWithTxs)
	e.POST("/tx/broadcast", handlers.TxAPI.BroadcastTx)
	e.POST("/tx/simulate", handlers.TxAPI.SimulateTx)
}

// RegisterUI registers the page routes for HTMX
func RegisterUI(e *echo.Echo) {
	e.GET("/", handlers.Pages.Home, common.UseHTMX)
	e.GET("/about", handlers.Pages.About, common.UseHTMX)
	e.GET("/ecosystem", handlers.Pages.Ecosystem, common.UseHTMX)
	e.GET("/research", handlers.Pages.Research, common.UseHTMX)
	e.GET("/login", handlers.Pages.Login, common.UseHTMX)
	e.GET("/register", handlers.Pages.Register, common.UseHTMX)
	e.GET("/chats", handlers.Pages.Chats, common.UseHTMX)
	e.GET("/changelog", handlers.Pages.Changelog, common.UseHTMX)
	e.GET("*", handlers.Pages.Root, common.UseHTMX)
}
