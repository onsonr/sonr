package models

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
)

const fingerprintDivider = "v=0;fingerprint="
const authDivider = "._auth."

// ** ─── HSRecord Constructers ────────────────────────────────────────────────────────
func NewAuthHSRecord(prefix string, name string, fingerprint string) *HSRecord {
	return &HSRecord{
		Ttl:   5,
		Type:  "TXT",
		Host:  fmt.Sprintf("%s%s%s", prefix, authDivider, name),
		Value: fmt.Sprintf("%s%s", fingerprintDivider, fingerprint),
	}
}

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
		return extractName(hs.Host)
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
func NewNamebaseRequest(record *HSRecord, isDelete bool) ([]byte, error) {
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

	// Marshal JSON
	bytes, err := protojson.Marshal(request)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
