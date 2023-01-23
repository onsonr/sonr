package types_test

import (
	"encoding/json"
	"testing"

	odid "github.com/ockam-network/did"
	uri "github.com/sonrhq/core/x/identity/types"
	"github.com/stretchr/testify/assert"
)

func TestDID_UnmarshalJSON(t *testing.T) {
	jsonTestSting := `"did:snr:123"`

	id := uri.DID{}
	err := json.Unmarshal([]byte(jsonTestSting), &id)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	if id.DID.Method != "snr" {
		t.Errorf("expected snr got %s", id.Method)
		return
	}
}

func TestDID_MarshalJSON(t *testing.T) {
	wrappedDid, err := odid.Parse("did:snr:123")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	id := uri.DID{*wrappedDid}
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	result, err := json.Marshal(id)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	if string(result) != `"did:snr:123"` {
		t.Errorf("expected \"did:snr:123\" got: %s", result)
	}

}

func TestParseDID(t *testing.T) {
	t.Run("parse a DID", func(t *testing.T) {
		id, err := uri.ParseDID("did:snr:123")

		if err != nil {
			t.Errorf("unexpected error: %s", err)
			return
		}

		if id.String() != "did:snr:123" {
			t.Errorf("expected parsed did to be 'did:snr:123', got: %s", id.String())
		}
	})
	t.Run("ok - parse a DID URL", func(t *testing.T) {
		id, err := uri.ParseDID("did:snr:123/path?query#fragment")
		assert.Equal(t, "did:snr:123/path?query#fragment", id.String(), "DID parses correctly")
		assert.Equal(t, "fragment", id.Fragment, "Fragment parses correctly")
		assert.Equal(t, "snr", id.Method, "Method parses currectly")
		assert.Equal(t, "path", id.Path, "Path parses currectly")
		assert.Equal(t, "123", id.ID, "ID parses currectly")
		assert.NoError(t, err)
	})

	t.Run("error - invalid DID", func(t *testing.T) {
		id, err := uri.ParseDID("invalidDID")
		assert.Nil(t, id)
		assert.EqualError(t, err, "invalid DID: input does not begin with 'did:' prefix")

	})
	t.Run("error - DID URL", func(t *testing.T) {
		id, err := uri.ParseDID("did:snr:123/path?query#fragment")

		assert.Nil(t, err)
		// After parsing the did, the parser should ignore everyhing after '/' as it is non valid did formation.
		assert.Len(t, id.PathSegments, 1)
		assert.Equal(t, id.IDStrings[0], "123")
	})
}
