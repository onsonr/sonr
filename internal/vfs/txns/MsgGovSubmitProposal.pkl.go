// Code generated from Pkl module `txns`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

type MsgGovSubmitProposal interface {
	Msg

	GetContent() *Proposal

	GetInitialDeposit() []*pkl.Object

	GetProposer() string
}

var _ MsgGovSubmitProposal = (*MsgGovSubmitProposalImpl)(nil)

// Gov module messages
type MsgGovSubmitProposalImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	Content *Proposal `pkl:"content"`

	InitialDeposit []*pkl.Object `pkl:"initialDeposit"`

	Proposer string `pkl:"proposer"`
}

// The type URL for the message
func (rcv *MsgGovSubmitProposalImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgGovSubmitProposalImpl) GetContent() *Proposal {
	return rcv.Content
}

func (rcv *MsgGovSubmitProposalImpl) GetInitialDeposit() []*pkl.Object {
	return rcv.InitialDeposit
}

func (rcv *MsgGovSubmitProposalImpl) GetProposer() string {
	return rcv.Proposer
}
