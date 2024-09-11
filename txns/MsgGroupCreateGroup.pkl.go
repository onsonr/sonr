// Code generated from Pkl module `transactions`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

type MsgGroupCreateGroup interface {
	Msg

	GetAdmin() string

	GetMembers() []*pkl.Object

	GetMetadata() string
}

var _ MsgGroupCreateGroup = (*MsgGroupCreateGroupImpl)(nil)

// Group module messages
type MsgGroupCreateGroupImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	Admin string `pkl:"admin"`

	Members []*pkl.Object `pkl:"members"`

	Metadata string `pkl:"metadata"`
}

// The type URL for the message
func (rcv *MsgGroupCreateGroupImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgGroupCreateGroupImpl) GetAdmin() string {
	return rcv.Admin
}

func (rcv *MsgGroupCreateGroupImpl) GetMembers() []*pkl.Object {
	return rcv.Members
}

func (rcv *MsgGroupCreateGroupImpl) GetMetadata() string {
	return rcv.Metadata
}
