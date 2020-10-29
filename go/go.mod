module github.com/ipfs-shipyard/gomobile-ipfs/go

go 1.14

require (
	github.com/ipfs/go-ipfs v0.7.0
	github.com/ipfs/go-ipfs-api v0.2.0
	github.com/ipfs/go-ipfs-config v0.9.0
	github.com/libp2p/go-libp2p v0.11.0
	github.com/libp2p/go-libp2p-core v0.6.1
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/pkg/errors v0.9.1
)

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.18.0 // required by go1.15
