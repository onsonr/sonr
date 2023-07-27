package highway

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/pkg/highway/routes"

	timeout "github.com/vearne/gin-timeout"
)

const serviceName = "Highway Protocol Service"
const serviceDescription = "Proxy for underlying blockchain protocol"

type highway struct {
	r *gin.Engine
}

func StartService() {
	gin.SetMode("release")
	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	prg := &highway{
		r: initGin(),
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

func (p highway) Start(s service.Service) error {
	return p.r.Run(local.HighwayHostPort())
}

func (p highway) Stop(s service.Service) error {
	return s.Stop()
}

func initGin() *gin.Engine {
	// init gin
	r := gin.Default()

	// add timeout middleware with 2 second duration
	defaultMsg := `{"code": -1, "msg":"http: Handler timeout"}`
	r.Use(timeout.Timeout(
		timeout.WithTimeout(local.HighwayRequestTimeout()),
		timeout.WithErrorHttpCode(http.StatusRequestTimeout), // optional
		timeout.WithDefaultMsg(defaultMsg),                   // optional
		timeout.WithCallBack(func(r *http.Request) {
			fmt.Println("timeout happen, url:", r.URL.String())
		})))
	routes.RegisterRoutes(r)
	return r
}
