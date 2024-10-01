package nebula

import (
	"context"
	"embed"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/models"
)

//go:embed assets
var embeddedFiles embed.FS

//go:embed assets/config.pkl
var config string

func getHTTPFS() (http.FileSystem, error) {
	fsys, err := fs.Sub(embeddedFiles, "assets")
	if err != nil {
		return nil, err
	}
	return http.FS(fsys), nil
}

// UseAssets is a middleware that serves static files from the embedded assets
func UseAssets(e *echo.Echo) error {
	err := models.LoadFromString(context.Background(), config)
	if err != nil {
		return err
	}
	fsys, err := getHTTPFS()
	if err != nil {
		return err
	}
	assets := http.FileServer(fsys)
	e.GET("/", echo.WrapHandler(assets))
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", assets)))
	return nil
}
