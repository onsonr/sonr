package common

import "log"

type MotorCallback interface {
	OnDiscover(data []byte)
	OnSuccess(data interface{})
	OnError(err error)
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

func (cb *defaultCallback) OnEvent(event string, data interface{}) {
	log.Println("ERROR: MotorCallback not implemented.")
}

func (cb *defaultCallback) OnSuccess(data interface{}) {
	log.Println("ERROR: MotorCallback not implemented.")
}

func (cb *defaultCallback) OnError(err error) {
	log.Println("ERROR: MotorCallback not implemented.")
}
