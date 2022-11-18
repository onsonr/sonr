package motor

import (
	"fmt"

	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	v1 "github.com/sonr-io/sonr/third_party/types/motor/api/v1/service/v1"
	rt "github.com/sonr-io/sonr/x/registry/types"
	"github.com/stretchr/testify/assert"
)

type callback struct {
	handler func(v1.LinkingEvent, *motorNodeImpl)
	motor   *motorNodeImpl
}

func (c callback) OnDiscover(data []byte)    {}
func (c callback) OnWalletEvent(data []byte) {}

func (c callback) OnLinking(data []byte) {
	ev := v1.LinkingEvent{}
	if err := ev.Unmarshal(data); err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	c.handler(ev, c.motor)
}

func (suite *MotorTestSuite) Test_OnboardDevice() {
	for ev := range suite.callback.walletEventChannel {
		shouldBreak := false
		switch ev.Type {
		case common.WALLET_EVENT_TYPE_VAULT_CREATE_END:
			shouldBreak = true
		case common.WALLET_EVENT_TYPE_VAULT_CREATE_ERROR:
			fmt.Printf("wallet error: %s", ev.ErrorMessage)
			suite.T().FailNow()
		default:
			continue
		}

		if shouldBreak {
			break
		}
	}

	// setup callback for onboarding device
	errChan := make(chan error)
	res := make(chan *rt.WhoIs)

	handler := func(ev v1.LinkingEvent, m *motorNodeImpl) {
		if ev.Type != v1.LinkingEventType_LINKING_EVENT_TYPE_LINKING_COMPLETE {
			errChan <- fmt.Errorf("link failed: %s", ev.Type)
			return
		}

		resp, err := m.OnboardDevice(mt.OnboardDeviceRequest{
			AccountId: ev.AuthInfo.Address,
			Password:  "password123",
			AesPskKey: ev.AuthInfo.AesPskKey,
		})
		assert.NoError(suite.T(), err, "onboard device")
		if err != nil {
			errChan <- err
			return
		}

		res <- resp.WhoIs
	}

	// setup the new motor
	newMotor := new(motorNodeImpl)
	_newMotor, err := EmptyMotor(&mt.InitializeRequest{
		DeviceId:   "new_device",
		ClientMode: mt.ClientMode_ENDPOINT_LOCAL,
	}, callback{
		motor:   newMotor,
		handler: handler,
	})
	*newMotor = *_newMotor
	assert.NoError(suite.T(), err, "empty motor")

	// err = mtu.SetupTestAddressWithKeys(newMotor)
	// assert.NoError(suite.T(), err, "setup new motor"))

	// perform device linking
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

	// await onboarding result
	select {
	case err := <-errChan:
		fmt.Printf("errChan: %s", err)
	case whoIs := <-res:
		assert.Equal(suite.T(), whoIs.Owner, motor.Address, "onboard succeeds")
	}
}
