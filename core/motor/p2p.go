package motor

import (
	mt "github.com/sonr-hq/sonr/core/motor/types/bind/v1"
)

// A method on the MotorNode struct. It takes a request of type `mt.ConnectRequest` and returns a
// response of type `mt.ConnectResponse` and an error.
func (mi *MotorNode) Connect(req *mt.ConnectRequest) (*mt.ConnectResponse, error) {
	err := mi.Node.Connect(req.GetMultiaddr())
	if err != nil {
		return nil, err
	}
	return &mt.ConnectResponse{
		Success: true,
	}, nil
}
