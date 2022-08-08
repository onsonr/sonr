package ipns_test

import (
	"fmt"
	"testing"

	"github.com/sonr-io/sonr/pkg/ipns"
	"github.com/stretchr/testify/assert"
)

func Test_IPNS(t *testing.T) {
	t.Run("Should create ipns record", func(t *testing.T) {
		rec, err := ipns.New()
		assert.NoError(t, err)
		rec.Builder.WithCid("QmZWD55U2SDp9uQ5m8hS77EdavpnatTcBMDAkEEKnPWGbn")
		err = rec.CreateRecord()
		assert.NoError(t, err)
		srv := rec.Builder.BuildService()
		assert.NotNil(t, srv)
		fmt.Print(srv.ID)
	})
}
