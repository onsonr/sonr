package core

import (
	"runtime"
	"time"

	v1 "github.com/sonr-io/sonr/internal/motor/x/discover/v1"
)

type GetPeerFunc func() *v1.Peer

func DefaultGetPeerFunc() GetPeerFunc {
	return func() *v1.Peer {
		return &v1.Peer{
			Os:       runtime.GOOS,
			Arch:     runtime.GOARCH,
			LastSeen: time.Now().Unix(),
		}
	}
}
