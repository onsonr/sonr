package channel

import (
	ct "github.com/sonr-io/blockchain/x/channel/types"
)

// handleStoreMessages method listens to Pubsub Messages for room
func (b *channel) handleChannelMessages() {
	// Loop Messages
	for {
		// Get next message
		buf, err := b.messagesSub.Next(b.ctx)
		if err != nil {
			return
		}

		// Unmarshal Message Data
		msg := &ct.ChannelMessage{}
		err = msg.Unmarshal(buf.Data)
		if err != nil {
			logger.Errorf("failed to Unmarshal Message from pubsub.Message")
			return
		}

		// Push Message to Channel
		b.messages <- msg
	}
}

// serve handles the serving of the beam
func (b *channel) serve() {
	for {
		select {
		case <-b.ctx.Done():
			logger.Debugf("Closing Beam (%s)", b.label)
			return
		}
	}
}
