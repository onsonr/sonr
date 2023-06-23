package keeper

import (
	"context"
	"crypto/rand"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/x/vault/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// This is a function in the `keeper` package that sends an inbox message to a specified recipient. It takes in a context and a `SendWalletMailRequest` object as input, and returns a `SendWalletMailResponse` object and an error (if any) as output. The function writes the message
// to the recipient's inbox using the `WriteInbox` function of the `Keeper` struct, and returns a success response with a message ID.
func (k Keeper) SendWalletMail(goCtx context.Context, req *types.SendWalletMailRequest) (*types.SendWalletMailResponse, error) {
	// err := k.WriteInbox(req.To, req.GetMail())
	// if err != nil {
	// 	return nil, types.ErrInboxWrite
	// }
	return &types.SendWalletMailResponse{
		Success: true,
	}, nil
}

// `func (k Keeper) ReadWalletMails(goCtx context.Context, req *types.ReadWalletMailsRequest) (*types.ReadWalletMailsResponse, error)` is a function in the `keeper` package that reads all the inbox messages of a specified recipient. It takes in a context and a
// `ReadWalletMailsRequest` object as input, and returns a `ReadWalletMailsResponse` object and an error (if any) as output. The function reads the messages from the recipient's inbox using the `ReadInbox` function of the `Keeper` struct, and returns a response with all the
// messages in the inbox.
func (k Keeper) ReadWalletMail(goCtx context.Context, req *types.ReadWalletMailRequest) (*types.ReadWalletMailResponse, error) {
	// _, err := k.ReadInbox(req.Creator)
	// if err != nil {
	// 	return nil, types.ErrInboxRead
	// }
	return &types.ReadWalletMailResponse{}, nil
}

func (k Keeper) ClaimableWalletAll(goCtx context.Context, req *types.QueryAllClaimableWalletRequest) (*types.QueryAllClaimableWalletResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var claimableWallets []types.ClaimableWallet
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	claimableWalletStore := prefix.NewStore(store, types.KeyPrefix(types.ClaimableWalletKey))

	pageRes, err := query.Paginate(claimableWalletStore, req.Pagination, func(key []byte, value []byte) error {
		var claimableWallet types.ClaimableWallet
		if err := k.cdc.Unmarshal(value, &claimableWallet); err != nil {
			return err
		}

		claimableWallets = append(claimableWallets, claimableWallet)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClaimableWalletResponse{ClaimableWallet: claimableWallets, Pagination: pageRes}, nil
}

func (k Keeper) ClaimableWallet(goCtx context.Context, req *types.QueryGetClaimableWalletRequest) (*types.QueryGetClaimableWalletResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	claimableWallet, found := k.GetClaimableWallet(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetClaimableWalletResponse{ClaimableWallet: claimableWallet}, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                            Helper Utility Functions                            ||
// ! ||--------------------------------------------------------------------------------||

var _ types.QueryServer = Keeper{}

// ChallengeLength - Length of bytes to generate for a challenge.¡¡
const ChallengeLength = 32

// createChallenge creates a new challenge that should be signed and returned by the authenticator. The spec recommends
// using at least 16 bytes with 100 bits of entropy. We use 32 bytes.
func createChallenge() (challenge protocol.URLEncodedBase64, err error) {
	challenge = make([]byte, ChallengeLength)

	if _, err = rand.Read(challenge); err != nil {
		return nil, err
	}
	return challenge, nil
}

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
