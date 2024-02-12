package types

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"

	servicev1 "github.com/sonrhq/sonr/api/sonr/service/module/v1"
)

type AuthenticatorAttachment = protocol.AuthenticatorAttachment

type COSEAlgorithmIdentifier = webauthncose.COSEAlgorithmIdentifier

type CredentialDescriptor = protocol.CredentialDescriptor

type PublicKeyCredential = protocol.PublicKeyCredential

type PublicKeyCredentialRequestOptions = protocol.PublicKeyCredentialRequestOptions

type PublicKeyCredentialCreationOptions = protocol.PublicKeyCredentialCreationOptions

type ServiceRecord = servicev1.ServiceRecord

type URLEncodedBase64 = protocol.URLEncodedBase64

type UserEntity = protocol.UserEntity
