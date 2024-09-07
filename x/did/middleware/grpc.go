package middleware

import (
	"context"
	"fmt"
	"net"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type ClientInfo struct {
	Authority   string
	ContentType string
	UserAgent   string
	Hostname    string
	IPAddress   string
}

func ExtractClientInfo(ctx context.Context) (*ClientInfo, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get metadata from context")
	}

	info := &ClientInfo{}

	// Extract authority, content-type, and user-agent
	if authority := md.Get("authority"); len(authority) > 0 {
		info.Authority = authority[0]
	}
	if contentType := md.Get("content-type"); len(contentType) > 0 {
		info.ContentType = contentType[0]
	}
	if userAgent := md.Get("user-agent"); len(userAgent) > 0 {
		info.UserAgent = userAgent[0]
	}

	// Extract hostname and IP address
	p, ok := peer.FromContext(ctx)
	if ok {
		if tcpAddr, ok := p.Addr.(*net.TCPAddr); ok {
			info.IPAddress = tcpAddr.IP.String()

			// Try to get hostname
			names, err := net.LookupAddr(info.IPAddress)
			if err == nil && len(names) > 0 {
				info.Hostname = strings.TrimSuffix(names[0], ".")
			}
		}
	}

	return info, nil
}
