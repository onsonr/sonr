package schemas_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sonr-io/sonr/internal/schemas"
	"github.com/sonr-io/sonr/pkg/client"

	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
)

func GenerateKeyForDID() string {
	return uuid.New().String()
}

func CreateMockHeirachyThreeLevel(creator string) []*st.WhatIs {
	whatIss := make([]*st.WhatIs, 0)

	did_one := fmt.Sprintf("did:snr:%s", GenerateKeyForDID())
	mockWhatIs := &st.WhatIs{
		Did: did_one,
		Schema: &st.Schema{
			Owner:  creator,
			Did:    did_one,
			Label:  "testing schema",
			Fields: make([]*st.SchemaField, 0),
		},
		Creator:   creator,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	mockWhatIs.Schema.Fields = append(mockWhatIs.Schema.Fields, &st.SchemaField{
		Name: "field-1",
		FieldKind: &st.SchemaFieldKind{
			Kind: st.Kind_INT,
		},
	})
	mockWhatIs.Schema.Fields = append(mockWhatIs.Schema.Fields, &st.SchemaField{
		Name: "field-2",
		FieldKind: &st.SchemaFieldKind{
			Kind: st.Kind_FLOAT,
		},
	})
	mockWhatIs.Schema.Fields = append(mockWhatIs.Schema.Fields, &st.SchemaField{
		Name: "field-3",
		FieldKind: &st.SchemaFieldKind{
			Kind: st.Kind_LIST,
			ListKind: &st.SchemaFieldKind{
				Kind: st.Kind_INT,
			},
		},
	})

	mockWhatIs.Schema.Fields = append(mockWhatIs.Schema.Fields, &st.SchemaField{
		Name: "field-4",
		FieldKind: &st.SchemaFieldKind{
			Kind: st.Kind_STRING,
		},
	})
	mockWhatIs.Schema.Fields = append(mockWhatIs.Schema.Fields, &st.SchemaField{
		Name: "field-5",
		FieldKind: &st.SchemaFieldKind{
			Kind: st.Kind_LIST,
			ListKind: &st.SchemaFieldKind{
				Kind: st.Kind_STRING,
			},
		},
	})
	whatIss = append(whatIss, mockWhatIs)

	did_two := fmt.Sprintf("did:snr:%s", GenerateKeyForDID())
	mockWhatIs_2 := &st.WhatIs{
		Did: did_two,
		Schema: &st.Schema{
			Did:    did_two,
			Label:  "testing schema",
			Owner:  creator,
			Fields: make([]*st.SchemaField, 0),
		},
		Creator:   creator,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	mockWhatIs_2.Schema.Fields = append(mockWhatIs_2.Schema.Fields, &st.SchemaField{
		Name: "message",
		FieldKind: &st.SchemaFieldKind{
			Kind: st.Kind_STRING,
		},
	})

	mockWhatIs_2.Schema.Fields = append(mockWhatIs_2.Schema.Fields, &st.SchemaField{
		Name: "icon",
		FieldKind: &st.SchemaFieldKind{
			Kind: st.Kind_INT,
		},
	})

	mockWhatIs_2.Schema.Fields = append(mockWhatIs_2.Schema.Fields, &st.SchemaField{
		Name: "type",
		FieldKind: &st.SchemaFieldKind{
			Kind: st.Kind_INT,
		},
	})

	mockWhatIs_2.Schema.Fields = append(mockWhatIs_2.Schema.Fields, &st.SchemaField{
		Name: "sub",
		FieldKind: &st.SchemaFieldKind{
			Kind:    st.Kind_LINK,
			LinkDid: did_one,
		},
	})
	whatIss = append(whatIss, mockWhatIs_2)

	did_three := fmt.Sprintf("did:snr:%s", GenerateKeyForDID())
	mockWhatIs_3 := &st.WhatIs{
		Did: did_three,
		Schema: &st.Schema{
			Did:    did_three,
			Label:  "testing schema",
			Owner:  creator,
			Fields: make([]*st.SchemaField, 0),
		},
		Creator:   creator,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	mockWhatIs_3.Schema.Fields = append(mockWhatIs_3.Schema.Fields, &st.SchemaField{
		Name: "id",
		FieldKind: &st.SchemaFieldKind{
			Kind: st.Kind_INT,
		},
	})

	mockWhatIs_3.Schema.Fields = append(mockWhatIs_3.Schema.Fields, &st.SchemaField{
		Name: "name",
		FieldKind: &st.SchemaFieldKind{
			Kind: st.Kind_STRING,
		},
	})

	mockWhatIs_3.Schema.Fields = append(mockWhatIs_3.Schema.Fields, &st.SchemaField{
		Name: "data",
		FieldKind: &st.SchemaFieldKind{
			Kind:    st.Kind_LINK,
			LinkDid: did_two,
		},
	})

	whatIss = append(whatIss, mockWhatIs_3)

	return whatIss
}

func CreateMocks(creator string, did string) (*st.WhatIs, *st.Schema) {
	mockWhatIs := &st.WhatIs{
		Did: did,
		Schema: &st.Schema{
			Did:    did,
			Label:  "testing schema",
			Fields: make([]*st.SchemaField, 0),
		},
		Creator:   creator,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	return mockWhatIs, mockWhatIs.Schema
}

func Test_IPLD_Nodes(t *testing.T) {
	store := &schemas.ReadStoreImpl{
		Client: client.NewClient(mt.ClientMode_ENDPOINT_BETA),
	}
	t.Run("Should build Nodes and store in map", func(t *testing.T) {
		whatIs, def := CreateMocks("snr12345", "did:snr:1234")
		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-1",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_INT,
			},
		})
		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-2",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_FLOAT,
			},
		})

		schema := schemas.New(store, whatIs)

		doc := map[string]interface{}{
			"field-1": 1,
			"field-2": 2.0,
		}
		err := schema.BuildNodesFromDefinition("doc", whatIs.Did, doc)
		assert.NoError(t, err)

		n, err := schema.GetNode()
		assert.NoError(t, err)

		assert.NotNil(t, n)
	})

	t.Run("Should build Nodes from definition", func(t *testing.T) {
		whatIs, def := CreateMocks("snr12345", "did:snr:1234")

		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-1",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_INT,
			},
		})
		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-2",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_FLOAT,
			},
		})
		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-3",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_LIST,
				ListKind: &st.SchemaFieldKind{
					Kind: st.Kind_INT,
				},
			},
		})

		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-4",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_STRING,
			},
		})
		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-5",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_LIST,
				ListKind: &st.SchemaFieldKind{
					Kind: st.Kind_STRING,
				},
			},
		})

		schema := schemas.New(store, whatIs)

		doc := map[string]interface{}{
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
		err := schema.BuildNodesFromDefinition("doc", whatIs.Did, doc)
		assert.NoError(t, err)

		n, err := schema.GetNode()
		assert.NoError(t, err)

		assert.NotNil(t, n)
	})

	t.Run("Should build Nodes from definition, should encode and decode correctly (JSON)", func(t *testing.T) {
		whatIs, def := CreateMocks("snr12345", "did:snr:1234")

		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-1",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_INT,
			},
		})
		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-2",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_FLOAT,
			},
		})

		schema := schemas.New(store, whatIs)
		doc := map[string]interface{}{
			"field-1": 1,
			"field-2": 2.0,
		}
		err := schema.BuildNodesFromDefinition("doc", whatIs.Did, doc)
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

		document, err := n.LookupByString(st.IPLD_DOCUMENT)
		assert.NoError(t, err)

		found, err := document.LookupByString("field-1")
		assert.NoError(t, err)
		assert.NotNil(t, found)
	})

	t.Run("Should build Nodes from definition, should encode and decode correctly (JSON)", func(t *testing.T) {
		whatIs, def := CreateMocks("snr12345", "did:snr:1234")

		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-1",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_INT,
			},
		})
		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-2",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_FLOAT,
			},
		})

		schema := schemas.New(store, whatIs)
		doc := map[string]interface{}{
			"field-1": 1,
			"field-2": 2.0,
		}
		err := schema.BuildNodesFromDefinition("doc", whatIs.Did, doc)
		assert.NoError(t, err)

		enc, err := schema.EncodeDagJson()
		assert.NoError(t, err)
		assert.NotNil(t, enc)

		err = schema.DecodeDagJson(enc)
		assert.NoError(t, err)

		n, err := schema.GetNode()
		assert.NoError(t, err)

		document, err := n.LookupByString(st.IPLD_DOCUMENT)
		assert.NoError(t, err)

		found, err := document.LookupByString("field-1")
		assert.NoError(t, err)
		assert.NotNil(t, found)
	})

	t.Run("Should build Nodes from definition, should encode and decode correctly (CBOR)", func(t *testing.T) {
		whatIs, def := CreateMocks("snr12345", "did:snr:1234")

		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-1",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_INT,
			},
		})
		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-2",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_FLOAT,
			},
		})

		schema := schemas.New(store, whatIs)
		doc := map[string]interface{}{
			"field-1": 1,
			"field-2": 2.0,
		}
		err := schema.BuildNodesFromDefinition("doc", whatIs.Did, doc)
		assert.NoError(t, err)

		enc, err := schema.EncodeDagCbor()
		assert.NoError(t, err)
		assert.NotNil(t, enc)

		err = schema.DecodeDagCbor(enc)
		assert.NoError(t, err)

		n, err := schema.GetNode()
		assert.NoError(t, err)

		document, err := n.LookupByString(st.IPLD_DOCUMENT)
		assert.NoError(t, err)

		found, err := document.LookupByString("field-1")
		assert.NoError(t, err)
		assert.NotNil(t, found)
	})

	t.Run("Should throw invalid error with mismatching definitions", func(t *testing.T) {
		whatIs, def := CreateMocks("snr12345", "did:snr:1234")

		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-1",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_INT,
			},
		})
		def.Fields = append(def.Fields, &st.SchemaField{
			Name: "field-2",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_STRING,
			},
		})

		schema := schemas.New(store, whatIs)
		doc := map[string]interface{}{
			"field-1": 1,
			"field-4": 2.0,
		}
		err := schema.BuildNodesFromDefinition("doc", whatIs.Did, doc)
		assert.Error(t, err)
	})
}

func Test_List_Types(t *testing.T) {
	store := &schemas.ReadStoreImpl{
		Client: client.NewClient(mt.ClientMode_ENDPOINT_BETA),
	}
	t.Run("Should build Nodes and store in map", func(t *testing.T) {
		whatIs, _ := CreateMocks("snr12345", "did:snr:1234")
		whatIs.Schema.Fields = append(whatIs.Schema.Fields, &st.SchemaField{
			Name: "field-1",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_LIST,
				ListKind: &st.SchemaFieldKind{
					Kind: st.Kind_INT,
				},
			},
		})

		whatIs.Schema.Fields = append(whatIs.Schema.Fields, &st.SchemaField{
			Name: "field-2",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_LIST,
				ListKind: &st.SchemaFieldKind{
					Kind: st.Kind_FLOAT,
				},
			},
		})

		whatIs.Schema.Fields = append(whatIs.Schema.Fields, &st.SchemaField{
			Name: "field-3",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_LIST,
				ListKind: &st.SchemaFieldKind{
					Kind: st.Kind_STRING,
				},
			},
		})

		whatIs.Schema.Fields = append(whatIs.Schema.Fields, &st.SchemaField{
			Name: "field-4",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_LIST,
				ListKind: &st.SchemaFieldKind{
					Kind: st.Kind_BOOL,
				},
			},
		})

		whatIs.Schema.Fields = append(whatIs.Schema.Fields, &st.SchemaField{
			Name: "field-5",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_LIST,
				ListKind: &st.SchemaFieldKind{
					Kind: st.Kind_BYTES,
				},
			},
		})

		whatIs.Schema.Fields = append(whatIs.Schema.Fields, &st.SchemaField{
			Name: "field-6",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_LIST,
				ListKind: &st.SchemaFieldKind{
					Kind: st.Kind_LIST,
					ListKind: &st.SchemaFieldKind{
						Kind: st.Kind_STRING,
					},
				},
			},
		})

		whatIs.Schema.Fields = append(whatIs.Schema.Fields, &st.SchemaField{
			Name: "field-7",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_LIST,
				ListKind: &st.SchemaFieldKind{
					Kind: st.Kind_LIST,
					ListKind: &st.SchemaFieldKind{
						Kind: st.Kind_LIST,
						ListKind: &st.SchemaFieldKind{
							Kind: st.Kind_INT,
						},
					},
				},
			},
		})

		schema := schemas.New(store, whatIs)

		doc := map[string]interface{}{
			"field-1": []int32{
				1, 2, 3, 4,
			},
			"field-2": []float32{
				2.1, 2.2, 3.1, 3.2,
			},
			"field-3": []string{
				"1", "2", "3", "4",
			},
			"field-4": []bool{
				true, true, false,
			},
			"field-5": [][]byte{
				[]byte("hello"),
				[]byte("world"),
			},
			"field-6": [][]string{
				{
					"hello",
				},
				{
					"world",
				},
			},
			"field-7": [][][]int64{
				{
					{
						1, 2, 4,
					},
					{
						1, 2, 4,
					},
				},
			},
		}

		err := schema.BuildNodesFromDefinition("doc", whatIs.Did, doc)
		assert.NoError(t, err)

		n, err := schema.GetNode()
		assert.NoError(t, err)

		assert.NotNil(t, n)
	})

	t.Run("Should validate link types in arrays", func(t *testing.T) {
		whatIs, _ := CreateMocks("snr12345", "did:snr:1234")
		whatIss := CreateMockHeirachyThreeLevel("snr12345")
		store := &schemas.ReadStoreImpl{
			Client: client.NewClient(mt.ClientMode_ENDPOINT_BETA),
		}

		for _, wi := range whatIss {
			b, err := wi.Marshal()
			if err != nil {
				assert.Error(t, err)
			}
			store.GetCache()[wi.Did] = b
		}
		whatIs.Schema.Fields = append(whatIs.Schema.Fields, &st.SchemaField{
			Name: "field-1",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_LIST,
				ListKind: &st.SchemaFieldKind{
					Kind:    st.Kind_LINK,
					LinkDid: whatIss[2].Did,
				},
			},
		})

		doc := map[string]interface{}{
			"field-1": []map[string]interface{}{
				{
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
				},
			},
		}

		schema := schemas.New(store, whatIs)

		err := schema.BuildNodesFromDefinition("doc", whatIs.Did, doc)
		assert.NoError(t, err)

		node, err := schema.GetNode()
		assert.NoError(t, err)
		assert.NotNil(t, node)
	})

	t.Run("Should throw error if object does not match link structure", func(t *testing.T) {
		whatIs, _ := CreateMocks("snr12345", "did:snr:1234")
		whatIss := CreateMockHeirachyThreeLevel("snr12345")
		store := &schemas.ReadStoreImpl{
			Client: client.NewClient(mt.ClientMode_ENDPOINT_BETA),
		}

		for _, wi := range whatIss {
			b, err := wi.Marshal()
			if err != nil {
				assert.Error(t, err)
			}
			store.GetCache()[wi.Did] = b
		}
		whatIs.Schema.Fields = append(whatIs.Schema.Fields, &st.SchemaField{
			Name: "field-1",
			FieldKind: &st.SchemaFieldKind{
				Kind: st.Kind_LIST,
				ListKind: &st.SchemaFieldKind{
					Kind:    st.Kind_LINK,
					LinkDid: whatIss[2].Did,
				},
			},
		})

		doc := map[string]interface{}{
			"field-1": []map[string]interface{}{
				{
					"id":   1,
					"name": "asAASD",
					"data": map[string]interface{}{
						"icon":    "heyy",
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
				},
			},
		}

		schema := schemas.New(store, whatIs)

		err := schema.BuildNodesFromDefinition("doc", whatIs.Did, doc)
		assert.Error(t, err)
		assert.NotNil(t, err)
	})
}

func Test_Sub_Schemas(t *testing.T) {
	whatIss := CreateMockHeirachyThreeLevel("snr12345")
	store := &schemas.ReadStoreImpl{
		Client: client.NewClient(mt.ClientMode_ENDPOINT_BETA),
	}

	for _, wi := range whatIss {
		b, err := wi.Marshal()
		if err != nil {
			assert.Error(t, err)
		}
		store.GetCache()[wi.Did] = b
	}
	t.Run("multi level sub schema should load into internal module", func(t *testing.T) {

		schema := schemas.New(store, whatIss[2])

		doc := map[string]interface{}{
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

		err := schema.BuildNodesFromDefinition("doc", whatIss[2].Did, doc)

		assert.NoError(t, err)
		bytes, err := schema.EncodeDagJson()
		assert.NoError(t, err)
		err = schema.DecodeDagJson(bytes)
		assert.NoError(t, err)
	})

	t.Run("multi level sub schema should error with invalid types", func(t *testing.T) {
		schema := schemas.New(store, whatIss[2])

		doc := map[string]interface{}{
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

		err := schema.BuildNodesFromDefinition("doc", whatIss[2].Did, doc)

		assert.Error(t, err)
	})
}
