package motor

import (
	"fmt"
	"testing"

	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/stretchr/testify/assert"
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
const SCHEMA_DID string = "did:snr:QmZLKGrTcUAKsUVUZ5e72rAWRg1Y1SzRJqWqcXaDqjFUqm"

func Test_ObjectBuilder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req := mt.LoginRequest{
			Did:      ADDR,
			Password: "password123",
		}

		m, _ := EmptyMotor(&mt.InitializeRequest{
			DeviceId: "test_device",
		}, common.DefaultCallback())
		_, err := m.Login(req)
		assert.NoError(t, err, "login succeeds")

		// query WhatIs so it's cached
		_, err = m.GetClient().QueryWhatIs(m.GetDID().String(), SCHEMA_DID)
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

		_, err = builder.Build()
		assert.NoError(t, err, "builds successfully")

		result, err := builder.Upload()
		assert.NoError(t, err, "upload succeeds")

		assert.Equal(t, "Player 1", result.Reference.Label)
	})
}
