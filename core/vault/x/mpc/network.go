package mpc

import (
	"context"

	"github.com/kataras/golog"
	"github.com/sonr-hq/sonr/internal/node"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/pkg/protocol"
)

func handlerLoopTopic(id party.ID, mpc protocol.Handler, ps node.TopicHandler) {
	ctx := context.Background()
	for {
		select {
		case msg := <-ps.Messages():
			var m protocol.Message
			err := m.UnmarshalBinary(msg)
			checkErr(ctx, err)
			if mpc.CanAccept(&m) {
				mpc.Accept(&m)
			} else {
				mpc.Stop()
			}

		case msg, ok := <-mpc.Listen():
			if !ok {
				mpc.Stop()
				return
			}

			buf, err := msg.MarshalBinary()
			checkErr(ctx, err)
			err = ps.Publish(buf)
			checkErr(ctx, err)
		case <-ctx.Done():
			return
		}
	}
}

func checkErr(ctx context.Context, err error) {
	if err != nil {
		golog.Warnf("error: %v", err) // nolint: errcheck
		ctx.Done()
	}
}
