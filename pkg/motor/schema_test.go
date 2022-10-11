package motor

import (
	"fmt"
	"log"

	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
)

func (suite *MotorTestSuite) Test_CreateSchema() {
	fmt.Println("Test_CreateSchema")
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
	fmt.Printf("success: %s\n", resp.WhatIs)
}

func (suite *MotorTestSuite) Test_QuerySchema() {
	fmt.Println("Test_QuerySchema")
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

	// CREATE DONE, TRY QUERY
	qReq := mt.QueryWhatIsRequest{
		Creator: suite.motorWithKeys.Address,
		Did:     resp.WhatIs.Did,
	}

	qresp, err := suite.motorWithKeys.QueryWhatIs(qReq)
	assert.NoError(suite.T(), err, "query response succeeds")
	assert.Equal(suite.T(), resp.WhatIs.Did, qresp.WhatIs.Did)
}

func (suite *MotorTestSuite) Test_QuerySchemaByCreator() {
	fmt.Println("Test_QuerySchemaByCreator")
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

	qReq := mt.QueryWhatIsByCreatorRequest{
		Creator: resp.WhatIs.Creator,
	}

	qresp, err := suite.motorWithKeys.QueryWhatIsByCreator(qReq)
	assert.NoError(suite.T(), err, "query response succeeds")
	if err != nil {
		log.Fatal(err)
	}

	if qresp.Schemas != nil {
		fmt.Println(qresp.Schemas)
	} else {
		fmt.Println("no schemas.")
	}
}

func (suite *MotorTestSuite) Test_QuerySchemaByDid() {
	fmt.Println("Test_QuerySchemaByDid")
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

	// CREATE DONE, TRY QUERY
	qresp, err := suite.motorWithKeys.QueryWhatIsByDid(resp.WhatIs.Did)
	assert.NoError(suite.T(), err, "query response succeeds")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(qresp)
}
