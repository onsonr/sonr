package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/app/gateway/handlers/htmx"
	modulesapi "github.com/sonrhq/sonr/app/gateway/handlers/modules"
)

// Mount all routes to the router
func Mount(r chi.Router) {
	r.Mount(consoleEndpoints())
	r.Mount(walletEndpoints())
	r.Mount(apiEndpoints())
	r.Mount(sseEndpoints())
}

func consoleEndpoints() (string, chi.Router) {
	r := chi.NewRouter()
	consoleHandler := htmx.ConsoleHandler{}
	r.Get("/", consoleHandler.IndexPage)
	return "/", r
}

func walletEndpoints() (string, chi.Router) {
	r := chi.NewRouter()
	dashHandler := htmx.WalletHandler{}
	r.Get("/", dashHandler.IndexPage)
	return "/wallet", r
}

func apiEndpoints() (string, chi.Router) {
	r := chi.NewRouter()
	bankHandler := modulesapi.BankHandler{}
	govHandler := modulesapi.GovHandler{}
	nodeHandler := modulesapi.NodeHandler{}
	stakeHandler := modulesapi.StakingHandler{}

	// Node endpoints
	r.Get("/balance/{address}", bankHandler.GetAllBalances)
	r.Get("/balance/{address}/spendable", bankHandler.GetSpendableBalances)
	r.Get("/balance/{address}/{denom}", bankHandler.GetBalance)
	r.Get("/balance/{address}/{denom}/spendable", bankHandler.GetSpendableBalancesByDenom)
	r.Get("/block", nodeHandler.GetLatestBlock)
	r.Get("/block/{height}", nodeHandler.GetBlockByHeight)
	r.Get("/constitution", govHandler.GetConstitution)
	r.Get("/delegators/{delegatorAddr}", stakeHandler.GetDelegatorDelegations)
	r.Get("/delegators/{delegatorAddr}/unbonding", stakeHandler.GetDelegatorUnbondingDelegations)
	r.Get("/delegators/{delegatorAddr}/validators", stakeHandler.GetDelegatorValidators)
	r.Get("/delegators/{delegatorAddr}/validators/{validatorAddr}", stakeHandler.GetDelegation)
	r.Get("/delegators/{delegatorAddr}/validators/{validatorAddr}/unbonding", stakeHandler.GetUnbondingDelegation)
	r.Get("/delegators/{delegatorAddr}/validators/{srcValidatorAddr}/redelegate/{dstValidatorAddr}", stakeHandler.GetRedelegations)
	r.Get("/health", nodeHandler.GetNodeInfo)
	r.Get("/history/{height}", stakeHandler.GetHistoricalInfo)
	r.Get("/proposals", govHandler.GetProposals)
	r.Get("/proposals/{proposalId}", govHandler.GetProposal)
	r.Get("/proposals/{proposalId}/deposits", govHandler.GetDeposits)
	r.Get("/proposals/{proposalId}/deposits/{depositor}", govHandler.GetDeposit)
	r.Get("/proposals/{proposalId}/tally", govHandler.GetTally)
	r.Get("/proposals/{proposalId}/votes", govHandler.GetVotes)
	r.Get("/proposals/{proposalId}/votes/{voter}", govHandler.GetVote)
	r.Get("/staking/{validatorAddr}", stakeHandler.GetValidator)
	r.Get("/supply", bankHandler.GetTotalSupply)
	r.Get("/supply/{denom}", bankHandler.GetSupplyOf)
	r.Get("/syncing", nodeHandler.GetSyncing)
	r.Get("/validators", stakeHandler.GetValidators)
	r.Get("/validators/{validatorAddr}", stakeHandler.GetValidator)
	r.Get("/validators/{validatorAddr}/delegations", stakeHandler.GetValidatorDelegations)
	// Final endpoint
	return "/api", r
}

func sseEndpoints() (string, chi.Router) {
	r := chi.NewRouter()
	// moduleHandler := htmx.ModuleHandler{}
	// r.Get("/", moduleHandler.IndexPage)
	return "/events", r
}
