package clients

import (
	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	"github.com/labstack/echo/v4"
	didv1 "github.com/onsonr/sonr/api/did/v1"
	dwnv1 "github.com/onsonr/sonr/api/dwn/v1"
	svcv1 "github.com/onsonr/sonr/api/svc/v1"
)

func BankQueryClient(c echo.Context) (bankv1beta1.QueryClient, error) {
	conn, err := GetClientConn(c)
	if err != nil {
		return nil, err
	}
	return bankv1beta1.NewQueryClient(conn), nil
}

func DIDQueryClient(c echo.Context) (didv1.QueryClient, error) {
	conn, err := GetClientConn(c)
	if err != nil {
		return nil, err
	}
	return didv1.NewQueryClient(conn), nil
}

func DWNQueryClient(c echo.Context) (dwnv1.QueryClient, error) {
	conn, err := GetClientConn(c)
	if err != nil {
		return nil, err
	}
	return dwnv1.NewQueryClient(conn), nil
}

func SVCQueryClient(c echo.Context) (svcv1.QueryClient, error) {
	conn, err := GetClientConn(c)
	if err != nil {
		return nil, err
	}
	return svcv1.NewQueryClient(conn), nil
}
