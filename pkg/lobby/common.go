package lobby

import (
	"fmt"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/api"
	common "github.com/sonr-io/core/pkg/common"
)

var (
	logger = golog.Default.Child("protocols/lobby")
)

// LobbyOption is a function that modifies the Lobby options.
type LobbyOption func(*options)

// DisableAutoPush disables the auto push of the Lobby for OLC
func DisableAutoPush() LobbyOption {
	return func(o *options) {
		o.autoPushEnabled = false
	}
}

// WithLocation sets the location of the Lobby for OLC
func WithLocation(l *common.Location) LobbyOption {
	return func(o *options) {
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
	return func(o *options) {
		o.interval = i
	}
}

// options is a collection of options for the Lobby.
type options struct {
	location        *common.Location
	interval        time.Duration
	autoPushEnabled bool
}

// defaultOptions returns the default options for the Lobby.
func defaultOptions() *options {
	return &options{
		location:        api.DefaultLocation(),
		interval:        time.Second * 5,
		autoPushEnabled: true,
	}
}

// createOlc Creates a new Olc from Location
func createOlc(l *common.Location) string {
	code := l.OLC()
	if code == "" {
		logger.Error("Failed to Determine OLC Code, set to Global")
		code = "global"
	}
	logger.Debug("Calculated OLC for Location: " + code)
	return fmt.Sprintf("sonr/topic/%s", code)
}

type LobbyEvent struct {
	ID     peer.ID
	Peer   *common.Peer
	isExit bool
}

func newLobbyEvent(i peer.ID, p *common.Peer) *LobbyEvent {
	if p == nil {
		return &LobbyEvent{
			ID:     i,
			isExit: true,
		}
	}
	return &LobbyEvent{
		ID:     i,
		Peer:   p,
		isExit: false,
	}
}
