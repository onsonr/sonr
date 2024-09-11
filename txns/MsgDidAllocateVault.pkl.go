// Code generated from Pkl module `transactions`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

type MsgDidAllocateVault interface {
	Msg

	GetAuthority() string

	GetSubject() string

	GetToken() *pkl.Object
}

var _ MsgDidAllocateVault = (*MsgDidAllocateVaultImpl)(nil)

type MsgDidAllocateVaultImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	Authority string `pkl:"authority"`

	Subject string `pkl:"subject"`

	Token *pkl.Object `pkl:"token"`
}

// The type URL for the message
func (rcv *MsgDidAllocateVaultImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgDidAllocateVaultImpl) GetAuthority() string {
	return rcv.Authority
}

func (rcv *MsgDidAllocateVaultImpl) GetSubject() string {
	return rcv.Subject
}

func (rcv *MsgDidAllocateVaultImpl) GetToken() *pkl.Object {
	return rcv.Token
}
