package actor

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/crypto/mpc"
	"github.com/ipfs/boxo/path"
)

// controllerActor is the spawned actor for the controller which manages the signing of messages
type controllerActor struct {
	VKS     kss.Val
	UKS     kss.User
	daedKey []byte
	Path    path.Path
}

// Implement the Receive method for message processing
func (state *controllerActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *Request:
		state.runSignProtocol(context, msg.Data)
		return
	}
}

func (c *controllerActor) runSignProtocol(ctx actor.Context, msg []byte) {
	uSign, err := c.UKS.GetSignFunc(msg)
	if err != nil {
		ctx.Respond(err)
	}
	vSign, err := c.VKS.GetSignFunc(msg)
	if err != nil {
		ctx.Respond(err)
	}
	sig, err := mpc.RunSignProtocol(vSign, uSign)
	if err != nil {
		ctx.Respond(err)
	}
	ctx.Respond(newSignResponse(sig))
}
