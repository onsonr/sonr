package object_test

import (
	"fmt"
	"testing"

	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/object"
	"github.com/sonr-io/sonr/pkg/schemas"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func CreateMockSchemaDefinition() (st.SchemaDefinition, map[string]interface{}) {
	obj := make(map[string]interface{})

	def := st.SchemaDefinition{
		Creator: "snr123456",
		Label:   "testing schema",
		Fields:  make([]*st.SchemaKindDefinition, 0),
	}
	for i := 1; i < 10000; i++ {
		name := fmt.Sprintf("field-%d", i)
		if i%2 == 0 {
			def.Fields = append(def.Fields, &st.SchemaKindDefinition{
				Name:  name,
				Field: st.SchemaKind_INT,
			})
			obj[name] = i
		} else if i%3 == 0 {
			def.Fields = append(def.Fields, &st.SchemaKindDefinition{
				Name:  name,
				Field: st.SchemaKind_BOOL,
			})
			obj[name] = true
		} else if i%7 == 0 {
			def.Fields = append(def.Fields, &st.SchemaKindDefinition{
				Name:  name,
				Field: st.SchemaKind_FLOAT,
			})
			obj[name] = 123.456
		} else if i%13 == 0 {
			def.Fields = append(def.Fields, &st.SchemaKindDefinition{
				Name:  name,
				Field: st.SchemaKind_STRING,
			})
			obj[name] = fmt.Sprintf("%d", i)
		} else {
			def.Fields = append(def.Fields, &st.SchemaKindDefinition{
				Name:  name,
				Field: st.SchemaKind_BYTES,
			})
			obj[name] = []byte("Hello-world")
		}
	}

	return def, obj
}
func Test_Object(t *testing.T) {
	config := object.Config{}
	config.WithSchemaImplementation(schemas.New("https://api.ipfs.sonr.ws", client.ConnEndpointType_LOCAL))
	config.WithStorageEndpoint("https://api.ipfs.sonr.ws")
	obj := object.New(&config)

	t.Run("Should upload object", func(t *testing.T) {
		def, jsonData := CreateMockSchemaDefinition()
		res, err := obj.CreateObject("testing", def.Fields, jsonData)
		t.Error(err)
		fmt.Print(res)

		data, err := obj.GetObject(res.Definition.Cid)
		fmt.Print(data)
	})
}
