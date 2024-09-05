// Code generated from Pkl module `transactions`. DO NOT EDIT.
package transactions

import "github.com/apple/pkl-go/pkl"

type MsgDidAuthorize interface {
	Msg

	GetAuthority() string

	GetController() string

	GetAddress() string

	GetOrigin() string

	GetToken() *pkl.Object
}

var _ MsgDidAuthorize = (*MsgDidAuthorizeImpl)(nil)

type MsgDidAuthorizeImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	Authority string `pkl:"authority"`

	Controller string `pkl:"controller"`

	Address string `pkl:"address"`

	Origin string `pkl:"origin"`

	Token *pkl.Object `pkl:"token"`
}

// The type URL for the message
func (rcv *MsgDidAuthorizeImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgDidAuthorizeImpl) GetAuthority() string {
	return rcv.Authority
}

func (rcv *MsgDidAuthorizeImpl) GetController() string {
	return rcv.Controller
}

func (rcv *MsgDidAuthorizeImpl) GetAddress() string {
	return rcv.Address
}

func (rcv *MsgDidAuthorizeImpl) GetOrigin() string {
	return rcv.Origin
}

func (rcv *MsgDidAuthorizeImpl) GetToken() *pkl.Object {
	return rcv.Token
}
