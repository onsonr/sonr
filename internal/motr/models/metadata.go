package models

import (
	"encoding/json"

	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/accumulator"
	"github.com/di-dao/sonr/crypto/secret"
	"github.com/di-dao/sonr/pkg/fs"
	"github.com/ipfs/boxo/path"
)

// Accumulators represents accumulated data of secret data of a resource.
type Accumulators map[string][]byte

// Marshal returns the JSON encoding of the Accumulators
func (i Accumulators) Marshal() ([]byte, error) {
	return json.Marshal(i)
}

// Unmarshal parses the JSON-encoded data and stores the result in the Accumulators
func (i *Accumulators) Unmarshal(data []byte) error {
	return json.Unmarshal(data, i)
}

// GetAccumulator Unmarshals the accumulator.
func (i Accumulators) GetAccumulator(key string) (*accumulator.Accumulator, error) {
	return secret.UnmarshalAccumulator(i[key])
}

// SetAccumulator marshals the accumulator and stores it.
func (i Accumulators) SetAccumulator(key string, acc *accumulator.Accumulator) error {
	raw, err := secret.MarshalAccumulator(acc)
	if err != nil {
		return err
	}
	i[key] = raw
	return nil
}

// Metadata represents metadata of a resource
type Metadata struct {
	PublicKey              crypto.PublicKey `json:"publicKey"`
	Path                   path.Path        `json:"path"`
	Accumulators           Accumulators     `json:"accumulators"`
	Address                string           `json:"address"`
	ChainID                string           `json:"chainId"`
	ValAddress             string           `json:"valAddress"`
	RPCEndpoint            string           `json:"rpcEndpoint"`
	APIEndpoint            string           `json:"apiEndpoint"`
	IPFSEndpoint           string           `json:"ipfsEndpoint"`
	PeerID                 string           `json:"peerId"`
	SupportedDenominations []string         `json:"supportedDenominations"`
	Number                 int              `json:"number"`
}

// Marshal returns the JSON encoding of the Metadata
func (i *Metadata) Marshal() ([]byte, error) {
	return json.Marshal(i)
}

// Unmarshal parses the JSON-encoded data and stores the result in the Metadata
func (i *Metadata) Unmarshal(data []byte) error {
	return json.Unmarshal(data, i)
}

// Save writes the JSON encoding of the Metadata to the provided file
func (i *Metadata) Save(file fs.File) error {
	bz, err := i.Marshal()
	if err != nil {
		return err
	}
	return file.Write(bz)
}
