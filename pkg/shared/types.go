package shared

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

// GRPCConn is a gRPC client connection.
type GRPCConn = *grpc.ClientConn

// EchoFunc is a function that takes an Echo instance and returns nothing.
type EchoFunc = func(c echo.Context) error
