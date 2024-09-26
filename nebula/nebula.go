package nebula

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

//go:embed assets
var embeddedFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		return http.FS(os.DirFS("assets"))
	}

	fsys, err := fs.Sub(embeddedFiles, "assets")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func UseAssets(e *echo.Echo) echo.HandlerFunc {
	assets := http.FileServer(getFileSystem(true))
	return echo.WrapHandler(assets)
}
