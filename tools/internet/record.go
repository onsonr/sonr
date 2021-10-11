package internet

import (
	"fmt"
	"strings"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
)

type RecordCategory int

const (
	Category_NONE RecordCategory = iota
	Category_AUTH
	Category_NAME
)

func (c RecordCategory) IsAuth() bool {
	return c == Category_AUTH
}

func (c RecordCategory) IsName() bool {
	return c == Category_NAME
}

func (c RecordCategory) IsNone() bool {
	return c == Category_NONE
}

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

// NamebaseRequest for either Adding or Removing DNS Records
type NamebaseRequest struct {
	// Records to be added to DNS Table
	Records []Record `json:"records"`

	// DeleteRecords are to be deleted from DNS Table
	DeleteRecords []DeleteRecord `json:"deleteRecords"`
}

// NewNamebaseRequest creates a new NamebaseRequest for adding records
func NewNBAddRequest(records ...Record) NamebaseRequest {
	return NamebaseRequest{
		Records:       records,
		DeleteRecords: make([]DeleteRecord, 0),
	}
}

// NewNBDeleteRequest creates a new NamebaseRequest for deleting records
func NewNBDeleteRequest(records ...DeleteRecord) NamebaseRequest {
	return NamebaseRequest{
		Records:       make([]Record, 0),
		DeleteRecords: records,
	}
}

// NamebaseResponse is JSON Response for NamebaseRequest
type NamebaseResponse struct {
	// Success is true if the request was successful
	Success bool `json:"success"`

	// Records is the list of records from GET request
	Records []Record `json:"records"`
}

// Print prints the NamebaseResponse
func (nr *NamebaseResponse) Print() {
	// Loop through all records
	for _, record := range nr.Records {
		record.ToPrint()
	}
}

// DeleteRecord is for Removing Records in Request
type DeleteRecord struct {
	// Type is the type of record to be deleted
	Type string `json:"type"`

	// Host is the hostname of the record to be deleted
	Host string `json:"host"`
}

// NewNamebaseDeleteRecord creates a new DeleteRecord
func NewNBDeleteRecord(host string) DeleteRecord {
	return DeleteRecord{
		Type: "TXT",
		Host: host,
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
		TTL:   5,
		Type:  "TXT",
		Host:  host,
		Value: value,
	}
}

// NewNBAuthRecord creates a new Record with Auth Signing
func NewNBAuthRecord(prefix string, name string, fingerprint string) Record {
	// Return Record
	return Record{
		TTL:   5,
		Type:  "TXT",
		Host:  fmt.Sprintf("%s%s%s", prefix, AUTH_DIVIDER, name),
		Value: fmt.Sprintf("%s%s", FINGERPRINT_DIVIDER, fingerprint),
	}
}

// NewNBAuthRecord creates a new Record with Auth Signing
func NewNBNameRecord(publicKey string, name string) Record {
	// Return Record
	return Record{
		TTL:   5,
		Type:  "TXT",
		Host:  name,
		Value: publicKey,
	}
}

// Fingerprint is the fingerprint for the Auth Record
func (r Record) Fingerprint() string {
	// Check for TXT Record
	if r.Type != "TXT" {
		return ""
	}

	// Split Value
	sections := strings.Split(r.Value, ";")
	last := sections[len(sections)-1]
	vals := strings.Split(last, "=")
	return vals[0]
}

// Prefix is the prefix for the Auth Record
func (r Record) Name() string {
	// Check for TXT Record
	if r.Type != "TXT" {
		return ""
	}

	// Return Prefix
	vals := strings.Split(r.Host, ".")
	return vals[len(vals)-1]
}

func (r Record) PeerID() (peer.ID, error) {
	// Check for TXT Record
	if r.Type != "TXT" {
		return peer.ID(""), errors.New("not a TXT record")
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
	// Check for TXT Record
	if r.Type != "TXT" {
		return ""
	}

	// Return Prefix
	vals := strings.Split(r.Host, ".")
	return vals[0]
}

func (r Record) PubKey() (crypto.PubKey, error) {
	// Check for TXT Record
	if r.Type != "TXT" {
		return nil, errors.New("not a TXT record")
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

func (r Record) PubKeyBuffer() ([]byte, error) {
	// Check for TXT Record
	if r.Type != "TXT" {
		return nil, errors.New("not a TXT record")
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

// ToPrint prints the Record to stdout
func (r Record) ToPrint() {
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
			println("--- NAME ---")
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
