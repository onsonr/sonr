package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	types "cosmossdk.io/api/cosmos/gov/v1"
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  API Endpoints                                 ||
// ! ||--------------------------------------------------------------------------------||

// GovAPI is a handler for the gov module
var GovAPI = govAPI{}

// govAPI is a handler for the gov module
type govAPI struct{}

// GetConstitution returns the constitution
func (h govAPI) GetConstitution(w http.ResponseWriter, r *http.Request) {
	res, err := middleware.GovClient(r, w).Constitution(r.Context(), &types.QueryConstitutionRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetProposal returns a proposal
func (h govAPI) GetProposal(w http.ResponseWriter, r *http.Request) {
	proposalIDStr := chi.URLParam(r, "proposalId")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
	res, err := middleware.GovClient(r, w).Proposal(r.Context(), &types.QueryProposalRequest{ProposalId: i})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetProposals returns all proposals
func (h govAPI) GetProposals(w http.ResponseWriter, r *http.Request) {
	res, err := middleware.GovClient(r, w).Proposals(r.Context(), &types.QueryProposalsRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetVote returns a vote
func (h govAPI) GetVote(w http.ResponseWriter, r *http.Request) {
	proposalIDStr := chi.URLParam(r, "proposalId")
	voterStr := chi.URLParam(r, "voter")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
	res, err := middleware.GovClient(r, w).Vote(r.Context(), &types.QueryVoteRequest{ProposalId: i, Voter: voterStr})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetVotes returns all votes for a proposal
func (h govAPI) GetVotes(w http.ResponseWriter, r *http.Request) {
	proposalIDStr := chi.URLParam(r, "proposalId")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
	res, err := middleware.GovClient(r, w).Votes(r.Context(), &types.QueryVotesRequest{ProposalId: i})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetDeposit returns a deposit
func (h govAPI) GetDeposit(w http.ResponseWriter, r *http.Request) {
	proposalIDStr := chi.URLParam(r, "proposalId")
	depositorStr := chi.URLParam(r, "depositor")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
	res, err := middleware.GovClient(r, w).Deposit(r.Context(), &types.QueryDepositRequest{ProposalId: i, Depositor: depositorStr})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetDeposits returns all deposits for a proposal
func (h govAPI) GetDeposits(w http.ResponseWriter, r *http.Request) {
	proposalIDStr := chi.URLParam(r, "proposalId")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
	res, err := middleware.GovClient(r, w).Deposits(r.Context(), &types.QueryDepositsRequest{ProposalId: i})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetTally returns the tally for a proposal
func (h govAPI) GetTally(w http.ResponseWriter, r *http.Request) {
	proposalIDStr := chi.URLParam(r, "proposalId")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
	res, err := middleware.GovClient(r, w).TallyResult(r.Context(), &types.QueryTallyResultRequest{ProposalId: i})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// RegisterRoutes registers the gov routes
func (h govAPI) RegisterRoutes(r chi.Router) {
	r.Get("/constitution", h.GetConstitution)
	r.Get("/proposals", h.GetProposals)
	r.Get("/proposals/{proposalId}", h.GetProposal)
	r.Get("/proposals/{proposalId}/deposits", h.GetDeposits)
	r.Get("/proposals/{proposalId}/deposits/{depositor}", h.GetDeposit)
	r.Get("/proposals/{proposalId}/tally", h.GetTally)
	r.Get("/proposals/{proposalId}/votes", h.GetVotes)
	r.Get("/proposals/{proposalId}/votes/{voter}", h.GetVote)
}
