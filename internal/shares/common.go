package shares

import (
	"fmt"

	"github.com/sonrhq/sonr/crypto/core/protocol"
)

// For DKG bob starts first. For refresh and sign, Alice starts first.
func runIteratedProtocol(firstParty protocol.Iterator, secondParty protocol.Iterator) error {
	var (
		message *protocol.Message
		aErr    error
		bErr    error
	)

	for aErr != protocol.ErrProtocolFinished || bErr != protocol.ErrProtocolFinished {
		// Crank each protocol forward one iteration
		message, bErr = firstParty.Next(message)
		if bErr != nil && bErr != protocol.ErrProtocolFinished {
			return bErr
		}

		message, aErr = secondParty.Next(message)
		if aErr != nil && aErr != protocol.ErrProtocolFinished {
			return aErr
		}
	}
	return checkProtocolErrors(aErr, bErr)
}

func checkProtocolErrors(aErr, bErr error) error {
	if aErr == protocol.ErrProtocolFinished && bErr == protocol.ErrProtocolFinished {
		return nil
	}
	if aErr != nil && bErr != nil {
		return fmt.Errorf("both parties failed: %v, %v", aErr, bErr)
	}
	if aErr != nil {
		return fmt.Errorf("alice failed: %v", aErr)
	}
	if bErr != nil {
		return fmt.Errorf("bob failed: %v", bErr)
	}
	return nil
}
