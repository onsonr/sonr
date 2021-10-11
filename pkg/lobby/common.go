package lobby

import (
	"time"

	"github.com/kataras/golog"
	common "github.com/sonr-io/core/internal/common"
)

var (
	logger = golog.Child("protocols/lobby")
)

type GetPeerFunc func() (*common.Peer, error)

// LobbyOption is a function that modifies the Lobby options.
type LobbyOption func(*lobbyOptions)

// WithLocation sets the location of the Lobby for OLC
func WithLocation(l *common.Location) LobbyOption {
	return func(o *lobbyOptions) {
		if o.location != nil {
			if o.location.GetLatitude() != 0 && o.location.GetLongitude() != 0 {
				logger.Info("Skipping Location Set")
			} else {
				o.location = l
			}
		} else {
			o.location = l
		}
	}
}

// lobbyOptions is a collection of options for the Lobby.
type lobbyOptions struct {
	location *common.Location
	peerFunc GetPeerFunc
	interval time.Duration
}

func defaultLobbyOptions() *lobbyOptions {
	return &lobbyOptions{
		location: common.DefaultLocation(),
		interval: time.Second * 5,
	}
}
