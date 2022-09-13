module github.com/sonr-io/sonr

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.45.5
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/fxamacker/cbor/v2 v2.4.0
	github.com/gabriel-vasile/mimetype v1.4.0
	github.com/gin-gonic/gin v1.7.7
	github.com/gogo/protobuf v1.3.3
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/google/open-location-code/go v0.0.0-20220120191843-cafb35c0d74d
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hsanjuan/ipfs-lite v1.4.0
	github.com/ignite-hq/cli v0.22.0
	github.com/ipfs/go-cid v0.2.0
	github.com/ipfs/go-datastore v0.5.1
	github.com/ipfs/go-ipfs-api v0.3.0
	github.com/ipfs/go-ipns v0.1.2
	github.com/ipld/go-ipld-prime v0.16.0
	github.com/kataras/go-events v0.0.3
	github.com/kataras/golog v0.1.7
	github.com/lestrrat-go/jwx/v2 v2.0.3
	github.com/libp2p/go-libp2p v0.20.1
	github.com/libp2p/go-libp2p-connmgr v0.4.0
	github.com/libp2p/go-libp2p-core v0.16.1
	github.com/libp2p/go-libp2p-discovery v0.7.0
	github.com/libp2p/go-libp2p-kad-dht v0.16.0
	github.com/libp2p/go-libp2p-pubsub v0.7.0
	github.com/libp2p/go-msgio v0.2.0
	github.com/marstr/guid v1.1.0
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-multiaddr v0.5.0
	github.com/ockam-network/did v0.1.3
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.2
	github.com/shengdoushi/base58 v1.0.0
	github.com/sonr-io/keyring v1.2.2-0.20220701165800-b4428a8aad16
	github.com/sonr-io/multi-party-sig v0.7.0
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.2
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe
	github.com/swaggo/gin-swagger v1.5.0
	github.com/tendermint/spn v0.2.1-0.20220609194312-7833ecf4454a
	github.com/tendermint/starport v0.19.5
	github.com/tendermint/tendermint v0.34.19
	github.com/tendermint/tm-db v0.6.7
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
	golang.org/x/mobile v0.0.0-20220518205345-8578da9835fd
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f // indirect
	golang.org/x/sys v0.0.0-20220909162455-aba9fc2a8ff2 // indirect
	golang.org/x/term v0.0.0-20220722155259-a9ba230a4035 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/99designs/keyring => github.com/sonr-io/keyring v1.2.2-0.20220701001003-9f5deb6c197a
	github.com/elastic/gosigar => github.com/sonr-io/gosigar v0.14.3-0.20220630215707-0c344b060447
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
