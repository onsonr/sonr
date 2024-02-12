package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  API Endpoints                                 ||
// ! ||--------------------------------------------------------------------------------||

// StakingAPI is a handler for the staking module
var StakingAPI = stakingAPI{}

// stakingAPI is a handler for the staking module
type stakingAPI struct{}

// GetDelegation returns a delegation
func (h stakingAPI) GetDelegation(c echo.Context) error {
	delegatorAddr := c.Param("delegatorAddr")
	validatorAddr := c.Param("validatorAddr")
	resp, err := middleware.StakingClient(c).Delegation(c.Request().Context(), &stakingv1beta1.QueryDelegationRequest{
		DelegatorAddr: delegatorAddr,
		ValidatorAddr: validatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetUnbondingDelegation returns an unbonding delegation
func (h stakingAPI) GetUnbondingDelegation(c echo.Context) error {
	delegatorAddr := c.Param("delegatorAddr")
	validatorAddr := c.Param("validatorAddr")
	resp, err := middleware.StakingClient(c).UnbondingDelegation(c.Request().Context(), &stakingv1beta1.QueryUnbondingDelegationRequest{
		DelegatorAddr: delegatorAddr,
		ValidatorAddr: validatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetDelegatorDelegations returns all delegations for a delegator
func (h stakingAPI) GetDelegatorDelegations(c echo.Context) error {
	delegatorAddr := c.Param("delegatorAddr")
	resp, err := middleware.StakingClient(c).DelegatorDelegations(c.Request().Context(), &stakingv1beta1.QueryDelegatorDelegationsRequest{
		DelegatorAddr: delegatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetDelegatorUnbondingDelegations returns all unbonding delegations for a delegator
func (h stakingAPI) GetDelegatorUnbondingDelegations(c echo.Context) error {
	delegatorAddr := c.Param("delegatorAddr")
	resp, err := middleware.StakingClient(c).DelegatorUnbondingDelegations(c.Request().Context(), &stakingv1beta1.QueryDelegatorUnbondingDelegationsRequest{
		DelegatorAddr: delegatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetRedelegations returns all redelegations for a delegator
func (h stakingAPI) GetRedelegations(c echo.Context) error {
	delegatorAddr := c.Param("delegatorAddr")
	srcValidatorAddr := c.Param("srcValidatorAddr")
	dstValidatorAddr := c.Param("dstValidatorAddr")
	resp, err := middleware.StakingClient(c).Redelegations(c.Request().Context(), &stakingv1beta1.QueryRedelegationsRequest{
		DelegatorAddr:    delegatorAddr,
		SrcValidatorAddr: srcValidatorAddr,
		DstValidatorAddr: dstValidatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetValidator returns a validator
func (h stakingAPI) GetValidator(c echo.Context) error {
	validatorAddr := c.Param("validatorAddr")
	resp, err := middleware.StakingClient(c).Validator(c.Request().Context(), &stakingv1beta1.QueryValidatorRequest{
		ValidatorAddr: validatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetValidators returns all validators
func (h stakingAPI) GetValidators(c echo.Context) error {
	resp, err := middleware.StakingClient(c).Validators(c.Request().Context(), &stakingv1beta1.QueryValidatorsRequest{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetValidatorDelegations returns all delegations for a validator
func (h stakingAPI) GetValidatorDelegations(c echo.Context) error {
	validatorAddr := c.Param("validatorAddr")
	resp, err := middleware.StakingClient(c).ValidatorDelegations(c.Request().Context(), &stakingv1beta1.QueryValidatorDelegationsRequest{
		ValidatorAddr: validatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetDelegatorValidators returns all validators for a delegator
func (h stakingAPI) GetDelegatorValidators(c echo.Context) error {
	delegatorAddr := c.Param("delegatorAddr")
	resp, err := middleware.StakingClient(c).DelegatorValidators(c.Request().Context(), &stakingv1beta1.QueryDelegatorValidatorsRequest{
		DelegatorAddr: delegatorAddr,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetHistoricalInfo returns historical info
func (h stakingAPI) GetHistoricalInfo(c echo.Context) error {
	heightStr := c.Param("height")
	height, _ := strconv.ParseInt(heightStr, 10, 64)
	resp, err := middleware.StakingClient(c).HistoricalInfo(c.Request().Context(), &stakingv1beta1.QueryHistoricalInfoRequest{
		Height: height,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// RegisterRoutes registers the staking routes
func (h stakingAPI) RegisterRoutes(e *echo.Echo) {
	e.GET("/delegators/:delegatorAddr", h.GetDelegatorDelegations)
	e.GET("/delegators/:delegatorAddr/unbonding", h.GetDelegatorUnbondingDelegations)
	e.GET("/delegators/:delegatorAddr/validators", h.GetDelegatorValidators)
	e.GET("/delegators/:delegatorAddr/validators/:validatorAddr", h.GetDelegation)
	e.GET("/delegators/:delegatorAddr/validators/:validatorAddr/unbonding", h.GetUnbondingDelegation)
	e.GET("/delegators/:delegatorAddr/validators/:srcValidatorAddr/redelegate/:dstValidatorAddr", h.GetRedelegations)
	e.GET("/history/{height}", h.GetHistoricalInfo)
	e.GET("/validators", h.GetValidators)
	e.GET("/validators/:validatorAddr", h.GetValidator)
	e.GET("/validators/:validatorAddr/delegations", h.GetValidatorDelegations)
}
