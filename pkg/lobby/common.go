package lobby

import (
	"fmt"
	"time"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/host"
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

// defaultLobbyOptions returns the default options for the Lobby.
func defaultLobbyOptions() *lobbyOptions {
	return &lobbyOptions{
		location:        api.GetLocation(),
		interval:        time.Second * 5,
		autoPushEnabled: true,
	}
}

// checkParams Checks if Non-nil Parameters were passed
func checkParams(host *host.SNRHost) error {
	if host == nil {
		logger.Error("Host provided is nil", ErrParameters)
		return ErrParameters
	}
	return host.HasRouting()
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
