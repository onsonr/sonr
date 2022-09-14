package schemas_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sonr-io/sonr/internal/schemas"
	"github.com/sonr-io/sonr/pkg/client"

	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
)

func GenerateKeyForDID() string {
	return uuid.New().String()
}

func CreateMockHeirachyThreeLevel(creator string) []*st.WhatIs {
	whatIss := make([]*st.WhatIs, 0)

	did_one := fmt.Sprintf("did:snr: %s", GenerateKeyForDID())
	mockWhatIs := st.WhatIs{
		Did: did_one,
		Schema: &st.SchemaDefinition{
			Creator: creator,
			Did:     did_one,
			Label:   "testing schema",
			Fields:  make([]*st.SchemaKindDefinition, 0),
		},
		Creator:   creator,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	mockWhatIs.Schema.Fields = append(mockWhatIs.Schema.Fields, &st.SchemaKindDefinition{
		Name:  "field-1",
		Field: st.SchemaKind_INT,
	})
	mockWhatIs.Schema.Fields = append(mockWhatIs.Schema.Fields, &st.SchemaKindDefinition{
		Name:  "field-2",
		Field: st.SchemaKind_FLOAT,
	})
	mockWhatIs.Schema.Fields = append(mockWhatIs.Schema.Fields, &st.SchemaKindDefinition{
		Name:  "field-3",
		Field: st.SchemaKind_LIST,
	})

	mockWhatIs.Schema.Fields = append(mockWhatIs.Schema.Fields, &st.SchemaKindDefinition{
		Name:  "field-4",
		Field: st.SchemaKind_STRING,
	})
	mockWhatIs.Schema.Fields = append(mockWhatIs.Schema.Fields, &st.SchemaKindDefinition{
		Name:  "field-5",
		Field: st.SchemaKind_LIST,
	})
	whatIss = append(whatIss, &mockWhatIs)

	did_two := fmt.Sprintf("did:snr: %s", GenerateKeyForDID())
	mockWhatIs_2 := st.WhatIs{
		Did: did_two,
		Schema: &st.SchemaDefinition{
			Did:     did_two,
			Label:   "testing schema",
			Creator: creator,
			Fields:  make([]*st.SchemaKindDefinition, 0),
		},
		Creator:   creator,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	mockWhatIs_2.Schema.Fields = append(mockWhatIs_2.Schema.Fields, &st.SchemaKindDefinition{
		Name:  "message",
		Field: st.SchemaKind_STRING,
	})

	mockWhatIs_2.Schema.Fields = append(mockWhatIs_2.Schema.Fields, &st.SchemaKindDefinition{
		Name:  "icon",
		Field: st.SchemaKind_INT,
	})

	mockWhatIs_2.Schema.Fields = append(mockWhatIs_2.Schema.Fields, &st.SchemaKindDefinition{
		Name:  "type",
		Field: st.SchemaKind_INT,
	})

	mockWhatIs_2.Schema.Fields = append(mockWhatIs_2.Schema.Fields, &st.SchemaKindDefinition{
		Name:     "sub",
		Field:    st.SchemaKind_LINK,
		LinkKind: st.LinkKind_SCHEMA,
		Link:     did_one,
	})
	whatIss = append(whatIss, &mockWhatIs_2)

	did_three := fmt.Sprintf("did:snr: %s", GenerateKeyForDID())
	mockWhatIs_3 := st.WhatIs{
		Did: did_three,
		Schema: &st.SchemaDefinition{
			Did:     did_three,
			Label:   "testing schema",
			Creator: creator,
			Fields:  make([]*st.SchemaKindDefinition, 0),
		},
		Creator:   creator,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	mockWhatIs_3.Schema.Fields = append(mockWhatIs_3.Schema.Fields, &st.SchemaKindDefinition{
		Name:  "id",
		Field: st.SchemaKind_INT,
	})

	mockWhatIs_3.Schema.Fields = append(mockWhatIs_3.Schema.Fields, &st.SchemaKindDefinition{
		Name:  "name",
		Field: st.SchemaKind_STRING,
	})

	mockWhatIs_3.Schema.Fields = append(mockWhatIs_3.Schema.Fields, &st.SchemaKindDefinition{
		Name:     "data",
		Field:    st.SchemaKind_LINK,
		LinkKind: st.LinkKind_SCHEMA,
		Link:     did_two,
	})

	whatIss = append(whatIss, &mockWhatIs_3)

	return whatIss
}

func CreateMocks(creator string, did string) (st.WhatIs, st.SchemaDefinition) {
	mockWhatIs := st.WhatIs{
		Did: did,
		Schema: &st.SchemaDefinition{
			Did:    did,
			Label:  "testing schema",
			Fields: make([]*st.SchemaKindDefinition, 0),
		},
		Creator:   creator,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	return mockWhatIs, *mockWhatIs.Schema
}

func Test_IPLD_Nodes(t *testing.T) {
	store := &schemas.ReadStoreImpl{
		Client: client.NewClient(client.ConnEndpointType_LOCAL),
	}
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

		schema := schemas.New(store, &whatIs)

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

		schema := schemas.New(store, &whatIs)

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

		schema := schemas.New(store, &whatIs)
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

		schema := schemas.New(store, &whatIs)
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

		schema := schemas.New(store, &whatIs)
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

		schema := schemas.New(store, &whatIs)
		obj := map[string]interface{}{
			"field-1": 1,
			"field-4": 2.0,
		}
		err := schema.BuildNodesFromDefinition(obj)
		assert.Error(t, err)
	})
}

func Test_Sub_Schemas(t *testing.T) {
	whatIss := CreateMockHeirachyThreeLevel("snr12345")
	store := &schemas.ReadStoreImpl{
		Client: client.NewClient(client.ConnEndpointType_LOCAL),
	}

	for _, wi := range whatIss {
		b, err := wi.Schema.Marshal()
		if err != nil {
			assert.Error(t, err)
		}
		store.GetCache()[wi.Did] = b
	}
	t.Run("multi level sub schema should load into internal module", func(t *testing.T) {

		schema := schemas.New(store, whatIss[2])

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
		schema := schemas.New(store, whatIss[2])

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
