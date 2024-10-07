package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/onsonr/sonr/internal/session"
	"github.com/onsonr/sonr/pkg/nebula"
	"github.com/onsonr/sonr/pkg/nebula/router"
)

type Server struct {
	*echo.Echo
}

func New() *Server {
	s := &Server{Echo: echo.New()}
	s.Logger.SetLevel(log.INFO)
	s.Use(session.UseSession)

	// Configure the server
	if err := nebula.UseAssets(s.Echo); err != nil {
		s.Logger.Fatal(err)
	}

	s.GET("/", router.Home)
	s.GET("/login", router.Login)
	s.GET("/register", router.Register)
	// s.GET("/profile", router.Profile)
	// s.GET("/authorize", router.Authorize)
	return s
}

func (s *Server) Start() {
	if err := s.Echo.Start(":1323"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
