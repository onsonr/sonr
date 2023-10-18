package types

import (
	"encoding/hex"
	"errors"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/go-webauthn/webauthn/protocol"
	"lukechampine.com/blake3"
)

const (
	// Minimum length of base colon separated DID
	MIN_BASE_PART_LENGTH = 3

	// Maximum length of base colon separated DID
	MAX_BASE_PART_LENGTH = 4
)

var (
	ErrBaseNotFound              = errors.New("Unable to determine base did of provided string.")
	ErrFragmentAndQuery          = errors.New("Unable to create new DID. Fragment and Query are mutually exclusive")
	ErrParseInvalid              = errors.New("Unable to parse string into DID, invalid format.")
	DidForbiddenSymbolsRegexp, _ = regexp.Compile(`^[^&\\]+$`)
)

func ToIdentifier(str string) string {
	if str[len(str)-1] == ':' {
		return str[:len(str)-1] + "/"
	}
	return str + "/"
}
func ToNetwork(str string) string {
	return str + ":"
}
func ToFragment(str string) string {
	return "#" + str
}

func ToQuery(str string) string {
	return "?" + str
}

// IndexOf returns the index of the first instance of a value in a slice
func IndexOf(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

// Contains returns true if the string is in the slice
func Contains(vs []string, t string) bool {
	return IndexOf(vs, t) >= 0
}

// Filter returns a new slice containing all strings from the slice that satisfy the predicate
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Complement returns a new slice containing all strings from the slice that do not satisfy the predicate
func Complement(vs []string, ts []string) []string {
	return Filter(vs, func(s string) bool {
		return !Contains(ts, s)
	})
}

// ContainsString returns true if this string contains target string
func ContainsString(s string, t string) bool {
	vs := strings.Split(s, "")
	for _, v := range vs {
		if v == t {
			return true
		}
	}
	return false
}

// ContainsFragment checks if a DID has a fragment in the full string
func ContainsFragment(didUrl string) bool {
	return ContainsString(didUrl, "#")
}

// ContainsModule checks if a core service module is present in the DID
func ContainsModule(didUrl string) bool {
	// Split parts
	parts := strings.Split(didUrl, "/")
	if len(parts) < MIN_BASE_PART_LENGTH || len(parts) > MAX_BASE_PART_LENGTH {
		return false
	}

	// Return default network
	return true
}

// ContainsPath returns true if a DID has a path in the full string
func ContainsPath(didUrl string) bool {
	return ContainsString(didUrl, "/")
}

// ContainsQuery checks if a DID has a query in the full string
func ContainsQuery(didUrl string) bool {
	return ContainsString(didUrl, "?")
}

// IsFragment checks if a DID fragment is valid
func IsFragment(didUrl string) bool {
	if didUrl[0] == '#' {
		return true
	}
	return false
}

// IsPath returns true if a DID has a path in the full string
func IsPath(didUrl string) bool {
	if didUrl[0] == '/' {
		return true
	}
	return false
}

// IsQuery checks if a DID query is valid
func IsQuery(didUrl string) bool {
	if didUrl[0] == '?' {
		return true
	}
	return false
}

// IsValidDid checks if a DID is valid
func IsValidDid(did string) bool {
	if !DidForbiddenSymbolsRegexp.MatchString(did) {
		return false
	}

	return strings.HasPrefix(did, "did:snr:")
}

// ExtractBase extracts the did base (did:snr:<network>:<address>) or (did:snr:address)
func ExtractBase(did string) (bool, string) {
	parts := strings.Split(did, ":")
	finalIdx := len(parts) - 1
	if len(parts) < MIN_BASE_PART_LENGTH || len(parts) > MAX_BASE_PART_LENGTH {
		return false, ""
	}

	base := strings.Join(parts[:finalIdx], "")
	return true, base
}

// ExtractFragment splits a DID URL and pulls the fragment
func ExtractFragment(didUrl string) (bool, string) {
	fragments := strings.Split(didUrl, "#")
	if len(fragments) < 2 {
		return false, ""
	}
	return true, fragments[1]
}

// ExtractIdentifier extracts the identifier from a DID
func ExtractIdentifier(did string) (bool, string) {
	if ok, base := ExtractBase(did); ok {
		parts := strings.Split(base, ":")
		return true, parts[len(parts)-1]
	}
	return false, ""
}

// ExtractPath splits a DID URL and pulls the path
func ExtractPath(didUrl string) (bool, string) {
	if ok, base := ExtractBase(didUrl); ok {
		parts := strings.Split(base, "/")
		if len(parts) < 2 {
			return false, ""
		}
		return true, strings.Join(parts[1:], "/")
	}
	paths := strings.Split(didUrl, "/")
	if len(paths) < 2 {
		return false, ""
	}

	// Get Full Path
	fullPath := strings.Join(paths[1:], "/")
	withoutFinalItemPath := strings.Join(paths[1:len(paths)-1], "/")
	if ContainsFragment(fullPath) {
		return true, withoutFinalItemPath
	}

	if ContainsQuery(fullPath) {
		return true, withoutFinalItemPath
	}
	return true, fullPath
}

// ExtractQuery splits a DID URL and pulls the query
func ExtractQuery(didUrl string) (bool, string) {
	query := strings.Split(didUrl, "?")
	if len(query) < 2 {
		return false, ""
	}
	return true, query[1]
}

// contains returns true if the given string is in the given slice.
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// WARNING: This method is used only for module simulation tests. Do not implement this method across different types in the
// package. ConvertAccAddressToDid converts an AccAddress to a DID
func ConvertAccAddressToDid(address interface{}) string {
	fn := func(address string) string {
		return "did:snr:" + address
	}
	// check if string or sdk.AccAddress
	switch address.(type) {
	case string:
		return fn(address.(string))
	case sdk.AccAddress:
		return fn(address.(sdk.AccAddress).String())
	default:
		return ""
	}
}

// WARNING: This method is used only for module simulation tests. Do not implement this method across different types in the
// package. ConvertDidToAccAddress converts a DID to an AccAddress
func ConvertDidToAccAddress(did string) (sdk.AccAddress, error) {
	if ok, base := ExtractBase(did); ok {
		parts := strings.Split(base, ":")
		return sdk.AccAddressFromBech32(parts[len(parts)-1])
	}
	return nil, errors.New("invalid did")
}

// blake3HashHex hashes the input string using the Blake3 algorithm and returns the
// hash as a hex-encoded string.
func blake3HashHex(input string) string {
	outputSize := 32 // Output size in bytes (32 bytes = 256 bits)
	key := []byte{}  // Empty key for keyless hashing

	hasher := blake3.New(outputSize, key)
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

// makeCredentialFromCreationData creates a new Credential from a ParsedCredentialCreationData and contains all needed information about a WebAuthn credential for storage.
// This is then used to create a VerificationMethod for the DID Document.
func makeCredentialFromCreationData(c *protocol.ParsedCredentialCreationData) *Credential {
	transport := make([]string, 0)
	for _, t := range c.Response.Transports {
		transport = append(transport, string(t))
	}

	newCredential := &Credential{
		Id:              c.Response.AttestationObject.AuthData.AttData.CredentialID,
		PublicKey:       c.Response.AttestationObject.AuthData.AttData.CredentialPublicKey,
		AttestationType: c.Response.AttestationObject.Format,
		Transport:       transport,

		Authenticator: &Authenticator{
			Aaguid:     c.Response.AttestationObject.AuthData.AttData.AAGUID,
			SignCount:  c.Response.AttestationObject.AuthData.Counter,
			Attachment: string(c.AuthenticatorAttachment),
		},
	}
	return newCredential
}

// makeCredentialFromAssertionData creates a new Credential from a ParsedCredentialAssertionData and contains all needed information about a WebAuthn credential for storage.
// This is then used to create a VerificationMethod for the DID Document.
func makeCredentialFromAssertionData(c *protocol.ParsedCredentialAssertionData) *Credential {
	return &Credential{
		Id:        c.RawID,
		PublicKey: c.Response.AuthenticatorData.AttData.CredentialPublicKey,
		Authenticator: &Authenticator{
			Aaguid:    c.Response.AuthenticatorData.AttData.AAGUID,
			SignCount: c.Response.AuthenticatorData.Counter,
		},
	}
}
