package motor

import (
	"testing"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/sonr-io/sonr/x/registry/types"
	"github.com/stretchr/testify/assert"
)

func Test_CreateAccount(t *testing.T) {
	_, err := CreateAccount(
		[]byte("mystrongpassword"),
		[]byte("somerandomdscpubkey"),
	)
	assert.NoError(t, err, "wallet generation succeeds")
}

func Test_DocumentSerialization(t *testing.T) {
	m, err := setupDefault()
	assert.NoError(t, err, "setup succeeds")

	m.DIDDoc.AddService(did.Service{
		ID:   ssi.MustParseURI("https://vault.sonr.ws"),
		Type: "vault",
		ServiceEndpoint: map[string]string{
			"cid": "asdfoasdfklasdjfashfk",
		},
	})
	marshalled, err := m.DIDDoc.MarshalJSON()
	assert.NoError(t, err, "marshals")

	des := did.BlankDocument()
	err = des.UnmarshalJSON(marshalled)
	assert.NoError(t, err, "unmarshals")

	regDes, err := types.NewDIDDocumentFromPkg(des)
	assert.NoError(t, err, "converts")

	assert.Equal(t, len(m.DIDDoc.GetServices()), len(regDes.Service))
}
