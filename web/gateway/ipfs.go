package gateway

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/labstack/echo/v4"
)

type Gateway struct {
	api *rpc.HttpApi
}

func New() (*Gateway, error) {
	api, err := rpc.NewLocalApi()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to local IPFS node: %w", err)
	}

	return &Gateway{
		api: api,
	}, nil
}

func (g *Gateway) Handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		fullPath := c.Request().URL.Path
		segments := strings.Split(strings.Trim(fullPath, "/"), "/")

		if len(segments) == 0 {
			return redirectOnError(c)
		}

		cid := segments[0]
		if !looksLikeCID(cid) {
			return redirectOnError(c)
		}

		remainingPath := "/"
		if len(segments) > 1 {
			remainingPath = "/" + strings.Join(segments[1:], "/")
		}

		// Always require a path
		if remainingPath == "/" {
			return redirectOnError(c)
		}

		ipfsPath, err := path.NewPath(fmt.Sprintf("/ipfs/%s%s", cid, remainingPath))
		if err != nil {
			return redirectOnError(c)
		}

		ctx, cancel := context.WithTimeout(c.Request().Context(), 30*time.Second)
		defer cancel()

		node, err := g.api.Unixfs().Get(ctx, ipfsPath)
		if err != nil {
			c.Logger().Errorf("IPFS error: %v", err)
			return redirectOnError(c)
		}

		switch n := node.(type) {
		case files.File:
			if err := streamFile(c, n, cid, remainingPath); err != nil {
				c.Logger().Errorf("Streaming error: %v", err)
				return redirectOnError(c)
			}
			return nil
		case files.Directory:
			return redirectOnError(c)
		default:
			return redirectOnError(c)
		}
	}
}

func getContentType(path string, defaultType string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".html", ".htm":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	case ".json":
		return "application/json"
	case ".xml":
		return "application/xml"
	case ".pdf":
		return "application/pdf"
	case ".zip":
		return "application/zip"
	case ".mp4":
		return "video/mp4"
	case ".mp3":
		return "audio/mpeg"
	case ".woff":
		return "font/woff"
	case ".woff2":
		return "font/woff2"
	case ".ttf":
		return "font/ttf"
	default:
		return defaultType
	}
}

func streamFile(c echo.Context, file files.File, cid string, filePath string) error {
	// Get file size if possible
	stat, err := file.Size()
	if err == nil {
		c.Response().Header().Set("Content-Length", fmt.Sprintf("%d", stat))
	}

	// Set content type based on file extension first
	contentType := getContentType(filePath, "")

	// If no content type found by extension, detect from content
	if contentType == "" {
		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		contentType = http.DetectContentType(buffer)

		// Reset file pointer after reading for content type detection
		if seeker, ok := file.(io.Seeker); ok {
			_, err = seeker.Seek(0, io.SeekStart)
			if err != nil {
				return err
			}
		}
	}

	// Set headers
	c.Response().Header().Set("Content-Type", contentType)
	c.Response().Header().Set("ETag", fmt.Sprintf(`"%s"`, cid))
	c.Response().Header().Set("Cache-Control", "public, max-age=29030400, immutable")
	c.Response().Header().Set("X-Content-Type-Options", "nosniff")
	c.Response().Header().Set("X-Frame-Options", "DENY")
	c.Response().Header().Set("X-XSS-Protection", "1; mode=block")

	// Stream the file
	return c.Stream(http.StatusOK, contentType, file)
}

func redirectOnError(c echo.Context) error {
	return c.Redirect(http.StatusFound, "http://localhost:3000")
}

func looksLikeCID(s string) bool {
	return strings.HasPrefix(s, "Qm") || // v0 CID
		strings.HasPrefix(s, "bafy") || // v1 CID
		strings.HasPrefix(s, "b") // other base32 v1 CIDs
}
