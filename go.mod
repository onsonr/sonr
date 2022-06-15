module github.com/sonr-io/sonr

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.45.5
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/docker/docker v20.10.16+incompatible // indirect
	github.com/frankban/quicktest v1.14.3 // indirect
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
	github.com/ipfs/go-cid v0.2.0
	github.com/ipfs/go-datastore v0.5.1
	github.com/ipld/go-ipld-prime v0.16.0
	github.com/kataras/go-events v0.0.3
	github.com/kataras/golog v0.1.7
	github.com/lestrrat-go/jwx v1.2.25
	github.com/lib/pq v1.10.5 // indirect
	github.com/libp2p/go-libp2p v0.20.1
	github.com/libp2p/go-libp2p-connmgr v0.4.0
	github.com/libp2p/go-libp2p-core v0.16.1
	github.com/libp2p/go-libp2p-discovery v0.7.0
	github.com/libp2p/go-libp2p-kad-dht v0.16.0
	github.com/libp2p/go-libp2p-mplex v0.8.0 // indirect
	github.com/libp2p/go-libp2p-pubsub v0.7.0
	// github.com/libp2p/go-libp2p-webrtc-direct v0.0.0-20220406105718-41a3151f0a37
	github.com/libp2p/go-msgio v0.2.0
	github.com/miekg/dns v1.1.49 // indirect
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-multiaddr v0.5.0
	github.com/ockam-network/did v0.1.3
	github.com/onsi/gomega v1.17.0 // indirect
	github.com/opencontainers/image-spec v1.0.3-0.20211202183452-c5a74bcca799 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.2
	github.com/sergi/go-diff v1.2.0 // indirect
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
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/mobile v0.0.0-20220518205345-8578da9835fd
	golang.org/x/net v0.0.0-20220524220425-1d687d428aca // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/term v0.0.0-20220526004731-065cf7ba2467 // indirect
	golang.org/x/text v0.3.8-0.20211004125949-5bd84dd9b33b // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v2 v2.4.0
	nhooyr.io/websocket v1.8.7 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
