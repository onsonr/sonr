package keeper_test

import (
	"testing"

	"github.com/google/uuid"
	m "github.com/sonr-io/sonr/pkg/motor"
	mtu "github.com/sonr-io/sonr/testutil/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/stretchr/testify/require"
)

func TestAllocateBucketMsgServer(t *testing.T) {
	// setup motor
	motor, err := m.EmptyMotor(&mt.InitializeRequest{DeviceId: "test_device"}, common.DefaultCallback())
	require.NoError(t, err)
	err = mtu.SetupTestAddressWithKeys(motor)
	require.NoError(t, err)

	creator := motor.GetAddress()
		
	// generate bucket
	uuid := uuid.New().String()
	genBucketResp, err := motor.GenerateBucket(mt.GenerateBucketRequest{Uuid: uuid, Name: "Test Bucket", Creator: creator, Bucket: &types.BucketConfig{Uuid: uuid, Creator: creator, Name: "Test Bucket"}})
	require.NoError(t, err)
	require.NotEmpty(t, genBucketResp.DidDocument.Service[0].ServiceEndpoint)
}