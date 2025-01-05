package crypto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDIDKey(t *testing.T) {
	str := "did:key:z6Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z"
	d, err := Parse(str)
	require.NoError(t, err)
	require.Equal(t, str, d.String())
}

func TestMustParseDIDKey(t *testing.T) {
	str := "did:key:z6Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z"
	require.NotPanics(t, func() {
		d := MustParse(str)
		require.Equal(t, str, d.String())
	})
	str = "did:key:z7Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z"
	require.Panics(t, func() {
		MustParse(str)
	})
}

func TestEquivalence(t *testing.T) {
	undef0 := DID{}
	undef1 := Undef

	did0, err := Parse("did:key:z6Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z")
	require.NoError(t, err)
	did1, err := Parse("did:key:z6Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z")
	require.NoError(t, err)

	require.True(t, undef0 == undef1)
	require.False(t, undef0 == did0)
	require.True(t, did0 == did1)
	require.False(t, undef1 == did1)
}
