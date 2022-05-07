// FORKED FROM: https://github.com/matrix-org/dendrite/blob/main/cmd/dendrite-demo-libp2p/p2pdendrite.go.
// However, modified to utilize our HostImpl struct.

package matrix

import (
	"context"

	"github.com/matrix-org/dendrite/appservice"
	asapi "github.com/matrix-org/dendrite/appservice/api"
	"github.com/matrix-org/dendrite/federationapi"
	fsapi "github.com/matrix-org/dendrite/federationapi/api"
	"github.com/matrix-org/dendrite/keyserver"
	keyapi "github.com/matrix-org/dendrite/keyserver/api"
	"github.com/matrix-org/dendrite/roomserver"
	rsapi "github.com/matrix-org/dendrite/roomserver/api"
	"github.com/matrix-org/dendrite/setup"
	"github.com/matrix-org/dendrite/setup/base"
	"github.com/matrix-org/dendrite/setup/config"
	"github.com/matrix-org/dendrite/setup/mscs"
	"github.com/matrix-org/dendrite/userapi"
	usrapi "github.com/matrix-org/dendrite/userapi/api"
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
	Host   host.SonrHost
	Config config.Dendrite

	Base *base.BaseDendrite

	// Stores
	keydb     *gomatrixserverlib.KeyRing
	fedClient *gomatrixserverlib.FederationClient
	client    *gomatrixserverlib.Client
	accountDB us.Database
	monolith  setup.Monolith

	// Services
	provider *publicRoomsProvider
	keyAPI   keyapi.KeyInternalAPI
	rsAPI    rsapi.RoomserverInternalAPI
	asAPI    asapi.AppServiceQueryAPI
	fsAPI    fsapi.FederationInternalAPI
	userAPI  usrapi.UserInternalAPI
}

func New(ctx context.Context, host host.SonrHost) (*MatrixProtocol, error) {
	// Get Config
	cfg, err := defaultConfig(host)
	if err != nil {
		return nil, err
	}

	base := base.NewBaseDendrite(&cfg, host.Config().MatrixServerName)

	p := &MatrixProtocol{
		ctx:    ctx,
		Host:   host,
		Config: cfg,
		Base:   base,
	}

	keyAPI := keyserver.NewInternalAPI(p.Base, &p.Config.KeyServer, p.fedClient)
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

	if err := mscs.Enable(p.Base, &p.monolith); err != nil {
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
