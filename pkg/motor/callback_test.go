package motor

import (
	"testing"

	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/stretchr/testify/assert"
)

type testCallback struct {
	WalletEventExec bool
}

func (tc *testCallback) OnDiscover(data []byte) {}

func (tc *testCallback) OnWalletEvent(data []byte) {
	tc.WalletEventExec = true
}

func (tc *testCallback) ClearState() {
	tc.WalletEventExec = false
}

func (tc *testCallback) GetWalletEventState() bool {
	return tc.WalletEventExec
}

func Test_Callbacks(t *testing.T) {
	cb := &testCallback{}
	// setup motor
	mtr, err := EmptyMotor(&mt.InitializeRequest{
		DeviceId:   "test_device",
		ClientMode: mt.ClientMode_ENDPOINT_BETA,
	}, cb)

	assert.NoError(t, err)

	mtr.triggerWalletEvent(common.WalletEvent{
		Type: common.WALLET_EVENT_TYPE_DID_DOCUMENT_CREATE_START,
	})
	assert.Equal(t, cb.GetWalletEventState(), true)
}

func Test_FailedInitializeRequest(t *testing.T) {
	cb := &testCallback{}
	// setup empty motor
	mtr, err := EmptyMotor(&mt.InitializeRequest{}, cb)
	// assert there is an error
	assert.Error(t, err)
	// assert the callback was not executed
	assert.Equal(t, cb.GetWalletEventState(), false)
	assert.Nil(t, mtr)
}

func Test_EmptyWalletEvent(t *testing.T) {
	cb := &testCallback{}
	// setup empty motor
	mtr, err := EmptyMotor(&mt.InitializeRequest{
		DeviceId:   "test_device",
		ClientMode: mt.ClientMode_ENDPOINT_BETA,
	}, cb)

	assert.NoError(t, err)

	mtr.triggerWalletEvent(common.WalletEvent{})
	assert.Equal(t, cb.GetWalletEventState(), true)
}
