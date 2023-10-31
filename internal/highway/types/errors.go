package types

import (
	"fmt"
	"strings"
)

var (
	// ErrNullWebauthnCredential is returned when a null credential is provided
	ErrNullWebauthnCredential = fmt.Errorf("provided webauthn credential is null")

	// ErrNullVaultKeyshare is returned when a null keyshare is provided
	ErrNullVaultKeyshare = fmt.Errorf("provided vault keyshare is null")

	// ErrInvalidCredentailCount is returned when the credential count is not 1
	ErrInvalidCredentailCount = fmt.Errorf("expected 1 credential formatted as json")

	// ErrJWTExpired is returned when the jwt has expired
	ErrJWTExpired = fmt.Errorf("jwt has expired")

	// ErrInvalidEmailController is returned when the controller is not a valid email controller
	ErrInvalidEmailController = fmt.Errorf("controller is not a valid email controller containing 'did:email:'")

	// ErrInvalidEmailAddress is returned when the email address is not properly formatted
	ErrInvalidEmailAddress = fmt.Errorf("email address is not properly formatted")

	// ErrFailedToRenameKeyshare is returned when the keyshare could not be renamed
	ErrFailedToRenameKeyshare = fmt.Errorf("failed to rename keyshare")

	// ErrNotFoundInPendingSet is returned when the did is not in the pending set for this vault
	ErrNotFoundInPendingSet = fmt.Errorf("this did is not in the pending set for this vault")

	// ErrNotFoundInControlledSet is returned when the did is not in the controlled set for this vault
	ErrNotFoundInControlledSet = fmt.Errorf("this did is not in the controlled set for this vault")

	// ErrInvalidKeyshareIndex is returned when the keyshare index is not valid
	ErrInvalidKeyshareIndex = fmt.Errorf("keyshare index is not valid")
)

// IsValidEmailAddress returns true if the email is valid
func IsValidEmailAddress(email string) bool {
	ptrs := strings.Split(email, "@")
	if len(ptrs) != 2 {
		return false
	}
	if !strings.Contains(ptrs[1], ".") {
		return false
	}
	return true
}

// IsValidEmailController returns true if the email did format is valid
func IsValidEmailController(emailDid string) bool {
	return strings.HasPrefix(emailDid, "did:email:")
}

// IsValidUnclaimedKeyshare returns true if the keyshare is valid for its DID format
func IsValidUnclaimedKeyshare(keyshare string) bool {
	ptrs := strings.Split(keyshare, ":")
	if len(ptrs) != 3 {
		return false
	}
	ksFrag := strings.Split(ptrs[2], "#")[1]
	if !strings.Contains(ksFrag, "ucw") {
		return false
	}
	return true
}
