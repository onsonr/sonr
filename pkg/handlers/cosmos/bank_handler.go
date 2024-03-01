package handlers

import (
	"net/http"

	types "cosmossdk.io/api/cosmos/bank/v1beta1"
	"github.com/labstack/echo/v4"
	"github.com/sonrhq/sonr/pkg/middleware/shared"
)

// BankAPI is a handler for the bank module
var BankAPI = bankAPI{}

// bankAPI is a handler for the bank module
type bankAPI struct{}

// GetAllBalances returns all balances for an address
func (h bankAPI) GetAllBalances(c echo.Context) error {
	address := c.Param("address")
	resp, err := shared.Clients(c).Bank().AllBalances(c.Request().Context(), &types.QueryAllBalancesRequest{
		Address: address,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetBalance returns a balance for an address and denom
func (h bankAPI) GetBalance(c echo.Context) error {
	address := c.Param("address")
	denom := c.Param("denom")
	resp, err := shared.Clients(c).Bank().Balance(c.Request().Context(), &types.QueryBalanceRequest{Address: address, Denom: denom})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetTotalSupply returns the total supply
func (h bankAPI) GetTotalSupply(c echo.Context) error {
	resp, err := shared.Clients(c).Bank().TotalSupply(c.Request().Context(), &types.QueryTotalSupplyRequest{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetSupplyOf returns the supply of a denom
func (h bankAPI) GetSupplyOf(c echo.Context) error {
	denom := c.Param("denom")
	resp, err := shared.Clients(c).Bank().SupplyOf(c.Request().Context(), &types.QuerySupplyOfRequest{Denom: denom})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetSpendableBalances returns the spendable balances for an address
func (h bankAPI) GetSpendableBalances(c echo.Context) error {
	address := c.Param("address")
	resp, err := shared.Clients(c).Bank().SpendableBalances(c.Request().Context(), &types.QuerySpendableBalancesRequest{Address: address})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetSpendableBalancesByDenom returns the spendable balances for an address and denom
func (h bankAPI) GetSpendableBalancesByDenom(c echo.Context) error {
	address := c.Param("address")
	denom := c.Param("denom")
	resp, err := shared.Clients(c).Bank().SpendableBalanceByDenom(c.Request().Context(), &types.QuerySpendableBalanceByDenomRequest{Address: address, Denom: denom})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
