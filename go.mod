module github.com/sonr-io/core

go 1.16

require (
	berty.tech/berty/v2 v2.268.0
	github.com/PuerkitoBio/goquery v1.6.1
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/golang/protobuf v1.5.1
	github.com/google/open-location-code/go v0.0.0-20201229230907-d47d9f9b95e9
	github.com/h2non/filetype v1.1.0
	github.com/libp2p/go-libp2p v0.13.0
	github.com/libp2p/go-libp2p-core v0.8.5
	github.com/libp2p/go-libp2p-discovery v0.5.0
	github.com/libp2p/go-libp2p-gorpc v0.1.2
	github.com/libp2p/go-libp2p-kad-dht v0.11.1
	github.com/libp2p/go-libp2p-pubsub v0.4.1
	github.com/libp2p/go-libp2p-swarm v0.4.0
	github.com/libp2p/go-libp2p-tls v0.1.3
	github.com/libp2p/go-msgio v0.0.6
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/pkg/errors v0.9.1
	github.com/tyler-smith/go-bip39 v1.1.0
	golang.org/x/net v0.0.0-20210316092652-d523dce5a7f4
	google.golang.org/protobuf v1.26.0
)

replace (
	bazil.org/fuse => bazil.org/fuse v0.0.0-20200117225306-7b5117fecadc // specific version for iOS building
	github.com/agl/ed25519 => github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412 // latest commit before the author shutdown the repo; see https://github.com/golang/go/issues/20504
	github.com/ipfs/go-ipfs => github.com/Jorropo/go-ipfs v0.4.20-0.20201127133049-9632069f4448 // temporary, see https://github.com/ipfs/go-ipfs/issues/7791
	github.com/libp2p/go-libp2p-rendezvous => github.com/berty/go-libp2p-rendezvous v0.0.0-20201028141428-5b2e7e8ff19a // use berty fork of go-libp2p-rendezvous
	github.com/libp2p/go-libp2p-swarm => github.com/Jorropo/go-libp2p-swarm v0.4.2 // temporary, see https://github.com/libp2p/go-libp2p-swarm/pull/227
	github.com/peterbourgon/ff/v3 => github.com/moul/ff/v3 v3.0.1 // temporary, see https://github.com/peterbourgon/ff/pull/67, https://github.com/peterbourgon/ff/issues/68
	golang.org/x/mobile => github.com/aeddi/mobile v0.0.2 // temporary, see https://github.com/golang/mobile/pull/58
)
