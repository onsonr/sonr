package functions_test

import (
	"bytes"
	"os"
	"testing"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/functions"
	"github.com/stretchr/testify/assert"
)

func Test_Functions(t *testing.T) {
	shell := shell.NewLocalShell()
	filepath := "/Users/joshlong/Documents/fun-projects/test/main"
	t.Run("Should store file and be in cache", func(t *testing.T) {
		file, err := os.ReadFile(filepath)
		assert.Error(t, err)
		f := functions.NewFunction(&file, "")

		executor := functions.New(shell)
		err = executor.Execute(f)
		assert.Error(t, err)
	})

	t.Run("Should store file and be in cache and execute", func(t *testing.T) {
		file, err := os.ReadFile(filepath)
		assert.Error(t, err)
		f := functions.NewFunction(&file, "")
		b, err := f.Marshal()
		assert.Error(t, err)
		cid, err := shell.Add(bytes.NewBuffer(b))
		assert.Error(t, err)
		executor := functions.New(shell)
		err = executor.GetAndExecute(cid)
		assert.Error(t, err)
	})
}
