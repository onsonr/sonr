// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

type Client interface {
	GetChainId() string

	GetKeyringBackend() string

	GetOutput() string

	GetNode() string

	GetBroadcastMode() string

	GetApiUrl() string

	GetAddressPrefix() string
}
