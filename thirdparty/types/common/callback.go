package common

import "log"

type MotorCallback interface {
	OnDiscover(data []byte)
	OnWalletCreated(ok bool)
}

type defaultCallback struct {
	MotorCallback
}

func DefaultCallback() MotorCallback {
	return &defaultCallback{}
}

func (cb *defaultCallback) OnDiscover(data []byte) {
	log.Println("ERROR: MotorCallback not implemented.")
}

func (cb *defaultCallback) OnWalletCreated(ok bool) {
	log.Println("ERROR: MotorCallback not implemented.")
}
