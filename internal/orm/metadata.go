package orm

import (
	"context"
	"encoding/json"

	"github.com/onsonr/hway/crypto"
	fs "github.com/onsonr/hway/internal/vfs"
	"github.com/ipfs/boxo/path"
)

// Metadata represents metadata of a resource
type Metadata struct {
	PublicKey              crypto.PublicKey `json:"publicKey"`
	Path                   path.Path        `json:"path"`
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

// CreateMetadata returns a new Metadata
func CreateMetadata(ctx context.Context) *Metadata {
	return &Metadata{}
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
