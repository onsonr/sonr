package local

import (
	"os"
	"path/filepath"

	"github.com/sonrhq/core/x/identity/types"
)

// Context is a struct that holds the current context of the application.
type Context struct {
	grpcApiEndpoint        string
	rpcApiEndpoint         string
	highwayServerPort      string
	TlsCertPath            string
	TlsKeyPath             string
	GlobalKvKsStore        string
	GlobalInboxDocsStore   string
	GlobalInboxEventsStore string
	HighwayMode            string
	HomeDir                string
	NodeHome               string
	IPFSRepoPath           string
	Rendevouz              string
	BsMultiaddrs           []string
}

func NewContext() Context {
	params := types.DefaultParams()
	return Context{
		grpcApiEndpoint:        currGrpcEndpoint(),
		rpcApiEndpoint:         currRpcEndpoint(),
		highwayServerPort:      getServerPort(),
		TlsCertPath:            getTLSCert(),
		TlsKeyPath:             getTLSKey(),
		GlobalKvKsStore:        params.DidMethodName + "-" + params.DidMethodVersion + "/keyvalue#ks",
		GlobalInboxDocsStore:   params.DidMethodName + "-" + params.DidMethodVersion + "/docsstore#inbox",
		GlobalInboxEventsStore: params.DidMethodName + "-" + params.DidMethodVersion + "/eventlog#inbox",
		HighwayMode:            "fiber",
		HomeDir:                filepath.Join(getHomeDir()),
		NodeHome:               filepath.Join(getHomeDir(), ".sonr"),
		IPFSRepoPath:           filepath.Join(getHomeDir(), ".sonr", "adapters", "ipfs"),
		Rendevouz:              defaultRendezvousString,
		BsMultiaddrs:           defaultBootstrapMultiaddrs,
	}
}

func (c Context) IsDev() bool {
	return os.Getenv("ENVIRONMENT") == "dev"
}

func (c Context) IsProd() bool {
	return !c.IsDev()
}

func (c Context) HasTlsCert() bool {
	return c.TlsCertPath != "" && c.TlsKeyPath != "" && c.IsProd()
}

func (c Context) IsHighwayFiber() bool {
	return c.HighwayMode == "fiber"
}

func (c Context) IsHighwayConnect() bool {
	return c.HighwayMode == "connect"
}

func (c Context) GrpcEndpoint() string {
	return c.grpcApiEndpoint
}

func (c Context) RpcEndpoint() string {
	return c.rpcApiEndpoint
}

func (c Context) HighwayPort() string {
	return c.highwayServerPort
}
