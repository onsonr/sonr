// Code generated from Pkl module `txns`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

type MsgStakingCreateValidator interface {
	Msg

	GetDescription() *pkl.Object

	GetCommission() *pkl.Object

	GetMinSelfDelegation() string

	GetDelegatorAddress() string

	GetValidatorAddress() string

	GetPubkey() *pkl.Object

	GetValue() *pkl.Object
}

var _ MsgStakingCreateValidator = (*MsgStakingCreateValidatorImpl)(nil)

// Staking module messages
type MsgStakingCreateValidatorImpl struct {
	// The type URL for the message
	TypeUrl string `pkl:"typeUrl"`

	Description *pkl.Object `pkl:"description"`

	Commission *pkl.Object `pkl:"commission"`

	MinSelfDelegation string `pkl:"minSelfDelegation"`

	DelegatorAddress string `pkl:"delegatorAddress"`

	ValidatorAddress string `pkl:"validatorAddress"`

	Pubkey *pkl.Object `pkl:"pubkey"`

	Value *pkl.Object `pkl:"value"`
}

// The type URL for the message
func (rcv *MsgStakingCreateValidatorImpl) GetTypeUrl() string {
	return rcv.TypeUrl
}

func (rcv *MsgStakingCreateValidatorImpl) GetDescription() *pkl.Object {
	return rcv.Description
}

func (rcv *MsgStakingCreateValidatorImpl) GetCommission() *pkl.Object {
	return rcv.Commission
}

func (rcv *MsgStakingCreateValidatorImpl) GetMinSelfDelegation() string {
	return rcv.MinSelfDelegation
}

func (rcv *MsgStakingCreateValidatorImpl) GetDelegatorAddress() string {
	return rcv.DelegatorAddress
}

func (rcv *MsgStakingCreateValidatorImpl) GetValidatorAddress() string {
	return rcv.ValidatorAddress
}

func (rcv *MsgStakingCreateValidatorImpl) GetPubkey() *pkl.Object {
	return rcv.Pubkey
}

func (rcv *MsgStakingCreateValidatorImpl) GetValue() *pkl.Object {
	return rcv.Value
}
