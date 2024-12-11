// Code generated from Pkl module `sonr.conf.Motr`. DO NOT EDIT.
package motr

type Config struct {
	IpfsGatewayUrl string `pkl:"ipfsGatewayUrl" json:"ipfsGatewayUrl,omitempty"`

	MotrToken string `pkl:"motrToken" json:"motrToken,omitempty"`

	MotrAddress string `pkl:"motrAddress" json:"motrAddress,omitempty"`

	SonrApiUrl string `pkl:"sonrApiUrl" json:"sonrApiUrl,omitempty"`

	SonrRpcUrl string `pkl:"sonrRpcUrl" json:"sonrRpcUrl,omitempty"`

	SonrChainId string `pkl:"sonrChainId" json:"sonrChainId,omitempty"`

	VaultSchema *Schema `pkl:"vaultSchema" json:"vaultSchema,omitempty"`
}
