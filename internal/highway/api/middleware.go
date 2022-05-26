package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HighwayMiddleware struct {
	definition gin.HandlerFunc // middleware definition
	disabled   bool            // should be loaded on registration
}

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

var middlewareDefs []HighwayMiddleware = nil

func (s *HighwayServer) RegisterMiddleWare() {
	// Add middleware here to be loaded into Highway's Router
	middlewareDefs = []HighwayMiddleware{
		{
			definition: func(ctx *gin.Context) {
				token := ctx.GetHeader("Authorization")
				if token == "" {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
						Message: "Authorization token not found",
					})
				}
				error := s.JwtToken.BuildJWTParseMiddleware(token)()

				if error != nil {
					logger.Errorf("Error while processing authorization header: %s", error.Error())
					ctx.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
						Message: error.Error(),
					})
					return
				}

				ctx.Next()
			},
			disabled: true,
		},
	}

	for _, mw := range middlewareDefs {
		if !mw.disabled {
			s.Router.Use(mw.definition)
		}
	}
}
