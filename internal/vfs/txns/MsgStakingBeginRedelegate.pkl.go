// Code generated from Pkl module `txns`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

type MsgStakingBeginRedelegate interface {
	Msg

	GetDelegatorAddress() string

	GetValidatorSrcAddress() string

	GetValidatorDstAddress() string

	GetAmount() *pkl.Object
}

var _ MsgStakingBeginRedelegate = (*MsgStakingBeginRedelegateImpl)(nil)

type MsgStakingBeginRedelegateImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	DelegatorAddress string `pkl:"delegatorAddress"`

	ValidatorSrcAddress string `pkl:"validatorSrcAddress"`

	ValidatorDstAddress string `pkl:"validatorDstAddress"`

	Amount *pkl.Object `pkl:"amount"`
}

// The type URL for the message
func (rcv *MsgStakingBeginRedelegateImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgStakingBeginRedelegateImpl) GetDelegatorAddress() string {
	return rcv.DelegatorAddress
}

func (rcv *MsgStakingBeginRedelegateImpl) GetValidatorSrcAddress() string {
	return rcv.ValidatorSrcAddress
}

func (rcv *MsgStakingBeginRedelegateImpl) GetValidatorDstAddress() string {
	return rcv.ValidatorDstAddress
}

func (rcv *MsgStakingBeginRedelegateImpl) GetAmount() *pkl.Object {
	return rcv.Amount
}
