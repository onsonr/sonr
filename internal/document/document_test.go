package document

import (
	"fmt"
	"testing"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/internal/schemas"
	"github.com/sonr-io/sonr/pkg/client"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
)

func CreateMockSchemaDefinition() (*st.WhatIs, map[string]interface{}) {
	doc := make(map[string]interface{})
	wi := &st.WhatIs{
		Did:     "did:snr:123456",
		Creator: "snr123456",
		Schema: &st.Schema{
			Did:    "did:snr:123456",
			Owner:  "snr123456",
			Label:  "testing schema",
			Fields: make([]*st.SchemaField, 0),
		},
	}

	for i := 1; i < 10000; i++ {
		name := fmt.Sprintf("field-%d", i)
		if i%2 == 0 {
			wi.Schema.Fields = append(wi.Schema.Fields, &st.SchemaField{
				Name: name,
				FieldKind: &st.SchemaFieldKind{
					Kind: st.Kind_INT,
				},
			})
			doc[name] = i
		} else if i%3 == 0 {
			wi.Schema.Fields = append(wi.Schema.Fields, &st.SchemaField{
				Name: name,
				FieldKind: &st.SchemaFieldKind{
					Kind: st.Kind_BOOL,
				},
			})
			doc[name] = true
		} else if i%7 == 0 {
			wi.Schema.Fields = append(wi.Schema.Fields, &st.SchemaField{
				Name: name,
				FieldKind: &st.SchemaFieldKind{
					Kind: st.Kind_FLOAT,
				},
			})
			doc[name] = 123.456
		} else if i%13 == 0 {
			wi.Schema.Fields = append(wi.Schema.Fields, &st.SchemaField{
				Name: name,
				FieldKind: &st.SchemaFieldKind{
					Kind: st.Kind_STRING,
				},
			})
			doc[name] = fmt.Sprintf("%d", i)
		} else {
			wi.Schema.Fields = append(wi.Schema.Fields, &st.SchemaField{
				Name: name,
				FieldKind: &st.SchemaFieldKind{
					Kind: st.Kind_BYTES,
				},
			})
			doc[name] = []byte("Hello-world")
		}
	}

	return wi, doc
}
func Test_Document(t *testing.T) {
	t.Skip("Skipping test in CI")
	store := &schemas.ReadStoreImpl{
		Client: client.NewClient(mt.ClientMode_ENDPOINT_BETA),
	}
	config := Config{}
	def, jsonData := CreateMockSchemaDefinition()
	schema := schemas.New(store, def)
	config.WithStorage(shell.NewShell("localhost:5001"))
	config.WithSchemaImpl(schema)

	doc := NewWithConfig(&config)

	t.Run("Should upload document", func(t *testing.T) {
		res, err := doc.CreateDocument("testing", schema.GetDID(), jsonData)
		assert.NoError(t, err)
		fmt.Print(res)

		data, err := doc.GetDocument(res.Cid)
		assert.NoError(t, err)
		assert.NotNil(t, data)
	})
}
