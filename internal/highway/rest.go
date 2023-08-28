package highway

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"

	timeout "github.com/vearne/gin-timeout"

	"github.com/sonrhq/core/config"
	"github.com/sonrhq/core/internal/highway/routes"
	"github.com/sonrhq/core/internal/highway/types"
)

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
	routes.RegisterRoutes(r)
	return r
}

type highway struct {
	r    *gin.Engine
	conf *service.Config
}

func (p highway) Start(s service.Service) error {
	fmt.Printf("Starting Highway at :8080")
	return p.r.Run(":8080")
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
