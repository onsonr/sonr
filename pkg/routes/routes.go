package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/handlers"
	"github.com/sonrhq/sonr/pkg/middleware"
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
	e.GET("/check/identifier/:id", handlers.AuthAPI.CheckIdentifier)
	e.GET("/service/:origin", handlers.ServiceAPI.QueryOrigin)
	e.GET("/service/:origin/login/:username/start", handlers.ServiceAPI.StartLogin)
	e.POST("/service/:origin/login/:username/finish", handlers.ServiceAPI.FinishLogin)
	e.GET("/service/:origin/register/:username/start", handlers.ServiceAPI.StartRegistration)
	e.POST("/service/:origin/register/:username/finish", handlers.ServiceAPI.FinishRegistration)
	e.GET("/tx/:txHash", handlers.TxAPI.GetTx)
	e.GET("/tx/block/:height", handlers.TxAPI.GetBlockWithTxs)
	e.POST("/tx/broadcast", handlers.TxAPI.BroadcastTx)
	e.POST("/tx/simulate", handlers.TxAPI.SimulateTx)
}

// RegisterStaticAssets registers the static asset routes
func RegisterStaticAssets(e *echo.Echo, assetsDir string) {
	e.Static("/*", assetsDir)
}

// RegisterHTMXPages registers the page routes for HTMX
func RegisterHTMXPages(e *echo.Echo) {
	e.GET("/", handlers.Pages.Index, middleware.UseHTMX)
	e.GET("/_panels/home", handlers.Pages.Home, middleware.UseHTMX)
	e.GET("/_panels/chat", handlers.Pages.Chat, middleware.UseHTMX)
	e.GET("/_panels/wallet", handlers.Pages.Wallet, middleware.UseHTMX)
	e.GET("/_panels/status", handlers.Pages.Status, middleware.UseHTMX)
	e.GET("/_panels/governance", handlers.Pages.Governance, middleware.UseHTMX)
	e.GET("/_panels/console", handlers.Pages.Console, middleware.UseHTMX)
	e.GET("/error-404", handlers.Pages.Error, middleware.UseHTMX)
}

// RegisterHTMXModals registers the modal routes for HTMX
func RegisterHTMXModals(e *echo.Echo) {
	e.GET("/swap", handlers.Modals.Swap, middleware.UseHTMX)
	e.GET("/deposit", handlers.Modals.Deposit, middleware.UseHTMX)
	e.GET("/settings", handlers.Modals.Settings, middleware.UseHTMX)
	e.GET("/share", handlers.Modals.Share, middleware.UseHTMX)
}
