package host

import "context"

func (sh *SonrHost) managePeers() {
	for peer := range sh.Channel {
		if peer.ID == sh.Host.ID() {
			continue
		}
		println("Found peer:", peer.String())

		err := sh.Host.Connect(context.Background(), peer)

		if err != nil {
			println("Error connecting to peer")
			continue
		} else {
			println("Connected to:", peer.String())
		}
	}
}
