package schemas_test

import (
	"testing"
	"time"

	"github.com/sonr-io/sonr/internal/motor/x/schemas"
	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
)

func CreateMocks(creator string, did string) (st.WhatIs, st.SchemaDefinition) {
	mockWhatIs := st.WhatIs{
		Did: did,
		Schema: &st.SchemaReference{
			Did:   did,
			Label: "testing schema",
			Cid:   "asdasd12312",
		},
		Creator:   creator,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}
	def := st.SchemaDefinition{
		Creator: "snr123456",
		Label:   "testing schema",
		Fields:  make(map[string]st.SchemaKind),
	}

	return mockWhatIs, def
}

func Test_IPLD_Nodes(t *testing.T) {
	schema := schemas.New()
	t.Run("Should build Nodes and store in map", func(t *testing.T) {
		_, def := CreateMocks("snr12345", "did:snr:1234")
		def.Fields["field-1"] = st.SchemaKind_INT
		def.Fields["field-2"] = st.SchemaKind_FLOAT
		obj := map[string]interface{}{
			"field-1": 1,
			"field-2": 2.0,
		}
		node, err := schema.BuildNodesFromDefinition(&def, obj)
		assert.NoError(t, err)

		assert.NotNil(t, node)
	})

	t.Run("Should build Nodes from definition", func(t *testing.T) {
		_, def := CreateMocks("snr12345", "did:snr:1234")
		def.Fields["field-1"] = st.SchemaKind_INT
		def.Fields["field-2"] = st.SchemaKind_FLOAT
		obj := map[string]interface{}{
			"field-1": 1,
			"field-2": 2.0,
		}
		node, err := schema.BuildNodesFromDefinition(&def, obj)
		assert.NoError(t, err)
		assert.NotNil(t, node)
	})

	t.Run("Should build Nodes from definition, should encode and decode correctly (JSON)", func(t *testing.T) {
		_, def := CreateMocks("snr12345", "did:snr:1234")
		def.Fields["field-1"] = st.SchemaKind_INT
		def.Fields["field-2"] = st.SchemaKind_FLOAT
		obj := map[string]interface{}{
			"field-1": 1,
			"field-2": 2.0,
		}
		node, err := schema.BuildNodesFromDefinition(&def, obj)
		assert.NoError(t, err)
		assert.NotNil(t, node)

		enc, err := schema.EncodeDagJson(node)
		assert.NoError(t, err)
		assert.NotNil(t, enc)
		dec, err := schema.DecodeDagJson(enc)
		assert.NoError(t, err)
		found, err := dec.LookupByString("field-1")
		assert.NoError(t, err)
		assert.NotNil(t, found)
	})

	t.Run("Should throw invalid error with mismatching definitions", func(t *testing.T) {
		_, def := CreateMocks("snr12345", "did:snr:1234")

		def.Fields["field-1"] = st.SchemaKind_INT
		def.Fields["field-2"] = st.SchemaKind_FLOAT
		obj := map[string]interface{}{
			"field-1": 1,
			"field-4": 2.0,
		}
		_, err := schema.BuildNodesFromDefinition(&def, obj)
		assert.Error(t, err)
	})
}
