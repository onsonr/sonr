package schemas_test

import (
	"bytes"
	"testing"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/internal/schemas"

	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
)

func CreateMockHeirachyThreeLevel(creator string, did string) (st.WhatIs, st.SchemaDefinition) {

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

	buf, err := def.Marshal()
	if err != nil {
		panic("Unable to serialize test data")
	}

	cid, err := shell.NewLocalShell().Add(bytes.NewReader(buf))

	if err != nil {
		panic("error while persisting mock data")
	}

	commentDef := st.SchemaDefinition{
		Creator: "snr1234",
		Label:   "MY App Comment",
		Fields:  make([]*st.SchemaKindDefinition, 0),
	}

	commentDef.Fields = append(commentDef.Fields, &st.SchemaKindDefinition{
		Name:  "message",
		Field: st.SchemaKind_STRING,
	})

	commentDef.Fields = append(commentDef.Fields, &st.SchemaKindDefinition{
		Name:  "icon",
		Field: st.SchemaKind_INT,
	})

	commentDef.Fields = append(commentDef.Fields, &st.SchemaKindDefinition{
		Name:  "type",
		Field: st.SchemaKind_INT,
	})

	commentDef.Fields = append(commentDef.Fields, &st.SchemaKindDefinition{
		Name:     "sub",
		Field:    st.SchemaKind_LINK,
		LinkKind: st.LinkKind_SCHEMA,
		Link:     cid,
	})

	commentBuf, err := commentDef.Marshal()
	if err != nil {
		panic("Unable to serialize test data")
	}

	commentCid, err := shell.NewLocalShell().Add(bytes.NewReader(commentBuf))

	if err != nil {
		panic("error while attempting to persist mocks")
	}

	topDef := st.SchemaDefinition{
		Creator: "snr1234",
		Label:   "MY App Comment",
		Fields:  make([]*st.SchemaKindDefinition, 0),
	}

	topDef.Fields = append(topDef.Fields, &st.SchemaKindDefinition{
		Name:  "id",
		Field: st.SchemaKind_INT,
	})

	topDef.Fields = append(topDef.Fields, &st.SchemaKindDefinition{
		Name:  "name",
		Field: st.SchemaKind_STRING,
	})

	topDef.Fields = append(topDef.Fields, &st.SchemaKindDefinition{
		Name:     "data",
		Field:    st.SchemaKind_LINK,
		LinkKind: st.LinkKind_SCHEMA,
		Link:     commentCid,
	})

	return mockWhatIs, topDef
}

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

func Test_Sub_Schemas(t *testing.T) {
	t.Skip("Skipping for CI")
	whatIs, def := CreateMockHeirachyThreeLevel("snr12345", "did:snr:1234")

	t.Run("multi level sub schema should load into internal module", func(t *testing.T) {
		schema := schemas.New(def.Fields, &whatIs)

		obj := map[string]interface{}{
			"id":   1,
			"name": "asAASD",
			"data": map[string]interface{}{
				"icon":    1,
				"message": "asdasd",
				"type":    2,
				"sub": map[string]interface{}{
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
				},
			},
		}

		err := schema.BuildNodesFromDefinition(obj)

		assert.NoError(t, err)
		bytes, err := schema.EncodeDagJson()
		assert.NoError(t, err)
		err = schema.DecodeDagJson(bytes)
		assert.NoError(t, err)
	})

	t.Run("multi level sub schema should error with invalid types", func(t *testing.T) {
		schema := schemas.New(def.Fields, &whatIs)

		obj := map[string]interface{}{
			"id":   1,
			"name": "asAASD",
			"data": map[string]interface{}{
				"icon":    1,
				"message": "hello/tworld",
				"type":    "bad_value",
				"sub": map[string]interface{}{
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
				},
			},
		}

		err := schema.BuildNodesFromDefinition(obj)

		assert.Error(t, err)
	})
}
