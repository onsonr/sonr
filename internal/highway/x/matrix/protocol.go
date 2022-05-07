// FORKED FROM: https://github.com/matrix-org/dendrite/blob/main/cmd/dendrite-demo-libp2p/p2pdendrite.go.
// However, modified to utilize our HostImpl struct.

package matrix

import (
	"context"

	"github.com/matrix-org/dendrite/appservice"
	"github.com/matrix-org/dendrite/federationapi"
	"github.com/matrix-org/dendrite/keyserver"
	"github.com/matrix-org/dendrite/roomserver"
	"github.com/matrix-org/dendrite/setup"
	"github.com/matrix-org/dendrite/setup/base"
	"github.com/matrix-org/dendrite/setup/config"
	"github.com/matrix-org/dendrite/setup/mscs"
	"github.com/matrix-org/dendrite/userapi"
	us "github.com/matrix-org/dendrite/userapi/storage"
	"github.com/matrix-org/gomatrixserverlib"
	"github.com/sirupsen/logrus"
	"github.com/sonr-io/sonr/pkg/host"
)

const (
	InstanceName = "snr/matrix"
)

type MatrixProtocol struct {
	ctx    context.Context
	node   host.SonrHost
	config config.Dendrite

	Base *base.BaseDendrite

	// Stores
	keydb     *gomatrixserverlib.KeyRing
	fedClient *gomatrixserverlib.FederationClient
	client    *gomatrixserverlib.Client
	accountDB us.Database
}

func New(ctx context.Context, host host.SonrHost, name string) (*MatrixProtocol, error) {
	// Get Config
	cfg, err := defaultConfig(host)
	if err != nil {
		return nil, err
	}

	p := &MatrixProtocol{
		ctx:    ctx,
		node:   host,
		config: cfg,
		Base:   base.NewBaseDendrite(&cfg, name),
	}

	keyAPI := keyserver.NewInternalAPI(p.Base, &p.config.KeyServer, p.fedClient)
	rsAPI := roomserver.NewInternalAPI(
		p.Base,
	)

	userAPI := userapi.NewInternalAPI(p.Base, p.accountDB, &cfg.UserAPI, nil, keyAPI, rsAPI, p.Base.PushGatewayHTTPClient())
	keyAPI.SetUserAPI(userAPI)

	asAPI := appservice.NewInternalAPI(p.Base, userAPI, rsAPI)
	rsAPI.SetAppserviceAPI(asAPI)
	fsAPI := federationapi.NewInternalAPI(
		p.Base, p.fedClient, rsAPI, p.Base.Caches, nil, true,
	)
	keyRing := fsAPI.KeyRing()
	rsAPI.SetFederationAPI(fsAPI, keyRing)
	provider := newPublicRoomsProvider(host.Pubsub(), rsAPI)
	err = provider.Start()
	if err != nil {
		panic("failed to create new public rooms provider: " + err.Error())
	}

	monolith := setup.Monolith{
		Config:    &p.config,
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
	if err := mscs.Enable(p.Base, &monolith); err != nil {
		logrus.WithError(err).Fatalf("Failed to enable MSCs")
	}

	// httpRouter := mux.NewRouter().SkipClean(true).UseEncodedPath()
	// httpRouter.PathPrefix(httputil.InternalPathPrefix).Handler(baseDendrite.InternalAPIMux)
	// httpRouter.PathPrefix(httputil.PublicClientPathPrefix).Handler(baseDendrite.PublicClientAPIMux)
	// httpRouter.PathPrefix(httputil.PublicMediaPathPrefix).Handler(baseDendrite.PublicMediaAPIMux)
	// embed.Embed(httpRouter, *instancePort, "Yggdrasil Demo")

	// libp2pRouter := mux.NewRouter().SkipClean(true).UseEncodedPath()
	// libp2pRouter.PathPrefix(httputil.PublicFederationPathPrefix).Handler(baseDendrite.PublicFederationAPIMux)
	// libp2pRouter.PathPrefix(httputil.PublicKeyPathPrefix).Handler(baseDendrite.PublicKeyAPIMux)
	// libp2pRouter.PathPrefix(httputil.PublicMediaPathPrefix).Handler(baseDendrite.PublicMediaAPIMux)

	return p, nil
}
