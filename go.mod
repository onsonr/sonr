module github.com/sonr-io/sonr

go 1.16

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/ipfs/go-ipfs-blockstore => github.com/ipfs/go-ipfs-blockstore v1.1.2
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)

require (
	github.com/cosmos/cosmos-sdk v0.44.5
	github.com/cosmos/ibc-go v1.2.2
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/duo-labs/webauthn v0.0.0-20220330035159-03696f3d4499
	github.com/duo-labs/webauthn.io v0.0.0-20200929144140-c031a3e0f95d
	github.com/fxamacker/cbor/v2 v2.4.0
	github.com/gabriel-vasile/mimetype v1.4.0
	github.com/gin-gonic/gin v1.7.7
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/open-location-code/go v0.0.0-20210504205230-1796878d947c
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hsanjuan/ipfs-lite v1.3.0
	github.com/ipfs/go-cid v0.1.0
	github.com/ipfs/go-datastore v0.5.1
	github.com/ipld/go-ipld-prime v0.14.2
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
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-multiaddr v0.5.0
	github.com/ockam-network/did v0.1.4-0.20210103172416-02ae01ce06d8
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/shengdoushi/base58 v1.0.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.3.0
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.1
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.4.2
	github.com/swaggo/swag v1.8.1
	github.com/tendermint/spm v0.1.9
	github.com/tendermint/starport v0.19.5
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.4
	go.buf.build/grpc/go/sonr-io/blockchain v1.3.3
	go.buf.build/grpc/go/sonr-io/core v1.3.19
	golang.org/x/mobile v0.0.0-20220414153400-ce6a79cf6a13
	google.golang.org/genproto v0.0.0-20220317150908-0efb43f6373e
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/square/go-jose.v2 v2.6.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/google/go-cmp v0.5.7 // indirect
	golang.org/x/oauth2 v0.0.0-20220309155454-6242fa91716a // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
