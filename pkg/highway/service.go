package highway

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"github.com/pelletier/go-toml/v2"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/pkg/highway/routes"

	tmcfg "github.com/cometbft/cometbft/config"
	timeout "github.com/vearne/gin-timeout"
)

const serviceName = "Highway Protocol Service"
const serviceDescription = "Proxy for underlying blockchain protocol"

type program struct {
	cfg *tmcfg.Config
	r  *gin.Engine
}

func Start() error {
	r := initGin()
	return r.Run()
}

func StartService() {
   gin.SetMode("release")
	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	cfg, err := getTmConfig()
	if err != nil {
		fmt.Println("Cannot get the config: " + err.Error())
	}
	prg := &program{
		cfg: cfg,
		r:   initGin(),
	}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		fmt.Println("Cannot create the service: " + err.Error())
	}
	err = s.Run()
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
}

func (p program) Start(s service.Service) error {
	if p.cfg != nil && p.cfg.RPC != nil && p.cfg.RPC.TLSCertFile != "" && p.cfg.RPC.TLSKeyFile != "" {
		return p.r.RunTLS(":8080", p.cfg.RPC.TLSCertFile, p.cfg.RPC.TLSKeyFile)
	}
	return p.r.Run()
}

func (p program) Stop(s service.Service) error {
	return s.Stop()
}

func getTmConfig() (*tmcfg.Config, error) {
	var config tmcfg.Config
	configFile := local.Context().ConfigTomlPath
	bz, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = toml.Unmarshal(bz, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func initGin() *gin.Engine {
	// init gin
	r := gin.Default()

	// add timeout middleware with 2 second duration
	defaultMsg := `{"code": -1, "msg":"http: Handler timeout"}`
	r.Use(timeout.Timeout(
		timeout.WithTimeout(5*time.Second),
		timeout.WithErrorHttpCode(http.StatusRequestTimeout), // optional
		timeout.WithDefaultMsg(defaultMsg),                   // optional
		timeout.WithCallBack(func(r *http.Request) {
			fmt.Println("timeout happen, url:", r.URL.String())
		})))
	routes.RegisterRoutes(r)
	return r
}
