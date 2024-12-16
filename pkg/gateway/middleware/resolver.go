package middleware

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common"
)

var grpcEndpoint = ""

// CurrentBlock returns the current block
func CurrentBlock(c echo.Context) uint64 {
	if grpcEndpoint == "" {
		return 0
	}
	qc, err := common.NewNodeClient(grpcEndpoint)
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
func GetBankParams() (*common.BankParamsResponse, error) {
	if grpcEndpoint == "" {
		return nil, errors.New("grpc endpoint not set")
	}
	c, err := common.NewBankClient(grpcEndpoint)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(bgCtx(), &common.BankParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDIDParams returns the DID params
func GetDIDParams() (*common.DIDParamsResponse, error) {
	if grpcEndpoint == "" {
		return nil, errors.New("grpc endpoint not set")
	}
	c, err := common.NewDIDClient(grpcEndpoint)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(bgCtx(), &common.DIDParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDWNParams returns the DWN params
func GetDWNParams() (*common.DWNParamsResponse, error) {
	if grpcEndpoint == "" {
		return nil, errors.New("grpc endpoint not set")
	}
	c, err := common.NewDWNClient(grpcEndpoint)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(bgCtx(), &common.DWNParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetNodeStatus returns the node status
func GetNodeStatus() (*common.StatusResponse, error) {
	if grpcEndpoint == "" {
		return nil, errors.New("grpc endpoint not set")
	}
	c, err := common.NewNodeClient(grpcEndpoint)
	if err != nil {
		return nil, err
	}
	resp, err := c.Status(bgCtx(), &common.StatusRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSVCParams returns the SVC params
func GetSVCParams() (*common.SVCParamsResponse, error) {
	if grpcEndpoint == "" {
		return nil, errors.New("grpc endpoint not set")
	}
	c, err := common.NewSVCClient(grpcEndpoint)
	if err != nil {
		return nil, err
	}
	resp, err := c.Params(bgCtx(), &common.SVCParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
