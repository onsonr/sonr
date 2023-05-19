package cosmos

import (
	"github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	models "github.com/sonrhq/core/x/vault/types"
)

// ! ||-------------------------------------------------------------------------------||
// ! ||                                  Stake Config                                 ||
// ! ||-------------------------------------------------------------------------------||
// TxStakingOpt is a function that sets optional parameters for a staking message.
type TxStakingOpt func(*StakingTransactionConfig)

// StakingTransactionConfig is a struct that holds the configuration for a staking transaction.
type StakingTransactionConfig struct {
	Identity        string
	Website         string
	Moniker         string
	SecurityContact string
	StakeAmount     types.Coin
	DelegateAmount  types.Coin
	CreationHeight  int64
	ValidatorAddr   string
}

// defaultConfig is a struct that holds the default configuration for a staking transaction.
func defaultConfig() *StakingTransactionConfig {
	return &StakingTransactionConfig{
		Identity:        "",
		Website:         "",
		Moniker:         "default",
		SecurityContact: "",
		StakeAmount:     types.NewCoin("usnr", types.NewInt(500000000000)),
		DelegateAmount:  types.NewCoin("usnr", types.NewInt(100000)),
		CreationHeight:  0,
	}
}

// GetStakingDescription returns a staking description from a staking transaction configuration.
func (stc *StakingTransactionConfig) GetStakingDescription() stakingtypes.Description {
	return stakingtypes.Description{
		Moniker:         stc.Moniker,
		Identity:        stc.Identity,
		Website:         stc.Website,
		SecurityContact: stc.SecurityContact,
	}
}

// GetMinSelfDelegation returns the minimum self delegation from a staking transaction configuration.
func (stc *StakingTransactionConfig) GetMinSelfDelegation() types.Int {
	return types.NewInt(500000000000)
}

// GetCommission returns the commission from a staking transaction configuration.
func (stc *StakingTransactionConfig) GetCommission() stakingtypes.CommissionRates {
	rate := types.MustNewDecFromStr("0.05")
	maxRate := types.MustNewDecFromStr("0.075")
	maxChangeRate := types.MustNewDecFromStr("0.01")
	return stakingtypes.CommissionRates{
		Rate:          rate,
		MaxRate:       maxRate,
		MaxChangeRate: maxChangeRate,
	}
}

// The function returns a CreateValidatorOpt that sets the identity of a staking message's description.
func WithDescriptionIdentity(identity string) TxStakingOpt {
	return func(mcv *StakingTransactionConfig) {
		mcv.Identity = identity
	}
}

// The function returns a CreateValidatorOpt that sets the moniker of a staking message's description.
func WithDescriptionMoniker(moniker string) TxStakingOpt {
	return func(mcv *StakingTransactionConfig) {
		mcv.Moniker = moniker
	}
}

// The function sets the website field of a staking message's description.
func WithDescriptionWebsite(website string) TxStakingOpt {
	return func(mcv *StakingTransactionConfig) {
		mcv.Website = website
	}
}

// The function WithSecurityContact sets the security contact for a staking validator.
func WithSecurityContact(securityContact string) TxStakingOpt {
	return func(mcv *StakingTransactionConfig) {
		mcv.SecurityContact = securityContact
	}
}

// The function sets the stake amount for a validator in a staking message.
func WithStakeAmount(amount int) TxStakingOpt {
	return func(mcv *StakingTransactionConfig) {
		mcv.StakeAmount = types.NewCoin("usnr", types.NewInt(int64(amount)))
	}
}

// The function sets the delegate amount for a validator in a staking message.
func WithDelegateAmount(amount int) TxStakingOpt {
	return func(mcv *StakingTransactionConfig) {
		mcv.DelegateAmount = types.NewCoin("usnr", types.NewInt(int64(amount)))
	}
}

// The function sets the creation height for a validator in a staking message.
func WithCreationHeight(height int64) TxStakingOpt {
	return func(mcv *StakingTransactionConfig) {
		mcv.CreationHeight = height
	}
}

// ! ||----------------------------------------------------------------------||
// ! ||                               Messages                               ||
// ! ||----------------------------------------------------------------------||

// This function builds and returns a staking message to create a validator with optional parameters.
func BuildMsgCreateValidator(acc models.Account, options ...TxStakingOpt) *stakingtypes.MsgCreateValidator {
	conf := defaultConfig()

	anyPk, err := acc.PubKey().Secp256k1AnyProto()
	if err != nil {
		return nil
	}
	for _, opt := range options {
		opt(conf)
	}
	msg := &stakingtypes.MsgCreateValidator{
		Description:       conf.GetStakingDescription(),
		Commission:        conf.GetCommission(),
		MinSelfDelegation: conf.GetMinSelfDelegation(),
		DelegatorAddress:  acc.Address(),
		ValidatorAddress:  acc.Address(),
		Value:             conf.StakeAmount,
	}
	msg.Pubkey = anyPk
	return msg
}

// This function builds and returns a staking message to edit a validator with optional parameters.
func BuildMsgEditValidator(acc models.Account, options ...TxStakingOpt) *stakingtypes.MsgEditValidator {
	conf := defaultConfig()
	for _, opt := range options {
		opt(conf)
	}
	msg := &stakingtypes.MsgEditValidator{
		Description:      conf.GetStakingDescription(),
		ValidatorAddress: acc.Address(),
	}
	return msg
}

// The function builds a message to delegate a certain amount to a validator for a given account.
func BuildMsgDelegate(acc models.Account, valAddr string, options ...TxStakingOpt) *stakingtypes.MsgDelegate {
	conf := defaultConfig()
	for _, opt := range options {
		opt(conf)
	}
	msg := &stakingtypes.MsgDelegate{
		DelegatorAddress: acc.Address(),
		ValidatorAddress: valAddr,
		Amount:           conf.DelegateAmount,
	}
	return msg
}

// The function builds a message to undelegate a certain amount from a validator for a given account.
func BuildMsgUndelegate(acc models.Account, valAddr string, options ...TxStakingOpt) *stakingtypes.MsgUndelegate {
	conf := defaultConfig()
	for _, opt := range options {
		opt(conf)
	}
	msg := &stakingtypes.MsgUndelegate{
		DelegatorAddress: acc.Address(),
		ValidatorAddress: valAddr,
		Amount:           conf.DelegateAmount,
	}
	return msg
}

// The function builds a message to cancel an unbonding delegation for a given account and validator
// address with optional configuration parameters.
func BuildMsgCancelUndelegate(acc models.Account, valAddr string, options ...TxStakingOpt) *stakingtypes.MsgCancelUnbondingDelegation {
	conf := defaultConfig()
	for _, opt := range options {
		opt(conf)
	}
	msg := &stakingtypes.MsgCancelUnbondingDelegation{
		DelegatorAddress: acc.Address(),
		ValidatorAddress: valAddr,
		Amount:           conf.DelegateAmount,
		CreationHeight:   conf.CreationHeight,
	}
	return msg
}
