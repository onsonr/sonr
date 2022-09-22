package motor

import (
	"fmt"
	"log"

	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
)

func (suite *MotorTestSuite) Test_CreateSchema() {
	suite.T().Skip()
	pskKey := loadKey(fmt.Sprintf("psk%s", suite.accountAddress))
	fmt.Printf("psk: %x\n", pskKey)
	if pskKey == nil || len(pskKey) != 32 {
		suite.T().Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		Did:      suite.accountAddress,
		Password: "password123",
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(suite.T(), err, "login succeeds")

	// LOGIN DONE, TRY TO CREATE SCHEMA
	createSchemaRequest := mt.CreateSchemaRequest{
		Label: "TestUser",
		Fields: map[string]st.SchemaKind{
			"email":     st.SchemaKind_STRING,
			"firstName": st.SchemaKind_STRING,
			"age":       st.SchemaKind_INT,
		},
	}
	resp, err := m.CreateSchema(createSchemaRequest)
	assert.NoError(suite.T(), err, "schema created successfully")
	fmt.Printf("success: %s\n", resp.WhatIs)
}

func (suite *MotorTestSuite) Test_QuerySchema() {
	suite.T().Skip()
	pskKey := loadKey(fmt.Sprintf("psk%s", suite.accountAddress))
	fmt.Printf("psk: %x\n", pskKey)
	if pskKey == nil || len(pskKey) != 32 {
		suite.T().Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		Did:      suite.accountAddress,
		Password: "password123",
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(suite.T(), err, "login succeeds")

	// LOGIN DONE, TRY TO QUERY SCHEMA
	createSchemaRequest := mt.CreateSchemaRequest{
		Label: "TestUser",
		Fields: map[string]st.SchemaKind{
			"email":     st.SchemaKind_STRING,
			"firstName": st.SchemaKind_STRING,
			"age":       st.SchemaKind_INT,
		},
	}
	resp, err := m.CreateSchema(createSchemaRequest)
	assert.NoError(suite.T(), err, "schema created successfully")

	// CREATE DONE, TRY QUERY
	qReq := mt.QueryWhatIsRequest{
		Creator: m.Address,
		Did:     resp.WhatIs.Did,
	}

	qresp, err := m.QueryWhatIs(qReq)
	assert.NoError(suite.T(), err, "query response succeeds")
	assert.Equal(suite.T(), resp.WhatIs.Did, qresp.WhatIs.Did)
}

func (suite *MotorTestSuite) Test_QuerySchemaByCreator() {
	suite.T().Skip()
	pskKey := loadKey(fmt.Sprintf("psk%s", suite.accountAddress))
	fmt.Printf("psk: %x\n", pskKey)
	if pskKey == nil || len(pskKey) != 32 {
		suite.T().Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		Did:      suite.accountAddress,
		Password: "password123",
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(suite.T(), err, "login succeeds")

	// CREATE DONE, TRY QUERY
	qReq := mt.QueryWhatIsByCreatorRequest{
		Creator: "did:snr:1r77u6dnsavm0094l2075zaqk2qval68mu0papc",
	}

	qresp, err := m.QueryWhatIsByCreator(qReq)
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
	suite.T().Skip()
	pskKey := loadKey(fmt.Sprintf("psk%s", suite.accountAddress))
	fmt.Printf("psk: %x\n", pskKey)
	if pskKey == nil || len(pskKey) != 32 {
		suite.T().Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		Did:      suite.accountAddress,
		Password: "password123",
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(suite.T(), err, "login succeeds")

	// CREATE DONE, TRY QUERY
	qresp, err := m.QueryWhatIsByDid("did:snr:Qme2eF6tp63kzjz6UDxmc9xkuthJaMBTb1bmB7Km65F5VM")
	assert.NoError(suite.T(), err, "query response succeeds")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(qresp)
}

func findItem(arr []*mt.QueryResultItem, target string) string {
	for _, item := range arr {
		if item.Did == target {
			return item.GetDid()
		}
	}
	return ""
}
