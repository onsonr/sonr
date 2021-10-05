package lobby

import (
	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
)

func checkParams(host *host.SNRHost, loc *common.Location, em *state.Emitter) error {
	if host == nil {
		return logger.Error("Host provided is nil", ErrParameters)
	}
	if loc == nil {
		return logger.Error("Location provided is nil", ErrParameters)
	}
	if em == nil {
		return logger.Error("Emitter provided is nil", ErrParameters)
	}
	return nil
}
