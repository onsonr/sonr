# macaroon
--
    import "gopkg.in/macaroon.v2"

The macaroon package implements macaroons as described in the paper "Macaroons:
Cookies with Contextual Caveats for Decentralized Authorization in the Cloud"
(http://theory.stanford.edu/~ataly/Papers/macaroons.pdf)

See the macaroon bakery packages at http://godoc.org/gopkg.in/macaroon-bakery.v2
for higher level services and operations that use macaroons.

## Usage

```go
const (
	TraceInvalid = TraceOpKind(iota)

	// TraceMakeKey represents the operation of calculating a
	// fixed length root key from the variable length input key.
	TraceMakeKey

	// TraceHash represents a keyed hash operation with one
	// or two values. If there is only one value, it will be in Data1.
	TraceHash

	// TraceBind represents the operation of binding a discharge macaroon
	// to its primary macaroon. Data1 holds the signature of the primary
	// macaroon.
	TraceBind

	// TraceFail represents a verification failure. If present, this will always
	// be the last operation in a trace.
	TraceFail
)
```

#### func  Base64Decode

```go
func Base64Decode(data []byte) ([]byte, error)
```
Base64Decode base64-decodes the given data. It accepts both standard and URL
encodings, both padded and unpadded.

#### type Caveat

```go
type Caveat struct {
	// Id holds the id of the caveat. For first
	// party caveats this holds the condition;
	// for third party caveats this holds the encrypted
	// third party caveat.
	Id []byte

	// VerificationId holds the verification id. If this is
	// non-empty, it's a third party caveat.
	VerificationId []byte

	// For third-party caveats, Location holds the
	// ocation hint. Note that this is not signature checked
	// as part of the caveat, so should only
	// be used as a hint.
	Location string
}
```

Caveat holds a first person or third party caveat.

#### type Macaroon

```go
type Macaroon struct {
}
```

Macaroon holds a macaroon. See Fig. 7 of
http://theory.stanford.edu/~ataly/Papers/macaroons.pdf for a description of the
data contained within. Macaroons are mutable objects - use Clone as appropriate
to avoid unwanted mutation.

#### func  New

```go
func New(rootKey, id []byte, loc string, version Version) (*Macaroon, error)
```
New returns a new macaroon with the given root key, identifier, location and
version.

#### func (*Macaroon) AddFirstPartyCaveat

```go
func (m *Macaroon) AddFirstPartyCaveat(condition []byte) error
```
AddFirstPartyCaveat adds a caveat that will be verified by the target service.

#### func (*Macaroon) AddThirdPartyCaveat

```go
func (m *Macaroon) AddThirdPartyCaveat(rootKey, caveatId []byte, loc string) error
```
AddThirdPartyCaveat adds a third-party caveat to the macaroon, using the given
shared root key, caveat id and location hint. The caveat id should encode the
root key in some way, either by encrypting it with a key known to the third
party or by holding a reference to it stored in the third party's storage.

#### func (*Macaroon) Bind

```go
func (m *Macaroon) Bind(sig []byte)
```
Bind prepares the macaroon for being used to discharge the macaroon with the
given signature sig. This must be used before it is used in the discharges
argument to Verify.

#### func (*Macaroon) Caveats

```go
func (m *Macaroon) Caveats() []Caveat
```
Caveats returns the macaroon's caveats. This method will probably change, and
it's important not to change the returned caveat.

#### func (*Macaroon) Clone

```go
func (m *Macaroon) Clone() *Macaroon
```
Clone returns a copy of the receiving macaroon.

#### func (*Macaroon) Id

```go
func (m *Macaroon) Id() []byte
```
Id returns the id of the macaroon. This can hold arbitrary information.

#### func (*Macaroon) Location

```go
func (m *Macaroon) Location() string
```
Location returns the macaroon's location hint. This is not verified as part of
the macaroon.

#### func (*Macaroon) MarshalBinary

```go
func (m *Macaroon) MarshalBinary() ([]byte, error)
```
MarshalBinary implements encoding.BinaryMarshaler by formatting the macaroon
according to the version specified by MarshalAs.

#### func (*Macaroon) MarshalJSON

```go
func (m *Macaroon) MarshalJSON() ([]byte, error)
```
MarshalJSON implements json.Marshaler by marshaling the macaroon in JSON format.
The serialisation format is determined by the macaroon's version.

#### func (*Macaroon) SetLocation

```go
func (m *Macaroon) SetLocation(loc string)
```
SetLocation sets the location associated with the macaroon. Note that the
location is not included in the macaroon's hash chain, so this does not change
the signature.

#### func (*Macaroon) Signature

```go
func (m *Macaroon) Signature() []byte
```
Signature returns the macaroon's signature.

#### func (*Macaroon) TraceVerify

```go
func (m *Macaroon) TraceVerify(rootKey []byte, discharges []*Macaroon) ([]Trace, error)
```
TraceVerify verifies the signature of the macaroon without checking any of the
first party caveats, and returns a slice of Traces holding the operations used
when verifying the macaroons.

Each element in the returned slice corresponds to the operation for one of the
argument macaroons, with m at index 0, and discharges at 1 onwards.

#### func (*Macaroon) UnmarshalBinary

```go
func (m *Macaroon) UnmarshalBinary(data []byte) error
```
UnmarshalBinary implements encoding.BinaryUnmarshaler. It accepts both V1 and V2
binary encodings.

#### func (*Macaroon) UnmarshalJSON

```go
func (m *Macaroon) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements json.Unmarshaller by unmarshaling the given macaroon in
JSON format. It accepts both V1 and V2 forms encoded forms, and also a
base64-encoded JSON string containing the binary-marshaled macaroon.

After unmarshaling, the macaroon's version will reflect the version that it was
unmarshaled as.

#### func (*Macaroon) Verify

```go
func (m *Macaroon) Verify(rootKey []byte, check func(caveat string) error, discharges []*Macaroon) error
```
Verify verifies that the receiving macaroon is valid. The root key must be the
same that the macaroon was originally minted with. The check function is called
to verify each first-party caveat - it should return an error if the condition
is not met.

The discharge macaroons should be provided in discharges.

Verify returns nil if the verification succeeds.

#### func (*Macaroon) VerifySignature

```go
func (m *Macaroon) VerifySignature(rootKey []byte, discharges []*Macaroon) ([]string, error)
```
VerifySignature verifies the signature of the given macaroon with respect to the
root key, but it does not validate any first-party caveats. Instead it returns
all the applicable first party caveats on success.

The caller is responsible for checking the returned first party caveat
conditions.

#### func (*Macaroon) Version

```go
func (m *Macaroon) Version() Version
```
Version returns the version of the macaroon.

#### type Slice

```go
type Slice []*Macaroon
```

Slice defines a collection of macaroons. By convention, the first macaroon in
the slice is a primary macaroon and the rest are discharges for its third party
caveats.

#### func (Slice) MarshalBinary

```go
func (s Slice) MarshalBinary() ([]byte, error)
```
MarshalBinary implements encoding.BinaryMarshaler.

#### func (*Slice) UnmarshalBinary

```go
func (s *Slice) UnmarshalBinary(data []byte) error
```
UnmarshalBinary implements encoding.BinaryUnmarshaler. It accepts all known
binary encodings for the data - all the embedded macaroons need not be encoded
in the same format.

#### type Trace

```go
type Trace struct {
	RootKey []byte
	Ops     []TraceOp
}
```

Trace holds all toperations involved in verifying a macaroon, and the root key
used as the initial verification key. This can be useful for debugging macaroon
implementations.

#### func (Trace) Results

```go
func (t Trace) Results() [][]byte
```
Results returns the output from all operations in the Trace. The result from
ts.Ops[i] will be in the i'th element of the returned slice. When a trace has
resulted in a failure, the last element will be nil.

#### type TraceOp

```go
type TraceOp struct {
	Kind  TraceOpKind `json:"kind"`
	Data1 []byte      `json:"data1,omitempty"`
	Data2 []byte      `json:"data2,omitempty"`
}
```

TraceOp holds one possible operation when verifying a macaroon.

#### func (TraceOp) Result

```go
func (op TraceOp) Result(input []byte) []byte
```
Result returns the result of computing the given operation with the given input
data. If op is TraceFail, it returns nil.

#### type TraceOpKind

```go
type TraceOpKind int
```

TraceOpKind represents the kind of a macaroon verification operation.

#### func (TraceOpKind) String

```go
func (k TraceOpKind) String() string
```
String returns a string representation of the operation.

#### type Version

```go
type Version uint16
```

Version specifies the version of a macaroon. In version 1, the macaroon id and
all caveats must be UTF-8-compatible strings, and the size of any part of the
macaroon may not exceed approximately 64K. In version 2, all field may be
arbitrary binary blobs.

```go
const (
	// V1 specifies version 1 macaroons.
	V1 Version = 1

	// V2 specifies version 2 macaroons.
	V2 Version = 2

	// LatestVersion holds the latest supported version.
	LatestVersion = V2
)
```

#### func (Version) String

```go
func (v Version) String() string
```
String returns a string representation of the version; for example V1 formats as
"v1".
