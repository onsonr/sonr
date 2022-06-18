module github.com/sonr-io/sonr

go 1.16

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)

require (
	github.com/cosmos/cosmos-sdk v0.45.5
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/fxamacker/cbor/v2 v2.4.0
	github.com/gabriel-vasile/mimetype v1.4.0
	github.com/gin-gonic/gin v1.7.7
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/open-location-code/go v0.0.0-20220120191843-cafb35c0d74d
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hsanjuan/ipfs-lite v1.4.0
	github.com/ignite-hq/cli v0.22.0
	github.com/ipfs/go-cid v0.1.0
	github.com/ipfs/go-datastore v0.5.1
	github.com/ipld/go-ipld-prime v0.16.0
	github.com/kataras/go-events v0.0.3
	github.com/kataras/golog v0.1.7
	github.com/lestrrat-go/jwx v1.2.25
	github.com/libp2p/go-libp2p v0.19.1
	github.com/libp2p/go-libp2p-connmgr v0.3.1
	github.com/libp2p/go-libp2p-core v0.15.1
	github.com/libp2p/go-libp2p-discovery v0.6.0
	github.com/libp2p/go-libp2p-kad-dht v0.15.0
	github.com/libp2p/go-libp2p-pubsub v0.6.0
	// github.com/libp2p/go-libp2p-webrtc-direct v0.0.0-20220406105718-41a3151f0a37
	github.com/libp2p/go-msgio v0.2.0
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-multiaddr v0.5.0
	github.com/ockam-network/did v0.1.3
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/shengdoushi/base58 v1.0.0
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.1
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe
	github.com/swaggo/gin-swagger v1.5.0
	github.com/swaggo/swag v1.8.1
	github.com/tendermint/spn v0.2.1-0.20220609194312-7833ecf4454a
	github.com/tendermint/starport v0.19.5
	github.com/tendermint/tendermint v0.34.19
	github.com/tendermint/tm-db v0.6.7
	go.buf.build/grpc/go/sonr-io/blockchain v1.3.7
	go.buf.build/grpc/go/sonr-io/motor v1.3.1
	golang.org/x/mobile v0.0.0-20200801112145-973feb4309de
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/btcsuite/btcd v0.22.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/libp2p/go-libp2p-asn-util v0.2.0 // indirect
	github.com/libp2p/go-libp2p-resource-manager v0.3.0 // indirect
	github.com/libp2p/go-reuseport v0.2.0 // indirect
	github.com/libp2p/go-yamux/v3 v3.1.2 // indirect
	github.com/lucas-clemente/quic-go v0.27.1 // indirect
	github.com/multiformats/go-multistream v0.3.1 // indirect
	github.com/prometheus/client_golang v1.12.1
)
