package shares

import (
	"fmt"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/sonrhq/sonr/crypto/core/curves"
	"github.com/sonrhq/sonr/crypto/core/protocol"
	"github.com/sonrhq/sonr/crypto/tecdsa/dklsv1"
)

type privateRefresh struct {
	aliceDkg     *dklsv1.AliceDkg
	aliceRefresh *dklsv1.AliceRefresh
	bobPID       *actor.PID
}

func (s *privateRefresh) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *protocol.Message:
		res, err := s.aliceRefresh.Next(msg)
		if err == protocol.ErrProtocolFinished {
			res, err := s.aliceRefresh.Result(protocol.Version1)
			if err != nil {
				context.Respond(err)
				return
			}
			refresh, err := dklsv1.DecodeAliceRefreshResult(res)
			if err != nil {
				context.Respond(err)
				return
			}

			context.Respond(refresh)
			return
		}
		if err != nil {
			context.Respond(err)
			return
		}
		ctx.Send(s.bobPID, res)
	}
}

type publicRefresh struct {
	bobDkg     *dklsv1.BobDkg
	bobRefresh *dklsv1.BobRefresh
	alicePID   *actor.PID
}

func (s *publicRefresh) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *protocol.Message:
		res, err := s.bobRefresh.Next(msg)
		if err == protocol.ErrProtocolFinished {
			res, err := s.bobRefresh.Result(protocol.Version1)
			if err != nil {
				context.Respond(err)
				return
			}
			refresh, err := dklsv1.DecodeBobRefreshResult(res)
			if err != nil {
				context.Respond(err)
				return
			}
			context.Respond(refresh)
		}
		if err != nil {
			context.Respond(err)
			return
		}
		ctx.Send(s.alicePID, res)
	}
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                           Refresh Actor Configuration                          ||
// ! ||--------------------------------------------------------------------------------||

// ApplyPrivateRefresh applies the spawn options in order to create a private share actor
func (c *options) ApplyPrivateRefresh(opts ...SpawnOption) func() actor.Actor {
	for _, o := range opts {
		o(c)
	}
	newFunc := func() actor.Actor {
		p := &privateRefresh{
			aliceDkg: dklsv1.NewAliceDkg(K_DEFAULT_MPC_CURVE, protocol.Version1),
		}
		return p
	}
	refresh, err := dklsv1.NewAliceRefresh(curves.P256(), c.DecodedMessage, protocol.Version1)
	if err != nil {
		fmt.Println(err)
		return newFunc
	}
	loadFunc := func() actor.Actor {
		p := &privateRefresh{
			aliceDkg:     dklsv1.NewAliceDkg(K_DEFAULT_MPC_CURVE, protocol.Version1),
			aliceRefresh: refresh,
		}
		return p
	}
	return loadFunc
}

// ApplyPublicRefresh applies the spawn options in order to create a public share actor
func (c *options) ApplyPublicRefresh(opts ...SpawnOption) func() actor.Actor {
	for _, o := range opts {
		o(c)
	}
	newFunc := func() actor.Actor {
		p := &publicRefresh{
			bobDkg: dklsv1.NewBobDkg(K_DEFAULT_MPC_CURVE, protocol.Version1),
		}
		return p
	}

	refresh, err := dklsv1.NewBobRefresh(curves.P256(), c.DecodedMessage, protocol.Version1)
	if err != nil {
		fmt.Println(err)
		return newFunc
	}

	loadFunc := func() actor.Actor {
		p := &publicRefresh{
			bobDkg:     dklsv1.NewBobDkg(K_DEFAULT_MPC_CURVE, protocol.Version1),
			bobRefresh: refresh,
		}
		return p
	}
	return loadFunc
}
