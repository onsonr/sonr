// Code generated from Pkl module `transactions`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

func init() {
	pkl.RegisterMapping("transactions", Transactions{})
	pkl.RegisterMapping("transactions#Proposal", Proposal{})
	pkl.RegisterMapping("transactions#MsgGovSubmitProposal", MsgGovSubmitProposalImpl{})
	pkl.RegisterMapping("transactions#MsgGovVote", MsgGovVoteImpl{})
	pkl.RegisterMapping("transactions#MsgGovDeposit", MsgGovDepositImpl{})
	pkl.RegisterMapping("transactions#MsgGroupCreateGroup", MsgGroupCreateGroupImpl{})
	pkl.RegisterMapping("transactions#MsgGroupSubmitProposal", MsgGroupSubmitProposalImpl{})
	pkl.RegisterMapping("transactions#MsgGroupVote", MsgGroupVoteImpl{})
	pkl.RegisterMapping("transactions#MsgStakingCreateValidator", MsgStakingCreateValidatorImpl{})
	pkl.RegisterMapping("transactions#MsgStakingDelegate", MsgStakingDelegateImpl{})
	pkl.RegisterMapping("transactions#MsgStakingUndelegate", MsgStakingUndelegateImpl{})
	pkl.RegisterMapping("transactions#MsgStakingBeginRedelegate", MsgStakingBeginRedelegateImpl{})
	pkl.RegisterMapping("transactions#MsgDidUpdateParams", MsgDidUpdateParamsImpl{})
	pkl.RegisterMapping("transactions#MsgDidAllocateVault", MsgDidAllocateVaultImpl{})
	pkl.RegisterMapping("transactions#MsgDidProveWitness", MsgDidProveWitnessImpl{})
	pkl.RegisterMapping("transactions#MsgDidSyncVault", MsgDidSyncVaultImpl{})
	pkl.RegisterMapping("transactions#MsgDidRegisterController", MsgDidRegisterControllerImpl{})
	pkl.RegisterMapping("transactions#MsgDidAuthorize", MsgDidAuthorizeImpl{})
	pkl.RegisterMapping("transactions#MsgDidRegisterService", MsgDidRegisterServiceImpl{})
	pkl.RegisterMapping("transactions#TxBody", TxBody{})
}
