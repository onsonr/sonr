package middleware

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/labstack/echo/v4"
)

type IPFSContext struct {
	echo.Context
	ipfs *rpc.HttpApi
}

func IPFSMiddleware(client *rpc.HttpApi) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &IPFSContext{
				Context: c,
				ipfs:    client,
			}
			return next(cc)
		}
	}
}

func GetIPFSClient(c echo.Context) (*rpc.HttpApi, error) {
	cc, ok := c.(*IPFSContext)
	if !ok {
		return nil, errors.New("not an IPFSContext")
	}
	if cc.ipfs == nil {
		return nil, errors.New("no IPFS client")
	}
	return cc.ipfs, nil
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
