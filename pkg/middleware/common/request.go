package common

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func Requests(c echo.Context) requests {
	return requests{
		Context: c,
	}
}

type requests struct {
	echo.Context
}

func (r requests) PathIs(path string) bool {
	// remove trailing slash
	cp := strings.TrimSuffix(path, "/")
	return strings.Contains(r.Request().URL.Path, cp)
}
