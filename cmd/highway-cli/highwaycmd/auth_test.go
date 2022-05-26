package highwaycmd

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/duo-labs/webauthn/protocol"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/stretchr/testify/assert"
)

func Test_saveCredential(t *testing.T) {
	expected, err := webauthn.MakeNewCredential(&protocol.ParsedCredentialCreationData{})
	assert.NoError(t, err, "Webauthn initializes")

	path := "/tmp/sonr/highway-cli/login.json"
	err = saveCredential(expected, path)
	assert.NoError(t, err, "Credential saves successfully")

	// verify that the file was saved
	assert.FileExists(t, path, "Credential file exists")

	// verify that the file reads
	credBuf, err := os.ReadFile(path)
	assert.NoError(t, err, "Credential file reads successfully")

	var actual webauthn.Credential
	err = json.Unmarshal(credBuf, &actual)
	assert.NoError(t, err, "Credential file unmarshals successfully")

	assert.Equal(t, *expected, actual, "Read credential should match saved credential")

	// cleanup
	assert.NoError(t, os.Remove(path), "Cleanup succeeds")
}
