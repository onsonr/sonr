package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/pkg/common"

	"github.com/pkg/errors"
)

const (
	// 53 is the default DNS port
	DIAL_TIMEOUT        = time.Millisecond * time.Duration(10000)
	HDNS_RESOLVER_ONE   = "103.196.38.38:53" // Hardcoded Public DNS Resolver for HDNS #1
	HDNS_RESOLVER_TWO   = "103.196.38.39:53" // Hardcoded Public DNS Resolver for HDNS #2
	HDNS_RESOLVER_THREE = "103.196.38.40:53" // Hardcoded Public DNS Resolver for HDNS #3
	DOMAIN              = "snr"
	API_ADDRESS         = "https://www.namebase.io/api/v0/"
	DNS_ENDPOINT        = "dns/domains/snr/nameserver"
	FINGERPRINT_DIVIDER = "v=0;fingerprint="
	AUTH_DIVIDER        = "._auth."
	NB_DNS_API_URL      = API_ADDRESS + DNS_ENDPOINT
)

// Error Definitions
var (
	ErrRecordCountMismatch = errors.New("Number of TXT records dont match Number of TTLs")
	ErrMultipleRecords     = errors.New("Multiple TXT records found for Query")
	ErrEmptyTXT            = errors.New("Empty TXT Record")
	ErrHDNSResolve         = errors.New("Failed to dial all three public HDNS resolvers.")
)

// DomainMap returns map with host as key and recordValue as value.
type DomainMap map[string]string

type RecordCategory int

const (
	Category_NONE RecordCategory = iota
	Category_AUTH
	Category_NAME
)

// IsAuth returns true if the Record is an Auth Record
func (c RecordCategory) IsAuth() bool {
	return c == Category_AUTH
}

// IsName returns true if the Record is a Name Record
func (c RecordCategory) IsName() bool {
	return c == Category_NAME
}

// IsNone returns true if the Record is a None Record
func (c RecordCategory) IsNone() bool {
	return c == Category_NONE
}

// String returns the string representation of the Record
func (c RecordCategory) String() string {
	switch c {
	case Category_AUTH:
		return "AUTH"
	case Category_NAME:
		return "NAME"
	default:
		return "NONE"
	}
}

// Record is a DNS Record
type Record struct {
	// Type is the type of record to be deleted
	Type string `json:"type"`

	// Host is the hostname of the record to be deleted
	Host string `json:"host"`

	// Value is the value of the record
	Value string `json:"value"`

	// TTL is the time to live of the record
	TTL int `json:"ttl"`

	// Category is the determined Sonr Category of Record
	Category RecordCategory
}

// FindRecordCategory determines the Sonr Category of Record
func FindRecordCategory(host, value string) RecordCategory {
	// Check for Auth Record
	if checkRecordForAuth(host, value) {
		return Category_AUTH
	}

	// Check for Name Record
	if checkRecordForSNID(host, value) {
		return Category_NAME
	}

	// Return None
	return Category_NONE
}

// NewNBAuthRecord creates a new Record with Auth Signing
func NewNBRecord(host string, value string) Record {
	// Return Record
	return Record{
		TTL:      5,
		Type:     "TXT",
		Host:     host,
		Value:    value,
		Category: FindRecordCategory(host, value),
	}
}

// NewNBAuthRecord creates a new Record with Auth Signing
func NewNBAuthRecord(prefix string, name string, fingerprint string) Record {
	// Return Record
	return Record{
		TTL:      5,
		Type:     "TXT",
		Host:     fmt.Sprintf("%s%s%s", prefix, AUTH_DIVIDER, name),
		Value:    fmt.Sprintf("%s%s", FINGERPRINT_DIVIDER, fingerprint),
		Category: Category_AUTH,
	}
}

// NewNBAuthRecord creates a new Record with Auth Signing
func NewNBNameRecord(publicKey string, name string) Record {
	// Return Record
	return Record{
		TTL:      5,
		Type:     "TXT",
		Host:     name,
		Value:    publicKey,
		Category: Category_NAME,
	}
}

// IsAuth returns true if the Record is an Auth Record
func (r Record) IsAuth() bool {
	return r.Category.IsAuth()
}

// IsName returns true if the Record is a Name Record
func (r Record) IsName() bool {
	return r.Category.IsName()
}

// ComparePeerID compares the PeerID of the Record with the given PeerID
func (r Record) ComparePeerID(id peer.ID) bool {
	// Check peer record
	pid, err := r.PeerID()
	if err != nil {
		logger.Error("Failed to extract PeerID from PublicKey", err)
		return false
	}
	return pid == id
}

// Fingerprint is the fingerprint for the Auth Record
func (r Record) Fingerprint() string {
	// Check for SNR Record
	if err := checkSnrRecord(r); err != nil {
		return ""
	}

	// Split Value
	sections := strings.Split(r.Value, ";")
	last := sections[len(sections)-1]
	vals := strings.Split(last, "=")
	return vals[0]
}

// Prefix is the prefix for the Auth/Name Record
func (r Record) Name() string {
	// Check for SNR Record
	if err := checkSnrRecord(r); err != nil {
		return ""
	}

	// Split Value for Auth
	if r.IsAuth() {
		// Return Prefix
		vals := strings.Split(r.Host, ".")
		return vals[len(vals)-1]
	}

	// Name record is just host
	return strings.ToLower(r.Host)
}

// Peer returns Peer from Record
func (r Record) Peer() (*common.Peer, error) {
	if r.IsName() {
		id, err := r.PeerID()
		if err != nil {
			return nil, err
		}

		pubBuf, err := r.PubKeyBuffer()
		if err != nil {
			return nil, err
		}

		return &common.Peer{
			PeerID:    id.String(),
			PublicKey: pubBuf,
			SName:     r.Name(),
			Profile: &common.Profile{
				FirstName: "Anonymous",
				LastName:  "Peer",
				SName:     r.Name(),
			},
		}, nil
	}
	return nil, errors.New("Not a Sonr Name Record")
}

// PeerID is the PeerID for the Name Record
func (r Record) PeerID() (peer.ID, error) {
	// Check for SNR Record
	if err := checkSnrRecord(r); err != nil {
		return peer.ID(""), err
	}

	// Verify Public Key
	pub, err := r.PubKey()
	if err != nil {
		return peer.ID(""), err
	}
	return peer.IDFromPublicKey(pub)
}

// Prefix is the prefix for the Auth Record
func (r Record) Prefix() string {
	// Check for SNR Record
	if err := checkSnrRecord(r); err != nil {
		return ""
	}

	// Return Prefix
	vals := strings.Split(r.Host, ".")
	return vals[0]
}

// PubKey is the Public Key for the Name Record
func (r Record) PubKey() (crypto.PubKey, error) {
	// Check for SNR Record
	if err := checkSnrRecord(r); err != nil {
		return nil, err
	}

	// Verify Public Key
	if err := verifyStringPubKey(r.Value); err != nil {
		return nil, err
	}

	// Get Buffer from Value
	buf, err := base64StrToBuffer(r.Value)
	if err != nil {
		return nil, err
	}
	return bufferToPubKey(buf)
}

// PubKeyBuffer is the Public Key for the Name Record as a Buffer
func (r Record) PubKeyBuffer() ([]byte, error) {
	// Check for SNR Record
	if err := checkSnrRecord(r); err != nil {
		return nil, err
	}

	// Verify Public Key
	if err := verifyStringPubKey(r.Value); err != nil {
		return nil, err
	}

	// Get Buffer from Value
	buf, err := base64StrToBuffer(r.Value)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// ToDeleteRecord converts a Record to a DeleteRecord
func (r Record) ToDeleteRecord() DeleteRecord {
	return DeleteRecord{
		Type: r.Type,
		Host: r.Host,
	}
}

// Print prints the Record to stdout
func (r Record) Print() {
	println("[NB.Record]")
	println("--- DNS ---")
	println(fmt.Sprintf("\t Host: %s", r.Host))
	println(fmt.Sprintf("\t Type: %s", r.Type))
	println(fmt.Sprintf("\t Value: %s", r.Value))
	println(fmt.Sprintf("\t TTL: %d", r.TTL))
	println("\n")
	// Check for TXT Record
	if r.Type == "TXT" {
		// Print by Category
		switch r.Category {
		case Category_AUTH:
			println("--- [AUTH] Properties ---")
			println(fmt.Sprintf("\t Name: %s", r.Name()))
			println(fmt.Sprintf("\t Prefix: %s", r.Prefix()))
			println(fmt.Sprintf("\t Fingerprint: %s", r.Fingerprint()))
		case Category_NAME:
			println("--- [NAME] Properties ---")
			peerid, _ := r.PeerID()
			println(fmt.Sprintf("\t Name: %s", r.Name()))
			println(fmt.Sprintf("\t PeerID: %s", peerid))
			println(fmt.Sprintf("\t Public Key: %s", r.Value))
		}
	}
}

func checkRecordForAuth(host, value string) bool {
	// Check for Auth Divider in Host for Record
	if !strings.Contains(host, AUTH_DIVIDER) {
		return false
	}

	// Check for Fingerprint Divider in Value for Record
	if !strings.Contains(value, FINGERPRINT_DIVIDER) {
		return false
	}
	return true
}

func checkRecordForSNID(host, value string) bool {
	// Check for Auth Divider in Host for Record
	if checkRecordForAuth(host, AUTH_DIVIDER) {
		return false
	}

	// Check for Fingerprint Divider in Value for Record
	if strings.Contains(value, FINGERPRINT_DIVIDER) {
		return false
	}

	// Verify Public Key
	if err := verifyStringPubKey(value); err != nil {
		logger.Error("Failed to verify public key: %s", err)
		return false
	}
	return true
}

func base64StrToBuffer(str string) ([]byte, error) {
	// Decode the key
	buf, err := crypto.ConfigDecodeKey(str)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to decode PubKey from String")
	}
	return buf, nil
}

func bufferToPubKey(buf []byte) (crypto.PubKey, error) {
	// Decode the key
	pub, err := crypto.UnmarshalPublicKey(buf)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to decode PubKey from Buffer")
	}
	return pub, nil
}

func verifyStringPubKey(str string) error {
	buf, err := base64StrToBuffer(str)
	if err != nil {
		return err
	}

	// Get Public Key from Buffer
	_, err = crypto.UnmarshalPublicKey(buf)
	return errors.WithMessage(err, "Failed to unmarshal PubKey from Bytes")
}

// Records is a slice of Record
type Records []Record

// Len returns the length of the Record Slice
func (rs Records) Len() int {
	return len(rs)
}

// GetAuthRecord returns the Auth Record from the Record Slice
func (rs Records) GetAuthRecord() (Record, error) {
	// Check for Auth Record
	for _, r := range rs {
		if r.Category == Category_AUTH {
			return r, nil
		}
	}

	// No Auth Record Found
	return Record{}, errors.New("no auth record found")
}

// GetNameRecord returns the Name Record from the Record Slice
func (rs Records) GetNameRecord() (Record, error) {
	// Check for Name Record
	for _, r := range rs {
		if r.Category == Category_NAME {
			return r, nil
		}
	}

	// No Name Record Found
	return Record{}, errors.New("no name record found")
}

// checkSnrRecord checks if the record is a SNR Record
func checkSnrRecord(r Record) error {
	// Check for TXT Record
	if r.Type != "TXT" {
		err := errors.New("not a TXT record")
		logger.Error("Failed to get Value from Record", golog.Fields{"error": err})
		return err
	}

	// Check Category
	c := FindRecordCategory(r.Host, r.Value)
	if c != Category_NAME && c != Category_AUTH {
		err := errors.New("Record does not have category")
		logger.Error("Failed to get Value from Record", golog.Fields{"error": err})
		return err
	}
	return nil
}
