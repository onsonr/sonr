package keeper_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sonrhq/sonr/x/service"
)

func TestUpdateParams(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	testCases := []struct {
		name         string
		request      *service.MsgUpdateParams
		expectErrMsg string
	}{
		{
			name: "set invalid authority (not an address)",
			request: &service.MsgUpdateParams{
				Authority: "foo",
			},
			expectErrMsg: "invalid authority address",
		},
		{
			name: "set invalid authority (not defined authority)",
			request: &service.MsgUpdateParams{
				Authority: f.addrs[1].String(),
			},
			expectErrMsg: fmt.Sprintf("unauthorized, authority does not match the module's authority: got %s, want %s", f.addrs[1].String(), f.k.GetAuthority()),
		},
		{
			name: "set valid params",
			request: &service.MsgUpdateParams{
				Authority: f.k.GetAuthority(),
				Params:    service.Params{},
			},
			expectErrMsg: "",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(_ *testing.T) {
			_, err := f.msgServer.UpdateParams(f.ctx, tc.request)
			if tc.expectErrMsg != "" {
				require.Error(err)
				require.ErrorContains(err, tc.expectErrMsg)
			} else {
				require.NoError(err)
			}
		})
	}
}

func TestCreateRecord(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	testCases := []struct {
		name         string
		request      *service.MsgCreateRecord
		expectErrMsg string
	}{
		{
			name: "create with empty owner",
			request: &service.MsgCreateRecord{
				Authority: "",
				Name:      "test record",
				Origin:    "test origin",
			},
			expectErrMsg: "owner cannot be empty",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		_, err := f.msgServer.CreateRecord(f.ctx, tc.request)
		if tc.expectErrMsg != "" {
			require.Error(err)
			require.Contains(err.Error(), tc.expectErrMsg)
		} else {
			require.NoError(err)
		}
	}
}
