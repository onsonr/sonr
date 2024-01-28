package motor

import (
	_ "golang.org/x/mobile/bind"
	// mtr "github.com/sonr-io/sonr/pkg/motor"
	// "github.com/sonr-io/sonr/pkg/motor/x/document"
	// ct "github.com/sonr-io/sonr/pkg/types/common"
	// rt "github.com/sonr-io/sonr/x/registry/types"
)

type MotorCallback interface {
	OnDiscover(data []byte)
	OnWalletEvent(data []byte)
	OnLinking(data []byte)
}

func Init(buf []byte, cb MotorCallback) ([]byte, error) {
	// Unmarshal the request
	// var req mt.InitializeRequest
	// if err := req.Unmarshal(buf); err != nil {
	// 	return nil, err
	// }

	// // Create Motor instance
	// mtr, err := mtr.EmptyMotor(&req, cb)
	// if err != nil {
	// 	return nil, err
	// }
	// instance = mtr

	// // init docBuilders
	// docBuilders = make(map[string]*document.DocumentBuilder)

	// // Return Initialization Response
	// resp := mt.InitializeResponse{
	// 	Success: true,
	// }

	// if req.AuthInfo != nil {
	// 	if res, err := instance.Login(mt.LoginRequest{
	// 		AccountId: req.AuthInfo.Did,
	// 		Password:  req.AuthInfo.Password,
	// 	}); err == nil {
	// 		return res.Marshal()
	// 	}
	// }
	// return resp.Marshal()
	return nil, nil
}

func Connect(buf []byte) ([]byte, error) {
	// Unmarshal the request
	// var req mt.ConnectRequest
	// if err := req.Unmarshal(buf); err != nil {
	// 	return nil, err
	// }

	// // Connect to the network
	// if err := instance.Connect(req); err != nil {
	// 	return nil, err
	// }

	// // Return Connect Response
	// resp := mt.ConnectResponse{
	// 	Success: true,
	// }
	// return resp.Marshal()
	return nil, nil
}
