// Code generated from Pkl module `txns`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

type MsgGovDeposit interface {
	Msg

	GetProposalId() int

	GetDepositor() string

	GetAmount() []*pkl.Object
}

var _ MsgGovDeposit = (*MsgGovDepositImpl)(nil)

type MsgGovDepositImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	ProposalId int `pkl:"proposalId"`

	Depositor string `pkl:"depositor"`

	Amount []*pkl.Object `pkl:"amount"`
}

// The type URL for the message
func (rcv *MsgGovDepositImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgGovDepositImpl) GetProposalId() int {
	return rcv.ProposalId
}

func (rcv *MsgGovDepositImpl) GetDepositor() string {
	return rcv.Depositor
}

func (rcv *MsgGovDepositImpl) GetAmount() []*pkl.Object {
	return rcv.Amount
}
