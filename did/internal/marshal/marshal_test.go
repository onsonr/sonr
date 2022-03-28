package marshal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeDocument(t *testing.T) {
	t.Run("pluralize", func(t *testing.T) {
		t.Run("string, plural", func(t *testing.T) {
			actual, _ := NormalizeDocument([]byte(`{"message": "Hello, World"}`), Plural("message"))
			assert.JSONEq(t, `{"message": ["Hello, World"]}`, string(actual))
		})
		t.Run("slice, plural", func(t *testing.T) {
			actual, _ := NormalizeDocument([]byte(`{"message": ["Hello, World"]}`), Plural("message"))
			assert.JSONEq(t, `{"message": ["Hello, World"]}`, string(actual))
		})
	})
	t.Run("alias", func(t *testing.T) {
		actual, _ := NormalizeDocument([]byte(`{"message": "Hello, World"}`), KeyAlias("message", "msg"))
		assert.JSONEq(t, `{"msg": "Hello, World"}`, string(actual))
	})
	t.Run("first alias, then plural", func(t *testing.T) {
		actual, _ := NormalizeDocument([]byte(`{"message": "Hello, World"}`), KeyAlias("message", "msg"), Plural("msg"))
		assert.JSONEq(t, `{"msg": ["Hello, World"]}`, string(actual))
	})
}
