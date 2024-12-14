package resolver

import (
	"net/http"

	"github.com/labstack/echo/v4"
	config "github.com/onsonr/sonr/pkg/config/hway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientsContext struct {
	echo.Context
	addr string
}

func GetClientConn(c echo.Context) (*grpc.ClientConn, error) {
	cc, ok := c.(*ClientsContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "ClientsContext not found")
	}
	grpcConn, err := grpc.NewClient(cc.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to dial gRPC")
	}
	return grpcConn, nil
}

func Middleware(env config.Hway) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ClientsContext{Context: c, addr: env.GetSonrGrpcUrl()}
			return next(cc)
		}
	}
}
