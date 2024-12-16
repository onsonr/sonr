package middleware

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common"
)

// CurrentBlock returns the current block
func CurrentBlock(c echo.Context) uint64 {
	cc, ok := c.(*GatewayContext)
	if !ok {
		return 0
	}
	qc, err := common.NewNodeClient(cc.grpcAddr)
	if err != nil {
		return 0
	}
	resp, err := qc.Status(bgCtx(), &common.StatusRequest{})
	if err != nil {
		return 0
	}
	return resp.GetHeight()
}

// GetBankParams returns the bank params
func GetBankParams(c echo.Context) (*common.BankParamsResponse, error) {
	cc, ok := c.(*GatewayContext)
	if !ok {
		return nil, errors.New("gateway context not found")
	}
	cl, err := common.NewBankClient(cc.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := cl.Params(bgCtx(), &common.BankParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDIDParams returns the DID params
func GetDIDParams(c echo.Context) (*common.DIDParamsResponse, error) {
	cc, ok := c.(*GatewayContext)
	if !ok {
		return nil, errors.New("gateway context not found")
	}
	cl, err := common.NewDIDClient(cc.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := cl.Params(bgCtx(), &common.DIDParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDWNParams returns the DWN params
func GetDWNParams(c echo.Context) (*common.DWNParamsResponse, error) {
	cc, ok := c.(*GatewayContext)
	if !ok {
		return nil, errors.New("gateway context not found")
	}

	cl, err := common.NewDWNClient(cc.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := cl.Params(bgCtx(), &common.DWNParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetNodeStatus returns the node status
func GetNodeStatus(c echo.Context) (*common.StatusResponse, error) {
	cc, ok := c.(*GatewayContext)
	if !ok {
		return nil, errors.New("gateway context not found")
	}

	cl, err := common.NewNodeClient(cc.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := cl.Status(bgCtx(), &common.StatusRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSVCParams returns the SVC params
func GetSVCParams(c echo.Context) (*common.SVCParamsResponse, error) {
	cc, ok := c.(*GatewayContext)
	if !ok {
		return nil, errors.New("gateway context not found")
	}

	cl, err := common.NewSVCClient(cc.grpcAddr)
	if err != nil {
		return nil, err
	}
	resp, err := cl.Params(bgCtx(), &common.SVCParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
