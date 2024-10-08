package nebula

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed assets
var embeddedFiles embed.FS

func getHTTPFS() (http.FileSystem, error) {
	fsys, err := fs.Sub(embeddedFiles, "assets")
	if err != nil {
		return nil, err
	}
	return http.FS(fsys), nil
}

// UseAssets is a middleware that serves static files from the embedded assets
func UseAssets(e *echo.Echo) error {
	fsys, err := getHTTPFS()
	if err != nil {
		return err
	}
	assets := http.FileServer(fsys)
	e.GET("/", echo.WrapHandler(assets))
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", assets)))
	return nil
}
