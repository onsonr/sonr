package handler

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/sonrhq/core/internal/crypto"
	mdw "github.com/sonrhq/core/internal/highway/middleware"

	walletpb "github.com/sonrhq/core/types/highway/wallet/v1"
)

// WalletAPI is the alias for the Highway Wallet Service Server.
type WalletAPI = walletpb.WalletServiceServer

// WalletHandler is the handler for the authentication service
type WalletHandler struct {
	cctx client.Context
}

// CreateAccount creates a new account with a given coin type and name.
//
// @Summary Create an account
// @Description Creates a new account with a given coin type and name.
// @Accept  json
// @Produce  json
// @Param   coinType path string true "Coin Type Name"
// @Success 200 {object} map[string]interface{} "Account Info"
// @Router /createAccount/{coinType} [post]
func (a *WalletHandler) CreateAccount(ctx context.Context, req *walletpb.CreateAccountRequest) (*walletpb.CreateAccountResponse, error) {
	cont, err := mdw.UseControllerAccount(req.Jwt)
	if err != nil {
		return nil, err
	}

	accInfo, err := cont.CreateWallet(req.GetCoinType())
	if err != nil {
		return nil, err
	}
	return &walletpb.CreateAccountResponse{
		Address:  accInfo.Address,
		CoinType: req.GetCoinType(),
		Owner:    cont.Account().Address,
	}, nil
}

// GetAccount returns an account's details given its DID.
//
// @Summary Get an account's details
// @Description Returns an account's details given its DID.
// @Accept  json
// @Produce  json
// @Param   did path string true "DID of Account"
// @Success 200 {object} map[string]interface{} "Account Info"
// @Router /getAccount/{did} [get]
func (a *WalletHandler) GetAccount(ctx context.Context, req *walletpb.GetAccountRequest) (*walletpb.GetAccountResponse, error) {
	cont, err := mdw.UseControllerAccount(req.Jwt)
	if err != nil {
		return nil, err
	}
	accInfo, err := cont.GetWallet(req.Address)
	if err != nil {
		return nil, err
	}
	return &walletpb.GetAccountResponse{
		Address: accInfo.Address,
	}, nil
}

// ListAccounts returns a list of wallet accounts given a coin type.
//
// @Summary List wallet accounts
// @Description Returns a list of wallet accounts given a coin type.
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "Accounts List"
// @Router /listAccounts [get]
func (a *WalletHandler) ListAccounts(ctx context.Context, req *walletpb.ListAccountsRequest) (*walletpb.ListAccountsResponse, error) {
	cont, err := mdw.UseControllerAccount(req.Jwt)
	if err != nil {
		return nil, err
	}
	return &walletpb.ListAccountsResponse{
		Success: true,
		Accounts: cont.ListWallets(),
	}, nil
}

// SignWithAccount signs a message with an account given its DID. Requires the JWT of their Keyshare.
//
// @Summary Sign a message with an account
// @Description Signs a message with an account given its DID. Requires the JWT of their Keyshare.
// @Accept  json
// @Produce  json
// @Param   did path string true "DID of Account"
// @Param   msg query string true "Message to Sign"
// @Success 200 {object} map[string]interface{} "Signature Info"
// @Router /signWithAccount/{did} [post]
func (a *WalletHandler) SignMessage(ctx context.Context, req *walletpb.SignMessageRequest) (*walletpb.SignMessageResponse, error) {
	cont, err := mdw.UseControllerAccount(req.Jwt)
	if err != nil {
		return nil, err
	}
	sig, err := cont.SignWithWallet("did", []byte{0x00})
	if err != nil {
		return nil, err
	}
	return &walletpb.SignMessageResponse{
		Success:  true,
		Signature: sig,
		Message:   string(req.Message),
	}, nil
}

// VerifyWithAccount verifies a signature with an account.
//
// @Summary Verify a signature with an account
// @Description Verifies a signature with an account.
// @Accept  json
// @Produce  json
// @Param   did path string true "DID of Account"
// @Param   msg query string true "Message"
// @Param   sig query string true "Signature"
// @Success 200 {object} map[string]interface{} "Verification Result"
// @Router /verifyWithAccount/{did} [post]
func (a *WalletHandler) VerifySignature(ctx context.Context, req *walletpb.VerifySignatureRequest) (*walletpb.VerifySignatureResponse, error) {
	cont, err := mdw.UseControllerAccount(req.Jwt)
	if err != nil {
		return nil, err
	}
	msg, err := crypto.Base64Decode(req.Address)
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Base64Decode(req.Address)
	if err != nil {
		return nil, err
	}
	valid, err := cont.VerifyWithWallet(req.Address, msg, sig)
	if err != nil {
		return nil, err
	}
	return &walletpb.VerifySignatureResponse{
		MessageVerified: valid,
		Message: 	   string(req.Message),
	}, nil
}

// ExportWallet returns the encoded Sonr Wallet structure with an encrypted keyshare, which can be opened
// with the user's password, within Sonr Clients.
//
// @Summary Export Wallet
// @Description Returns the encoded Sonr Wallet structure with an encrypted keyshare.
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "Wallet Export Info"
// @Router /exportWallet [get]
func (a *WalletHandler) ExportWallet(ctx context.Context, req *walletpb.ExportWalletRequest) (*walletpb.ExportWalletResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
