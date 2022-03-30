package user

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrNameTooShort    = errors.New("Provided username is too short")
	ErrUsernameInvalid = errors.New("Provided username contains invalid characters")
	re                 = regexp.MustCompile("^[a-z0-9]*$")
)

// CleanName checks if the username is available
func CleanNameForSuffix(name string) string {
	if strings.Contains(name, ".snr") {
		return name
	}
	return name + ".snr"
}

// ValidateName validates the username for valid characters and length
func ValidateName(name string) error {
	// Check for valid length
	if len(name) < 3 {
		return ErrNameTooShort
	}

	// Check for valid characters
	if !re.MatchString(name) {
		return errors.New("Username contains invalid characters")
	}
	return nil
}
