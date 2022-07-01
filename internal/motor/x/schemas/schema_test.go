package schemas_test

import (
	"testing"
	"time"

	"github.com/sonr-io/sonr/internal/motor/x/schemas"
	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
)

func Test_IPLD_Schemas(t *testing.T) {
	schema := schemas.New()
	t.Run("Should build Schema from definition", func(t *testing.T) {
		mockWhatIs := st.WhatIs{
			Did: "did:snr:1234",
			Schema: &st.SchemaReference{
				Did:   "did:snr:1234",
				Label: "testing schema",
				Cid:   "asdasd12312",
			},
			Creator:   "snr123456",
			Timestamp: time.Now().Unix(),
			IsActive:  true,
		}
		def := st.SchemaDefinition{
			Creator: "snr123456",
			Label:   "testing schema",
			Fields:  make(map[string]st.SchemaKind),
		}
		def.Fields["field-1"] = st.SchemaKind_INT
		def.Fields["field-2"] = st.SchemaKind_FLOAT

		node, err := schema.BuildNodesFromDefinition(mockWhatIs.Did, &def)
		assert.NoError(t, err)
		assert.NotNil(t, node)

	})
}
