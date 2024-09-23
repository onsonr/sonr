package types

import (
	"cosmossdk.io/x/nft"
)

func (n DIDNamespace) ID() string {
	switch n {
	case DIDNamespace_DID_NAMESPACE_DWN:
		return "did:dwn:0"
	case DIDNamespace_DID_NAMESPACE_SONR:
		return "did:sonr:0"
	case DIDNamespace_DID_NAMESPACE_BITCOIN:
		return "did:btc:0"
	case DIDNamespace_DID_NAMESPACE_ETHEREUM:
		return "did:eth:0"
	case DIDNamespace_DID_NAMESPACE_IBC:
		return "did:ibc:0"
	case DIDNamespace_DID_NAMESPACE_WEBAUTHN:
		return "did:authn:0"
	case DIDNamespace_DID_NAMESPACE_SERVICE:
		return "did:web:0"
	case DIDNamespace_DID_NAMESPACE_IPFS:
		return "did:ipfs:0"
	}
	return ""
}

func (n DIDNamespace) Name() string {
	switch n {
	case DIDNamespace_DID_NAMESPACE_DWN:
		return "DecentralizedWebNode"
	case DIDNamespace_DID_NAMESPACE_SONR:
		return "SonrNetwork"
	case DIDNamespace_DID_NAMESPACE_BITCOIN:
		return "BitcoinNetwork"
	case DIDNamespace_DID_NAMESPACE_ETHEREUM:
		return "EthereumNetwork"
	case DIDNamespace_DID_NAMESPACE_IBC:
		return "IBCNetwork"
	case DIDNamespace_DID_NAMESPACE_WEBAUTHN:
		return "WebAuthentication"
	case DIDNamespace_DID_NAMESPACE_SERVICE:
		return "DecentrlizedService"
	case DIDNamespace_DID_NAMESPACE_IPFS:
		return "IPFSStorage"
	}
	return ""
}

func (n DIDNamespace) Symbol() string {
	switch n {
	case DIDNamespace_DID_NAMESPACE_DWN:
		return "DWN"
	case DIDNamespace_DID_NAMESPACE_SONR:
		return "SONR"
	case DIDNamespace_DID_NAMESPACE_BITCOIN:
		return "BTC"
	case DIDNamespace_DID_NAMESPACE_ETHEREUM:
		return "ETH"
	case DIDNamespace_DID_NAMESPACE_IBC:
		return "IBC"
	case DIDNamespace_DID_NAMESPACE_WEBAUTHN:
		return "WEBAUTHN"
	case DIDNamespace_DID_NAMESPACE_SERVICE:
		return "SERVICE"
	case DIDNamespace_DID_NAMESPACE_IPFS:
		return "IPFS"
	}
	return ""
}

func (n DIDNamespace) Description() string {
	switch n {
	case DIDNamespace_DID_NAMESPACE_DWN:
		return "DWN Service Provider"
	case DIDNamespace_DID_NAMESPACE_SONR:
		return "Sonr Network Gateway"
	case DIDNamespace_DID_NAMESPACE_BITCOIN:
		return "Bitcoin Network Gateway"
	case DIDNamespace_DID_NAMESPACE_ETHEREUM:
		return "Ethereum Network Gateway"
	case DIDNamespace_DID_NAMESPACE_IBC:
		return "IBC Network Gateway"
	case DIDNamespace_DID_NAMESPACE_WEBAUTHN:
		return "Web Authentication Key"
	case DIDNamespace_DID_NAMESPACE_SERVICE:
		return "Decentrlized Service"
	case DIDNamespace_DID_NAMESPACE_IPFS:
		return "Data Storage on IPFS"
	}
	return ""
}

func (n DIDNamespace) GetNFTClass() *nft.Class {
	return &nft.Class{
		Id:          n.ID(),
		Name:        n.Name(),
		Symbol:      n.Symbol(),
		Description: n.Description(),
	}
}
