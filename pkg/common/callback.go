package common

import (
	"log"
)

type NodeCallback interface {
	OnDiscover(data []byte)
	OnLinking(data []byte)
}

type defaultCallback struct{}

func DefaultCallback() NodeCallback {
	return &defaultCallback{}
}

func (cb *defaultCallback) OnDiscover(data []byte) {
	log.Println(ErrDefaultStillImplemented.Error())
}

func (cb *defaultCallback) OnLinking(data []byte) {
	log.Println("ERROR: MotorCallback not implemented.")
}
