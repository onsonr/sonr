package types_test

import (
	"testing"

	"github.com/sonrhq/core/x/vault/types"
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

				ClaimableWalletList: []types.ClaimableWallet{
					{
						Index: 0,
					},
					{
						Index: 1,
					},
				},
				ClaimableWalletCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated claimableWallet",
			genState: &types.GenesisState{
				ClaimableWalletList: []types.ClaimableWallet{
					{
						Index: 0,
					},
					{
						Index: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid claimableWallet count",
			genState: &types.GenesisState{
				ClaimableWalletList: []types.ClaimableWallet{
					{
						Index: 1,
					},
				},
				ClaimableWalletCount: 0,
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
