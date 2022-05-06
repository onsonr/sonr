package did

import (
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"
)

var (
	test_id = "WIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0I"
	logger  = golog.Child("test/builder")
)

func Test_Build(t *testing.T) {

	logger.Info("Testing dummy Fragment builder")
	didUrl, err := NewDID(test_id, WithNetwork("testnet"), WithPathSegments("test"))
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "did:sonr:testnet:WIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0I/test", didUrl.String())
	logger.Info(didUrl.String())

	logger.Info("Testing Channels Fragment builder")
	didUrl2, err := NewDID(test_id, WithPathSegments("beam", "channels"))
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "did:sonr:WIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0I/beam/channels", didUrl2.String())
	logger.Info(didUrl2.String())

	logger.Info("Testing Objects Query builder")
	didUrl3, err := NewDID(test_id, WithQuery("?Profile"), WithPathSegments("beam", "objects"))
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "did:sonr:WIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0I/beam/objects?Profile", didUrl3.String())
	logger.Info(didUrl3.String())

	logger.Info("Testing Device Key Id Fragment builder")
	didUrl4, err := NewDID(test_id, WithFragment("#iphone_14_5"))
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "did:sonr:WIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0I#iphone_14_5", didUrl4.String())
	logger.Info(didUrl3.String())
}
