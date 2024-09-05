// Code generated from Pkl module `transactions`. DO NOT EDIT.
package transactions

type MsgGroupVote interface {
	Msg

	GetProposalId() int

	GetVoter() string

	GetOption() int

	GetMetadata() string

	GetExec() int
}

var _ MsgGroupVote = (*MsgGroupVoteImpl)(nil)

type MsgGroupVoteImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	ProposalId int `pkl:"proposalId"`

	Voter string `pkl:"voter"`

	Option int `pkl:"option"`

	Metadata string `pkl:"metadata"`

	Exec int `pkl:"exec"`
}

// The type URL for the message
func (rcv *MsgGroupVoteImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgGroupVoteImpl) GetProposalId() int {
	return rcv.ProposalId
}

func (rcv *MsgGroupVoteImpl) GetVoter() string {
	return rcv.Voter
}

func (rcv *MsgGroupVoteImpl) GetOption() int {
	return rcv.Option
}

func (rcv *MsgGroupVoteImpl) GetMetadata() string {
	return rcv.Metadata
}

func (rcv *MsgGroupVoteImpl) GetExec() int {
	return rcv.Exec
}
