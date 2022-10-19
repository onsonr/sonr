package common

import (
	"log"
)

type MotorCallback interface {
	OnDiscover(data []byte)
	OnWalletEvent(data []byte)
}

type defaultCallback struct{}

func DefaultCallback() MotorCallback {
	return &defaultCallback{}
}

func (cb *defaultCallback) OnDiscover(data []byte) {
	log.Println(ErrDefaultStillImplemented.Error())
}

func (cb *defaultCallback) OnWalletEvent(data []byte) {
	event := WalletEvent{}
	err := event.Unmarshal(data)

	if err != nil {
		log.Printf("error while unmarshalling event \n%s\n", err.Error())
	}
	log.Printf("type: %s\nErrored: %t\nError Message: %s\nMessage: %s\nCompleted: %t\n", event.Type, event.HasErrored, event.ErrorMessage, event.Message, event.Completed)
}
