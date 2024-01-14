package shares

import (
	"fmt"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/sonrhq/sonr/crypto/core/protocol"
	"github.com/sonrhq/sonr/crypto/tecdsa/dklsv1"
	"golang.org/x/crypto/sha3"
)

type privateSign struct {
	aliceDkg  *dklsv1.AliceDkg
	aliceSign *dklsv1.AliceSign
	bobPID    *actor.PID
}

func (s *privateSign) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *protocol.Message:
		res, err := s.aliceSign.Next(msg)
		if err == protocol.ErrProtocolFinished {
			msg, err := s.aliceSign.Result(protocol.Version1)
			if err != nil {
				context.Respond(err)
			}
			context.Respond(msg)
		}
		if err != nil {
			context.Respond(err)
		}
		ctx.Send(s.bobPID, res)
	}
}

type publicSign struct {
	bobDkg   *dklsv1.BobDkg
	bobSign  *dklsv1.BobSign
	alicePID *actor.PID
}

func (s *publicSign) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *protocol.Message:
		res, err := s.bobSign.Next(msg)
		if err == protocol.ErrProtocolFinished {
			msg, err := s.bobSign.Result(protocol.Version1)
			if err != nil {
				context.Respond(err)
			}
			sig, err := dklsv1.DecodeSignature(msg)
			if err != nil {
				context.Respond(err)
			}
			context.Respond(sig)
		}
		if err != nil {
			context.Respond(err)
		}
		ctx.Send(s.alicePID, res)
	}
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                            Sign Actor configuration                            ||
// ! ||--------------------------------------------------------------------------------||

// ApplyPrivateGen applies the spawn options in order to create a private share actor
func (c *options) ApplyPrivateSign(msg []byte, opts ...SpawnOption) func() actor.Actor {
	for _, o := range opts {
		o(c)
	}
	newFunc := func() actor.Actor {
		p := &privateSign{
			aliceDkg: dklsv1.NewAliceDkg(K_DEFAULT_MPC_CURVE, protocol.Version1),
		}
		return p
	}
	signAlice, err := dklsv1.NewAliceSign(K_DEFAULT_MPC_CURVE, sha3.New256(), msg, c.DecodedMessage, protocol.Version1)
	if err != nil {
		fmt.Println(err)
		return newFunc
	}
	loadFunc := func() actor.Actor {
		p := &privateSign{
			aliceDkg:  dklsv1.NewAliceDkg(K_DEFAULT_MPC_CURVE, protocol.Version1),
			aliceSign: signAlice,
		}
		return p
	}
	return loadFunc
}

// ApplyPublicGen applies the spawn options in order to create a public share actor
func (c *options) ApplyPublicSign(msg []byte, opts ...SpawnOption) func() actor.Actor {
	for _, o := range opts {
		o(c)
	}
	newFunc := func() actor.Actor {
		p := &publicSign{
			bobDkg: dklsv1.NewBobDkg(K_DEFAULT_MPC_CURVE, protocol.Version1),
		}
		return p
	}
	signBob, err := dklsv1.NewBobSign(K_DEFAULT_MPC_CURVE, sha3.New256(), msg, c.DecodedMessage, protocol.Version1)
	if err != nil {
		fmt.Println(err)
		return newFunc
	}

	loadFunc := func() actor.Actor {
		p := &publicSign{
			bobDkg:  dklsv1.NewBobDkg(K_DEFAULT_MPC_CURVE, protocol.Version1),
			bobSign: signBob,
		}
		return p
	}
	return loadFunc
}
