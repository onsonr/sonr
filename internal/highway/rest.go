package highway

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	highlightGin "github.com/highlight/highlight/sdk/highlight-go/middleware/gin"
	"github.com/kardianos/service"

	timeout "github.com/vearne/gin-timeout"

	"github.com/sonrhq/core/config"
	"github.com/sonrhq/core/internal/highway/routes"
	"github.com/sonrhq/core/internal/highway/types"
)

// @title           Sonr Highway Protocol API
// @version         1.0
// @description     API for the Sonr Highway Protocol, a peer-to-peer identity and asset management system.
// @termsOfService  <URL_to_your_terms_of_service>
// @contact.name    Sonr API Support
// @contact.url     <URL_to_your_support>
// @contact.email   <your_support_email>
// @license.name    <Your_License_Name>
// @license.url     <URL_to_license>
// @host            <host_address>:<port>
// @BasePath        /api/v1
func initGin() *gin.Engine {
	gin.SetMode("release")
	// init gin
	r := gin.Default()

	// add timeout middleware with 2 second duration
	defaultMsg := `{"code": -1, "msg":"http: Handler timeout"}`
	r.Use(timeout.Timeout(
		timeout.WithTimeout(config.HighwayRequestTimeout()),
		timeout.WithErrorHttpCode(http.StatusRequestTimeout), // optional
		timeout.WithDefaultMsg(defaultMsg),                   // optional
		timeout.WithCallBack(func(r *http.Request) {
			fmt.Println("timeout happen, url:", r.URL.String())
		})))
	r.Use(highlightGin.Middleware())
	routes.RegisterRoutes(r)
	return r
}

type highway struct {
	r    *gin.Engine
	conf *service.Config
}

func (p highway) Start(s service.Service) error {
	fmt.Printf("Starting Highway at %s", config.HighwayHostAddress())
	go p.r.Run(":8080")
	return nil
}

func (p highway) Stop(s service.Service) error {
	return s.Stop()
}

func runHighway() error {
	h := &highway{
		r: initGin(),
		conf: &service.Config{
			Name:        types.HighwayServiceName,
			DisplayName: types.HighwayServiceDisplayName,
			Description: types.HighwayServiceDescription,
		},
	}

	s, err := service.New(h, h.conf)
	if err != nil {
		return err
	}
	err = s.Run()
	if err != nil {
		return err
	}
	return nil
}
