package highway

import (
	"fmt"
	// swagger embed files
	// gin-swagger middleware

	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"

	"github.com/sonrhq/core/config"
	"github.com/sonrhq/core/internal/highway/types"
)

type highway struct {
	r    *gin.Engine
	conf *service.Config
}

func (p highway) Start(s service.Service) error {
	fmt.Printf("Starting Highway at %s", config.HighwayHostAddress())
	return p.r.Run(
		":8080",
	)
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

type icefirekv struct {
	conf *service.Config
}

func (p icefirekv) Start(s service.Service) error {
	fmt.Println("Starting IceFire KV")
	return s.Start()
}

func (p icefirekv) Stop(s service.Service) error {
	fmt.Println("Stopping IceFire KV")
	return s.Stop()
}

func runIcefireKv(exec string, args []string) error {
	if exec == "" {
		fmt.Println("No executable found for IceFire KV")
		return nil
	}
	fmt.Println("Starting IceFire KV")
	ikv := &icefirekv{
		conf: &service.Config{
			Name:        types.IceFireKVServiceName,
			DisplayName: types.IceFireKVServiceDisplayName,
			Description: types.IceFireKVServiceDescription,
			Executable:  exec,
			Arguments:   args,
		},
	}

	s, err := service.New(ikv, ikv.conf)
	if err != nil {
		return err
	}
	err = s.Run()
	if err != nil {
		panic(err)
	}
	return nil
}

type icefiresql struct {
	conf *service.Config
}

func (p icefiresql) Start(s service.Service) error {
	fmt.Println("Starting IceFire SQL")
	return s.Start()
}

func (p icefiresql) Stop(s service.Service) error {
	fmt.Println("Stopping IceFire SQL")
	return s.Stop()
}

func runIcefireSQL(exec string, args []string) error {
	if exec == "" {
		fmt.Println("No executable found for IceFire SQL")
		return nil
	}
	fmt.Println("Starting IceFire SQL")
	ikv := &icefiresql{
		conf: &service.Config{
			Name:        types.IceFireKVServiceName,
			DisplayName: types.IceFireKVServiceDisplayName,
			Description: types.IceFireKVServiceDescription,
			Executable:  exec,
			Arguments:   args,
		},
	}

	s, err := service.New(ikv, ikv.conf)
	if err != nil {
		return err
	}
	err = s.Run()
	if err != nil {
		panic(err)
	}
	return nil
}
