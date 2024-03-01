package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/middleware/shared"
	"github.com/sonrhq/sonr/x/service"
)

// ServiceAPI is a handler for the staking module
var ServiceAPI = serviceAPI{}

// serviceAPI is a handler for the staking module
type serviceAPI struct{}

// QueryOrigin returns the service for the origin host
func (h serviceAPI) QueryOrigin(c echo.Context) error {
	origin := c.Param("origin")
	resp, err := service.GetRecordByOrigin(shared.GrpcClientConn(c), origin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// QueryCredentials returns the service credentials for a user of the origin host
func (h serviceAPI) QueryCredentials(c echo.Context) error {
	origin := c.Param("origin")
	handle := c.Param("handle")
	resp, err := service.GetCredentialsByHandle(shared.GrpcClientConn(c), handle, origin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// StartRegistration returns credential creation options for the origin host
func (h serviceAPI) StartRegistration(c echo.Context) error {
	handleStr := c.Param("handle")
	origin := c.Param("origin")
	resp, err := service.GetRecordByOrigin(shared.GrpcClientConn(c), origin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	opts := service.GetCredentialCreationOptions(resp, protocol.UserEntity{
		DisplayName: handleStr,
		ID:          []byte("vc.Address"),
	})
	return c.JSON(http.StatusOK, opts)
}

// FinishRegistration returns the result of the credential creation
func (h serviceAPI) FinishRegistration(c echo.Context) error {
	origin := c.Param("origin")
	resp, err := service.GetRecordByOrigin(shared.GrpcClientConn(c), origin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var credential protocol.PublicKeyCredential
	if err := json.NewDecoder(c.Request().Body).Decode(&credential); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

// StartLogin returns credential request options for the origin host
func (h serviceAPI) StartLogin(c echo.Context) error {
	origin := c.Param("origin")
	resp, err := service.GetRecordByOrigin(shared.GrpcClientConn(c), origin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	opts := service.GetCredentialRequestOptions(resp, []protocol.CredentialDescriptor{})
	return c.JSON(http.StatusOK, opts)
}

// FinishLogin returns the result of the credential request
func (h serviceAPI) FinishLogin(c echo.Context) error {
	origin := c.Param("origin")
	resp, err := service.GetRecordByOrigin(shared.GrpcClientConn(c), origin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var credential protocol.PublicKeyCredential
	if err := json.NewDecoder(c.Request().Body).Decode(&credential); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
