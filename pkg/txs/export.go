package txs

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	grouptypes "github.com/cosmos/cosmos-sdk/x/group"
	domaintypes "github.com/sonr-io/core/x/domain/types"
	identitytypes "github.com/sonr-io/core/x/identity/types"
	servicetypes "github.com/sonr-io/core/x/service/types"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                                 Cosmos: x/bank                                 ||
// ! ||--------------------------------------------------------------------------------||
type MsgSend = banktypes.MsgSend
type MsgMultiSend = banktypes.MsgMultiSend

// ! ||--------------------------------------------------------------------------------||
// ! ||                                 Cosmos: x/group                                ||
// ! ||--------------------------------------------------------------------------------||
type MsgCreateGroup = grouptypes.MsgCreateGroup
type MsgUpdateGroupMembers = grouptypes.MsgUpdateGroupMembers
type MsgUpdateGroupAdmin = grouptypes.MsgUpdateGroupAdmin
type MsgUpdateGroupMetadata = grouptypes.MsgUpdateGroupMetadata
type MsgCreateGroupPolicy = grouptypes.MsgCreateGroupPolicy
type MsgUpdateGroupPolicyAdmin = grouptypes.MsgUpdateGroupPolicyAdmin
type MsgCreateGroupWithPolicy = grouptypes.MsgCreateGroupWithPolicy
type MsgUpdateGroupPolicyDecisionPolicy = grouptypes.MsgUpdateGroupPolicyDecisionPolicy
type MsgUpdateGroupPolicyMetadata = grouptypes.MsgUpdateGroupPolicyMetadata
type MsgSubmitProposal = grouptypes.MsgSubmitProposal
type MsgWithdrawProposal = grouptypes.MsgWithdrawProposal
type MsgVote = grouptypes.MsgVote
type MsgExec = grouptypes.MsgExec
type MsgLeaveGroup = grouptypes.MsgLeaveGroup

// ! ||--------------------------------------------------------------------------------||
// ! ||                                 Sonr: x/domain                                 ||
// ! ||--------------------------------------------------------------------------------||
type MsgCreateUsernameRecords = domaintypes.MsgCreateUsernameRecords
type MsgUpdateUsernameRecords = domaintypes.MsgUpdateUsernameRecords
type MsgDeleteUsernameRecords = domaintypes.MsgDeleteUsernameRecords

// ! ||--------------------------------------------------------------------------------||
// ! ||                                 Sonr: x/service                                ||
// ! ||--------------------------------------------------------------------------------||
type MsgCreateServiceRecord = servicetypes.MsgCreateServiceRecord
type MsgUpdateServiceRecord = servicetypes.MsgUpdateServiceRecord
type MsgDeleteServiceRecord = servicetypes.MsgDeleteServiceRecord

// ! ||--------------------------------------------------------------------------------||
// ! ||                                Sonr: x/identity                                ||
// ! ||--------------------------------------------------------------------------------||
type MsgCreateControllerAccount = identitytypes.MsgCreateControllerAccount
type MsgUpdateControllerAccount = identitytypes.MsgUpdateControllerAccount
type MsgDeleteControllerAccount = identitytypes.MsgDeleteControllerAccount
type MsgCreateEscrowAccount = identitytypes.MsgCreateEscrowAccount
type MsgUpdateEscrowAccount = identitytypes.MsgUpdateEscrowAccount
type MsgDeleteEscrowAccount = identitytypes.MsgDeleteEscrowAccount
type MsgRegisterIdentity = identitytypes.MsgRegisterIdentity

