package types


import (
	"encoding"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/sonr-io/sonr/x/identity/types/ssi"

	odid "github.com/ockam-network/did"
)

var _ fmt.Stringer = DID{}
var _ encoding.TextMarshaler = DID{}

// DIDContextV1 contains the JSON-LD context for a DID Document
const DIDContextV1 = "https://www.w3.org/ns/did/v1"

// DIDContextV1URI returns DIDContextV1 as a URI
func DIDContextV1URI() ssi.URI {
	return ssi.MustParseURI(DIDContextV1)
}

// DID represents a Decentralized Identifier as specified by the DID Core specification (https://www.w3.org/TR/did-core/#identifier).
type DID struct {
	odid.DID
}

func New(addr string) (*DID, error) {
	return ParseDID(fmt.Sprintf("did:snr:%s", addr))
}

// Empty checks whether the DID is set or not
func (d DID) Empty() bool {
	return d.Method == ""
}

// String returns the DID as formatted string.
func (d DID) String() string {
	return d.DID.String()
}

// MarshalText implements encoding.TextMarshaler
func (d DID) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// Equals checks whether the DID is exactly equal to another DID
// The check is case sensitive.
func (d DID) Equals(other DID) bool {
	return d.String() == other.String()
}

// UnmarshalJSON unmarshals a DID encoded as JSON string, e.g.:
// "did:snr:c0dc584345da8a0e1e7a584aa4a36c30ebdb79d907aff96fe0e90ee972f58a17"
func (d *DID) UnmarshalJSON(bytes []byte) error {
	var didString string
	err := json.Unmarshal(bytes, &didString)
	if err != nil {
		return ErrInvalidDID.wrap(err)
	}
	tmp, err := odid.Parse(didString)
	if err != nil {
		return ErrInvalidDID.wrap(err)
	}
	d.DID = *tmp
	return nil
}

// Address returns the Sonr Account address from the given DID
func (d *DID) Address() string {
	return d.IDStrings[1] + d.IDStrings[2]
}

// MarshalJSON marshals the DID to a JSON string
func (d DID) MarshalJSON() ([]byte, error) {
	didAsString := d.DID.String()
	return json.Marshal(didAsString)
}

// URI converts the DID to an URI.
// URIs are used in Verifiable Credentials
func (d DID) URI() ssi.URI {
	return ssi.URI{
		URL: url.URL{
			Scheme:   "did",
			Opaque:   fmt.Sprintf("%s:%s", d.Method, d.ID),
			Fragment: d.Fragment,
		},
	}
}

// ParseDID parses a raw DID.
// If the input contains a path, query or fragment, use the ParseDIDURL instead.
// If it can't be parsed, an error is returned.
func ParseDID(input string) (*DID, error) {
	ockDid, err := odid.Parse(input)
	if err != nil {
		return nil, ErrInvalidDID.wrap(err)
	}

	return &DID{DID: *ockDid}, nil
}

// GetCreatorDid returns the creator did
func CreateDIDFromAccount(creator string) string {
	rawCreator := creator

	// Trim snr account prefix
	if strings.HasPrefix(rawCreator, "snr") {
		rawCreator = strings.TrimLeft(rawCreator, "snr")
	}

	// Trim cosmos account prefix
	if strings.HasPrefix(rawCreator, "cosmos") {
		rawCreator = strings.TrimLeft(rawCreator, "cosmos")
	}
	return fmt.Sprintf("did:snr:%s", rawCreator)
}

// must accepts a function like Parse and returns the value without error or panics otherwise.
func must(fn func(string) (*DID, error), input string) DID {
	v, err := fn(input)
	if err != nil {
		panic(err)
	}
	return *v
}

// MustParseDID is like ParseDID but panics if the string cannot be parsed.
// It simplifies safe initialization of global variables holding compiled UUIDs.
func MustParseDID(input string) DID {
	return must(ParseDID, input)
}

// ErrInvalidDID is returned when a parser function is supplied with a string that can't be parsed as DID.
var ErrInvalidDID = ParserError{msg: "invalid DID"}

// ParserError is used when returning DID-parsing related errors.
type ParserError struct {
	msg string
	err error
}

func (w ParserError) wrap(err error) error {
	return ParserError{msg: fmt.Sprintf("%s: %s", w.msg, err.Error()), err: err}
}

// Is checks whether the given error is a ParserError
func (w ParserError) Is(other error) bool {
	_, ok := other.(ParserError)
	return ok
}

// Unwrap returns the underlying error.
func (w ParserError) Unwrap() error {
	return w.err
}

// Error returns the message of the error.
func (w ParserError) Error() string {
	return w.msg
}

