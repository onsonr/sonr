// Code generated from Pkl module `transactions`. DO NOT EDIT.
package transactions

import "github.com/apple/pkl-go/pkl"

type MsgDidProveWitness interface {
	Msg

	GetAuthority() string

	GetProperty() string

	GetWitness() []int

	GetToken() *pkl.Object
}

var _ MsgDidProveWitness = (*MsgDidProveWitnessImpl)(nil)

type MsgDidProveWitnessImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	Authority string `pkl:"authority"`

	Property string `pkl:"property"`

	Witness []int `pkl:"witness"`

	Token *pkl.Object `pkl:"token"`
}

// The type URL for the message
func (rcv *MsgDidProveWitnessImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgDidProveWitnessImpl) GetAuthority() string {
	return rcv.Authority
}

func (rcv *MsgDidProveWitnessImpl) GetProperty() string {
	return rcv.Property
}

func (rcv *MsgDidProveWitnessImpl) GetWitness() []int {
	return rcv.Witness
}

func (rcv *MsgDidProveWitnessImpl) GetToken() *pkl.Object {
	return rcv.Token
}
