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
	srv, ctx := setupMsgServer(t)

	// setup motor
	motor, err := m.EmptyMotor(&mt.InitializeRequest{DeviceId: "test_device"}, common.DefaultCallback())
	require.NoError(t, err)
	err = mtu.SetupTestAddressWithKeys(motor)
	require.NoError(t, err)

	creator := motor.Address
	
	// create bucket
	defineBucketResp, err := srv.DefineBucket(ctx, &types.MsgDefineBucket{Creator: creator, Label: "Test Label"})
	require.NoError(t, err)
		
	// generate bucket
	uuid := uuid.New().String()
	genBucketResp, err := motor.GenerateBucket(mt.GenerateBucketRequest{Uuid: uuid, Name: "Test Bucket", Creator: creator, Bucket: defineBucketResp.GetBucket()})
	require.NoError(t, err)

	_, err = srv.AllocateBucket(ctx, &types.MsgAllocateBucket{Creator: creator, Bucket: defineBucketResp.GetBucket(), Cid: genBucketResp.GetCid()})
	require.NoError(t, err)
}