package did

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"testing"

	odid "github.com/ockam-network/did"
	"github.com/stretchr/testify/assert"
)

func TestDID_UnmarshalJSON(t *testing.T) {
	jsonTestSting := `"did:snr:123"`

	id := DID{}
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
	id := DID{*wrappedDid}
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
		id, err := ParseDID("did:snr:123")

		if err != nil {
			t.Errorf("unexpected error: %s", err)
			return
		}

		if id.String() != "did:snr:123" {
			t.Errorf("expected parsed did to be 'did:snr:123', got: %s", id.String())
		}
	})
	t.Run("ok - parse a DID URL", func(t *testing.T) {
		id, err := ParseDID("did:snr:123/path?query#fragment")
		assert.Equal(t, "did:snr:123/path?query#fragment", id.String())
		assert.NoError(t, err)
	})

	t.Run("error - invalid DID", func(t *testing.T) {
		id, err := ParseDID("invalidDID")
		assert.Nil(t, id)
		assert.EqualError(t, err, "invalid DID: input does not begin with 'did:' prefix")

	})
}

func TestMustParseDID(t *testing.T) {
	assert.Panics(t, func() {
		MustParseDID("did:snr:123/path?query#fragment")
	})
}

func TestDID_String(t *testing.T) {
	expected := "did:snr:123"
	id, _ := ParseDID(expected)
	assert.Equal(t, expected, fmt.Sprintf("%s", *id))
}

func TestDID_MarshalText(t *testing.T) {
	expected := "did:snr:123"
	id, _ := ParseDID(expected)
	actual, err := id.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, []byte(expected), actual)
}

func TestDID_Empty(t *testing.T) {
	t.Run("not empty for filled did", func(t *testing.T) {
		id, err := ParseDID("did:snr:123")
		if err != nil {
			t.Errorf("unexpected error: %s", err)
			return
		}
		assert.False(t, id.Empty())
	})

	t.Run("empty when just generated", func(t *testing.T) {
		id := DID{}
		assert.True(t, id.Empty())
	})
}

func TestDID_URI(t *testing.T) {
	id, err := ParseDID("did:snr:123")

	if !assert.NoError(t, err) {
		return
	}

	uri := id.URI()

	assert.Equal(t, id.String(), uri.String())
}

func TestError(t *testing.T) {
	actual := ErrInvalidDID.wrap(io.EOF)
	assert.True(t, errors.Is(actual, ErrInvalidDID))
	assert.True(t, errors.Is(actual, io.EOF))
	assert.False(t, errors.Is(actual, io.ErrShortBuffer))
}
