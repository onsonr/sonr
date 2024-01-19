package did

import (
	"fmt"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
)

// GetCoinTypeDIDMethod returns the DID method for a given coin type
func GetCoinTypeDIDMethod(coinType modulev1.CoinType) string {
	switch coinType {
	case modulev1.CoinType_COIN_TYPE_ATOM:
		return "did:atom:"
	case modulev1.CoinType_COIN_TYPE_AXELAR:
		return "did:axlr:"
	case modulev1.CoinType_COIN_TYPE_BITCOIN:
		return "did:btcr:"
	case modulev1.CoinType_COIN_TYPE_ETHEREUM:
		return "did:ethr:"
	case modulev1.CoinType_COIN_TYPE_EVMOS:
		return "did:evmos:"
	case modulev1.CoinType_COIN_TYPE_FILECOIN:
		return "did:filr:"
	case modulev1.CoinType_COIN_TYPE_JUNO:
		return "did:juno:"
	case modulev1.CoinType_COIN_TYPE_OSMO:
		return "did:osmos:"
	case modulev1.CoinType_COIN_TYPE_SOLANA:
		return "did:sol:"
	case modulev1.CoinType_COIN_TYPE_SONR:
		return "did:sonr:"
	case modulev1.CoinType_COIN_TYPE_STARGAZE:
		return "did:starz:"
	default:
		return "did:unspecified:"
	}
}

// GetDIDMethodCoinType returns the coin type for a given DID method
func GetDIDMethodCoinType(didMethod string) (modulev1.CoinType, error) {
	switch didMethod {
	case "did:atom:":
		return modulev1.CoinType_COIN_TYPE_ATOM, nil
	case "did:axlr:":
		return modulev1.CoinType_COIN_TYPE_AXELAR, nil
	case "did:btcr:":
		return modulev1.CoinType_COIN_TYPE_BITCOIN, nil
	case "did:ethr:":
		return modulev1.CoinType_COIN_TYPE_ETHEREUM, nil
	case "did:evmos:":
		return modulev1.CoinType_COIN_TYPE_EVMOS, nil
	case "did:filr:":
		return modulev1.CoinType_COIN_TYPE_FILECOIN, nil
	case "did:juno:":
		return modulev1.CoinType_COIN_TYPE_JUNO, nil
	case "did:osmos:":
		return modulev1.CoinType_COIN_TYPE_OSMO, nil
	case "did:sol:":
		return modulev1.CoinType_COIN_TYPE_SOLANA, nil
	case "did:sonr:":
		return modulev1.CoinType_COIN_TYPE_SONR, nil
	case "did:starz:":
		return modulev1.CoinType_COIN_TYPE_STARGAZE, nil
	default:
		return modulev1.CoinType_COIN_TYPE_UNSPECIFIED, fmt.Errorf("unsupported DID method")
	}
}

// GetCoinTypeHRP returns the HRP for a given coin type
func GetCoinTypeHRP(coinType modulev1.CoinType) string {
	switch coinType {
	case modulev1.CoinType_COIN_TYPE_ATOM:
		return "cosmos"
	case modulev1.CoinType_COIN_TYPE_AXELAR:
		return "axelar"
	case modulev1.CoinType_COIN_TYPE_BITCOIN:
		return "bc"
	case modulev1.CoinType_COIN_TYPE_ETHEREUM:
		return "eth"
	case modulev1.CoinType_COIN_TYPE_EVMOS:
		return "evmos"
	case modulev1.CoinType_COIN_TYPE_FILECOIN:
		return "fil"
	case modulev1.CoinType_COIN_TYPE_JUNO:
		return "juno"
	case modulev1.CoinType_COIN_TYPE_OSMO:
		return "osmo"
	case modulev1.CoinType_COIN_TYPE_SOLANA:
		return "sol"
	case modulev1.CoinType_COIN_TYPE_SONR:
		return "idx"
	case modulev1.CoinType_COIN_TYPE_STARGAZE:
		return "starz"
	default:
		return "0x"
	}
}
