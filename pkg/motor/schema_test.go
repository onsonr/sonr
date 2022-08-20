package motor

import (
	"context"
	"fmt"
	"testing"

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

	m := EmptyMotor("test_device")
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

	m := EmptyMotor("test_device")
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
	queryWhatIsRequest := mt.QueryWhatIsRequest{
		Creator: resp.WhatIs.Creator,
		Did:     resp.WhatIs.Did,
	}

	qresp, err := m.QueryWhatIs(context.Background(), queryWhatIsRequest)
	assert.NoError(t, err, "query response succeeds")

	assert.Equal(t, resp.WhatIs.Did, qresp.WhatIs.Did)
}
