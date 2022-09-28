package object_test

import (
	"fmt"
	"testing"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/internal/object"
	"github.com/sonr-io/sonr/internal/schemas"
	"github.com/sonr-io/sonr/pkg/client"
	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
)

func CreateMockSchemaDefinition() (*st.WhatIs, map[string]interface{}) {
	obj := make(map[string]interface{})
	wi := &st.WhatIs{
		Did:     "did:snr:123456",
		Creator: "snr123456",
		Schema: &st.SchemaDefinition{
			Did:     "did:snr:123456",
			Creator: "snr123456",
			Label:   "testing schema",
			Fields:  make([]*st.SchemaKindDefinition, 0),
		},
	}

	for i := 1; i < 10000; i++ {
		name := fmt.Sprintf("field-%d", i)
		if i%2 == 0 {
			wi.Schema.Fields = append(wi.Schema.Fields, &st.SchemaKindDefinition{
				Name:  name,
				Field: st.SchemaKind_INT,
			})
			obj[name] = i
		} else if i%3 == 0 {
			wi.Schema.Fields = append(wi.Schema.Fields, &st.SchemaKindDefinition{
				Name:  name,
				Field: st.SchemaKind_BOOL,
			})
			obj[name] = true
		} else if i%7 == 0 {
			wi.Schema.Fields = append(wi.Schema.Fields, &st.SchemaKindDefinition{
				Name:  name,
				Field: st.SchemaKind_FLOAT,
			})
			obj[name] = 123.456
		} else if i%13 == 0 {
			wi.Schema.Fields = append(wi.Schema.Fields, &st.SchemaKindDefinition{
				Name:  name,
				Field: st.SchemaKind_STRING,
			})
			obj[name] = fmt.Sprintf("%d", i)
		} else {
			wi.Schema.Fields = append(wi.Schema.Fields, &st.SchemaKindDefinition{
				Name:  name,
				Field: st.SchemaKind_BYTES,
			})
			obj[name] = []byte("Hello-world")
		}
	}

	return wi, obj
}
func Test_Object(t *testing.T) {
	t.Skip("Skipping test in CI")
	store := &schemas.ReadStoreImpl{
		Client: client.NewClient(client.ConnEndpointType_LOCAL),
	}
	config := object.Config{}
	def, jsonData := CreateMockSchemaDefinition()
	schema := schemas.New(store, def)
	config.WithStorage(shell.NewShell("localhost:5001"))
	config.WithSchemaImpl(schema)

	obj := object.NewWithConfig(&config)

	t.Run("Should upload object", func(t *testing.T) {
		res, err := obj.CreateObject("testing", jsonData)
		assert.NoError(t, err)
		fmt.Print(res)

		data, err := obj.GetObject(res.Reference.Cid)
		assert.NoError(t, err)
		assert.NotNil(t, data)
	})
}
