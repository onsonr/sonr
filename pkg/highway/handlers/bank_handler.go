package handlers

import (
	"encoding/json"
	"net/http"

	types "cosmossdk.io/api/cosmos/bank/v1beta1"
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  API Endpoints                                 ||
// ! ||--------------------------------------------------------------------------------||

// BankAPI is a handler for the bank module
var BankAPI = bankAPI{}

// bankAPI is a handler for the bank module
type bankAPI struct{}

// GetAllBalances returns all balances for an address
func (h bankAPI) GetAllBalances(c echo.Context) error {
	address := c.Param("address")
	resp, err := middleware.BankClient(c).AllBalances(c.Request().Context(), &types.QueryAllBalancesRequest{
		Address: address,
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

// GetBalance returns a balance for an address and denom
func (h bankAPI) GetBalance(c echo.Context) error {
	address := c.Param("address")
	denom := c.Param("denom")
	resp, err := middleware.BankClient(c).Balance(c.Request().Context(), &types.QueryBalanceRequest{Address: address, Denom: denom})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetTotalSupply returns the total supply
func (h bankAPI) GetTotalSupply(c echo.Context) error {
	resp, err := middleware.BankClient(c).TotalSupply(c.Request().Context(), &types.QueryTotalSupplyRequest{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetSupplyOf returns the supply of a denom
func (h bankAPI) GetSupplyOf(c echo.Context) error {
	denom := c.Param("denom")
	resp, err := middleware.BankClient(c).SupplyOf(c.Request().Context(), &types.QuerySupplyOfRequest{Denom: denom})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetSpendableBalances returns the spendable balances for an address
func (h bankAPI) GetSpendableBalances(c echo.Context) error {
	address := c.Param("address")
	resp, err := middleware.BankClient(c).SpendableBalances(c.Request().Context(), &types.QuerySpendableBalancesRequest{Address: address})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// GetSpendableBalancesByDenom returns the spendable balances for an address and denom
func (h bankAPI) GetSpendableBalancesByDenom(c echo.Context) error {
	address := c.Param("address")
	denom := c.Param("denom")
	resp, err := middleware.BankClient(c).SpendableBalanceByDenom(c.Request().Context(), &types.QuerySpendableBalanceByDenomRequest{Address: address, Denom: denom})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rBz)
}

// RegisterRoutes registers the bank routes
func (h bankAPI) RegisterRoutes(e *echo.Echo) {
	e.GET("/balance/:address", h.GetAllBalances)
	e.GET("/balance/:address/spendable", h.GetSpendableBalances)
	e.GET("/balance/:address/:denom", h.GetBalance)
	e.GET("/balance/:address/:denom/spendable", h.GetSpendableBalancesByDenom)
	e.GET("/supply", h.GetTotalSupply)
	e.GET("/supply/:denom", h.GetSupplyOf)
}
