// Code generated from Pkl module `dwn`. DO NOT EDIT.
package config

type Config struct {
	IpfsGatewayUrl string `pkl:"ipfsGatewayUrl" json:"ipfsGatewayUrl,omitempty"`

	MotrKeyshare string `pkl:"motrKeyshare" json:"motrKeyshare,omitempty"`

	MotrAddress string `pkl:"motrAddress" json:"motrAddress,omitempty"`

	SonrApiUrl string `pkl:"sonrApiUrl" json:"sonrApiUrl,omitempty"`

	SonrRpcUrl string `pkl:"sonrRpcUrl" json:"sonrRpcUrl,omitempty"`

	SonrChainId string `pkl:"sonrChainId" json:"sonrChainId,omitempty"`

	VaultSchema *Schema `pkl:"vaultSchema" json:"vaultSchema,omitempty"`
}
