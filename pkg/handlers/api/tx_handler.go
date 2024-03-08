package handlers

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/labstack/echo/v4"
<<<<<<< HEAD:pkg/handlers/tx_handler.go

=======
>>>>>>> master:pkg/handlers/api/tx_handler.go
	shared "github.com/sonrhq/sonr/pkg/middleware/common"
)

// TxAPI is a handler for the staking module
var TxAPI = txAPI{}

// txAPI is a handler for the staking module
type txAPI struct{}

// GetTx returns a transaction by hash
func (h txAPI) GetTx(c echo.Context) error {
<<<<<<< HEAD:pkg/handlers/tx_handler.go
	txHash := c.Param("hash")
=======
	txHash := c.Param("txHash")
>>>>>>> master:pkg/handlers/api/tx_handler.go
	resp, err := shared.Clients(c).Tx().GetTx(c.Request().Context(), &tx.GetTxRequest{
		Hash: txHash,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// GetBlockWithTxs returns a block with transactions
func (h txAPI) GetBlockWithTxs(c echo.Context) error {
	height, err := strconv.ParseInt(c.Param("height"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resp, err := shared.Clients(c).Tx().GetBlockWithTxs(c.Request().Context(), &tx.GetBlockWithTxsRequest{
		Height: height,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// BroadcastTx broadcasts a transaction
func (h txAPI) BroadcastTx(c echo.Context) error {
	var req tx.BroadcastTxRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resp, err := shared.Clients(c).Tx().BroadcastTx(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// SimulateTx simulates a transaction
func (h txAPI) SimulateTx(c echo.Context) error {
	var req tx.SimulateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resp, err := shared.Clients(c).Tx().Simulate(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
