package types

import (
	"fmt"
	"strings"
)

var (
	// ErrNullWebauthnCredential is returned when a null credential is provided
	ErrNullWebauthnCredential  = fmt.Errorf("provided webauthn credential is null")
	ErrNullVaultKeyshare       = fmt.Errorf("provided vault keyshare is null")
	ErrInvalidCredentailCount  = fmt.Errorf("expected 1 credential formatted as json")
	ErrJWTExpired              = fmt.Errorf("jwt has expired")
	ErrInvalidEmailController  = fmt.Errorf("controller is not a valid email controller containing 'did:email:'")
	ErrInvalidEmailAddress     = fmt.Errorf("email address is not properly formatted")
	ErrFailedToRenameKeyshare  = fmt.Errorf("failed to rename keyshare")
	ErrNotFoundInPendingSet    = fmt.Errorf("this did is not in the pending set for this vault")
	ErrNotFoundInControlledSet = fmt.Errorf("this did is not in the controlled set for this vault")
	ErrInvalidKeyshareIndex    = fmt.Errorf("keyshare index is not valid")
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
