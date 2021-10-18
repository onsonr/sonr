package device

import (
	"errors"

	"github.com/kataras/golog"
)

// Error definitions
var (
	logger = golog.Child("internal/device")
	// General errors
	ErrEmptyDeviceID = errors.New("Device ID cannot be empty")
	ErrMissingEnvVar = errors.New("Cannot set EnvVariable with empty value")

	// Directory errors
	ErrDirectoryInvalid = errors.New("Directory Type is invalid")
	ErrDirectoryUnset   = errors.New("Directory path has not been set")
	ErrDirectoryJoin    = errors.New("Failed to join directory path")
)

// // NewRecordPrefix returns a new device ID prefix for users HDNS records
// func NewRecordPrefix(sName string) (string, error) {
// 	// Check if the device ID is empty
// 	if deviceID == "" {
// 		return "", ErrEmptyDeviceID
// 	}

// 	// Check if the SName is empty
// 	if sName == "" {
// 		return "", errors.New("SName cannot by Empty or Less than 4 characters.")
// 	}
// 	val := fmt.Sprintf("%s:%s", deviceID, sName)
// 	return keychain.Primary.SignHmacWith(keychain.Account, val)
// }

// // VerifyRecordPrefix returns true if the prefix is valid for the device ID.
// func VerifyRecordPrefix(prefix string, sName string) bool {
// 	// Check if the prefix is empty
// 	if prefix == "" {
// 		logger.Warn("Empty Prefix Provided as Parameter")
// 		return false
// 	}

// 	// Check if the prefix is valid
// 	val := fmt.Sprintf("%s:%s", deviceID, sName)
// 	ok, err := keychain.Primary.VerifyHmacWith(keychain.Account, prefix, val)
// 	if err != nil {
// 		logger.Error("Failed to verify prefix", err)
// 		return false
// 	}
// 	return ok
// }
