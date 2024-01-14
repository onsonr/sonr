package coins

import (
	"fmt"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
)

// GetCoinTypeDIDMethod returns the DID method for a given coin type
func GetCoinTypeDIDMethod(coinType modulev1.CoinType) (string, error) {
	switch coinType {
	case modulev1.CoinType_COIN_TYPE_ATOM:
		return "did:atom:", nil
	case modulev1.CoinType_COIN_TYPE_AXELAR:
		return "did:axlr:", nil
	case modulev1.CoinType_COIN_TYPE_BITCOIN:
		return "did:btcr:", nil
	case modulev1.CoinType_COIN_TYPE_ETHEREUM:
		return "did:ethr:", nil
	case modulev1.CoinType_COIN_TYPE_EVMOS:
		return "did:evmos:", nil
	case modulev1.CoinType_COIN_TYPE_FILECOIN:
		return "did:filr:", nil
	case modulev1.CoinType_COIN_TYPE_JUNO:
		return "did:juno:", nil
	case modulev1.CoinType_COIN_TYPE_OSMO:
		return "did:osmos:", nil
	case modulev1.CoinType_COIN_TYPE_SOLANA:
		return "did:sol:", nil
	case modulev1.CoinType_COIN_TYPE_SONR:
		return "did:sonr:", nil
	case modulev1.CoinType_COIN_TYPE_STARGAZE:
		return "did:starz:", nil
	default:
		return "", fmt.Errorf("unsupported coin type")
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
