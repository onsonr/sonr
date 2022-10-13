package motor

import (
	"testing"

	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
)

func (suite *MotorTestSuite) Test_DocumentBuilder() {
	suite.T().Run("success", func(t *testing.T) {

		// create schema
		createAssociateRequest := mt.CreateSchemaRequest{
			Label: "Associate",
			Fields: map[string]*st.SchemaFieldKind{
				"name": {
					Kind: st.Kind_STRING,
				},
			},
		}

		resp, err := suite.motorWithKeys.CreateSchema(createAssociateRequest)
		assert.NoError(suite.T(), err, "associate schema created successfully")

		createInmateRequest := mt.CreateSchemaRequest{
			Label: "Inmate",
			Fields: map[string]*st.SchemaFieldKind{
				"name": {
					Kind: st.Kind_STRING,
				},
				"age": {
					Kind: st.Kind_INT,
				},
				"known_aliases": {
					Kind: st.Kind_LIST,
					ListKind: &st.SchemaFieldKind{
						Kind: st.Kind_STRING,
					},
				},
				"height": {
					Kind: st.Kind_FLOAT,
				},
				"mug_shot": {
					Kind: st.Kind_BYTES,
				},
				"associates": {
					Kind: st.Kind_LIST,
					ListKind: &st.SchemaFieldKind{
						Kind:    st.Kind_LINK,
						LinkDid: resp.WhatIs.Did,
					},
				},
				"at_large": {
					Kind: st.Kind_BOOL,
				},
			},
		}

		resp, err = suite.motorWithKeys.CreateSchema(createInmateRequest)
		assert.NoError(suite.T(), err, "schema created successfully")

		// query WhatIs so it's cached
		_, err = suite.motorWithKeys.QueryWhatIsByDid(resp.WhatIs.Did)
		assert.NoError(t, err, "query whatis")

		// upload object
		builder, err := suite.motorWithKeys.NewDocumentBuilder(resp.WhatIs.Did)
		assert.NoError(t, err, "object builder created successfully")

		builder.SetLabel("Billy the kid")
		err = builder.Set("name", "Billy")
		assert.NoError(t, err, "set name property")
		err = builder.Set("age", 24)
		assert.NoError(t, err, "set age property")
		err = builder.Set("known_aliases", []string{"Brayden", "Bob", "The kid"})
		assert.NoError(t, err, "set known_alias property")
		err = builder.Set("height", 180.5)
		assert.NoError(t, err, "set height property")
		err = builder.Set("mug_shot", []byte{0xef, 0xbe, 0x4e, 0x10, 0xef, 0xbe, 0x4e, 0x10})
		assert.NoError(t, err, "set mug_shot property")
		ints := make([]interface{}, 1)
		ints[0] = map[string]string{
			"name": "Nicky Bobby",
		}
		err = builder.Set("associates", ints)
		assert.NoError(t, err, "set associates property")
		err = builder.Set("at_large", true)
		assert.NoError(t, err, "set at_large property")

		_, err = builder.Build()
		assert.NoError(t, err, "builds successfully")

		result, err := builder.Upload()
		assert.NoError(t, err, "upload succeeds")

		assert.Equal(t, "Player 1", result.Document.Label)
	})
}
