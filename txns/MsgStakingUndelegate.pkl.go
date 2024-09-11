// Code generated from Pkl module `transactions`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

type MsgStakingUndelegate interface {
	Msg

	GetDelegatorAddress() string

	GetValidatorAddress() string

	GetAmount() *pkl.Object
}

var _ MsgStakingUndelegate = (*MsgStakingUndelegateImpl)(nil)

type MsgStakingUndelegateImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	DelegatorAddress string `pkl:"delegatorAddress"`

	ValidatorAddress string `pkl:"validatorAddress"`

	Amount *pkl.Object `pkl:"amount"`
}

// The type URL for the message
func (rcv *MsgStakingUndelegateImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgStakingUndelegateImpl) GetDelegatorAddress() string {
	return rcv.DelegatorAddress
}

func (rcv *MsgStakingUndelegateImpl) GetValidatorAddress() string {
	return rcv.ValidatorAddress
}

func (rcv *MsgStakingUndelegateImpl) GetAmount() *pkl.Object {
	return rcv.Amount
}
