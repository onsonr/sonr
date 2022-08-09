package motor

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	prt "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

/*
prt.CreateSchemaRequest{
	Label: "TestUser",
	Fields: map[string]prt.CreateSchemaRequest_SchemaKind{
		"email":     prt.CreateSchemaRequest_SCHEMA_KIND_STRING,
		"firstName": prt.CreateSchemaRequest_SCHEMA_KIND_STRING,
		"age":       prt.CreateSchemaRequest_SCHEMA_KIND_INT,
	},
}
*/
const SCHEMA_DID string = "did:snr:8a7c357c-f0c1-4f77-b8e3-f1f374d19951"

func Test_ObjectBuilder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
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

		// query WhatIs so it's cached
		_, err = m.QueryWhatIs(context.Background(), prt.QueryWhatIsRequest{
			Creator: m.GetDID().String(),
			Did:     SCHEMA_DID,
		})
		assert.NoError(t, err, "query whatis")

		// upload object
		builder, err := m.NewObjectBuilder(SCHEMA_DID)
		assert.NoError(t, err, "object builder created successfully")

		builder.SetLabel("Player 1")
		err = builder.Set("email", "player1@sonr.io")
		assert.NoError(t, err, "set email property")
		err = builder.Set("firstName", "Brayden")
		assert.NoError(t, err, "set firstName property")
		err = builder.Set("age", 24)
		assert.NoError(t, err, "set age property")

		toUpload, err := builder.Build()
		assert.NoError(t, err, "builds successfully")

		fmt.Println("toUpload")
		fmt.Println(toUpload)
	})
}
