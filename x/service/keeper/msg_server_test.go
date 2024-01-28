package keeper_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	modulev1 "github.com/sonrhq/sonr/api/sonr/service/module/v1"
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
				Authority:   "",
				Name:        "test record",
				Origin:      "test origin",
				Permissions: int32(modulev1.ServicePermissions_SERVICE_PERMISSIONS_BASE),
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

func TestUpdateRecord(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	// Setup or mock prerequisites here, such as inserting a record to update.

	testCases := []struct {
		name         string
		request      *service.MsgUpdateRecord
		expectErrMsg string
	}{
		{
			name: "update non-existent record",
			request: &service.MsgUpdateRecord{
				RecordId:    0,
				Authority:   f.addrs[0].String(),
				Name:        "updated name",
				Origin:      "updated origin",
				Permissions: int32(modulev1.ServicePermissions_SERVICE_PERMISSIONS_BASE),
			},
			expectErrMsg: "not found",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		_, err := f.msgServer.UpdateRecord(f.ctx, tc.request)
		if tc.expectErrMsg != "" {
			require.Error(err)
			require.Contains(err.Error(), tc.expectErrMsg)
		} else {
			require.NoError(err)
		}
	}
}

func TestDeleteRecord(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	// Setup or mock prerequisites here, such as inserting a record to delete.

	testCases := []struct {
		name         string
		request      *service.MsgDeleteRecord
		expectErrMsg string
	}{
		{
			name: "delete non-existent record",
			request: &service.MsgDeleteRecord{
				RecordId: 0,
			},
			expectErrMsg: "not found",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		_, err := f.msgServer.DeleteRecord(f.ctx, tc.request)
		if tc.expectErrMsg != "" {
			require.Error(err)
			require.Contains(err.Error(), tc.expectErrMsg)
		} else {
			require.NoError(err)
		}
	}
}
