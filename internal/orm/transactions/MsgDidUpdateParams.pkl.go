// Code generated from Pkl module `transactions`. DO NOT EDIT.
package transactions

import "github.com/apple/pkl-go/pkl"

type MsgDidUpdateParams interface {
	Msg

	GetAuthority() string

	GetParams() *pkl.Object

	GetToken() *pkl.Object
}

var _ MsgDidUpdateParams = (*MsgDidUpdateParamsImpl)(nil)

type MsgDidUpdateParamsImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	Authority string `pkl:"authority"`

	Params *pkl.Object `pkl:"params"`

	Token *pkl.Object `pkl:"token"`
}

// The type URL for the message
func (rcv *MsgDidUpdateParamsImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgDidUpdateParamsImpl) GetAuthority() string {
	return rcv.Authority
}

func (rcv *MsgDidUpdateParamsImpl) GetParams() *pkl.Object {
	return rcv.Params
}

func (rcv *MsgDidUpdateParamsImpl) GetToken() *pkl.Object {
	return rcv.Token
}
