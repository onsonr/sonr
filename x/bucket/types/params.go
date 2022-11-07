package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/joho/godotenv"
	"github.com/sonr-io/sonr/internal/projectpath"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

const (
	IPFS_GATEWAY_ADDR = "https://ipfs.sonr.ws"
	IPFS_API_ADDR     = "https://api.ipfs.sonr.ws"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	gateway_addr := IPFS_GATEWAY_ADDR
	api_addr := IPFS_API_ADDR
	env_path := filepath.Join(projectpath.Root, ".env")
	err := godotenv.Load(env_path)
	if err != nil {
		return Params{IpfsGateway: gateway_addr,
			IptsApiUrl: api_addr,
		}
	}
	return Params{IpfsGateway: os.Getenv("IPFS_ADDRESS"),
		IptsApiUrl: os.Getenv("IPFS_API_ADDRESS"),
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
