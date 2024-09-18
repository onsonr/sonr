// Code generated from Pkl module `txns`. DO NOT EDIT.
package txns

import "github.com/apple/pkl-go/pkl"

func init() {
	pkl.RegisterMapping("txns", Txns{})
	pkl.RegisterMapping("txns#Proposal", Proposal{})
	pkl.RegisterMapping("txns#SWT", SWT{})
	pkl.RegisterMapping("txns#MsgGovSubmitProposal", MsgGovSubmitProposalImpl{})
	pkl.RegisterMapping("txns#MsgGovVote", MsgGovVoteImpl{})
	pkl.RegisterMapping("txns#MsgGovDeposit", MsgGovDepositImpl{})
	pkl.RegisterMapping("txns#MsgGroupCreateGroup", MsgGroupCreateGroupImpl{})
	pkl.RegisterMapping("txns#MsgGroupSubmitProposal", MsgGroupSubmitProposalImpl{})
	pkl.RegisterMapping("txns#MsgGroupVote", MsgGroupVoteImpl{})
	pkl.RegisterMapping("txns#MsgStakingCreateValidator", MsgStakingCreateValidatorImpl{})
	pkl.RegisterMapping("txns#MsgStakingDelegate", MsgStakingDelegateImpl{})
	pkl.RegisterMapping("txns#MsgStakingUndelegate", MsgStakingUndelegateImpl{})
	pkl.RegisterMapping("txns#MsgStakingBeginRedelegate", MsgStakingBeginRedelegateImpl{})
	pkl.RegisterMapping("txns#MsgDidUpdateParams", MsgDidUpdateParamsImpl{})
	pkl.RegisterMapping("txns#MsgDidAllocateVault", MsgDidAllocateVaultImpl{})
	pkl.RegisterMapping("txns#MsgDidProveWitness", MsgDidProveWitnessImpl{})
	pkl.RegisterMapping("txns#MsgDidSyncVault", MsgDidSyncVaultImpl{})
	pkl.RegisterMapping("txns#MsgDidRegisterController", MsgDidRegisterControllerImpl{})
	pkl.RegisterMapping("txns#MsgDidAuthorize", MsgDidAuthorizeImpl{})
	pkl.RegisterMapping("txns#MsgDidRegisterService", MsgDidRegisterServiceImpl{})
	pkl.RegisterMapping("txns#TxBody", TxBody{})
}
