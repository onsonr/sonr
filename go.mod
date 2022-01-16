module github.com/sonr-io/core

go 1.16

require (
	git.mills.io/prologic/bitcask v1.0.0
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/gabriel-vasile/mimetype v1.3.1
	github.com/golang/protobuf v1.5.2
	github.com/google/open-location-code/go v0.0.0-20210504205230-1796878d947c
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/kataras/golog v0.1.7
	github.com/libp2p/go-libp2p v0.15.1
	github.com/libp2p/go-libp2p-connmgr v0.2.4
	github.com/libp2p/go-libp2p-core v0.9.0
	github.com/libp2p/go-libp2p-discovery v0.5.1
	github.com/libp2p/go-libp2p-kad-dht v0.13.1
	github.com/libp2p/go-libp2p-pubsub v0.5.4
	github.com/libp2p/go-msgio v0.0.6
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-multiaddr v0.4.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/sonr-io/sonr v0.0.1
	github.com/spf13/viper v1.9.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/starport v0.19.1
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/tools v0.1.8-0.20211022200916-316ba0b74098 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
