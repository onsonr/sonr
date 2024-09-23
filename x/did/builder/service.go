package builder

import (
	"crypto/rand"
	"strings"

	"github.com/onsonr/sonr/x/did/types"
)

// ChallengeLength - Length of bytes to generate for a challenge.
const ChallengeLength = 32

// CreateChallenge creates a new challenge that should be signed and returned by the authenticator. The spec recommends
// using at least 16 bytes with 100 bits of entropy. We use 32 bytes.
func CreateChallenge() (challenge URLEncodedBase64, err error) {
	challenge = make([]byte, ChallengeLength)

	if _, err = rand.Read(challenge); err != nil {
		return nil, err
	}

	return challenge, nil
}

type CredentialEntity struct {
	// A human-palatable name for the entity. Its function depends on what the PublicKeyCredentialEntity represents:
	//
	// When inherited by PublicKeyCredentialRpEntity it is a human-palatable identifier for the Relying Party,
	// intended only for display. For example, "ACME Corporation", "Wonderful Widgets, Inc." or "ОАО Примертех".
	//
	// When inherited by PublicKeyCredentialUserEntity, it is a human-palatable identifier for a user account. It is
	// intended only for display, i.e., aiding the user in determining the difference between user accounts with similar
	// displayNames. For example, "alexm", "alex.p.mueller@example.com" or "+14255551234".
	Name string `json:"name"`
}

func NewCredentialEntity(name string) CredentialEntity {
	return CredentialEntity{
		Name: name,
	}
}

type CredentialParameter struct {
	Type      CredentialType                `json:"type"`
	Algorithm types.COSEAlgorithmIdentifier `json:"alg"`
}

func NewCredentialParameter(ki *types.KeyInfo) CredentialParameter {
	return CredentialParameter{
		Type:      CredentialTypePublicKeyCredential,
		Algorithm: ki.Algorithm.CoseIdentifier(),
	}
}

func ExtractCredentialParameters(p *types.Params) []CredentialParameter {
	var keys []*types.KeyInfo
	for k, v := range p.AllowedPublicKeys {
		if strings.Contains(k, "webauthn") {
			keys = append(keys, v)
		}
	}
	var cparams []CredentialParameter
	for _, ki := range keys {
		cparams = append(cparams, NewCredentialParameter(ki))
	}
	return cparams
}

type RelyingPartyEntity struct {
	CredentialEntity

	// A unique identifier for the Relying Party entity, which sets the RP ID.
	ID string `json:"id"`
}

func NewRelayingParty(name string, origin string) RelyingPartyEntity {
	return RelyingPartyEntity{
		CredentialEntity: NewCredentialEntity(origin),
		ID:               origin,
	}
}

type UserEntity struct {
	CredentialEntity
	// A human-palatable name for the user account, intended only for display.
	// For example, "Alex P. Müller" or "田中 倫". The Relying Party SHOULD let
	// the user choose this, and SHOULD NOT restrict the choice more than necessary.
	DisplayName string `json:"displayName"`

	// ID is the user handle of the user account entity. To ensure secure operation,
	// authentication and authorization decisions MUST be made on the basis of this id
	// member, not the displayName nor name members. See Section 6.1 of
	// [RFC8266](https://www.w3.org/TR/webauthn/#biblio-rfc8266).
	ID any `json:"id"`
}

func NewUserEntity(name string, subject string, cid string) UserEntity {
	return UserEntity{
		CredentialEntity: NewCredentialEntity(name),
		DisplayName:      subject,
		ID:               cid,
	}
}
