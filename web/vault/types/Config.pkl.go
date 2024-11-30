// Code generated from Pkl module `sonr.motr.DWN`. DO NOT EDIT.
package types

type Config struct {
	IpfsGatewayUrl string `pkl:"ipfsGatewayUrl" json:"ipfsGatewayUrl,omitempty"`

	MotrToken string `pkl:"motrToken" json:"motrToken,omitempty"`

	MotrAddress string `pkl:"motrAddress" json:"motrAddress,omitempty"`

	SonrApiUrl string `pkl:"sonrApiUrl" json:"sonrApiUrl,omitempty"`

	SonrRpcUrl string `pkl:"sonrRpcUrl" json:"sonrRpcUrl,omitempty"`

	SonrChainId string `pkl:"sonrChainId" json:"sonrChainId,omitempty"`

	VaultSchema *Schema `pkl:"vaultSchema" json:"vaultSchema,omitempty"`
}
