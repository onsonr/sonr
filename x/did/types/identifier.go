package types

import (
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/di-dao/core/crypto/accumulator"
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
		return "", ErrInvalidEmailFormat
	}
	return EmailDID("did:email:" + hex.EncodeToString(Blake3Hash([]byte(email)))), nil
}

// GetPhoneDID returns the DID representation of the phone number
func GetPhoneDID(phone string) (PhoneDID, error) {
	re := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	if !re.MatchString(phone) {
		return "", ErrInvalidPhoneFormat
	}
	return PhoneDID("did:phone:" + hex.EncodeToString(Blake3Hash([]byte(phone)))), nil
}

// ConvertMapToPropertyList converts a map of accumulators to a list of properties
func ConvertMapToPropertyList(propertyMap map[string]*accumulator.Accumulator) ([]*Property, error) {
	properties := make([]*Property, 0, len(propertyMap))
	for k, v := range propertyMap {
		accBz, err := v.MarshalBinary()
		if err != nil {
			return nil, err
		}
		properties = append(properties, &Property{
			Key:         k,
			Accumulator: accBz,
		})
	}
	return properties, nil
}

// ConvertPropertyListToMap converts a list of properties to a map of accumulators
func ConvertPropertyListToMap(properties []*Property) (map[string]*accumulator.Accumulator, error) {
	propertyMap := make(map[string]*accumulator.Accumulator)
	for _, p := range properties {
		acc := &accumulator.Accumulator{}
		if err := acc.UnmarshalBinary(p.Accumulator); err != nil {
			return nil, err
		}
		propertyMap[p.Key] = acc
	}
	return propertyMap, nil
}
