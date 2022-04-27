package types

// DONTCOVER

import (
	fmt "fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/registry module sentinel errors
var (
	// Validator Errors
	ErrNameTooShort     = sdkerrors.Register(ModuleName, 1100, "Provided Sonr Name is Too Short")
	ErrNameInvalid      = sdkerrors.Register(ModuleName, 1101, "Provided Sonr Name contains invalid characters")
	ErrNameRegistered   = sdkerrors.Register(ModuleName, 1102, "Provided Sonr Name has already been registered")
	ErrInvalidWhoisType = sdkerrors.Register(ModuleName, 1103, "Returned whois type for name is not of required type")
	
	// COSEKey Errors
	ErrDecodeAttestedCredentialData = Error{err: "error decoding attested credential data"}
	ErrDecodeAuthenticatorData      = Error{err: "error decoding authenticator data"}
	ErrDecodeCOSEKey                = Error{err: "error decoding raw public key"}
	ErrECDAANotSupported            = Error{err: "ECDAA not supported"}
	ErrEncodeAttestedCredentialData = Error{err: "error encoding attested credential data"}
	ErrEncodeAuthenticatorData      = Error{err: "error encoding authenticator data"}
	ErrGenerateChallenge            = Error{err: "error generating challenge"}
	ErrMarshalAttestationObject     = Error{err: "error marshaling attestation object"}
	ErrOption                       = Error{err: "option error"}
	ErrNotImplemented               = Error{err: "not implemented"}
	ErrUnmarshalAttestationObject   = Error{err: "error unmarshaling attestation object"}
	ErrVerifyAttestation            = Error{err: "error verifying attestation"}
	ErrVerifyAuthentication         = Error{err: "error verifying authentication"}
	ErrVerifyClientExtensionOutput  = Error{err: "error verifying client extension output"}
	ErrVerifyRegistration           = Error{err: "error verifying registration"}
	ErrVerifySignature              = Error{err: "error verifying signature"}
)

//Error represents an error in a WebAuthn relying party operation
type Error struct {
	err     string
	wrapped error
}

//Error implements the error interface
func (e Error) Error() string {
	return e.err
}

//Unwrap allows for error unwrapping
func (e Error) Unwrap() error {
	return e.wrapped
}

//Wrap returns a new error which contains the provided error wrapped with this
//error
func (e Error) Wrap(err error) Error {
	n := e
	n.wrapped = err
	return n
}

//Is establishes equality for error types
func (e Error) Is(target error) bool {
	return e.Error() == target.Error()
}

//NewError returns a new Error with a custom message
func NewError(fmStr string, els ...interface{}) Error {
	return Error{
		err: fmt.Sprintf(fmStr, els...),
	}
}

//Categorical top-level errors
var ()
