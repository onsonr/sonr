package util

import (
	"regexp"
	"strings"


)

const (
	PublicKeyJwk       = "PublicKeyJwk"
	PublicKeyMultibase = "PublicKeyMultibase"
)

var VerificationMethodType = map[string]string{
	"JsonWebKey2020":             PublicKeyJwk,
	"Ed25519VerificationKey2020": PublicKeyMultibase,
}

var ValidNetworkPrefixes = []string{
	"mainnet",
	"testnet",
	"devnet",
}

var DidForbiddenSymbolsRegexp, _ = regexp.Compile(`^[^&\\]+$`)

// GetVerificationMethodType returns the verification method type
func GetVerificationMethodType(vmType string) string {
	return VerificationMethodType[vmType]
}

// IsDidFragment checks if a DID fragment is valid
func IsDidFragment(prefix string, didUrl string) bool {
	if !strings.Contains(didUrl, "#") {
		return false
	}

	if didUrl[0] == '#' {
		return true
	}

	did, _ := SplitDidUrlIntoDidAndFragment(didUrl)
	return IsValidDid(prefix, did)
}

// IsFullDidFragment checks if a DID fragment is for full string
func IsFullDidFragment(prefix string, didUrl string) bool {
	if !strings.Contains(didUrl, "#") {
		return false
	}

	did, _ := SplitDidUrlIntoDidAndFragment(didUrl)
	return IsValidDid(prefix, did)
}

// IsNotValidDIDArray checks if a DID array is invalid
func IsNotValidDIDArray(prefix string, array []string) (bool, int) {
	for i, did := range array {
		if !IsValidDid(prefix, did) {
			return true, i
		}
	}

	return false, 0
}

// IsNotValidDIDArrayFragment checks if a DID array is invalid
func IsNotValidDIDArrayFragment(prefix string, array []string) (bool, int) {
	for i, did := range array {
		if !IsDidFragment(prefix, did) {
			return true, i
		}
	}

	return false, 0
}

// IsValidDid checks if a DID is valid
func IsValidDid(prefix string, did string) bool {
	if len(did) == 0 {
		return false
	}

	if !DidForbiddenSymbolsRegexp.MatchString(did) {
		return false
	}

	// FIXME: Empty namespace must be allowed even if namespace is set in state
	// https://github.com/cheqd/cheqd-node/blob/main/architecture/adr-list/adr-002-cheqd-did-method.md#method-specific-identifier
	return strings.HasPrefix(did, prefix)
}

func IsValidNetworkPrefix(prefix string) bool {
	return Contains(ValidNetworkPrefixes, prefix)
}

// SplitDidUrlIntoDidAndFragment splits a DID URL into DID and fragment
func SplitDidUrlIntoDidAndFragment(didUrl string) (string, string) {
	fragments := strings.Split(didUrl, "#")
	return fragments[0], fragments[1]
}
