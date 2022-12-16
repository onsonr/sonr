package mpc

import (
	"github.com/kataras/golog"
	"github.com/sonr-hq/sonr/internal/node"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/pkg/protocol"
)

func handlerLoopTopic(id party.ID, h protocol.Handler, topicHandler node.TopicHandler) {

	for {
		select {
		case msg, ok := <-topicHandler.Messages():
			if !ok {
				return
			}
			var m protocol.Message
			err := m.UnmarshalBinary(msg)
			checkErr(err)
			h.Accept(&m)

		case msg, ok := <-h.Listen():
			if !ok {
				return
			}
			buf, err := msg.MarshalBinary()
			checkErr(err)
			err = topicHandler.Publish(buf)
			checkErr(err)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		golog.Warnf("error: %v", err) // nolint: errcheck
	}
}
