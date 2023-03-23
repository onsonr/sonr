package app

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/bufbuild/connect-go"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/sonrhq/core/internal/controller"
	"github.com/sonrhq/core/internal/resolver"
	v1 "github.com/sonrhq/core/types/highway/v1"
	highwayv1connect "github.com/sonrhq/core/types/highway/v1/highwayv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var _ highwayv1connect.AuthenticationHandler = (*Protocol)(nil)
var _ highwayv1connect.MpcHandler = (*Protocol)(nil)
var _ highwayv1connect.VaultHandler = (*Protocol)(nil)

var hway *Protocol

type Protocol struct {
	ctx client.Context
}

func AuthenticationHandler() (string, http.Handler) {
	return highwayv1connect.NewAuthenticationHandler(hway)
}

func MpcHandler() (string, http.Handler) {
	return highwayv1connect.NewMpcHandler(hway)
}

func VaultHandler() (string, http.Handler) {
	return highwayv1connect.NewVaultHandler(hway)
}

func RegisterHighway(ctx client.Context) {
	hway = &Protocol{ctx: ctx}
	mux := http.NewServeMux()
	mux.Handle(AuthenticationHandler())
	mux.Handle(MpcHandler())
	mux.Handle(VaultHandler())
	go hway.serveHTTP(mux)
}

func (p *Protocol) serveHTTP(mux *http.ServeMux) {
	http.ListenAndServe(
		getServerHost(),
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

func getServerHost() string {
	if host := os.Getenv("CONNECT_SERVER_ADDRESS"); host != "" {
		log.Printf("using CONNECT_SERVER_ADDRESS: %s", host)
		return host
	}
	return "localhost:8080"
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                             Authentication Handler                             ||
// ! ||--------------------------------------------------------------------------------||
func (p *Protocol) Keygen(ctx context.Context, req *connect.Request[v1.KeygenRequest]) (*connect.Response[v1.KeygenResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	service, err := resolver.GetService(ctx, req.Msg.Origin, resolver.SonrPublicRpcOrigin)
	if err != nil {
		return nil, err
	}
	cred, err := service.VerifyCreationChallenge(req.Msg.CredentialResponse)
	if err != nil {
		return nil, err
	}
	cont, err := controller.NewController(ctx, cred)
	if err != nil {
		return nil, err
	}
	return &connect.Response[v1.KeygenResponse]{
		Msg: &v1.KeygenResponse{
			Did:         cont.DidDocument().Id,
			DidDocument: cont.DidDocument(),
			Success:     true,
		},
	}, nil
}

func (p *Protocol) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *Protocol) QueryDocument(ctx context.Context, req *connect.Request[v1.QueryDocumentRequest]) (*connect.Response[v1.QueryDocumentResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *Protocol) QueryService(ctx context.Context, req *connect.Request[v1.QueryServiceRequest]) (*connect.Response[v1.QueryServiceResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	service, err := resolver.GetService(ctx, req.Msg.Origin, resolver.SonrPublicRpcOrigin)
	if err != nil {
		return nil, err
	}
	challenge, err := service.IssueChallenge()
	if err != nil {
		return nil, err
	}
	return &connect.Response[v1.QueryServiceResponse]{
		Msg: &v1.QueryServiceResponse{
			Challenge: string(challenge),
			Service:   service,
			RpName:    "Sonr",
			RpId:      service.Origin,
		},
	}, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                Accounts handler                                ||
// ! ||--------------------------------------------------------------------------------||

func (p *Protocol) CreateAccount(context.Context, *connect.Request[v1.CreateAccountRequest]) (*connect.Response[v1.CreateAccountResponse], error) {
	return nil, nil
}
func (p *Protocol) ListAccounts(context.Context, *connect.Request[v1.ListAccountsRequest]) (*connect.Response[v1.ListAccountsResponse], error) {
	return nil, nil
}
func (p *Protocol) GetAccount(context.Context, *connect.Request[v1.GetAccountRequest]) (*connect.Response[v1.GetAccountResponse], error) {
	return nil, nil
}
func (p *Protocol) DeleteAccount(context.Context, *connect.Request[v1.DeleteAccountRequest]) (*connect.Response[v1.DeleteAccountResponse], error) {
	return nil, nil
}
func (p *Protocol) SignMessage(context.Context, *connect.Request[v1.SignMessageRequest]) (*connect.Response[v1.SignMessageResponse], error) {
	return nil, nil
}
func (p *Protocol) VerifyMessage(context.Context, *connect.Request[v1.VerifyMessageRequest]) (*connect.Response[v1.VerifyMessageResponse], error) {
	return nil, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  Vault handler                                 ||
// ! ||--------------------------------------------------------------------------------||

func (p *Protocol) Add(context.Context, *connect.Request[v1.AddShareRequest]) (*connect.Response[v1.AddShareResponse], error) {
	return nil, nil
}
func (p *Protocol) Sync(context.Context, *connect.Request[v1.SyncShareRequest]) (*connect.Response[v1.SyncShareResponse], error) {
	return nil, nil
}
func (p *Protocol) Refresh(context.Context, *connect.Request[v1.RefreshShareRequest]) (*connect.Response[v1.RefreshShareResponse], error) {
	return nil, nil
}
