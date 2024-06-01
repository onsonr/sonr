package claims

import (
	"time"
)

// Returns the issued at, not before and expiration timescales for credential authentication. (now, now, 1h)
func getCredentialTimescale() (time.Time, time.Time, time.Time) {
	curr := time.Now()
	return curr, curr, curr.Add(time.Hour)
}

// Returns the issued at, not before and expiration timescales for controller authentication. (now, now, 30m)
func getControllerTimescale() (time.Time, time.Time, time.Time) {
	curr := time.Now()
	return curr, curr, curr.Add(time.Minute * 30)
}
