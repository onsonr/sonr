package handlers

import (
	"net/http"
	"strconv"

	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	"github.com/labstack/echo/v4"
	shared "github.com/sonrhq/sonr/pkg/middleware/common"
)

// StakingAPI is a handler for the staking module
var StakingAPI = stakingAPI{}

// stakingAPI is a handler for the staking module
type stakingAPI struct{}

// GetDelegation returns a delegation
func (h stakingAPI) GetDelegation(c echo.Context) error {
	delegatorAddr := c.Param("delegatorAddr")
	validatorAddr := c.Param("validatorAddr")
	resp, err := shared.Clients(c).Staking().Delegation(c.Request().Context(), &stakingv1beta1.QueryDelegationRequest{
		DelegatorAddr: delegatorAddr,
		ValidatorAddr: validatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetUnbondingDelegation returns an unbonding delegation
func (h stakingAPI) GetUnbondingDelegation(c echo.Context) error {
	delegatorAddr := c.Param("delegatorAddr")
	validatorAddr := c.Param("validatorAddr")
	resp, err := shared.Clients(c).Staking().UnbondingDelegation(c.Request().Context(), &stakingv1beta1.QueryUnbondingDelegationRequest{
		DelegatorAddr: delegatorAddr,
		ValidatorAddr: validatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetDelegatorDelegations returns all delegations for a delegator
func (h stakingAPI) GetDelegatorDelegations(c echo.Context) error {
	delegatorAddr := c.Param("delegatorAddr")
	resp, err := shared.Clients(c).Staking().DelegatorDelegations(c.Request().Context(), &stakingv1beta1.QueryDelegatorDelegationsRequest{
		DelegatorAddr: delegatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetDelegatorUnbondingDelegations returns all unbonding delegations for a delegator
func (h stakingAPI) GetDelegatorUnbondingDelegations(c echo.Context) error {
	delegatorAddr := c.Param("delegatorAddr")
	resp, err := shared.Clients(c).Staking().DelegatorUnbondingDelegations(c.Request().Context(), &stakingv1beta1.QueryDelegatorUnbondingDelegationsRequest{
		DelegatorAddr: delegatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetRedelegations returns all redelegations for a delegator
func (h stakingAPI) GetRedelegations(c echo.Context) error {
	delegatorAddr := c.Param("delegatorAddr")
	srcValidatorAddr := c.Param("srcValidatorAddr")
	dstValidatorAddr := c.Param("dstValidatorAddr")
	resp, err := shared.Clients(c).Staking().Redelegations(c.Request().Context(), &stakingv1beta1.QueryRedelegationsRequest{
		DelegatorAddr:    delegatorAddr,
		SrcValidatorAddr: srcValidatorAddr,
		DstValidatorAddr: dstValidatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetValidator returns a validator
func (h stakingAPI) GetValidator(c echo.Context) error {
	validatorAddr := c.Param("validatorAddr")
	resp, err := shared.Clients(c).Staking().Validator(c.Request().Context(), &stakingv1beta1.QueryValidatorRequest{
		ValidatorAddr: validatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetValidators returns all validators
func (h stakingAPI) GetValidators(c echo.Context) error {
	resp, err := shared.Clients(c).Staking().Validators(c.Request().Context(), &stakingv1beta1.QueryValidatorsRequest{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetValidatorDelegations returns all delegations for a validator
func (h stakingAPI) GetValidatorDelegations(c echo.Context) error {
	validatorAddr := c.Param("validatorAddr")
	resp, err := shared.Clients(c).Staking().ValidatorDelegations(c.Request().Context(), &stakingv1beta1.QueryValidatorDelegationsRequest{
		ValidatorAddr: validatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetDelegatorValidators returns all validators for a delegator
func (h stakingAPI) GetDelegatorValidators(c echo.Context) error {
	delegatorAddr := c.Param("delegatorAddr")
	resp, err := shared.Clients(c).Staking().DelegatorValidators(c.Request().Context(), &stakingv1beta1.QueryDelegatorValidatorsRequest{
		DelegatorAddr: delegatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetHistoricalInfo returns historical info
func (h stakingAPI) GetHistoricalInfo(c echo.Context) error {
	heightStr := c.Param("height")
	height, _ := strconv.ParseInt(heightStr, 10, 64)
	resp, err := shared.Clients(c).Staking().HistoricalInfo(c.Request().Context(), &stakingv1beta1.QueryHistoricalInfoRequest{
		Height: height,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
