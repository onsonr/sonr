// FORKED FROM: https://github.com/matrix-org/dendrite/blob/main/cmd/dendrite-demo-libp2p/main.go
// Integrated with our config file and utilizes HostImpl

package matrix

import (
	"fmt"
	"net/http"

	p2phttp "github.com/libp2p/go-libp2p-http"
	"github.com/matrix-org/dendrite/appservice"
	"github.com/matrix-org/dendrite/federationapi"
	"github.com/matrix-org/dendrite/keyserver"
	"github.com/matrix-org/dendrite/roomserver"
	"github.com/matrix-org/dendrite/setup"
	"github.com/matrix-org/dendrite/setup/config"

	"github.com/matrix-org/dendrite/userapi"
	"github.com/matrix-org/gomatrixserverlib"
	"github.com/sonr-io/sonr/pkg/host"
)

func defaultConfig(host host.SonrHost) (config.Dendrite, error) {
	// Get Private Key
	cfg := config.Dendrite{}
	priv, err := host.PrivateKey()
	if err != nil {
		return cfg, err
	}

	cfg.Defaults(true)
	cfg.Global.ServerName = gomatrixserverlib.ServerName(host.HostID().String())
	cfg.Global.PrivateKey = priv
	cfg.Global.KeyID = gomatrixserverlib.KeyID(fmt.Sprintf("ed25519:%s", InstanceName))
	cfg.FederationAPI.FederationMaxRetries = 6
	cfg.Global.JetStream.StoragePath = config.Path(fmt.Sprintf("%s/", InstanceName))
	cfg.UserAPI.AccountDatabase.ConnectionString = config.DataSource(fmt.Sprintf("file:%s-account.db", InstanceName))
	cfg.MediaAPI.Database.ConnectionString = config.DataSource(fmt.Sprintf("file:%s-mediaapi.db", InstanceName))
	cfg.SyncAPI.Database.ConnectionString = config.DataSource(fmt.Sprintf("file:%s-syncapi.db", InstanceName))
	cfg.RoomServer.Database.ConnectionString = config.DataSource(fmt.Sprintf("file:%s-roomserver.db", InstanceName))
	cfg.FederationAPI.Database.ConnectionString = config.DataSource(fmt.Sprintf("file:%s-federationapi.db", InstanceName))
	cfg.AppServiceAPI.Database.ConnectionString = config.DataSource(fmt.Sprintf("file:%s-appservice.db", InstanceName))
	cfg.KeyServer.Database.ConnectionString = config.DataSource(fmt.Sprintf("file:%s-e2ekey.db", InstanceName))
	cfg.MSCs.MSCs = []string{"msc2836"}
	cfg.MSCs.Database.ConnectionString = config.DataSource(fmt.Sprintf("file:%s-mscs.db", InstanceName))
	if err := cfg.Derive(); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func (p *MatrixProtocol) setupMonolith() error {
	keyAPI := keyserver.NewInternalAPI(p.Base, &p.Config.KeyServer, p.fedClient)
	rsAPI := roomserver.NewInternalAPI(
		p.Base,
	)

	userAPI := userapi.NewInternalAPI(p.Base, p.accountDB, &p.Config.UserAPI, nil, keyAPI, rsAPI, p.Base.PushGatewayHTTPClient())
	keyAPI.SetUserAPI(userAPI)

	asAPI := appservice.NewInternalAPI(p.Base, userAPI, rsAPI)
	rsAPI.SetAppserviceAPI(asAPI)
	fsAPI := federationapi.NewInternalAPI(
		p.Base, p.fedClient, rsAPI, p.Base.Caches, nil, true,
	)
	keyRing := fsAPI.KeyRing()
	rsAPI.SetFederationAPI(fsAPI, keyRing)
	provider := newPublicRoomsProvider(p.Host.Pubsub(), rsAPI)
	err := provider.Start()
	if err != nil {

		return err
	}

	monolith := setup.Monolith{
		Config:    &p.Config,
		AccountDB: p.accountDB,
		Client:    p.client,
		FedClient: p.fedClient,
		KeyRing:   keyRing,

		AppserviceAPI:          asAPI,
		FederationAPI:          fsAPI,
		RoomserverAPI:          rsAPI,
		UserAPI:                userAPI,
		KeyAPI:                 keyAPI,
		ExtPublicRoomsProvider: provider,
	}
	monolith.AddAllPublicRoutes(
		p.Base.ProcessContext,
		p.Base.PublicClientAPIMux,
		p.Base.PublicFederationAPIMux,
		p.Base.PublicKeyAPIMux,
		p.Base.PublicWellKnownAPIMux,
		p.Base.PublicMediaAPIMux,
		p.Base.SynapseAdminMux,
	)
	return nil
}

func (p *MatrixProtocol) setupClients() {
	p.accountDB = p.Base.CreateAccountsDB()

	fmt.Println("Running in libp2p federation mode")
	fmt.Println("Warning: Federation with non-libp2p homeservers will not work in this mode yet!")
	tr := &http.Transport{}
	tr.RegisterProtocol(
		"matrix",
		p2phttp.NewTransport(p.Host.Host(), p2phttp.ProtocolOption("/matrix")),
	)

	// Setup FederationClient
	p.fedClient = gomatrixserverlib.NewFederationClient(
		p.Base.Cfg.Global.ServerName, p.Base.Cfg.Global.KeyID,
		p.Base.Cfg.Global.PrivateKey,
		gomatrixserverlib.WithTransport(tr),
	)

	// Setup Client
	tr2 := &http.Transport{}
	tr.RegisterProtocol(
		"matrix",
		p2phttp.NewTransport(p.Host.Host(), p2phttp.ProtocolOption("/matrix")),
	)
	p.client = gomatrixserverlib.NewClient(
		gomatrixserverlib.WithTransport(tr2),
	)
}
