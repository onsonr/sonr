package motor

import (
	"context"
	"fmt"
	"testing"

	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
	prt "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

func Test_CreateSchema(t *testing.T) {
	pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
	fmt.Printf("psk: %x\n", pskKey)
	if pskKey == nil || len(pskKey) != 32 {
		t.Errorf("could not load psk key")
		return
	}

	req := prt.LoginRequest{
		Did:       ADDR,
		Password:  "password123",
		AesPskKey: pskKey,
	}

	m := EmptyMotor("test_device")
	_, err := m.Login(req)
	assert.NoError(t, err, "login succeeds")

	// LOGIN DONE, TRY TO CREATE SCHEMA
	createSchemaRequest := prt.CreateSchemaRequest{
		Label: "TestUser",
		Fields: map[string]prt.CreateSchemaRequest_SchemaKind{
			"email":     prt.CreateSchemaRequest_SCHEMA_KIND_STRING,
			"firstName": prt.CreateSchemaRequest_SCHEMA_KIND_STRING,
			"age":       prt.CreateSchemaRequest_SCHEMA_KIND_INT,
		},
	}
	resp, err := m.CreateSchema(createSchemaRequest)
	assert.NoError(t, err, "schema created successfully")

	whatIs := &st.WhatIs{}
	err = whatIs.Unmarshal(resp.WhatIs)
	assert.NoError(t, err, "unmarshal WhatIs")
	fmt.Printf("success: %s\n", whatIs)
}

func Test_QuerySchema(t *testing.T) {
	pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
	fmt.Printf("psk: %x\n", pskKey)
	if pskKey == nil || len(pskKey) != 32 {
		t.Errorf("could not load psk key")
		return
	}

	req := prt.LoginRequest{
		Did:       ADDR,
		Password:  "password123",
		AesPskKey: pskKey,
	}

	m := EmptyMotor("test_device")
	_, err := m.Login(req)
	assert.NoError(t, err, "login succeeds")

	// LOGIN DONE, TRY TO QUERY SCHEMA
	createSchemaRequest := prt.CreateSchemaRequest{
		Label: "TestUser",
		Fields: map[string]prt.CreateSchemaRequest_SchemaKind{
			"email":     prt.CreateSchemaRequest_SCHEMA_KIND_STRING,
			"firstName": prt.CreateSchemaRequest_SCHEMA_KIND_STRING,
			"age":       prt.CreateSchemaRequest_SCHEMA_KIND_INT,
		},
	}
	resp, err := m.CreateSchema(createSchemaRequest)
	assert.NoError(t, err, "schema created successfully")

	whatIs := &st.WhatIs{}
	err = whatIs.Unmarshal(resp.WhatIs)
	assert.NoError(t, err, "unmarshal WhatIs")

	// CREATE DONE, TRY QUERY
	queryWhatIsRequest := prt.QueryWhatIsRequest{
		Creator: whatIs.Creator,
		Did:     whatIs.Did,
	}

	qresp, err := m.QueryWhatIs(context.Background(), queryWhatIsRequest)
	assert.NoError(t, err, "query response succeeds")

	qwhatIs := &st.WhatIs{}
	err = qwhatIs.Unmarshal(qresp.WhatIs)
	assert.NoError(t, err, "unmarshal WhatIs")
	assert.Equal(t, whatIs.Did, qwhatIs.Did)
}
