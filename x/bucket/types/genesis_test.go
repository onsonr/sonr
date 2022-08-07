package types_test

import (
	"testing"

	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				WhereIsList: []types.WhereIs{
					{
						Did: "did:sonr:1",
					},
					{
						Did: "did:sonr:2",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated whereIs",
			genState: &types.GenesisState{
				WhereIsList: []types.WhereIs{
					{
						Did: "did:sonr:1",
					},
					{
						Did: "did:sonr:1",
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid whereIs count",
			genState: &types.GenesisState{
				WhereIsList: []types.WhereIs{
					{
						Did: "did:sonr:1",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
