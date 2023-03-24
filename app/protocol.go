package app

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bufbuild/connect-go"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/sonrhq/core/internal/controller"
	"github.com/sonrhq/core/internal/resolver"
	v1 "github.com/sonrhq/core/types/highway/v1"
	highway "github.com/sonrhq/core/types/highway/v1/highwayv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Protocol struct {
	ctx client.Context
	highway.UnimplementedMpcHandler
	highway.UnimplementedAuthenticationHandler
	highway.UnimplementedVaultHandler
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
	if hasTLSCert() {
		http.ListenAndServeTLS(
			fmt.Sprintf(":%s", getServerPort()),
			getTLSCert(),
			getTLSKey(),
			mux,
		)
	} else {
		http.ListenAndServe(
			fmt.Sprintf("%s:%s", getServerHost(), getServerPort()),
			h2c.NewHandler(mux, &http2.Server{}),
		)
	}
}

func getServerHost() string {
	if host := os.Getenv("CONNECT_SERVER_ADDRESS"); host != "" {
		log.Printf("using CONNECT_SERVER_ADDRESS: %s", host)
		return host
	}
	return "localhost"
}

func getServerPort() string {
	if port := os.Getenv("CONNECT_SERVER_PORT"); port != "" {
		log.Printf("using CONNECT_SERVER_PORT: %s", port)
		return port
	}
	return "8080"
}

func getTLSCert() string {
	if cert := os.Getenv("CONNECT_SERVER_TLS_CERT"); cert != "" {
		log.Printf("using CONNECT_SERVER_TLS_CERT: %s", cert)
		return cert
	}
	return ""
}

func getTLSKey() string {
	if key := os.Getenv("CONNECT_SERVER_TLS_KEY"); key != "" {
		log.Printf("using CONNECT_SERVER_TLS_KEY: %s", key)
		return key
	}
	return ""
}

func hasTLSCert() bool {
	return getTLSCert() != "" && getTLSKey() != "" && !isDev()
}

func isDev() bool {
	return os.Getenv("ENVIRONMENT") == "dev"
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                            Authorization interceptor                           ||
// ! ||--------------------------------------------------------------------------------||
const tokenHeader = "Sonr-KS"

func NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				// Send a token with client requests.
				req.Header().Set(tokenHeader, "sample")
			} else if req.Header().Get(tokenHeader) == "" {
				// Check token in handlers.
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                             Authentication Handler                             ||
// ! ||--------------------------------------------------------------------------------||

func AuthenticationHandler() (string, http.Handler) {
	return highway.NewAuthenticationHandler(hway)
}

func (p *Protocol) Keygen(ctx context.Context, req *connect.Request[v1.KeygenRequest]) (*connect.Response[v1.KeygenResponse], error) {
	service, err := resolver.GetService(ctx, req.Msg.GetOrigin())
	if err != nil {
		return nil, err
	}
	cred, err := service.VerifyCreationChallenge(req.Msg.GetCredentialResponse())
	if err != nil {
		return nil, err
	}
	cont, err := controller.NewController(ctx, cred)
	if err != nil {
		return nil, err
	}
	res := &connect.Response[v1.KeygenResponse]{
		Msg: &v1.KeygenResponse{
			Did:         cont.DidDocument().Id,
			DidDocument: cont.DidDocument(),
			Success:     true,
		},
	}
	res.Header().Set("Sonr-Version", "v0.6.0")
	return res, nil
}

func (p *Protocol) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	service, err := resolver.GetService(ctx, req.Msg.GetOrigin())
	if err != nil {
		return nil, err
	}
	//cred, err := service.VeriifyAssertionChallenge(resp string, cred *crypto.WebauthnCredential)
	fmt.Printf("Login service: %v", service)
	return nil, nil
}

func (p *Protocol) QueryDocument(ctx context.Context, req *connect.Request[v1.QueryDocumentRequest]) (*connect.Response[v1.QueryDocumentResponse], error) {
	doc, err := resolver.GetDID(ctx, req.Msg.GetDid())
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&v1.QueryDocumentResponse{
		Success:        (doc != nil),
		AccountAddress: doc.DIDIdentifier(),
		DidDocument:    doc,
	})
	res.Header().Set("Sonr-Version", "v0.6.0")
	return res, nil
}

func (p *Protocol) QueryService(ctx context.Context, req *connect.Request[v1.QueryServiceRequest]) (*connect.Response[v1.QueryServiceResponse], error) {
	request := req.Any()
	if request == nil {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			errors.New("no request provided"),
		)
	}

	// Get the origin from the request.
	service, err := resolver.GetService(ctx, req.Msg.GetOrigin())
	if err != nil {
		return nil, err
	}
	challenge, err := service.IssueChallenge()
	if err != nil {
		return nil, err
	}
	res := &connect.Response[v1.QueryServiceResponse]{
		Msg: &v1.QueryServiceResponse{
			Challenge: string(challenge),
			RpName:    "Sonr",
			RpId:      service.Origin,
		},
	}
	res.Header().Set("Sonr-Version", "v0.6.0")
	return res, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                Accounts handler                                ||
// ! ||--------------------------------------------------------------------------------||

func MpcHandler() (string, http.Handler) {
	return highway.NewMpcHandler(hway, connect.WithInterceptors(NewAuthInterceptor()))
}

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

func VaultHandler() (string, http.Handler) {
	return highway.NewVaultHandler(hway)
}

func (p *Protocol) Add(ctx context.Context, req *connect.Request[v1.AddShareRequest]) (*connect.Response[v1.AddShareResponse], error) {
	err := resolver.InsertRecord(req.Msg.Key, req.Msg.Value)
	if err != nil {
		return nil, err
	}
	res := &connect.Response[v1.AddShareResponse]{
		Msg: &v1.AddShareResponse{
			Success: true,
		},
	}
	res.Header().Set("Sonr-Version", "v0.6.0")
	return res, nil
}

func (p *Protocol) Sync(ctx context.Context, req *connect.Request[v1.SyncShareRequest]) (*connect.Response[v1.SyncShareResponse], error) {
	records, err := resolver.GetRecord(req.Msg.Key)
	if err != nil {
		return nil, err
	}
	res := &connect.Response[v1.SyncShareResponse]{
		Msg: &v1.SyncShareResponse{
			Key:     req.Msg.Key,
			Success: true,
			Value:   base64.StdEncoding.EncodeToString(records),
		},
	}
	res.Header().Set("Sonr-Version", "v0.6.0")
	return res, nil
}

func (p *Protocol) Refresh(context.Context, *connect.Request[v1.RefreshShareRequest]) (*connect.Response[v1.RefreshShareResponse], error) {
	return nil, nil
}
