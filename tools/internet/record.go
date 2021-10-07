package internet

import (
	"fmt"
	"strings"
)

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
	println("--- DNS Record ---")
	println(fmt.Sprintf("\t Host: %s", r.Host))
	println(fmt.Sprintf("\t Type: %s", r.Type))
	println(fmt.Sprintf("\t Value: %s", r.Value))
	println(fmt.Sprintf("\t TTL: %d", r.TTL))

	// Check for TXT Record
	if r.Type == "TXT" {
		println("--- Auth Crypto ---")
		println(fmt.Sprintf("\t Fingerprint: %s", r.Fingerprint()))
		println(fmt.Sprintf("\t Name: %s", r.Name()))
		println(fmt.Sprintf("\t Prefix: %s", r.Prefix()))
	}
}
