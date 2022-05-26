package api

import "github.com/gin-gonic/gin"

var middlewareDefs []gin.HandlerFunc = nil

func (s *HighwayServer) RegisterMiddleWare() {
	// Add middleware here to be loaded into Highway's Router
	middlewareDefs = []gin.HandlerFunc{
		func(ctx *gin.Context) {
			token := ctx.GetHeader("Authorization")
			if token != "" {
				error := s.JwtToken.BuildJWTParseMiddleware(token)()

				if error != nil {
					logger.Errorf("Error while processing authorization header: %s", error.Error())
				}
				ctx.Next()
			}
		},
	}

	for _, mw := range middlewareDefs {
		s.Router.Use(mw)
	}
}
