package shares

import (
	"encoding/json"
	"fmt"

	"github.com/asynkron/protoactor-go/actor"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
	"github.com/sonrhq/sonr/crypto/core/curves"
	"github.com/sonrhq/sonr/crypto/core/protocol"
	"github.com/sonrhq/sonr/crypto/tecdsa/dklsv1"
)

type privateGen struct {
	aliceDkg *dklsv1.AliceDkg
	bobPID   *actor.PID
}

func (s *privateGen) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *protocol.Message:
		res, err := s.aliceDkg.Next(msg)
		if err == protocol.ErrProtocolFinished {
			msg, err := s.aliceDkg.Result(protocol.Version1)
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

type publicGen struct {
	bobDkg   *dklsv1.BobDkg
	alicePID *actor.PID
}

func (s *publicGen) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *protocol.Message:
		res, err := s.bobDkg.Next(msg)
		if err == protocol.ErrProtocolFinished {
			msg, err := s.bobDkg.Result(protocol.Version1)
			if err != nil {
				context.Respond(err)
			}
			context.Respond(msg)
		}
		if err != nil {
			context.Respond(err)
		}
		context.Send(context.Sender(), res)
	}
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                   Actor Spawn                                  ||
// ! ||--------------------------------------------------------------------------------||

func Generate(rootDir string, coinType modulev1.CoinType) (*actor.PID, *actor.PID, error) {
	c := defaultOptions()
	pub, err := ctx.SpawnNamed(c.ApplyPublicGen(), "public")
	if err != nil {
		return nil, nil, err
	}
	priv, err := ctx.SpawnNamed(c.ApplyPrivateGen(pub), "private")
	if err != nil {
		return nil, nil, err
	}
	return priv, pub, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                               Spawn Configuration                              ||
// ! ||--------------------------------------------------------------------------------||
var ctx = actor.NewActorSystem().Root

// K_DEFAULT_MPC_CURVE is the default curve for the controller.
var K_DEFAULT_MPC_CURVE = curves.K256()

type SpawnOption func(c *options)

type options struct {
	DecodedMessage *protocol.Message
	Message        []byte
}

// defaultOptions returns the default options for the private share actor
func defaultOptions() *options {
	return &options{}
}

// Attempts to decode the output bytes into a protocol.Message
func WithOutputBytes(out []byte) SpawnOption {
	return func(c *options) {
		decodedMsg := &protocol.Message{}
		err := json.Unmarshal(out, decodedMsg)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.DecodedMessage = decodedMsg
	}
}

// ! ||-------------------------------------------------------------------------------||
// ! ||                            Gen Actor configuration                            ||
// ! ||-------------------------------------------------------------------------------||

// ApplyPrivateGen applies the spawn options in order to create a private share actor
func (c *options) ApplyPrivateGen(pubPid *actor.PID) *actor.Props {
	newFunc := func() actor.Actor {
		p := &privateGen{
			aliceDkg: dklsv1.NewAliceDkg(K_DEFAULT_MPC_CURVE, protocol.Version1),
			bobPID:   pubPid,
		}
		return p
	}
	return actor.PropsFromProducer(newFunc)
}

// ApplyPublicGen applies the spawn options in order to create a public share actor
func (c *options) ApplyPublicGen() *actor.Props {
	newFunc := func() actor.Actor {
		p := &publicGen{
			bobDkg: dklsv1.NewBobDkg(K_DEFAULT_MPC_CURVE, protocol.Version1),
		}
		return p
	}
	return actor.PropsFromProducer(newFunc)
}
