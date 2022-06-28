package motor

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	prt "go.buf.build/grpc/go/sonr-io/motor/registry/v1"
)

func Test_CreateAccount(t *testing.T) {
	req, err := json.Marshal(prt.CreateAccountRequest{
		Password:          "password",
		EcdsaDscKey:       []byte("somerandomdscpubkey"),
		EcdsaPresharedKey: []byte("somerandompskpubkey"),
	})
	assert.NoError(t, err, "create account request marshals")

	_, err = CreateAccount(req)
	assert.NoError(t, err, "wallet generation succeeds")
}
