module github.com/sonr-io/core

go 1.16

require (
	git.mills.io/prologic/bitcask v1.0.0
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/duo-labs/webauthn v0.0.0-20220330035159-03696f3d4499
	github.com/duo-labs/webauthn.io v0.0.0-20200929144140-c031a3e0f95d
	github.com/fxamacker/cbor/v2 v2.4.0
	github.com/gabriel-vasile/mimetype v1.4.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/google/open-location-code/go v0.0.0-20210504205230-1796878d947c
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/hsanjuan/ipfs-lite v1.3.0
	github.com/ipfs/go-cid v0.1.0
	github.com/ipfs/go-datastore v0.5.1
	github.com/ipfs/go-ipfs-files v0.0.9 // indirect
	github.com/ipfs/go-path v0.2.1 // indirect
	github.com/ipfs/interface-go-ipfs-core v0.5.2
	github.com/ipld/go-codec-dagpb v1.3.2 // indirect
	github.com/ipld/go-ipld-prime v0.14.2 // indirect
	github.com/kataras/golog v0.1.7
	github.com/lestrrat-go/jwx v1.2.20
	github.com/libp2p/go-libp2p v0.17.0
	github.com/libp2p/go-libp2p-connmgr v0.3.1
	github.com/libp2p/go-libp2p-core v0.13.0
	github.com/libp2p/go-libp2p-discovery v0.6.0
	github.com/libp2p/go-libp2p-http v0.2.1
	github.com/libp2p/go-libp2p-kad-dht v0.15.0
	github.com/libp2p/go-libp2p-pubsub v0.6.0
	github.com/libp2p/go-msgio v0.1.0
	github.com/matrix-org/dendrite v0.8.1
	github.com/matrix-org/gomatrixserverlib v0.0.0-20220408160933-cf558306b56f
	github.com/matrix-org/util v0.0.0-20200807132607-55161520e1d4
	github.com/multiformats/go-multiaddr v0.5.0
	github.com/ockam-network/did v0.1.4-0.20210103172416-02ae01ce06d8
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.1
	github.com/shengdoushi/base58 v1.0.0
	github.com/sirupsen/logrus v1.8.1
	github.com/sonr-io/blockchain v0.0.9
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/starport v0.19.5
	github.com/warpfork/go-testmark v0.9.0 // indirect
	github.com/whyrusleeping/cbor-gen v0.0.0-20210219115102-f37d292932f2 // indirect
	go.buf.build/grpc/go/sonr-io/blockchain v1.3.1
	go.buf.build/grpc/go/sonr-io/core v1.3.11
	go.uber.org/goleak v1.1.11 // indirect
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/ipfs/go-ipfs-blockstore => github.com/ipfs/go-ipfs-blockstore v1.1.2
