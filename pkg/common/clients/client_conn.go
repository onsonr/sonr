package clients

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
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
	grpcConn, err := grpc.NewClient(cc.addr, grpc.WithInsecure())
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to dial gRPC")
	}
	return grpcConn, nil
}

func GRPCClientsMiddleware(addr string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ClientsContext{Context: c, addr: addr}
			return next(cc)
		}
	}
}
