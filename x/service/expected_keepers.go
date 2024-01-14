package service

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/group"
)

// IdentityKeeper is the expected interface for the identity keeper.
type IdentityKeeper interface {
    // LinkCredential links a credential to an identity.
	GenerateIdentity(ctx context.Context) error
    LinkCredential(ctx context.Context, identityID string) error
	LinkPersona(ctx context.Context, identityID string) error
	RevokeAccount(ctx context.Context, identityID string) error
	RevokeIdentity(ctx context.Context, identityID string) error
	SignWithAccount(ctx context.Context, identityID string) error
	UnlinkCredential(ctx context.Context, identityID string) error
	UnlinkPersona(ctx context.Context, identityID string) error
	VerifyAccountSignature(ctx context.Context, identityID string) error
}

type GroupKeeper interface {
	CreateGroupWithPolicy(goCtx context.Context, req *group.MsgCreateGroupWithPolicy) (*group.MsgCreateGroupWithPolicyResponse, error)
	GroupMembers(goCtx context.Context, request *group.QueryGroupMembersRequest) (*group.QueryGroupMembersResponse, error)
	GroupPolicyInfo(goCtx context.Context, request *group.QueryGroupPolicyInfoRequest) (*group.QueryGroupPolicyInfoResponse, error)
	GroupsByAdmin(goCtx context.Context, request *group.QueryGroupsByAdminRequest) (*group.QueryGroupsByAdminResponse, error)
	GroupsByMember(goCtx context.Context, request *group.QueryGroupsByMemberRequest) (*group.QueryGroupsByMemberResponse, error)
	LeaveGroup(goCtx context.Context, req *group.MsgLeaveGroup) (*group.MsgLeaveGroupResponse, error)
	Proposal(goCtx context.Context, request *group.QueryProposalRequest) (*group.QueryProposalResponse, error)
	SubmitProposal(goCtx context.Context, req *group.MsgSubmitProposal) (*group.MsgSubmitProposalResponse, error)
	UpdateGroupMembers(goCtx context.Context, req *group.MsgUpdateGroupMembers) (*group.MsgUpdateGroupMembersResponse, error)
	UpdateGroupMetadata(goCtx context.Context, req *group.MsgUpdateGroupMetadata) (*group.MsgUpdateGroupMetadataResponse, error)
	Vote(goCtx context.Context, req *group.MsgVote) (*group.MsgVoteResponse, error)
	VotesByProposal(goCtx context.Context, request *group.QueryVotesByProposalRequest) (*group.QueryVotesByProposalResponse, error)
	WithdrawProposal(goCtx context.Context, req *group.MsgWithdrawProposal) (*group.MsgWithdrawProposalResponse, error)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
