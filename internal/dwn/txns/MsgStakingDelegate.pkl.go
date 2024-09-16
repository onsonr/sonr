// Code generated from Pkl module `txns`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

type MsgStakingDelegate interface {
	Msg

	GetDelegatorAddress() string

	GetValidatorAddress() string

	GetAmount() *pkl.Object
}

var _ MsgStakingDelegate = (*MsgStakingDelegateImpl)(nil)

type MsgStakingDelegateImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	DelegatorAddress string `pkl:"delegatorAddress"`

	ValidatorAddress string `pkl:"validatorAddress"`

	Amount *pkl.Object `pkl:"amount"`
}

// The type URL for the message
func (rcv *MsgStakingDelegateImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgStakingDelegateImpl) GetDelegatorAddress() string {
	return rcv.DelegatorAddress
}

func (rcv *MsgStakingDelegateImpl) GetValidatorAddress() string {
	return rcv.ValidatorAddress
}

func (rcv *MsgStakingDelegateImpl) GetAmount() *pkl.Object {
	return rcv.Amount
}
