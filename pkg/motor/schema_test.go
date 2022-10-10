package motor

import (
	"fmt"
	"log"
	"testing"

	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
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

	req := mt.LoginWithKeysRequest{
		AccountId: ADDR,
		Password:  "password123",
		AesPskKey: pskKey,
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.LoginWithKeys(req)
	assert.NoError(t, err, "login succeeds")

	// LOGIN DONE, TRY TO CREATE SCHEMA
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

	req := mt.LoginWithKeysRequest{
		AccountId: ADDR,
		Password:  "password123",
		AesPskKey: pskKey,
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.LoginWithKeys(req)
	assert.NoError(t, err, "login succeeds")

	// LOGIN DONE, TRY TO QUERY SCHEMA
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
	resp, err := m.CreateSchema(createSchemaRequest)
	assert.NoError(t, err, "schema created successfully")

	// CREATE DONE, TRY QUERY
	qReq := mt.QueryWhatIsRequest{
		Creator: m.Address,
		Did:     resp.WhatIs.Did,
	}

	qresp, err := m.QueryWhatIs(qReq)
	assert.NoError(t, err, "query response succeeds")
	assert.Equal(t, resp.WhatIs.Did, qresp.WhatIs.Did)
}

func Test_QuerySchemaByCreator(t *testing.T) {
	pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
	fmt.Printf("psk: %x\n", pskKey)
	if pskKey == nil || len(pskKey) != 32 {
		t.Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		AccountId: ADDR,
		Password:  "password123",
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(t, err, "login succeeds")

	// CREATE DONE, TRY QUERY
	qReq := mt.QueryWhatIsByCreatorRequest{
		Creator: "did:snr:1r77u6dnsavm0094l2075zaqk2qval68mu0papc",
	}

	qresp, err := m.QueryWhatIsByCreator(qReq)
	assert.NoError(t, err, "query response succeeds")
	if err != nil {
		log.Fatal(err)
	}

	if qresp.Schemas != nil {
		fmt.Println(qresp.Schemas)
	} else {
		fmt.Println("no schemas.")
	}
}

func Test_QuerySchemaByDid(t *testing.T) {
	pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
	fmt.Printf("psk: %x\n", pskKey)
	if pskKey == nil || len(pskKey) != 32 {
		t.Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		AccountId: ADDR,
		Password:  "password123",
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(t, err, "login succeeds")

	// CREATE DONE, TRY QUERY
	qresp, err := m.QueryWhatIsByDid("did:snr:Qme2eF6tp63kzjz6UDxmc9xkuthJaMBTb1bmB7Km65F5VM")
	assert.NoError(t, err, "query response succeeds")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(qresp)
}
