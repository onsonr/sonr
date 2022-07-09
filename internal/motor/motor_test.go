package motor

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/stretchr/testify/assert"
	prt "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

func Test_CreateAccount(t *testing.T) {
	aesKey, err := crypto.NewAesKey()
	assert.NoError(t, err, "generates aes key")

	req, err := json.Marshal(prt.CreateAccountRequest{
		Password:  "password123",
		AesDscKey: aesKey,
	})
	assert.NoError(t, err, "create account request marshals")
	m, _, err := New()
	assert.NoError(t, err, "creates motor node")

	_, err = m.CreateAccount(req)
	assert.NoError(t, err, "wallet generation succeeds")

	b := m.Balance()
	log.Println("balance:", b)

	// Print the address of the wallet
	log.Println("address:", m.Address)
}
