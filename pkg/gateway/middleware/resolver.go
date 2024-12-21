package middleware

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common"
)

// ParamsBank returns the bank params
func ParamsBank(c echo.Context) (*common.BankParamsResponse, error) {
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

// ParamsDID returns the DID params
func ParamsDID(c echo.Context) (*common.DIDParamsResponse, error) {
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

// ParamsDWN returns the DWN params
func ParamsDWN(c echo.Context) (*common.DWNParamsResponse, error) {
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

// ParamsSVC returns the SVC params
func ParamsSVC(c echo.Context) (*common.SVCParamsResponse, error) {
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

// StatusBlock returns the current block
func StatusBlock(c echo.Context) uint64 {
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

// StatusNode returns the node status
func StatusNode(c echo.Context) (*common.StatusResponse, error) {
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

// TxBroadcast broadcasts a transaction to the network
func TxBroadcast(c echo.Context) error {
	return nil
}

// TxEncode encodes a transaction
func TxEncode(c echo.Context) error {
	return nil
}

// TxDecode decodes a transaction
func TxDecode(c echo.Context) error {
	return nil
}

// TxSimulate simulates a transaction on the network
func TxSimulate(c echo.Context) error {
	return nil
}
