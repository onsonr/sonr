package random_test

import (
	"testing"

	"github.com/onsonr/hway/crypto/subtle/random"
)

func TestGetRandomBytes(t *testing.T) {
	for i := 0; i <= 32; i++ {
		buf := random.GetRandomBytes(uint32(i))
		if len(buf) != i {
			t.Errorf("length of the output doesn't match the input")
		}
	}
}
