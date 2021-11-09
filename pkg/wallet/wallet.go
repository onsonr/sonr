package wallet

import (
	"fmt"

	"github.com/sonr-io/core/pkg/device"
)

// NewRecordPrefix returns a new device ID prefix for users HDNS records
func NewRecordPrefix(sName string) (string, error) {
	// Check if the device ID is empty
	deviceId, err := device.ID()
	if err != nil {
		return "", err
	}

	val := fmt.Sprintf("%s:%s", deviceId, sName)
	return Sonr.SignHmacWith(Account, val)
}

// VerifyRecordPrefix returns true if the prefix is valid for the device ID.
func VerifyRecordPrefix(prefix string, sName string) bool {
	// Check if the prefix is empty
	if prefix == "" {
		logger.Warn("Empty Prefix Provided as Parameter")
		return false
	}

	// Check if the device ID is empty
	deviceId, err := device.ID()
	if err != nil {
		logger.Errorf("Failed to get device ID: %s", err)
		return false
	}

	// Check if the prefix is valid
	val := fmt.Sprintf("%s:%s", deviceId, sName)
	ok, err := Sonr.VerifyHmacWith(Account, prefix, val)
	if err != nil {
		logger.Errorf("%s - Failed to verify prefix", err)
		return false
	}
	return ok
}
