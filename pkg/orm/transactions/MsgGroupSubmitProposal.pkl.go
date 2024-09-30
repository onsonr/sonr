// Code generated from Pkl module `transactions`. DO NOT EDIT.
package transactions

import "github.com/apple/pkl-go/pkl"

type MsgGroupSubmitProposal interface {
	Msg

	GetGroupPolicyAddress() string

	GetProposers() []string

	GetMetadata() string

	GetMessages() []*pkl.Object

	GetExec() int
}

var _ MsgGroupSubmitProposal = (*MsgGroupSubmitProposalImpl)(nil)

type MsgGroupSubmitProposalImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	GroupPolicyAddress string `pkl:"groupPolicyAddress"`

	Proposers []string `pkl:"proposers"`

	Metadata string `pkl:"metadata"`

	Messages []*pkl.Object `pkl:"messages"`

	Exec int `pkl:"exec"`
}

// The type URL for the message
func (rcv *MsgGroupSubmitProposalImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgGroupSubmitProposalImpl) GetGroupPolicyAddress() string {
	return rcv.GroupPolicyAddress
}

func (rcv *MsgGroupSubmitProposalImpl) GetProposers() []string {
	return rcv.Proposers
}

func (rcv *MsgGroupSubmitProposalImpl) GetMetadata() string {
	return rcv.Metadata
}

func (rcv *MsgGroupSubmitProposalImpl) GetMessages() []*pkl.Object {
	return rcv.Messages
}

func (rcv *MsgGroupSubmitProposalImpl) GetExec() int {
	return rcv.Exec
}
