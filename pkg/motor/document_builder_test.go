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
		createSchemaRequest := mt.CreateSchemaRequest{
			Label: "TestUser",
			Fields: map[string]*st.SchemaFieldKind{
				"email": {
					Kind: st.Kind_STRING,
				},
				"firstName": {
					Kind: st.Kind_STRING,
				},
				"age": {
					Kind: st.Kind_INT,
				},
			},
		}

		resp, err := suite.motorWithKeys.CreateSchema(createSchemaRequest)
		assert.NoError(suite.T(), err, "schema created successfully")

		// query WhatIs so it's cached
		_, err = suite.motorWithKeys.QueryWhatIsByDid(resp.WhatIs.Did)
		assert.NoError(t, err, "query whatis")

		// upload object
		builder, err := suite.motorWithKeys.NewDocumentBuilder(resp.WhatIs.Did)
		assert.NoError(t, err, "object builder created successfully")

		builder.SetLabel("Player 1")
		err = builder.Set("email", "player1@sonr.io")
		assert.NoError(t, err, "set email property")
		err = builder.Set("firstName", "Brayden")
		assert.NoError(t, err, "set firstName property")
		err = builder.Set("age", 24)
		assert.NoError(t, err, "set age property")

		_, err = builder.Build()
		assert.NoError(t, err, "builds successfully")

		result, err := builder.Upload()
		assert.NoError(t, err, "upload succeeds")

		assert.Equal(t, "Player 1", result.Document.Label)
	})
}
