package exchange

import (
	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/host"
)

var (
	logger = golog.Child("protocols/exchange")
)

// checkParams Checks if Non-nil Parameters were passed
func checkParams(host *host.SNRHost) error {
	if host == nil {
		logger.Error("Host provided is nil", ErrParameters)
		return ErrParameters
	}
	return host.HasRouting()
}
