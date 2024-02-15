package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	types "cosmossdk.io/api/cosmos/gov/v1"
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// GovAPI is a handler for the gov module
var GovAPI = govAPI{}

// govAPI is a handler for the gov module
type govAPI struct{}

// GetConstitution returns the constitution
func (h govAPI) GetConstitution(c echo.Context) error {
	res, err := middleware.GovClient(c).Constitution(c.Request().Context(), &types.QueryConstitutionRequest{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetProposal returns a proposal
func (h govAPI) GetProposal(c echo.Context) error {
	proposalIDStr := c.Param("proposalId")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
	res, err := middleware.GovClient(c).Proposal(c.Request().Context(), &types.QueryProposalRequest{ProposalId: i})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetProposals returns all proposals
func (h govAPI) GetProposals(c echo.Context) error {
	res, err := middleware.GovClient(c).Proposals(c.Request().Context(), &types.QueryProposalsRequest{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetVote returns a vote
func (h govAPI) GetVote(c echo.Context) error {
	proposalIDStr := c.Param("proposalId")
	voterStr := c.Param("voter")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
	res, err := middleware.GovClient(c).Vote(c.Request().Context(), &types.QueryVoteRequest{ProposalId: i, Voter: voterStr})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetVotes returns all votes for a proposal
func (h govAPI) GetVotes(c echo.Context) error {
	proposalIDStr := c.Param("proposalId")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
	res, err := middleware.GovClient(c).Votes(c.Request().Context(), &types.QueryVotesRequest{ProposalId: i})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetDeposit returns a deposit
func (h govAPI) GetDeposit(c echo.Context) error {
	proposalIDStr := c.Param("proposalId")
	depositorStr := c.Param("depositor")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
	res, err := middleware.GovClient(c).Deposit(c.Request().Context(), &types.QueryDepositRequest{ProposalId: i, Depositor: depositorStr})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetDeposits returns all deposits for a proposal
func (h govAPI) GetDeposits(c echo.Context) error {
	proposalIDStr := c.Param("proposalId")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
	res, err := middleware.GovClient(c).Deposits(c.Request().Context(), &types.QueryDepositsRequest{ProposalId: i})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetTally returns the tally for a proposal
func (h govAPI) GetTally(c echo.Context) error {
	proposalIDStr := c.Param("proposalId")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
	res, err := middleware.GovClient(c).TallyResult(c.Request().Context(), &types.QueryTallyResultRequest{ProposalId: i})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}
