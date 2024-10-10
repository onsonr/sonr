// Code generated from Pkl module `transactions`. DO NOT EDIT.
package transactions

import "github.com/apple/pkl-go/pkl"

type MsgDidRegisterService interface {
	Msg

	GetController() string

	GetOriginUri() string

	GetScopes() *pkl.Object

	GetDescription() string

	GetServiceEndpoints() map[string]string

	GetMetadata() *pkl.Object

	GetToken() *pkl.Object
}

var _ MsgDidRegisterService = (*MsgDidRegisterServiceImpl)(nil)

type MsgDidRegisterServiceImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	Controller string `pkl:"controller"`

	OriginUri string `pkl:"originUri"`

	Scopes *pkl.Object `pkl:"scopes"`

	Description string `pkl:"description"`

	ServiceEndpoints map[string]string `pkl:"serviceEndpoints"`

	Metadata *pkl.Object `pkl:"metadata"`

	Token *pkl.Object `pkl:"token"`
}

// The type URL for the message
func (rcv *MsgDidRegisterServiceImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgDidRegisterServiceImpl) GetController() string {
	return rcv.Controller
}

func (rcv *MsgDidRegisterServiceImpl) GetOriginUri() string {
	return rcv.OriginUri
}

func (rcv *MsgDidRegisterServiceImpl) GetScopes() *pkl.Object {
	return rcv.Scopes
}

func (rcv *MsgDidRegisterServiceImpl) GetDescription() string {
	return rcv.Description
}

func (rcv *MsgDidRegisterServiceImpl) GetServiceEndpoints() map[string]string {
	return rcv.ServiceEndpoints
}

func (rcv *MsgDidRegisterServiceImpl) GetMetadata() *pkl.Object {
	return rcv.Metadata
}

func (rcv *MsgDidRegisterServiceImpl) GetToken() *pkl.Object {
	return rcv.Token
}
