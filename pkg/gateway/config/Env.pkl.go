// Code generated from Pkl module `sonr.hway.Env`. DO NOT EDIT.
package config

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Env interface {
	GetServePort() int

	GetConfigDir() string

	GetSqliteFile() string

	GetChainId() string

	GetIpfsGatewayUrl() string

	GetSonrApiUrl() string

	GetSonrGrpcUrl() string

	GetSonrRpcUrl() string
}

var _ Env = (*EnvImpl)(nil)

type EnvImpl struct {
	ServePort int `pkl:"servePort"`

	ConfigDir string `pkl:"configDir"`

	SqliteFile string `pkl:"sqliteFile"`

	ChainId string `pkl:"chainId"`

	IpfsGatewayUrl string `pkl:"ipfsGatewayUrl"`

	SonrApiUrl string `pkl:"sonrApiUrl"`

	SonrGrpcUrl string `pkl:"sonrGrpcUrl"`

	SonrRpcUrl string `pkl:"sonrRpcUrl"`
}

func (rcv *EnvImpl) GetServePort() int {
	return rcv.ServePort
}

func (rcv *EnvImpl) GetConfigDir() string {
	return rcv.ConfigDir
}

func (rcv *EnvImpl) GetSqliteFile() string {
	return rcv.SqliteFile
}

func (rcv *EnvImpl) GetChainId() string {
	return rcv.ChainId
}

func (rcv *EnvImpl) GetIpfsGatewayUrl() string {
	return rcv.IpfsGatewayUrl
}

func (rcv *EnvImpl) GetSonrApiUrl() string {
	return rcv.SonrApiUrl
}

func (rcv *EnvImpl) GetSonrGrpcUrl() string {
	return rcv.SonrGrpcUrl
}

func (rcv *EnvImpl) GetSonrRpcUrl() string {
	return rcv.SonrRpcUrl
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Env
func LoadFromPath(ctx context.Context, path string) (ret Env, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Env
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Env, error) {
	var ret EnvImpl
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
