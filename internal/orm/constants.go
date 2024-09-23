package orm

import (
	"github.com/onsonr/sonr/x/did/types"
)

type (
	AuthenticatorAttachment string
	AuthenticatorTransport  string
)

const (
	// Platform represents a platform authenticator is attached using a client device-specific transport, called
	// platform attachment, and is usually not removable from the client device. A public key credential bound to a
	// platform authenticator is called a platform credential.
	Platform AuthenticatorAttachment = "platform"

	// CrossPlatform represents a roaming authenticator is attached using cross-platform transports, called
	// cross-platform attachment. Authenticators of this class are removable from, and can "roam" among, client devices.
	// A public key credential bound to a roaming authenticator is called a roaming credential.
	CrossPlatform AuthenticatorAttachment = "cross-platform"
)

func ParseAuthenticatorAttachment(s string) AuthenticatorAttachment {
	switch s {
	case "platform":
		return Platform
	default:
		return CrossPlatform
	}
}

const (
	// USB indicates the respective authenticator can be contacted over removable USB.
	USB AuthenticatorTransport = "usb"

	// NFC indicates the respective authenticator can be contacted over Near Field Communication (NFC).
	NFC AuthenticatorTransport = "nfc"

	// BLE indicates the respective authenticator can be contacted over Bluetooth Smart (Bluetooth Low Energy / BLE).
	BLE AuthenticatorTransport = "ble"

	// SmartCard indicates the respective authenticator can be contacted over ISO/IEC 7816 smart card with contacts.
	//
	// WebAuthn Level 3.
	SmartCard AuthenticatorTransport = "smart-card"

	// Hybrid indicates the respective authenticator can be contacted using a combination of (often separate)
	// data-transport and proximity mechanisms. This supports, for example, authentication on a desktop computer using
	// a smartphone.
	//
	// WebAuthn Level 3.
	Hybrid AuthenticatorTransport = "hybrid"

	// Internal indicates the respective authenticator is contacted using a client device-specific transport, i.e., it
	// is a platform authenticator. These authenticators are not removable from the client device.
	Internal AuthenticatorTransport = "internal"
)

func ParseAuthenticatorTransport(s string) AuthenticatorTransport {
	switch s {
	case "usb":
		return USB
	case "nfc":
		return NFC
	case "ble":
		return BLE
	case "smart-card":
		return SmartCard
	case "hybrid":
		return Hybrid
	default:
		return Internal
	}
}

type AuthenticatorFlags byte

const (
	// FlagUserPresent Bit 00000001 in the byte sequence. Tells us if user is present. Also referred to as the UP flag.
	FlagUserPresent AuthenticatorFlags = 1 << iota // Referred to as UP

	// FlagRFU1 is a reserved for future use flag.
	FlagRFU1

	// FlagUserVerified Bit 00000100 in the byte sequence. Tells us if user is verified
	// by the authenticator using a biometric or PIN. Also referred to as the UV flag.
	FlagUserVerified

	// FlagBackupEligible Bit 00001000 in the byte sequence. Tells us if a backup is eligible for device. Also referred
	// to as the BE flag.
	FlagBackupEligible // Referred to as BE

	// FlagBackupState Bit 00010000 in the byte sequence. Tells us if a backup state for device. Also referred to as the
	// BS flag.
	FlagBackupState

	// FlagRFU2 is a reserved for future use flag.
	FlagRFU2

	// FlagAttestedCredentialData Bit 01000000 in the byte sequence. Indicates whether
	// the authenticator added attested credential data. Also referred to as the AT flag.
	FlagAttestedCredentialData

	// FlagHasExtensions Bit 10000000 in the byte sequence. Indicates if the authenticator data has extensions. Also
	// referred to as the ED flag.
	FlagHasExtensions
)

type AttestationFormat string

const (
	// AttestationFormatPacked is the "packed" attestation statement format is a WebAuthn-optimized format for
	// attestation. It uses a very compact but still extensible encoding method. This format is implementable by
	// authenticators with limited resources (e.g., secure elements).
	AttestationFormatPacked AttestationFormat = "packed"

	// AttestationFormatTPM is the TPM attestation statement format returns an attestation statement in the same format
	// as the packed attestation statement format, although the rawData and signature fields are computed differently.
	AttestationFormatTPM AttestationFormat = "tpm"

	// AttestationFormatAndroidKey is the attestation statement format for platform authenticators on versions "N", and
	// later, which may provide this proprietary "hardware attestation" statement.
	AttestationFormatAndroidKey AttestationFormat = "android-key"

	// AttestationFormatAndroidSafetyNet is the attestation statement format that Android-based platform authenticators
	// MAY produce an attestation statement based on the Android SafetyNet API.
	AttestationFormatAndroidSafetyNet AttestationFormat = "android-safetynet"

	// AttestationFormatFIDOUniversalSecondFactor is the attestation statement format that is used with FIDO U2F
	// authenticators.
	AttestationFormatFIDOUniversalSecondFactor AttestationFormat = "fido-u2f"

	// AttestationFormatApple is the attestation statement format that is used with Apple devices' platform
	// authenticators.
	AttestationFormatApple AttestationFormat = "apple"

	// AttestationFormatNone is the attestation statement format that is used to replace any authenticator-provided
	// attestation statement when a WebAuthn Relying Party indicates it does not wish to receive attestation information.
	AttestationFormatNone AttestationFormat = "none"
)

func ExtractAttestationFormats(p *types.Params) []AttestationFormat {
	var formats []AttestationFormat
	for _, v := range p.AttestationFormats {
		formats = append(formats, parseAttestationFormat(v))
	}
	return formats
}

func parseAttestationFormat(s string) AttestationFormat {
	switch s {
	case "packed":
		return AttestationFormatPacked
	case "tpm":
		return AttestationFormatTPM
	case "android-key":
		return AttestationFormatAndroidKey
	case "android-safetynet":
		return AttestationFormatAndroidSafetyNet
	case "fido-u2f":
		return AttestationFormatFIDOUniversalSecondFactor
	case "apple":
		return AttestationFormatApple
	case "none":
		return AttestationFormatNone
	default:
		return AttestationFormatPacked
	}
}

type CredentialType string

const (
	CredentialTypePublicKeyCredential CredentialType = "public-key"
)

type ConveyancePreference string

const (
	// PreferNoAttestation is a ConveyancePreference value.
	//
	// This value indicates that the Relying Party is not interested in authenticator attestation. For example, in order
	// to potentially avoid having to obtain user consent to relay identifying information to the Relying Party, or to
	// save a round trip to an Attestation CA or Anonymization CA.
	//
	// This is the default value.
	//
	// Specification: §5.4.7. Attestation Conveyance Preference Enumeration (https://www.w3.org/TR/webauthn/#dom-attestationconveyancepreference-none)
	PreferNoAttestation ConveyancePreference = "none"

	// PreferIndirectAttestation is a ConveyancePreference value.
	//
	// This value indicates that the Relying Party prefers an attestation conveyance yielding verifiable attestation
	// statements, but allows the client to decide how to obtain such attestation statements. The client MAY replace the
	// authenticator-generated attestation statements with attestation statements generated by an Anonymization CA, in
	// order to protect the user’s privacy, or to assist Relying Parties with attestation verification in a
	// heterogeneous ecosystem.
	//
	// Note: There is no guarantee that the Relying Party will obtain a verifiable attestation statement in this case.
	// For example, in the case that the authenticator employs self attestation.
	//
	// Specification: §5.4.7. Attestation Conveyance Preference Enumeration (https://www.w3.org/TR/webauthn/#dom-attestationconveyancepreference-indirect)
	PreferIndirectAttestation ConveyancePreference = "indirect"

	// PreferDirectAttestation is a ConveyancePreference value.
	//
	// This value indicates that the Relying Party wants to receive the attestation statement as generated by the
	// authenticator.
	//
	// Specification: §5.4.7. Attestation Conveyance Preference Enumeration (https://www.w3.org/TR/webauthn/#dom-attestationconveyancepreference-direct)
	PreferDirectAttestation ConveyancePreference = "direct"

	// PreferEnterpriseAttestation is a ConveyancePreference value.
	//
	// This value indicates that the Relying Party wants to receive an attestation statement that may include uniquely
	// identifying information. This is intended for controlled deployments within an enterprise where the organization
	// wishes to tie registrations to specific authenticators. User agents MUST NOT provide such an attestation unless
	// the user agent or authenticator configuration permits it for the requested RP ID.
	//
	// If permitted, the user agent SHOULD signal to the authenticator (at invocation time) that enterprise
	// attestation is requested, and convey the resulting AAGUID and attestation statement, unaltered, to the Relying
	// Party.
	//
	// Specification: §5.4.7. Attestation Conveyance Preference Enumeration (https://www.w3.org/TR/webauthn/#dom-attestationconveyancepreference-enterprise)
	PreferEnterpriseAttestation ConveyancePreference = "enterprise"
)

func ExtractConveyancePreference(p *types.Params) ConveyancePreference {
	switch p.ConveyancePreference {
	case "none":
		return PreferNoAttestation
	case "indirect":
		return PreferIndirectAttestation
	case "direct":
		return PreferDirectAttestation
	case "enterprise":
		return PreferEnterpriseAttestation
	default:
		return PreferNoAttestation
	}
}

type PublicKeyCredentialHints string

const (
	// PublicKeyCredentialHintSecurityKey is a PublicKeyCredentialHint that indicates that the Relying Party believes
	// that users will satisfy this request with a physical security key. For example, an enterprise Relying Party may
	// set this hint if they have issued security keys to their employees and will only accept those authenticators for
	// registration and authentication.
	//
	// For compatibility with older user agents, when this hint is used in PublicKeyCredentialCreationOptions, the
	// authenticatorAttachment SHOULD be set to cross-platform.
	PublicKeyCredentialHintSecurityKey PublicKeyCredentialHints = "security-key"

	// PublicKeyCredentialHintClientDevice is a PublicKeyCredentialHint that indicates that the Relying Party believes
	// that users will satisfy this request with a platform authenticator attached to the client device.
	//
	// For compatibility with older user agents, when this hint is used in PublicKeyCredentialCreationOptions, the
	// authenticatorAttachment SHOULD be set to platform.
	PublicKeyCredentialHintClientDevice PublicKeyCredentialHints = "client-device"

	// PublicKeyCredentialHintHybrid is a PublicKeyCredentialHint that indicates that the Relying Party believes that
	// users will satisfy this request with general-purpose authenticators such as smartphones. For example, a consumer
	// Relying Party may believe that only a small fraction of their customers possesses dedicated security keys. This
	// option also implies that the local platform authenticator should not be promoted in the UI.
	//
	// For compatibility with older user agents, when this hint is used in PublicKeyCredentialCreationOptions, the
	// authenticatorAttachment SHOULD be set to cross-platform.
	PublicKeyCredentialHintHybrid PublicKeyCredentialHints = "hybrid"
)

func ParsePublicKeyCredentialHints(s string) PublicKeyCredentialHints {
	switch s {
	case "security-key":
		return PublicKeyCredentialHintSecurityKey
	case "client-device":
		return PublicKeyCredentialHintClientDevice
	case "hybrid":
		return PublicKeyCredentialHintHybrid
	default:
		return ""
	}
}

type AttestedCredentialData struct {
	AAGUID       []byte `json:"aaguid"`
	CredentialID []byte `json:"credential_id"`

	// The raw credential public key bytes received from the attestation data.
	CredentialPublicKey []byte `json:"public_key"`
}

type ResidentKeyRequirement string

const (
	// ResidentKeyRequirementDiscouraged indicates the Relying Party prefers creating a server-side credential, but will
	// accept a client-side discoverable credential. This is the default.
	ResidentKeyRequirementDiscouraged ResidentKeyRequirement = "discouraged"

	// ResidentKeyRequirementPreferred indicates to the client we would prefer a discoverable credential.
	ResidentKeyRequirementPreferred ResidentKeyRequirement = "preferred"

	// ResidentKeyRequirementRequired indicates the Relying Party requires a client-side discoverable credential, and is
	// prepared to receive an error if a client-side discoverable credential cannot be created.
	ResidentKeyRequirementRequired ResidentKeyRequirement = "required"
)

func ParseResidentKeyRequirement(s string) ResidentKeyRequirement {
	switch s {
	case "discouraged":
		return ResidentKeyRequirementDiscouraged
	case "preferred":
		return ResidentKeyRequirementPreferred
	default:
		return ResidentKeyRequirementRequired
	}
}

type (
	AuthenticationExtensions    map[string]any
	UserVerificationRequirement string
)

const (
	// VerificationRequired User verification is required to create/release a credential
	VerificationRequired UserVerificationRequirement = "required"

	// VerificationPreferred User verification is preferred to create/release a credential
	VerificationPreferred UserVerificationRequirement = "preferred" // This is the default

	// VerificationDiscouraged The authenticator should not verify the user for the credential
	VerificationDiscouraged UserVerificationRequirement = "discouraged"
)
