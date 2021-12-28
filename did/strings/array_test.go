package strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComplement(t *testing.T) {
	cases := []struct {
		first    []string
		second   []string
		expected []string
	}{
		{[]string{}, []string{}, []string{}},
		{nil, []string{}, []string{}},
		{nil, nil, []string{}},
		{[]string{"1", "2"}, []string{"1", "2"}, []string{}},
		{[]string{"1", "2", "3"}, []string{}, []string{"1", "2", "3"}},
		{[]string{"1", "2", "3"}, nil, []string{"1", "2", "3"}},
		{[]string{"1", "2", "3"}, []string{"4", "5", "6"}, []string{"1", "2", "3"}},
		{[]string{"1", "2", "3"}, []string{"1", "5", "2"}, []string{"3"}},
		{[]string{"4", "1", "6", "2", "3"}, []string{"1", "5", "2"}, []string{"4", "6", "3"}},
	}

	for _, tc := range cases {
		actual := Complement(tc.first, tc.second)
		require.Equal(t, tc.expected, actual)
	}
}
