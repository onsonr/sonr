package motor

import (
	"context"
	"fmt"
	"testing"

	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/stretchr/testify/assert"
)

func Test_CreateBucket(t *testing.T) {
	pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
	if pskKey == nil || len(pskKey) != 32 {
		t.Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		Did:      ADDR,
		Password: "password123",
	}
	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(t, err, "login succeeds")

	createReq := mt.CreateBucketRequest{
		Creator:    ADDR,
		Label:      "my awesome bucket",
		Visibility: types.BucketVisibility_PUBLIC,
		Role:       types.BucketRole_USER,
		Content:    make([]*types.BucketItem, 0),
	}
	b, err := m.CreateBucket(context.Background(), createReq)
	assert.NoError(t, err)
	assert.NotNil(t, b)
}

func Test_GetBucket(t *testing.T) {
	pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
	if pskKey == nil || len(pskKey) != 32 {
		t.Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		Did:      ADDR,
		Password: "password123",
	}

	m, _ := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(t, err, "login succeeds")

	b, err := m.GetBucket("did:snr:e7360323120e4b70a5366984c994b536")
	assert.NoError(t, err, "get bucket")

	b.ResolveContent()
	for _, c := range b.GetBucketItems() {
		fmt.Println(c.Name)
	}
}
