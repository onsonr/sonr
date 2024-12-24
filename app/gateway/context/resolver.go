package context

import (
	"fmt"

	"github.com/onsonr/sonr/internal/common"
)

// ParamsBank returns the bank params
func (cc *GatewayContext) ParamsBank() (*common.BankParamsResponse, error) {
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
func (cc *GatewayContext) ParamsDID() (*common.DIDParamsResponse, error) {
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
func (cc *GatewayContext) ParamsDWN() (*common.DWNParamsResponse, error) {
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
func (cc *GatewayContext) ParamsSVC() (*common.SVCParamsResponse, error) {
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
func (cc *GatewayContext) StatusBlock() string {
	qc, err := common.NewNodeClient(cc.grpcAddr)
	if err != nil {
		return "-1"
	}
	resp, err := qc.Status(bgCtx(), &common.StatusRequest{})
	if err != nil {
		return "-1"
	}
	return fmt.Sprintf("%d", resp.GetHeight())
}

// StatusNode returns the node status
func (cc *GatewayContext) StatusNode() (*common.StatusResponse, error) {
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
func (cc *GatewayContext) TxBroadcast() error {
	return nil
}

// TxEncode encodes a transaction
func (cc *GatewayContext) TxEncode() error {
	return nil
}

// TxDecode decodes a transaction
func (cc *GatewayContext) TxDecode() error {
	return nil
}

// TxSimulate simulates a transaction on the network
func (cc *GatewayContext) TxSimulate() error {
	return nil
}
