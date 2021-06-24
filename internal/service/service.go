package service

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	net "github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
)

type ServiceHandler interface {
	OnInvite([]byte)
	OnReply(id peer.ID, data []byte)
	OnConfirmed(inv *md.InviteRequest)
}

type ServiceClient interface {
	Invite(id peer.ID, inv *md.InviteRequest) error
	Respond(rep *md.InviteResponse)
	Close()
}

type serviceClient struct {
	ServiceClient

	// Common
	ctx     context.Context
	handler ServiceHandler
	host    net.HostNode
	user    *md.User

	// Services
	Local  *AuthService
	Device *DeviceService
}

func NewService(ctx context.Context, h net.HostNode, u *md.User, sh ServiceHandler) (ServiceClient, *md.SonrError) {
	// Create Client
	client := &serviceClient{
		ctx:     ctx,
		handler: sh,
		host:    h,
		user:    u,
	}

	// Begin Local Service
	err := client.StartLocal()
	if err != nil {
		return nil, err
	}

	return client, nil
}
