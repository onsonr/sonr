package schemas_test

import (
	"testing"
	"time"

	"github.com/sonr-io/sonr/pkg/schemas"
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
		Fields:  make([]*st.SchemaKindDefinition, 0),
	}

	return mockWhatIs, def
}

func Test_IPLD_Nodes(t *testing.T) {

	t.Run("Should build Nodes and store in map", func(t *testing.T) {
		whatIs, def := CreateMocks("snr12345", "did:snr:1234")
		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-1",
			Field: st.SchemaKind_INT,
		})
		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-2",
			Field: st.SchemaKind_FLOAT,
		})

		schema := schemas.New(def.Fields, &whatIs)

		obj := map[string]interface{}{
			"field-1": 1,
			"field-2": 2.0,
		}
		err := schema.BuildNodesFromDefinition(obj)
		assert.NoError(t, err)

		n, err := schema.GetNode()
		assert.NoError(t, err)

		assert.NotNil(t, n)
	})

	t.Run("Should build Nodes from definition", func(t *testing.T) {
		whatIs, def := CreateMocks("snr12345", "did:snr:1234")

		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-1",
			Field: st.SchemaKind_INT,
		})
		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-2",
			Field: st.SchemaKind_FLOAT,
		})
		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-3",
			Field: st.SchemaKind_LIST,
		})

		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-4",
			Field: st.SchemaKind_STRING,
		})
		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-5",
			Field: st.SchemaKind_LIST,
		})

		schema := schemas.New(def.Fields, &whatIs)

		obj := map[string]interface{}{
			"field-1": 1,
			"field-2": 2.0,
			"field-3": []int{
				1, 2, 3, 4,
			},
			"field-4": "hey there",
			"field-5": []string{
				"hey",
				"there",
			},
		}
		err := schema.BuildNodesFromDefinition(obj)
		assert.NoError(t, err)

		n, err := schema.GetNode()
		assert.NoError(t, err)

		assert.NotNil(t, n)
	})

	t.Run("Should build Nodes from definition, should encode and decode correctly (JSON)", func(t *testing.T) {
		whatIs, def := CreateMocks("snr12345", "did:snr:1234")

		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-1",
			Field: st.SchemaKind_INT,
		})
		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-2",
			Field: st.SchemaKind_FLOAT,
		})

		schema := schemas.New(def.Fields, &whatIs)
		obj := map[string]interface{}{
			"field-1": 1,
			"field-2": 2.0,
		}
		err := schema.BuildNodesFromDefinition(obj)
		assert.NoError(t, err)
		n, err := schema.GetNode()
		assert.NoError(t, err)

		assert.NotNil(t, n)

		enc, err := schema.EncodeDagJson()
		assert.NoError(t, err)
		assert.NotNil(t, enc)
		err = schema.DecodeDagJson(enc)
		assert.NoError(t, err)

		n, err = schema.GetNode()
		assert.NoError(t, err)

		found, err := n.LookupByString("field-1")
		assert.NoError(t, err)
		assert.NotNil(t, found)
	})

	t.Run("Should build Nodes from definition, should encode and decode correctly (JSON)", func(t *testing.T) {
		whatIs, def := CreateMocks("snr12345", "did:snr:1234")

		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-1",
			Field: st.SchemaKind_INT,
		})
		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-2",
			Field: st.SchemaKind_FLOAT,
		})

		schema := schemas.New(def.Fields, &whatIs)
		obj := map[string]interface{}{
			"field-1": 1,
			"field-2": 2.0,
		}
		err := schema.BuildNodesFromDefinition(obj)
		assert.NoError(t, err)

		enc, err := schema.EncodeDagJson()
		assert.NoError(t, err)
		assert.NotNil(t, enc)
		err = schema.DecodeDagJson(enc)
		assert.NoError(t, err)
		n, err := schema.GetNode()
		assert.NoError(t, err)
		found, err := n.LookupByString("field-1")
		assert.NoError(t, err)
		assert.NotNil(t, found)
	})

	t.Run("Should build Nodes from definition, should encode and decode correctly (CBOR)", func(t *testing.T) {
		whatIs, def := CreateMocks("snr12345", "did:snr:1234")

		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-1",
			Field: st.SchemaKind_INT,
		})
		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-2",
			Field: st.SchemaKind_FLOAT,
		})

		schema := schemas.New(def.Fields, &whatIs)
		obj := map[string]interface{}{
			"field-1": 1,
			"field-2": 2.0,
		}
		err := schema.BuildNodesFromDefinition(obj)
		assert.NoError(t, err)

		enc, err := schema.EncodeDagCbor()
		assert.NoError(t, err)
		assert.NotNil(t, enc)
		err = schema.DecodeDagCbor(enc)
		assert.NoError(t, err)
		n, err := schema.GetNode()
		assert.NoError(t, err)
		found, err := n.LookupByString("field-1")
		assert.NoError(t, err)
		assert.NotNil(t, found)
	})

	t.Run("Should throw invalid error with mismatching definitions", func(t *testing.T) {
		whatIs, def := CreateMocks("snr12345", "did:snr:1234")

		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-1",
			Field: st.SchemaKind_INT,
		})
		def.Fields = append(def.Fields, &st.SchemaKindDefinition{
			Name:  "field-2",
			Field: st.SchemaKind_STRING,
		})

		schema := schemas.New(def.Fields, &whatIs)
		obj := map[string]interface{}{
			"field-1": 1,
			"field-4": 2.0,
		}
		err := schema.BuildNodesFromDefinition(obj)
		assert.Error(t, err)
	})
}
