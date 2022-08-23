package motor

import (
	"fmt"
	"testing"

	"github.com/sonr-io/sonr/thirdparty/types/common"
	mt "github.com/sonr-io/sonr/thirdparty/types/motor"
	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
)

func Test_CreateSchema(t *testing.T) {
	pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
	fmt.Printf("psk: %x\n", pskKey)
	if pskKey == nil || len(pskKey) != 32 {
		t.Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		Did:       ADDR,
		Password:  "password123",
		AesPskKey: pskKey,
	}

	m := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(t, err, "login succeeds")

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
	assert.NoError(t, err, "schema created successfully")
	fmt.Printf("success: %s\n", resp.WhatIs)
}

func Test_QuerySchema(t *testing.T) {
	pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
	fmt.Printf("psk: %x\n", pskKey)
	if pskKey == nil || len(pskKey) != 32 {
		t.Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		Did:       ADDR,
		Password:  "password123",
		AesPskKey: pskKey,
	}

	m := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(t, err, "login succeeds")

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
	assert.NoError(t, err, "schema created successfully")

	// CREATE DONE, TRY QUERY
	qReq := mt.QueryRequest{
		Query:  resp.WhatIs.Did,
		Module: common.BlockchainModule_SCHEMA,
		Kind:   common.EntityKind_DID,
	}

	qresp, err := m.Query(qReq)
	assert.NoError(t, err, "query response succeeds")
	r := findItem(qresp.Results, resp.WhatIs.Did)
	assert.Equal(t, resp.WhatIs.Did, r)
}

func findItem(arr []*mt.QueryResultItem, target string) string {
	for _, item := range arr {
		if item.Did == target {
			return item.GetDid()
		}
	}
	return ""
}
