package types

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	"github.com/go-webauthn/webauthn/protocol"
	identitytypes "github.com/sonrhq/core/x/identity/types"
)

type GroupKeeper interface {
	CreateGroup(goCtx context.Context, req *group.MsgCreateGroup) (*group.MsgCreateGroupResponse, error)
	CreateGroupPolicy(goCtx context.Context, req *group.MsgCreateGroupPolicy) (*group.MsgCreateGroupPolicyResponse, error)
	CreateGroupWithPolicy(goCtx context.Context, req *group.MsgCreateGroupWithPolicy) (*group.MsgCreateGroupWithPolicyResponse, error)
	GroupMembers(goCtx context.Context, request *group.QueryGroupMembersRequest) (*group.QueryGroupMembersResponse, error)
	GroupPolicyInfo(goCtx context.Context, request *group.QueryGroupPolicyInfoRequest) (*group.QueryGroupPolicyInfoResponse, error)
	GroupsByAdmin(goCtx context.Context, request *group.QueryGroupsByAdminRequest) (*group.QueryGroupsByAdminResponse, error)
	GroupsByMember(goCtx context.Context, request *group.QueryGroupsByMemberRequest) (*group.QueryGroupsByMemberResponse, error)
	LeaveGroup(goCtx context.Context, req *group.MsgLeaveGroup) (*group.MsgLeaveGroupResponse, error)
	Proposal(goCtx context.Context, request *group.QueryProposalRequest) (*group.QueryProposalResponse, error)
	PruneProposals(ctx sdk.Context) error
	SubmitProposal(goCtx context.Context, req *group.MsgSubmitProposal) (*group.MsgSubmitProposalResponse, error)
	UpdateGroupMembers(goCtx context.Context, req *group.MsgUpdateGroupMembers) (*group.MsgUpdateGroupMembersResponse, error)
	UpdateGroupMetadata(goCtx context.Context, req *group.MsgUpdateGroupMetadata) (*group.MsgUpdateGroupMetadataResponse, error)
	Vote(goCtx context.Context, req *group.MsgVote) (*group.MsgVoteResponse, error)
	VotesByProposal(goCtx context.Context, request *group.QueryVotesByProposalRequest) (*group.QueryVotesByProposalResponse, error)
	WithdrawProposal(goCtx context.Context, req *group.MsgWithdrawProposal) (*group.MsgWithdrawProposalResponse, error)
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// IdentityKeeper defines the expected interface needed to retrieve account balances.
type IdentityKeeper interface {
	CheckAlsoKnownAs(ctx sdk.Context, alias string) error
	AssignIdentity(ctx sdk.Context, ucw identitytypes.ClaimableWallet, cred *WebauthnCredential, alias string) (*identitytypes.Identity, error)
	GetAuthentication(ctx sdk.Context, reference string) (identitytypes.VerificationRelationship, bool)
	GetAssertion(ctx sdk.Context, reference string) (identitytypes.VerificationRelationship, bool)
	GetCapabilityInvocation(ctx sdk.Context, reference string) (invocation identitytypes.VerificationRelationship, found bool)
	GetCapabilityDelegation(ctx sdk.Context, reference string) (delegation identitytypes.VerificationRelationship, found bool)
	GetClaimableWallet(ctx sdk.Context, id uint64) (val identitytypes.ClaimableWallet, found bool)
	GetKeyAgreement(ctx sdk.Context, reference string) (agreement identitytypes.VerificationRelationship, found bool)

	NextUnclaimedWallet(ctx sdk.Context) (*identitytypes.ClaimableWallet, protocol.URLEncodedBase64, error)
	RegisterIdentity(goCtx context.Context, msg *identitytypes.MsgRegisterIdentity) (*identitytypes.MsgRegisterIdentityResponse, error)
	ResolveIdentity(ctx sdk.Context, did string) (identitytypes.DIDDocument, error)

	SetAuthentication(ctx sdk.Context, authentication identitytypes.VerificationRelationship)
	SetAssertion(ctx sdk.Context, assertion identitytypes.VerificationRelationship)
	SetCapabilityDelegation(ctx sdk.Context, delegation identitytypes.VerificationRelationship)
	SetCapabilityInvocation(ctx sdk.Context, invocation identitytypes.VerificationRelationship)
	SetKeyAgreement(ctx sdk.Context, agreement identitytypes.VerificationRelationship)
}
