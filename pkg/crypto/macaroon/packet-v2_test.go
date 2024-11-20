package macaroon

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/google/go-cmp/cmp"
)

var packetEquals = qt.CmpEquals(cmp.AllowUnexported(packetV1{}, packetV2{}))

var parsePacketV2Tests = []struct {
	about        string
	data         string
	expectPacket packetV2
	expectData   string
	expectError  string
}{{
	about: "EOS packet",
	data:  "\x00",
	expectPacket: packetV2{
		fieldType: fieldEOS,
	},
}, {
	about: "simple field",
	data:  "\x02\x03xyz",
	expectPacket: packetV2{
		fieldType: 2,
		data:      []byte("xyz"),
	},
}, {
	about:       "empty buffer",
	data:        "",
	expectError: "varint value extends past end of buffer",
}, {
	about:       "varint out of range",
	data:        "\xff\xff\xff\xff\xff\xff\x7f",
	expectError: "varint value out of range",
}, {
	about:       "varint way out of range",
	data:        "\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\x7f",
	expectError: "varint value out of range",
}, {
	about:       "unterminated varint",
	data:        "\x80",
	expectError: "varint value extends past end of buffer",
}, {
	about:       "field data too long",
	data:        "\x01\x02a",
	expectError: "field data extends past end of buffer",
}, {
	about:       "bad data length varint",
	data:        "\x01\xff",
	expectError: "varint value extends past end of buffer",
}}

func TestParsePacketV2(t *testing.T) {
	c := qt.New(t)
	for i, test := range parsePacketV2Tests {
		c.Logf("test %d: %v", i, test.about)
		data, p, err := parsePacketV2([]byte(test.data))
		if test.expectError != "" {
			c.Assert(err, qt.ErrorMatches, test.expectError)
			c.Assert(data, qt.IsNil)
			c.Assert(p, packetEquals, packetV2{})
		} else {
			c.Assert(err, qt.Equals, nil)
			c.Assert(p, packetEquals, test.expectPacket)
		}
	}
}

var parseSectionV2Tests = []struct {
	about string
	data  string

	expectData    string
	expectPackets []packetV2
	expectError   string
}{{
	about: "no packets",
	data:  "\x00",
}, {
	about: "one packet",
	data:  "\x02\x03xyz\x00",
	expectPackets: []packetV2{{
		fieldType: 2,
		data:      []byte("xyz"),
	}},
}, {
	about: "two packets",
	data:  "\x02\x03xyz\x07\x05abcde\x00",
	expectPackets: []packetV2{{
		fieldType: 2,
		data:      []byte("xyz"),
	}, {
		fieldType: 7,
		data:      []byte("abcde"),
	}},
}, {
	about:       "unterminated section",
	data:        "\x02\x03xyz\x07\x05abcde",
	expectError: "section extends past end of buffer",
}, {
	about:       "out of order fields",
	data:        "\x07\x05abcde\x02\x03xyz\x00",
	expectError: "fields out of order",
}, {
	about:       "bad packet",
	data:        "\x07\x05abcde\xff",
	expectError: "varint value extends past end of buffer",
}}

func TestParseSectionV2(t *testing.T) {
	c := qt.New(t)
	for i, test := range parseSectionV2Tests {
		c.Logf("test %d: %v", i, test.about)
		data, ps, err := parseSectionV2([]byte(test.data))
		if test.expectError != "" {
			c.Assert(err, qt.ErrorMatches, test.expectError)
			c.Assert(data, qt.IsNil)
			c.Assert(ps, qt.IsNil)
		} else {
			c.Assert(err, qt.Equals, nil)
			c.Assert(ps, packetEquals, test.expectPackets)
		}
	}
}
