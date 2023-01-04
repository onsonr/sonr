package common

import (
	"log"
)

// NodeCallback is an interface with three methods: OnDiscover, OnLinking, and OnTopicMessage.
// @property OnDiscover - This is called when a node is discovered. The data is the data that was sent
// by the node.
// @property OnLinking - This is called when a node is linking to the gateway.
// @property OnTopicMessage - This is the callback that will be called when a message is received on a
// topic.
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

func (cb *defaultCallback) OnTopicMessage(topic string, data []byte) {
	log.Println("ERROR: MotorCallback not implemented.")
}
