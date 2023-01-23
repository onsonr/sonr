package types_test

import (
	"testing"

	"github.com/sonr-hq/sonr/x/identity/types"
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
				DidDocumentList: []types.DidDocument{
					{
						ID: "0",
					},
					{
						ID: "1",
					},
				},
				DomainRecordList: []types.DomainRecord{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated didDocument",
			genState: &types.GenesisState{
				DidDocumentList: []types.DidDocument{
					{
						ID: "0",
					},
					{
						ID: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated DomainRecord",
			genState: &types.GenesisState{
				DomainRecordList: []types.DomainRecord{
					{
						Index: "0",
					},
					{
						Index: "0",
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
