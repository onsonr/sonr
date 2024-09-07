// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

type Sonr interface {
	Client

	GetRpcUrl() any
}

var _ Sonr = (*SonrImpl)(nil)

type SonrImpl struct {
	ChainId string `pkl:"chainId"`

	KeyringBackend string `pkl:"keyringBackend"`

	Output string `pkl:"output"`

	RpcUrl any `pkl:"rpcUrl"`

	BroadcastMode string `pkl:"broadcastMode"`

	ApiUrl string `pkl:"apiUrl"`

	AddressPrefix string `pkl:"addressPrefix"`

	Node string `pkl:"node"`
}

func (rcv *SonrImpl) GetChainId() string {
	return rcv.ChainId
}

func (rcv *SonrImpl) GetKeyringBackend() string {
	return rcv.KeyringBackend
}

func (rcv *SonrImpl) GetOutput() string {
	return rcv.Output
}

func (rcv *SonrImpl) GetRpcUrl() any {
	return rcv.RpcUrl
}

func (rcv *SonrImpl) GetBroadcastMode() string {
	return rcv.BroadcastMode
}

func (rcv *SonrImpl) GetApiUrl() string {
	return rcv.ApiUrl
}

func (rcv *SonrImpl) GetAddressPrefix() string {
	return rcv.AddressPrefix
}

func (rcv *SonrImpl) GetNode() string {
	return rcv.Node
}
