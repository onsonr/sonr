package motor

import (
	"context"
	"fmt"
	"sync"
	"testing"

	st "github.com/sonr-io/sonr/x/schema/types"
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
const SCHEMA_DID string = "did:snr:64ab70e9-0d4e-466a-aad4-55520b93b617"

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
		// this needs to be awaited here for race conditions
		// not always necessary since you may not use the schema right away
		var wg sync.WaitGroup
		wg.Add(1)
		_, err = m.QueryWhatIsWithSchemaCallback(context.Background(), prt.QueryWhatIsRequest{
			Creator: m.GetDID().String(),
			Did:     SCHEMA_DID,
		}, func(_ *st.SchemaDefinition, err error) {
			assert.NoError(t, err, "stored schema")
			wg.Done()
		})
		assert.NoError(t, err, "query whatis")
		wg.Wait()

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
