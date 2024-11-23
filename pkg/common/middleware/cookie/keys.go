package cookie

// Key is a type alias for string.
type Key string

const (
	// SessionID is the key for the session ID cookie.
	SessionID Key = "session.id"

	// SessionChallenge is the key for the session challenge cookie.
	SessionChallenge Key = "session.challenge"

	// SessionRole is the key for the session role cookie.
	SessionRole Key = "session.role"

	// SonrAddress is the key for the Sonr address cookie.
	SonrAddress Key = "sonr.address"

	// SonrKeyshare is the key for the Sonr address cookie.
	SonrKeyshare Key = "sonr.keyshare"

	// SonrDID is the key for the Sonr DID cookie.
	SonrDID Key = "sonr.did"

	// UserHandle is the key for the User Handle cookie.
	UserHandle Key = "user.handle"

	// VaultCID is the key for the Vault CID cookie.
	VaultCID Key = "vault.cid"

	// VaultSchema is the key for the Vault schema cookie.
	VaultSchema Key = "vault.schema"
)

// String returns the string representation of the CookieKey.
func (c Key) String() string {
	return string(c)
}
