package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  API Endpoints                                 ||
// ! ||--------------------------------------------------------------------------------||

// CometAPI is a handler for the node module
var CometAPI = cometAPI{}

// cometAPI is a handler for the node module
type cometAPI struct{}

// GetLatestBlock returns the latest block
func (h cometAPI) GetLatestBlock(e echo.Context) error {
	res, err := middleware.CometClient(e).GetLatestBlock(e.Request().Context(), &cmtservice.GetLatestBlockRequest{})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, rBz)
}

// GetNodeInfo returns the node info
func (h cometAPI) GetNodeInfo(e echo.Context) error {
	res, err := middleware.CometClient(e).GetNodeInfo(e.Request().Context(), &cmtservice.GetNodeInfoRequest{})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, rBz)
}

// GetSyncing returns the syncing status
func (h cometAPI) GetSyncing(e echo.Context) error {
	res, err := middleware.CometClient(e).GetSyncing(e.Request().Context(), &cmtservice.GetSyncingRequest{})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, rBz)
}

// GetBlockByHeight returns a block by height
func (h cometAPI) GetBlockByHeight(e echo.Context) error {
	heightStr := e.Param("height")
	i, _ := strconv.ParseInt(heightStr, 10, 64)
	res, err := middleware.CometClient(e).GetBlockByHeight(e.Request().Context(), &cmtservice.GetBlockByHeightRequest{Height: i})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, rBz)
}

// GetLatestValidatorSet returns the latest validator set
func (h cometAPI) GetLatestValidatorSet(e echo.Context) error {
	res, err := middleware.CometClient(e).GetLatestValidatorSet(e.Request().Context(), &cmtservice.GetLatestValidatorSetRequest{})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, rBz)
}

// GetValidatorSetByHeight returns a validator set by height
func (h cometAPI) GetValidatorSetByHeight(e echo.Context) error {
	heightStr := e.Param("height")
	i, _ := strconv.ParseInt(heightStr, 10, 64)
	res, err := middleware.CometClient(e).GetValidatorSetByHeight(e.Request().Context(), &cmtservice.GetValidatorSetByHeightRequest{Height: i})
	if err != nil {
		e.JSON(http.StatusInternalServerError, err.Error())
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, rBz)
}

// RegisterRoutes registers the node routes
func (h cometAPI) RegisterRoutes(e *echo.Echo) {
	e.GET("/block", h.GetLatestBlock)
	e.GET("/block/:height", h.GetBlockByHeight)
	e.GET("/health", h.GetNodeInfo)
	e.GET("/syncing", h.GetSyncing)
	e.GET("/validatorsets/latest", h.GetLatestValidatorSet)
	e.GET("/validatorsets/:height", h.GetValidatorSetByHeight)
}
