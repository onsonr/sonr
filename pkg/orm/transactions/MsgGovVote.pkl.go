// Code generated from Pkl module `transactions`. DO NOT EDIT.
package transactions

type MsgGovVote interface {
	Msg

	GetProposalId() int

	GetVoter() string

	GetOption() int
}

var _ MsgGovVote = (*MsgGovVoteImpl)(nil)

type MsgGovVoteImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	ProposalId int `pkl:"proposalId"`

	Voter string `pkl:"voter"`

	Option int `pkl:"option"`
}

// The type URL for the message
func (rcv *MsgGovVoteImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgGovVoteImpl) GetProposalId() int {
	return rcv.ProposalId
}

func (rcv *MsgGovVoteImpl) GetVoter() string {
	return rcv.Voter
}

func (rcv *MsgGovVoteImpl) GetOption() int {
	return rcv.Option
}
