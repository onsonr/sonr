package macaroon_test

import (
	"encoding/json"
	"testing"

	qt "github.com/frankban/quicktest"

	"github.com/onsonr/sonr/pkg/crypto/macaroon"
)

func TestMarshalUnmarshalMacaroonV1(t *testing.T) {
	c := qt.New(t)
	testMarshalUnmarshalWithVersion(c, macaroon.V1)
}

func TestMarshalUnmarshalMacaroonV2(t *testing.T) {
	c := qt.New(t)
	testMarshalUnmarshalWithVersion(c, macaroon.V2)
}

func testMarshalUnmarshalWithVersion(c *qt.C, vers macaroon.Version) {
	rootKey := []byte("secret")
	m := MustNew(rootKey, []byte("some id"), "a location", vers)

	// Adding the third party caveat before the first party caveat
	// tests a former bug where the caveat wasn't zeroed
	// before moving to the next caveat.
	err := m.AddThirdPartyCaveat([]byte("shared root key"), []byte("3rd party caveat"), "remote.com")
	c.Assert(err, qt.Equals, nil)

	err = m.AddFirstPartyCaveat([]byte("a caveat"))
	c.Assert(err, qt.Equals, nil)

	b, err := m.MarshalBinary()
	c.Assert(err, qt.Equals, nil)

	var um macaroon.Macaroon
	err = um.UnmarshalBinary(b)
	c.Assert(err, qt.Equals, nil)

	c.Assert(um.Location(), qt.Equals, m.Location())
	c.Assert(string(um.Id()), qt.Equals, string(m.Id()))
	c.Assert(um.Signature(), qt.DeepEquals, m.Signature())
	c.Assert(um.Caveats(), qt.DeepEquals, m.Caveats())
	c.Assert(um.Version(), qt.Equals, vers)
	um.SetVersion(m.Version())
	c.Assert(m, qt.DeepEquals, &um)
}

func TestMarshalBinaryRoundTrip(t *testing.T) {
	c := qt.New(t)
	// This data holds the V2 binary encoding of
	data := []byte(
		"\x02" +
			"\x01\x0ehttp://mybank/" +
			"\x02\x1cwe used our other secret key" +
			"\x00" +
			"\x02\x14account = 3735928559" +
			"\x00" +
			"\x01\x13http://auth.mybank/" +
			"\x02'this was how we remind auth of key/pred" +
			"\x04\x48\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xd3\x6e\xc5\x02\xe0\x58\x86\xd1\xf0\x27\x9f\x05\x5f\xa5\x25\x54\xd1\x6d\x16\xc1\xb1\x40\x74\xbb\xb8\x3f\xf0\xfd\xd7\x9d\xc2\xfe\x09\x8f\x0e\xd4\xa2\xb0\x91\x13\x0e\x6b\x5d\xb4\x6a\x20\xa8\x6b" +
			"\x00" +
			"\x00" +
			"\x06\x20\xd2\x7d\xb2\xfd\x1f\x22\x76\x0e\x4c\x3d\xae\x81\x37\xe2\xd8\xfc\x1d\xf6\xc0\x74\x1c\x18\xae\xd4\xb9\x72\x56\xbf\x78\xd1\xf5\x5c",
	)
	var m macaroon.Macaroon
	err := m.UnmarshalBinary(data)
	c.Assert(err, qt.Equals, nil)
	assertLibMacaroonsMacaroon(c, &m)
	c.Assert(m.Version(), qt.Equals, macaroon.V2)

	data1, err := m.MarshalBinary()
	c.Assert(err, qt.Equals, nil)
	c.Assert(data1, qt.DeepEquals, data)
}

func TestBinaryJSONRoundTripV1(t *testing.T) {
	c := qt.New(t)
	testBinaryJSONRoundTrip(c, macaroon.V1)
}

func TestBinaryJSONRoundTripV2(t *testing.T) {
	c := qt.New(t)
	testBinaryJSONRoundTrip(c, macaroon.V2)
}

func testBinaryJSONRoundTrip(c *qt.C, vers macaroon.Version) {
	m1 := MustNew([]byte("rootkey"), []byte("some id"), "a location", vers)
	err := m1.AddFirstPartyCaveat([]byte("a caveat"))
	c.Assert(err, qt.Equals, nil)
	err = m1.AddThirdPartyCaveat([]byte("shared root key"), []byte("3rd party caveat"), "remote.com")
	c.Assert(err, qt.Equals, nil)

	binData1, err := m1.MarshalBinary()
	c.Assert(err, qt.Equals, nil)

	jsonData1, err := json.Marshal(m1)
	c.Assert(err, qt.Equals, nil)

	var m2 *macaroon.Macaroon
	err = json.Unmarshal(jsonData1, &m2)
	c.Assert(err, qt.Equals, nil)

	binData2, err := m2.MarshalBinary()
	c.Assert(err, qt.Equals, nil)

	c.Assert(binData1, qt.DeepEquals, binData2)
}

func TestMarshalUnmarshalSliceV1(t *testing.T) {
	c := qt.New(t)
	testMarshalUnmarshalSliceWithVersion(c, macaroon.V1)
}

func TestMarshalUnmarshalSliceV2(t *testing.T) {
	c := qt.New(t)
	testMarshalUnmarshalSliceWithVersion(c, macaroon.V2)
}

func testMarshalUnmarshalSliceWithVersion(c *qt.C, vers macaroon.Version) {
	rootKey := []byte("secret")
	m1 := MustNew(rootKey, []byte("some id"), "a location", vers)
	m2 := MustNew(rootKey, []byte("some other id"), "another location", vers)

	err := m1.AddFirstPartyCaveat([]byte("a caveat"))
	c.Assert(err, qt.Equals, nil)
	err = m2.AddFirstPartyCaveat([]byte("another caveat"))
	c.Assert(err, qt.Equals, nil)

	macaroons := macaroon.Slice{m1, m2}

	b, err := macaroons.MarshalBinary()
	c.Assert(err, qt.Equals, nil)

	var unmarshaledMacs macaroon.Slice
	err = unmarshaledMacs.UnmarshalBinary(b)
	c.Assert(err, qt.Equals, nil)

	c.Assert(unmarshaledMacs, qt.HasLen, len(macaroons))
	for i, m := range macaroons {
		um := unmarshaledMacs[i]
		c.Assert(um.Location(), qt.Equals, m.Location())
		c.Assert(string(um.Id()), qt.Equals, string(m.Id()))
		c.Assert(um.Signature(), qt.DeepEquals, m.Signature())
		c.Assert(um.Caveats(), qt.DeepEquals, m.Caveats())
		c.Assert(um.Version(), qt.Equals, vers)
		um.SetVersion(m.Version())
	}
	c.Assert(macaroons, qt.DeepEquals, unmarshaledMacs)

	// Check that appending a caveat to the first does not
	// affect the second.
	for i := 0; i < 10; i++ {
		err = unmarshaledMacs[0].AddFirstPartyCaveat([]byte("caveat"))
		c.Assert(err, qt.Equals, nil)
	}
	unmarshaledMacs[1].SetVersion(macaroons[1].Version())
	c.Assert(unmarshaledMacs[1], qt.DeepEquals, macaroons[1])
	c.Assert(err, qt.Equals, nil)
}

func TestSliceRoundTripV1(t *testing.T) {
	c := qt.New(t)
	testSliceRoundTripWithVersion(c, macaroon.V1)
}

func TestSliceRoundTripV2(t *testing.T) {
	c := qt.New(t)
	testSliceRoundTripWithVersion(c, macaroon.V2)
}

func testSliceRoundTripWithVersion(c *qt.C, vers macaroon.Version) {
	rootKey := []byte("secret")
	m1 := MustNew(rootKey, []byte("some id"), "a location", vers)
	m2 := MustNew(rootKey, []byte("some other id"), "another location", vers)

	err := m1.AddFirstPartyCaveat([]byte("a caveat"))
	c.Assert(err, qt.Equals, nil)
	err = m2.AddFirstPartyCaveat([]byte("another caveat"))
	c.Assert(err, qt.Equals, nil)

	macaroons := macaroon.Slice{m1, m2}

	b, err := macaroons.MarshalBinary()
	c.Assert(err, qt.Equals, nil)

	var unmarshaledMacs macaroon.Slice
	err = unmarshaledMacs.UnmarshalBinary(b)
	c.Assert(err, qt.Equals, nil)

	marshaledMacs, err := unmarshaledMacs.MarshalBinary()
	c.Assert(err, qt.Equals, nil)

	c.Assert(b, qt.DeepEquals, marshaledMacs)
}

var base64DecodeTests = []struct {
	about       string
	input       string
	expect      string
	expectError string
}{{
	about:  "empty string",
	input:  "",
	expect: "",
}, {
	about:  "standard encoding, padded",
	input:  "Z29+IQ==",
	expect: "go~!",
}, {
	about:  "URL encoding, padded",
	input:  "Z29-IQ==",
	expect: "go~!",
}, {
	about:  "standard encoding, not padded",
	input:  "Z29+IQ",
	expect: "go~!",
}, {
	about:  "URL encoding, not padded",
	input:  "Z29-IQ",
	expect: "go~!",
}, {
	about:       "standard encoding, too much padding",
	input:       "Z29+IQ===",
	expectError: `illegal base64 data at input byte 8`,
}}

func TestBase64Decode(t *testing.T) {
	c := qt.New(t)
	for i, test := range base64DecodeTests {
		c.Logf("test %d: %s", i, test.about)
		out, err := macaroon.Base64Decode([]byte(test.input))
		if test.expectError != "" {
			c.Assert(err, qt.ErrorMatches, test.expectError)
		} else {
			c.Assert(err, qt.Equals, nil)
			c.Assert(string(out), qt.Equals, test.expect)
		}
	}
}
