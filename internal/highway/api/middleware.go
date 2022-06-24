package api

import (
	"github.com/gin-gonic/gin"
)

type HighwayMiddleware struct {
	definition gin.HandlerFunc // middleware definition
	disabled   bool            // should be loaded on registration
}

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

var (
	middlewareDefs []HighwayMiddleware = make([]HighwayMiddleware, 0)
)

func (s *HighwayServer) AddMiddlewareDefinition(middlewareDef HighwayMiddleware) {
	middlewareDefs = append(middlewareDefs, middlewareDef)
}

func (s *HighwayServer) AddMiddlewareDefinitions(middlewareDefs []HighwayMiddleware) {
	middlewareDefs = append(middlewareDefs, middlewareDefs...)
}

func (s *HighwayServer) RegisterMiddleWare() {

	for _, mw := range middlewareDefs {
		if !mw.disabled {
			s.Router.Use(mw.definition)
		}
	}
}
