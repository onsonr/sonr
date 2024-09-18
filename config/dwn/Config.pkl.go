// Code generated from Pkl module `dwn`. DO NOT EDIT.
package dwn

type Config struct {
	Ipfs *IPFS `pkl:"ipfs" json:"ipfs,omitempty"`

	Sonr *Sonr `pkl:"sonr" json:"sonr,omitempty"`

	Keyshare *string `pkl:"keyshare" json:"keyshare,omitempty"`

	Address *string `pkl:"address" json:"address,omitempty"`

	Vault *IndexedDB `pkl:"vault" json:"vault,omitempty"`
}
