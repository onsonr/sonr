package types

import (
	"testing"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDIDState_ValidateBasic(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		didState MsgCreateWhoIs
		valid    bool
	}{
		{
			desc:     "Valid creator did in message body",
			didState: *NewMsgCreateWhoIs("8yqegpg4qvare73hlpssyuaw7jgc0s4azag", nil, nil, WhoIsType_USER),
			valid:    true,
		},
		{
			desc:     "Valid creator did in message body",
			didState: *NewMsgCreateWhoIs("8yqegpg4qvare73h###lpssyuaw7jgc0s4azag", nil, nil, WhoIsType_USER),
			valid:    true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			tc.didState.GetCreatorDid()
			msgDid, err := did.ParseDID(tc.didState.GetCreatorDid())
			require.NoError(t, err)
			assert.NotNil(t, msgDid.DID.ID)
			assert.NotNil(t, msgDid.DID.Method)
		})
	}
}
