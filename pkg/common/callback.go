package common

import (
	"log"
)

type MotorCallback interface {
	OnDiscover(data []byte)
	OnLinking(data []byte)
}

type defaultCallback struct{}

func DefaultCallback() MotorCallback {
	return &defaultCallback{}
}

func (cb *defaultCallback) OnDiscover(data []byte) {
	log.Println(ErrDefaultStillImplemented.Error())
}

func (cb *defaultCallback) OnLinking(data []byte) {
	log.Println("ERROR: MotorCallback not implemented.")
}
