// Code generated from Pkl module `transactions`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

type MsgDidRegisterController interface {
	Msg

	GetAuthority() string

	GetCid() string

	GetOrigin() string

	GetAuthentication() []*pkl.Object

	GetToken() *pkl.Object
}

var _ MsgDidRegisterController = (*MsgDidRegisterControllerImpl)(nil)

type MsgDidRegisterControllerImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	Authority string `pkl:"authority"`

	Cid string `pkl:"cid"`

	Origin string `pkl:"origin"`

	Authentication []*pkl.Object `pkl:"authentication"`

	Token *pkl.Object `pkl:"token"`
}

// The type URL for the message
func (rcv *MsgDidRegisterControllerImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgDidRegisterControllerImpl) GetAuthority() string {
	return rcv.Authority
}

func (rcv *MsgDidRegisterControllerImpl) GetCid() string {
	return rcv.Cid
}

func (rcv *MsgDidRegisterControllerImpl) GetOrigin() string {
	return rcv.Origin
}

func (rcv *MsgDidRegisterControllerImpl) GetAuthentication() []*pkl.Object {
	return rcv.Authentication
}

func (rcv *MsgDidRegisterControllerImpl) GetToken() *pkl.Object {
	return rcv.Token
}
