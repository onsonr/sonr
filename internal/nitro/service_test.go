package nitro_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maunium.net/go/mautrix/appservice"
)

func TestClient_UnixSocket(t *testing.T) {
	as := appservice.Create()
	as.Host = appservice.HostConfig{
		Hostname: "matrix.sonr.dev",
		Port:     443,
	}

	reg := appservice.CreateRegistration()
	as.Registration = reg
	err := as.SetHomeserverURL("https://matrix.sonr.dev")
	assert.NoError(t, err)

	// Save the appservice
	t.Log(as)
}
