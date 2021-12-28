package did

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsDid(t *testing.T) {
	cases := []struct {
		valid bool
		did   string
	}{
		{true, "did:sonr:test:wyywywywyw"},
		{true, "did:sonr:test:wyywywywyw:sdadasda"},
		{false, "did1:sonr:test:wyywywywyw:sdadasda"},
		{false, "did:sonr2:test:wyywywywyw:sdadasda"},
		{false, "did:sonr:test4:wyywywywyw:sdadasda"},
		{false, ""},
		{false, "did:sonr"},
		{false, "did:sonr:test"},
		{false, "did:sonr:test:dsdasdad#weqweqwew"},
		{false, "did:sonr:test:sdasdasdasd/qeweqweqwee"},
		{false, "did:sonr:test:sdasdasdasd?=qeweqweqwee"},
		{false, "did:sonr:test:sdasdasdasd&qeweqweqwee"},
	}

	for _, tc := range cases {
		isDid := IsValidDid("did:sonr:test:", tc.did)

		if tc.valid {
			require.True(t, isDid)
		} else {
			require.False(t, isDid)
		}
	}
}
