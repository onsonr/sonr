package highway

import (
	context "context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/duo-labs/webauthn.io/session"
	"github.com/duo-labs/webauthn/protocol"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/fxamacker/cbor/v2"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	rtv1 "github.com/sonr-io/blockchain/x/registry/types"
	"github.com/sonr-io/core/channel"
	"github.com/sonr-io/core/device"

	"github.com/sonr-io/core/crypto"
	"github.com/sonr-io/core/did"
	"github.com/sonr-io/core/did/ssi"
	"github.com/sonr-io/core/highway/user"
	hn "github.com/sonr-io/core/host"
	"github.com/sonr-io/core/host/discover"
	"github.com/sonr-io/core/host/exchange"
	"github.com/sonr-io/core/util"
	v1 "go.buf.build/grpc/go/sonr-io/core/highway/v1"
	"google.golang.org/grpc"

	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

// Error Definitions
var (
	logger                 = golog.Default.Child("node/highway")
	ErrEmptyQueue          = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery        = errors.New("No SName or PeerID provided.")
	ErrMissingParam        = errors.New("Paramater is missing.")
	ErrProtocolsNotSet     = errors.New("Node Protocol has not been initialized.")
	ErrMethodUnimplemented = errors.New("Method is not implemented.")
)

// HighwayServer is the RPC Service for the Custodian Node.
type HighwayServer struct {
	v1.HighwayServer
	node   hn.HostImpl
	cosmos cosmosclient.Client

	// Properties
	ctx      context.Context
	listener net.Listener
	grpc     *grpc.Server
	router   *mux.Router
	*discover.DiscoverProtocol
	*exchange.ExchangeProtocol

	// Configuration
	auth         *webauthn.WebAuthn
	sessionStore *session.Store
	userDb       *user.UserDB
	// ipfs *storage.IPFSService

	// List of Entries
	channels map[string]channel.Channel
}

// NewHighwayServer creates a new Highway service stub for the node.
func NewHighway(ctx context.Context, opts ...hn.Option) (*HighwayServer, error) {
	// Create a new HostImpl
	r := mux.NewRouter()
	node, err := hn.NewHost(ctx, device.Role_HIGHWAY, opts...)
	if err != nil {
		return nil, err
	}

	// // Set IPFS Service
	// stub.ipfs, err = storage.Init()
	// if err != nil {
	// 	return nil, err
	// }

	// Get the Listener for the Host
	lst, err := node.Listener()
	if err != nil {
		return nil, err
	}

	// create an instance of cosmosclient
	cosmos, err := cosmosclient.New(ctx)
	if err != nil {
		return nil, err
	}

	// Create a WebAuthn instance
	web, err := webauthn.New(node.WebauthnConfig())
	if err != nil {
		return nil, err
	}

	// Create a new Session Store
	sessionStore, err := session.NewStore()
	if err != nil {
		return nil, err
	}

	// Create the RPC Service
	stub := &HighwayServer{
		node:         node,
		ctx:          ctx,
		grpc:         grpc.NewServer(),
		cosmos:       cosmos,
		listener:     lst,
		auth:         web,
		router:       r,
		sessionStore: sessionStore,
		userDb:       user.DB(),
	}

	r.HandleFunc("/register/begin/{username}", stub.BeginRegistration).Methods("GET")
	r.HandleFunc("/register/finish/{username}", stub.FinishRegistration).Methods("POST")
	r.HandleFunc("/login/begin/{username}", stub.BeginLogin).Methods("GET")
	r.HandleFunc("/login/finish/{username}", stub.FinishLogin).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

	// TODO Implement P2P Protocols for Sonr Network
	// // Set Discovery Protocol
	// stub.DiscoverProtocol, err = discover.New(ctx, node, stub)
	// if err != nil {
	// 	logger.Errorf("%s - Failed to start DiscoveryProtocol", err)
	// 	return nil, err
	// }

	// // Set Transmit Protocol
	// stub.ExchangeProtocol, err = exchange.New(ctx, node, stub)
	// if err != nil {
	// 	logger.Errorf("%s - Failed to start TransmitProtocol", err)
	// 	return nil, err
	// }

	// Register RPC Service
	v1.RegisterHighwayServer(stub.grpc, stub)
	return stub, nil
}

// Serve starts the RPC Service.
func (s *HighwayServer) Serve() {
	logger.Infof("Starting RPC Server on %s", s.listener.Addr().String())
	go s.serveCtxListener(s.ctx, s.listener)
}

// Serve serves the RPC Service on the given port.
func (s *HighwayServer) serveCtxListener(ctx context.Context, listener net.Listener) {
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	if err := s.grpc.Serve(listener); err != nil {
		logger.Errorf("%s - Failed to start HTTP server", err)
	}
	s.node.Persist()
}

func (s *HighwayServer) BeginRegistration(w http.ResponseWriter, r *http.Request) {
	// get username/friendly name
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		util.JsonResponse(w, fmt.Errorf("must supply a valid username i.e. foo@bar.com"), http.StatusBadRequest)
		return
	}

	// get user
	usr, err := s.userDb.GetUser(username)
	// user doesn't exist, create new user
	if err != nil {
		usr = user.NewUser(username, fmt.Sprintf("%s.snr", username))
		s.userDb.PutUser(usr)
	}
	// Updating the AuthenticatorSelection options.
	// See the struct declarations for values
	authSelect := protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.AuthenticatorAttachment("platform"),
		RequireResidentKey:      protocol.ResidentKeyUnrequired(),
		UserVerification:        protocol.VerificationRequired,
	}

	// Updating the ConveyencePreference options.
	// See the struct declarations for values
	conveyencePref := protocol.ConveyancePreference(protocol.PreferNoAttestation)
	registerOptions := func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
		credCreationOpts.CredentialExcludeList = usr.CredentialExcludeList()
	}

	// generate PublicKeyCredentialCreationOptions, session data
	options, sessionData, err := s.auth.BeginRegistration(
		usr,
		registerOptions,
		webauthn.WithAuthenticatorSelection(authSelect),
		webauthn.WithConveyancePreference(conveyencePref),
	)

	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// store session data as marshaled JSON
	err = s.sessionStore.SaveWebauthnSession("registration", sessionData, r, w)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.JsonResponse(w, options, http.StatusOK)
}

func (s *HighwayServer) FinishRegistration(w http.ResponseWriter, r *http.Request) {

	// get username
	vars := mux.Vars(r)
	username := vars["username"]

	// get user
	user, err := s.userDb.GetUser(username)
	// user doesn't exist
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// load the session data
	sessionData, err := s.sessionStore.GetWebauthnSession("registration", r)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	credential, err := s.auth.FinishRegistration(user, sessionData, r)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	coseKey := crypto.COSEKey{}
	err = cbor.Unmarshal(credential.PublicKey, &coseKey)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	pubKey, err := crypto.DecodePublicKey(&coseKey)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// account `alice` was initialized during `starport chain serve`
	accountName := "alice"

	// get account from the keyring by account name and return a bech32 address
	acc, err := s.cosmos.Account(accountName)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, "Failed to find blockchain account", http.StatusNotFound)
	}

	didStr := fmt.Sprintf("did:sonr:%s", acc.Address("snr"))
	id, err := did.ParseDID(didStr)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, "Failed to parse DID", http.StatusNotFound)
	}

	verf, err := did.NewVerificationMethod(*id, ssi.JsonWebKey2020, *id, pubKey)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, "Failed to create verification method", http.StatusNotFound)
	}

	log.Println(verf)

	// define a message to create a did
	msg := &rtv1.MsgRegisterName{
		Creator:         acc.Address("snr"),
		NameToRegister:  username,
		PublicKeyBuffer: credential.PublicKey,
	}
	log.Println(msg.String())

	// broadcast a transaction from account `alice` with the message to create a did
	// store response in txResp
	txResp, err := s.cosmos.BroadcastTx(accountName, msg)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, "Failed to broadcast to blockchain", http.StatusBadRequest)
	}
	log.Println(txResp.String())

	// Return response from broadcasting a transaction
	user.AddCredential(*credential)
	util.JsonResponse(w, txResp.String(), http.StatusOK)
}

func (s *HighwayServer) BeginLogin(w http.ResponseWriter, r *http.Request) {

	// get username
	vars := mux.Vars(r)
	username := vars["username"]

	// get user
	user, err := s.userDb.GetUser(username)

	// user doesn't exist
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// generate PublicKeyCredentialRequestOptions, session data
	options, sessionData, err := s.auth.BeginLogin(user)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// store session data as marshaled JSON
	err = s.sessionStore.SaveWebauthnSession("authentication", sessionData, r, w)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.JsonResponse(w, options, http.StatusOK)
}

func (s *HighwayServer) FinishLogin(w http.ResponseWriter, r *http.Request) {

	// get username
	vars := mux.Vars(r)
	username := vars["username"]

	// get user
	user, err := s.userDb.GetUser(username)

	// user doesn't exist
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// load the session data
	sessionData, err := s.sessionStore.GetWebauthnSession("authentication", r)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// in an actual implementation, we should perform additional checks on
	// the returned 'credential', i.e. check 'credential.Authenticator.CloneWarning'
	// and then increment the credentials counter
	_, err = s.auth.FinishLogin(user, sessionData, r)
	if err != nil {
		log.Println(err)
		util.JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// handle successful login
	util.JsonResponse(w, "Login Success", http.StatusOK)
}
