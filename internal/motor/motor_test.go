package motor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateAccount(t *testing.T) {
	_, err := CreateAccount()
	assert.NoError(t, err, "wallet generation succeeds")

}
