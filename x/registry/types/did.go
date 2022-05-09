package types

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
)

// GenerateNameDid generates a new did document
func GenerateApplicationDid(accountAddr, appToRegister string, cred *Credential) (*did.Document, error) {
	// Generate a new DID String
	baseDid, err := did.ParseDID(fmt.Sprintf("did:snr:%s", strings.TrimPrefix(accountAddr, "snr")))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse account address")
	}

	// Generate a new DID String
	appDid, err := did.ParseDID(fmt.Sprintf("did:snr:%s/%s", strings.TrimPrefix(accountAddr, "snr"), appToRegister))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse base DID")
	}

	// Get PubKey from Credential
	pubKey, err := cred.DecodePublicKey()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create new VerificationMethod
	verf, err := did.NewVerificationMethod(*appDid, ssi.JsonWebKey2020, *appDid, pubKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Define function to generate AlsoKnownAs field
	genAlsoKnownAs := func(ss ...string) []string {
		var r []string
		for _, v := range ss {
			if strings.HasPrefix(v, "/") {
				r = append(r, v)
			} else {
				r = append(r, fmt.Sprintf("/%s", v))
			}
		}
		return r
	}

	// Empty did document:
	doc := &did.Document{
		Context:     []ssi.URI{did.DIDContextV1URI()},
		ID:          *appDid,
		AlsoKnownAs: genAlsoKnownAs(appToRegister),
		Controller: []did.DID{
			*baseDid,
		},
	}

	// This adds the method to the VerificationMethod list and stores a reference to the assertion list
	doc.AddAssertionMethod(verf)
	return doc, nil
}

// GenerateNameDid generates a new did document
func GenerateNameDid(accountAddr, nameToRegister string, cred *Credential) (*did.Document, error) {
	// Generate a new DID String
	baseDid, err := did.ParseDID(fmt.Sprintf("did:snr:%s", strings.TrimPrefix(accountAddr, "snr")))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse base DID")
	}

	// Get PubKey from Credential
	pubKey, err := cred.DecodePublicKey()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create new VerificationMethod
	verf, err := did.NewVerificationMethod(*baseDid, ssi.JsonWebKey2020, *baseDid, pubKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Define function to generate AlsoKnownAs field
	genAlsoKnownAs := func(ss ...string) []string {
		var r []string
		for _, v := range ss {
			if strings.HasSuffix(v, ".snr") {
				r = append(r, v)
			} else {
				r = append(r, fmt.Sprintf("%s.snr", v))
			}
		}
		return r
	}

	// Empty did document:
	doc := &did.Document{
		Context:     []ssi.URI{did.DIDContextV1URI()},
		ID:          *baseDid,
		AlsoKnownAs: genAlsoKnownAs(nameToRegister),
		Controller: []did.DID{
			*baseDid,
		},
	}

	// This adds the method to the VerificationMethod list and stores a reference to the assertion list
	doc.AddAssertionMethod(verf)
	return doc, nil
}

// ValidateAppName checks if the given name to register is correct length and valid characters
func ValidateAppName(rs string) (string, error) {
	// Create trim suffix function
	trimSuffix := func(s string) string {
		return strings.TrimPrefix(s, "/")
	}

	// Trim suffix if it exists
	s := trimSuffix(rs)

	// Check for valid length
	if len(s) < 3 {
		return "", ErrNameTooShort
	}

	isAlpha := regexp.MustCompile(`^[0-9a-z]+$`).MatchString
	if !isAlpha(s) {
		return "", sdkerrors.Wrapf(ErrNameInvalid, "invalid name to register (%s)", s)
	}

	return s, nil
}

// ValidateName checks if the given name to register is correct length and valid characters
func ValidateName(rs string) (string, error) {
	// Create trim suffix function
	trimSuffix := func(s string) string {
		return strings.TrimSuffix(s, ".snr")
	}

	// Trim suffix if it exists
	s := trimSuffix(rs)

	// Check for valid length
	if len(s) < 3 {
		return "", ErrNameTooShort
	}

	isAlpha := regexp.MustCompile(`^[0-9a-z]+$`).MatchString
	if !isAlpha(s) {
		return "", sdkerrors.Wrapf(ErrNameInvalid, "invalid name to register (%s)", s)
	}

	return s, nil
}
