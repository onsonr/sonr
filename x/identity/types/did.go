package types

import (
	"encoding/hex"
	fmt "fmt"

	"github.com/okx/go-wallet-sdk/coins/cosmos"

	commonv1 "github.com/didao-org/sonr/api/common/v1"
)

// GetCoinTypeDIDMethod returns the DID method for a given coin type
func GetCoinTypeDIDMethod(coinType commonv1.CoinType) string {
	switch coinType {
	case commonv1.CoinType_COIN_TYPE_ATOM:
		return "did:atom:"
	case commonv1.CoinType_COIN_TYPE_AXELAR:
		return "did:axlr:"
	case commonv1.CoinType_COIN_TYPE_BITCOIN:
		return "did:btcr:"
	case commonv1.CoinType_COIN_TYPE_ETHEREUM:
		return "did:ethr:"
	case commonv1.CoinType_COIN_TYPE_EVMOS:
		return "did:evmos:"
	case commonv1.CoinType_COIN_TYPE_FILECOIN:
		return "did:filr:"
	case commonv1.CoinType_COIN_TYPE_JUNO:
		return "did:juno:"
	case commonv1.CoinType_COIN_TYPE_OSMO:
		return "did:osmos:"
	case commonv1.CoinType_COIN_TYPE_SOLANA:
		return "did:sol:"
	case commonv1.CoinType_COIN_TYPE_SONR:
		return "did:sonr:"
	case commonv1.CoinType_COIN_TYPE_STARGAZE:
		return "did:starz:"
	default:
		return "did:unspecified:"
	}
}

// GetDIDMethodCoinType returns the coin type for a given DID method
func GetDIDMethodCoinType(didMethod string) (commonv1.CoinType, error) {
	switch didMethod {
	case "did:atom:":
		return commonv1.CoinType_COIN_TYPE_ATOM, nil
	case "did:axlr:":
		return commonv1.CoinType_COIN_TYPE_AXELAR, nil
	case "did:btcr:":
		return commonv1.CoinType_COIN_TYPE_BITCOIN, nil
	case "did:ethr:":
		return commonv1.CoinType_COIN_TYPE_ETHEREUM, nil
	case "did:evmos:":
		return commonv1.CoinType_COIN_TYPE_EVMOS, nil
	case "did:filr:":
		return commonv1.CoinType_COIN_TYPE_FILECOIN, nil
	case "did:juno:":
		return commonv1.CoinType_COIN_TYPE_JUNO, nil
	case "did:osmos:":
		return commonv1.CoinType_COIN_TYPE_OSMO, nil
	case "did:sol:":
		return commonv1.CoinType_COIN_TYPE_SOLANA, nil
	case "did:sonr:":
		return commonv1.CoinType_COIN_TYPE_SONR, nil
	case "did:starz:":
		return commonv1.CoinType_COIN_TYPE_STARGAZE, nil
	default:
		return commonv1.CoinType_COIN_TYPE_UNSPECIFIED, fmt.Errorf("unsupported DID method")
	}
}

// GetCoinTypeHRP returns the HRP for a given coin type
func GetCoinTypeHRP(coinType commonv1.CoinType) string {
	switch coinType {
	case commonv1.CoinType_COIN_TYPE_ATOM:
		return "cosmos"
	case commonv1.CoinType_COIN_TYPE_AXELAR:
		return "axelar"
	case commonv1.CoinType_COIN_TYPE_BITCOIN:
		return "bc"
	case commonv1.CoinType_COIN_TYPE_ETHEREUM:
		return "eth"
	case commonv1.CoinType_COIN_TYPE_EVMOS:
		return "evmos"
	case commonv1.CoinType_COIN_TYPE_FILECOIN:
		return "fil"
	case commonv1.CoinType_COIN_TYPE_JUNO:
		return "juno"
	case commonv1.CoinType_COIN_TYPE_OSMO:
		return "osmo"
	case commonv1.CoinType_COIN_TYPE_SOLANA:
		return "sol"
	case commonv1.CoinType_COIN_TYPE_SONR:
		return "idx"
	case commonv1.CoinType_COIN_TYPE_STARGAZE:
		return "starz"
	default:
		return "0x"
	}
}

func GetAddressByPublicKey(pubKey []byte, coinType commonv1.CoinType) (string, error) {
	hrp := GetCoinTypeHRP(coinType)
	pubKeyHex := hex.EncodeToString(pubKey)
	return cosmos.GetAddressByPublicKey(pubKeyHex, hrp)
}
