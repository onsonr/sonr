package lobby

import (
	"time"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/api"
	common "github.com/sonr-io/core/pkg/common"
)

var (
	logger = golog.Child("protocols/lobby")
)

// LobbyOption is a function that modifies the Lobby options.
type LobbyOption func(*lobbyOptions)

// DisableAutoPush disables the auto push of the Lobby for OLC
func DisableAutoPush() LobbyOption {
	return func(o *lobbyOptions) {
		o.autoPushEnabled = false
	}
}

// WithLocation sets the location of the Lobby for OLC
func WithLocation(l *common.Location) LobbyOption {
	return func(o *lobbyOptions) {
		if o.location != nil {
			if o.location.GetLatitude() != 0 && o.location.GetLongitude() != 0 {
				logger.Debug("Skipping Location Set")
			} else {
				o.location = l
			}
		}
	}
}

// WithInterval sets the interval of the Lobby for OLC
func WithInterval(i time.Duration) LobbyOption {
	return func(o *lobbyOptions) {
		o.interval = i
	}
}

// lobbyOptions is a collection of options for the Lobby.
type lobbyOptions struct {
	location        *common.Location
	interval        time.Duration
	autoPushEnabled bool
}

func defaultLobbyOptions() *lobbyOptions {
	return &lobbyOptions{
		location:        api.GetLocation(),
		interval:        time.Second * 5,
		autoPushEnabled: true,
	}
}
