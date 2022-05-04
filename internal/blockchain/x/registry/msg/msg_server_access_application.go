package msg

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
)

func (k msgServer) AccessApplication(goCtx context.Context, msg *types.MsgAccessApplication) (*types.MsgAccessApplicationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Try getting name information from the store
	whois, found := k.GetWhoIs(ctx, msg.GetAppName())
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name not found in registry")
	}

	// If the message sender address doesn't match the name owner, throw an error
	if !(msg.Creator == whois.GetCreator()) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Registered name is owned by another address")
	}

	// If the whois type is not user, throw an error
	if whois.GetType() != types.WhoIs_Application {
		return nil, types.ErrInvalidWhoisType
	}

	// Create new session object
	session := &types.Session{
		BaseDid:    whois.Did,
		Whois:      &whois,
		Credential: msg.GetCredential(),
	}

	return &types.MsgAccessApplicationResponse{
		Session: session,
		WhoIs:   &whois,
	}, nil
}
