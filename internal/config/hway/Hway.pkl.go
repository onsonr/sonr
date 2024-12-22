// Code generated from Pkl module `sonr.net.Hway`. DO NOT EDIT.
package hway

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Hway interface {
	GetServePort() int

	GetSqliteFile() string

	GetChainId() string

	GetIpfsGatewayUrl() string

	GetSonrApiUrl() string

	GetSonrGrpcUrl() string

	GetSonrRpcUrl() string

	GetPsqlDSN() string

	GetTurnstileSiteKey() string
}

var _ Hway = (*HwayImpl)(nil)

type HwayImpl struct {
	ServePort int `pkl:"servePort"`

	SqliteFile string `pkl:"sqliteFile"`

	ChainId string `pkl:"chainId"`

	IpfsGatewayUrl string `pkl:"ipfsGatewayUrl"`

	SonrApiUrl string `pkl:"sonrApiUrl"`

	SonrGrpcUrl string `pkl:"sonrGrpcUrl"`

	SonrRpcUrl string `pkl:"sonrRpcUrl"`

	PsqlDSN string `pkl:"psqlDSN"`

	TurnstileSiteKey string `pkl:"turnstileSiteKey"`
}

func (rcv *HwayImpl) GetServePort() int {
	return rcv.ServePort
}

func (rcv *HwayImpl) GetSqliteFile() string {
	return rcv.SqliteFile
}

func (rcv *HwayImpl) GetChainId() string {
	return rcv.ChainId
}

func (rcv *HwayImpl) GetIpfsGatewayUrl() string {
	return rcv.IpfsGatewayUrl
}

func (rcv *HwayImpl) GetSonrApiUrl() string {
	return rcv.SonrApiUrl
}

func (rcv *HwayImpl) GetSonrGrpcUrl() string {
	return rcv.SonrGrpcUrl
}

func (rcv *HwayImpl) GetSonrRpcUrl() string {
	return rcv.SonrRpcUrl
}

func (rcv *HwayImpl) GetPsqlDSN() string {
	return rcv.PsqlDSN
}

func (rcv *HwayImpl) GetTurnstileSiteKey() string {
	return rcv.TurnstileSiteKey
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Hway
func LoadFromPath(ctx context.Context, path string) (ret Hway, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Hway
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Hway, error) {
	var ret HwayImpl
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
