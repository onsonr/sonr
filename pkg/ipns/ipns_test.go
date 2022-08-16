package ipns_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/ipns"
	"github.com/stretchr/testify/assert"
)

func Test_IPNS(t *testing.T) {
	shell := shell.NewLocalShell()
	t.Run("Should create ipns record", func(t *testing.T) {
		time_stamp := fmt.Sprintf("%d", time.Now().Unix())

		out_path := filepath.Join(os.TempDir(), time_stamp+".txt")

		defer os.Remove(out_path)

		rec, err := ipns.New()
		assert.NoError(t, err)
		rec.Builder.SetCid("bafyreihnj3feeesb6wmd46lmsvtwalvuckns647ghy44xn63lfsfed3ydm")
		err = rec.CreateRecord()
		assert.NoError(t, err)
		srv := rec.Builder.BuildService()
		assert.NotNil(t, srv)
		fmt.Print(srv.ID)
		id, err := ipns.Publish(shell, rec)
		assert.NoError(t, err)
		str, err := ipns.Resolve(shell, id)
		assert.NoError(t, err)
		assert.NotNil(t, str)
		err = shell.Get(str, out_path)
		assert.NoError(t, err)
		buf, err := os.ReadFile(out_path)
		assert.NoError(t, err)
		fmt.Print(string(buf))
	})
}
