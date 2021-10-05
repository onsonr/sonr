package exchange

import (
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
)

func checkParams(host *host.SNRHost, em *state.Emitter) error {
	if host == nil {
		return logger.Error("Host provided is nil", ErrParameters)
	}
	if em == nil {
		return logger.Error("Emitter provided is nil", ErrParameters)
	}
	return nil
}
