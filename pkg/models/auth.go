package models

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
)

// @ Constant Dividers
const fingerprintDivider = "v=0;fingerprint="
const authDivider = "._auth."

// @ Constant Name Filters
var blockedNames = []string{"elon", "vitalik", "prad", "rishi", "brax", "vt", "isa"}
var restrictedNames = []string{"root", "admin", "mail", "auth", "crypto", "id", "app", "beta", "alpha", "code", "ios", "android", "test", "node", "sonr"}

// ** ─── HSRecord Constructers ────────────────────────────────────────────────────────
func NewBlankHSRecord() *HSRecord {
	return &HSRecord{
		Ttl:   -1,
		Type:  "-",
		Host:  "-",
		Value: "-",
	}
}

// ** ─── HSRecord Methods ────────────────────────────────────────────────────────
// ^ Checks if Record is for Auth ^
func (hs *HSRecord) IsAuth() bool {
	return strings.Contains(hs.Host, authDivider)
}

// ^ Checks if Record is Blank ^
func (hs *HSRecord) IsBlank() bool {
	return hs.Ttl == -1 && hs.Host == "-" && hs.Type == "-" && hs.Value == "-"
}

// ^ Checks if Record Matches Name ^ //
func (hs *HSRecord) IsName(name string) bool {
	return hs.Name() == name
}

// ^ Checks if Record DOES NOT Match Name ^ //
func (hs *HSRecord) IsNotName(name string) bool {
	return hs.Name() != name
}

// ^ Returns Record Fingerprint if Transfer ^
func (hs *HSRecord) Fingerprint() []byte {
	if hs.IsAuth() {
		return extractFingerprint(hs.Value)
	}
	return nil
}

// ^ Returns Record Name if Transfer ^
func (hs *HSRecord) Name() string {
	if hs.IsAuth() {
		return strings.ToLower(extractName(hs.Host))
	}
	return ""
}

// ^ Returns Record Prefix if Transfer ^
func (hs *HSRecord) Prefix() string {
	if hs.IsAuth() {
		return extractPrefix(hs.Host)
	}
	return ""
}

// ^ Converts HSRecord to Map ^ //
func (hs *HSRecord) ToMap() map[string]interface{} {
	data := make(map[string]interface{})
	data["ttl"] = hs.Ttl
	data["type"] = hs.Type
	data["host"] = hs.Host
	data["value"] = hs.Value
	return data
}

// ** ─── HSRecord Helpers ────────────────────────────────────────────────────────
// @ Helper: Verifies Unrestricted
func checkUnrestricted(val string) bool {
	for _, v := range restrictedNames {
		if v == val {
			return false
		}
	}
	return true
}

// @ Helper: Verifies Unblocked
func checkUnblocked(val string) bool {
	for _, v := range blockedNames {
		if v == val {
			return false
		}
	}
	return true
}

// @ Helper: Extracts Fingerprint from `Host`
func extractFingerprint(value string) []byte {
	return []byte(substring(value, len(fingerprintDivider), len(value)))
}

// @ Helper: Extracts Prefix from `Host`
func extractPrefix(host string) string {
	idx := strings.Index(host, authDivider)
	return substring(host, 0, idx)
}

// @ Helper: Extracts SName from `Host`
func extractName(host string) string {
	idx := strings.Index(host, authDivider)
	return substring(host, idx+len(authDivider), len(host))
}

// @ Helper: Gets Substrings
func substring(input string, start int, length int) string {
	asRunes := []rune(input)
	if start >= len(asRunes) {
		return ""
	}
	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}
	return string(asRunes[start : start+length])
}

// ** ─── NamebaseResponse Management ────────────────────────────────────────────────────────
// ^ Method Returns New Namebase Response from JSON buffer ^ //
func NewNamebaseResponse(data []byte) (*NamebaseResponse, error) {
	response := &NamebaseResponse{}
	err := protojson.Unmarshal(data, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// ^ Method Returns New Namebase Request as JSON Bytes
func NewNamebaseRequest(record *HSRecord, isDelete bool) *NamebaseRequest {
	// Verify Non-Empty Name
	if record.Name() != "" {
		// Validate Name
		if checkUnblocked(record.Name()) && checkUnrestricted(record.Name()) {
			// Create Request
			request := &NamebaseRequest{
				Records:       []*HSRecord{},
				DeleteRecords: []*HSRecord{},
			}

			// Check for Delete or Create
			if isDelete {
				request.DeleteRecords = append(request.DeleteRecords, record)
			} else {
				request.Records = append(request.Records, record)
			}
			return request
		}
	}
	return nil
}

func (req *NamebaseRequest) JSON() ([]byte, error) {
	// Marshal JSON
	bytes, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// ** ─── UsernameResponse MANAGEMENT ────────────────────────────────────────────────────────
// Converts Username Response to User_Crypto
func (ur *AuthenticationResponse) ToUserCrypto() *User_Crypto {
	if ur.GetIsValid() {
		return &User_Crypto{
			Prefix:   ur.GetPrefix(),
			Mnemonic: ur.GetMnemonic(),
			SName:    ur.GetSName(),
		}
	}
	return nil
}

// Converts Authentication Request to HSRecord
func (ur *AuthenticationRequest) ToHSRecord(prefix string, fingerprint string) *HSRecord {
	return &HSRecord{
		Ttl:   5,
		Type:  "TXT",
		Host:  fmt.Sprintf("%s%s%s", prefix, authDivider, ur.SName),
		Value: fmt.Sprintf("%s%s", fingerprintDivider, fingerprint),
	}
}
