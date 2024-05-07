package did

import (
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/di-dao/core/x/did/types"
	"lukechampine.com/blake3"
)

// EmailDID is the DID method for email addresses
type EmailDID string

// Equals returns true if a provided email string is equal to the hash of the email DID
func (e EmailDID) Equals(email string) bool {
	other, err := GetEmailDID(email)
	if err != nil {
		return false
	}
	return strings.EqualFold(e.String(), other.String())
}

// String returns the string representation of the email DID
func (e EmailDID) String() string {
	return string(e)
}

// PhoneDID is the DID method for phone numbers
type PhoneDID string

// Equals returns true if a provided phone string is equal to the hash of the phone DID
func (p PhoneDID) Equals(phone string) bool {
	other, err := GetPhoneDID(phone)
	if err != nil {
		return false
	}
	return strings.EqualFold(p.String(), other.String())
}

// String returns the string representation of the phone DID
func (p PhoneDID) String() string {
	return string(p)
}

// Blake3Hash returns the blake3 hash of the input bytes
func Blake3Hash(bz []byte) []byte {
	bz32 := blake3.Sum256(bz)
	return bz32[:]
}

// GetEmailDID returns the DID representation of the email address
func GetEmailDID(email string) (EmailDID, error) {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return "", types.ErrInvalidEmailFormat
	}
	return EmailDID("did:email:" + hex.EncodeToString(Blake3Hash([]byte(email)))), nil
}

// GetPhoneDID returns the DID representation of the phone number
func GetPhoneDID(phone string) (PhoneDID, error) {
	re := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	if !re.MatchString(phone) {
		return "", types.ErrInvalidPhoneFormat
	}
	return PhoneDID("did:phone:" + hex.EncodeToString(Blake3Hash([]byte(phone)))), nil
}
