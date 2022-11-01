package motor

import (
	mtu "github.com/sonr-io/sonr/testutil/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/stretchr/testify/assert"
)

func (suite *MotorTestSuite) Test_OnboardDevice() {
	newMotor, err := EmptyMotor(&mt.InitializeRequest{
		DeviceId:   "new_device",
		ClientMode: mt.ClientMode_ENDPOINT_BETA,
	}, common.DefaultCallback())
	assert.NoError(suite.T(), err, "empty motor")

	err = mtu.SetupTestAddressWithKeys(newMotor)
	assert.NoError(suite.T(), err, "setup new motor")

	motor, ok := suite.motorWithKeys.(*motorNodeImpl)
	assert.True(suite.T(), ok, "motor is motorWithKeys")

	// connect test suite motor
	connectResp, err := suite.motorWithKeys.Connect(mt.ConnectRequest{
		EnableDiscovery: true,
		EnableTransmit:  true,
		EnableLinking:   true,
	})
	assert.NoError(suite.T(), err, "suite motor connects")
	assert.True(suite.T(), connectResp.Success, "suite connect succeeds")

	// connect newMotor. In practice these steps will create a QR code
	connectResp, err = newMotor.Connect(mt.ConnectRequest{
		EnableDiscovery: true,
		EnableTransmit:  true,
		EnableLinking:   true,
	})
	assert.NoError(suite.T(), err, "new motor connects")
	assert.True(suite.T(), connectResp.Success, "new connect succeeds")

	linkResp, err := newMotor.OpenLinking(mt.LinkingRequest{})
	assert.NoError(suite.T(), err, "motor opens linking")
	assert.True(suite.T(), linkResp.Success, "open linking succeeds")

	// Pair to new device and send PSK
	pairResp, err := motor.PairDevice(mt.PairingRequest{
		P2PAddrs: linkResp.P2PAddrs,
	})
	assert.NoError(suite.T(), err, "motor pairs device")
	assert.True(suite.T(), pairResp.Success, "pair device succeeds")
}
