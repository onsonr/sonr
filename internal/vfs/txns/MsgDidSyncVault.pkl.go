// Code generated from Pkl module `txns`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

type MsgDidSyncVault interface {
	Msg

	GetController() string

	GetToken() *pkl.Object
}

var _ MsgDidSyncVault = (*MsgDidSyncVaultImpl)(nil)

type MsgDidSyncVaultImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	Controller string `pkl:"controller"`

	Token *pkl.Object `pkl:"token"`
}

// The type URL for the message
func (rcv *MsgDidSyncVaultImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgDidSyncVaultImpl) GetController() string {
	return rcv.Controller
}

func (rcv *MsgDidSyncVaultImpl) GetToken() *pkl.Object {
	return rcv.Token
}
