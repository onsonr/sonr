package types

import (
	"encoding/hex"
	fmt "fmt"

	"github.com/okx/go-wallet-sdk/coins/cosmos"
)

// GetCoinTypeDIDMethod returns the DID method for a given coin type
func GetCoinTypeDIDMethod(coinType CoinType) string {
	switch coinType {
	case CoinType_COIN_TYPE_ATOM:
		return "did:atom:"
	case CoinType_COIN_TYPE_AXELAR:
		return "did:axlr:"
	case CoinType_COIN_TYPE_BITCOIN:
		return "did:btcr:"
	case CoinType_COIN_TYPE_ETHEREUM:
		return "did:ethr:"
	case CoinType_COIN_TYPE_EVMOS:
		return "did:evmos:"
	case CoinType_COIN_TYPE_FILECOIN:
		return "did:filr:"
	case CoinType_COIN_TYPE_JUNO:
		return "did:juno:"
	case CoinType_COIN_TYPE_OSMO:
		return "did:osmos:"
	case CoinType_COIN_TYPE_SOLANA:
		return "did:sol:"
	case CoinType_COIN_TYPE_SONR:
		return "did:sonr:"
	case CoinType_COIN_TYPE_STARGAZE:
		return "did:starz:"
	default:
		return "did:unspecified:"
	}
}

// GetDIDMethodCoinType returns the coin type for a given DID method
func GetDIDMethodCoinType(didMethod string) (CoinType, error) {
	switch didMethod {
	case "did:atom:":
		return CoinType_COIN_TYPE_ATOM, nil
	case "did:axlr:":
		return CoinType_COIN_TYPE_AXELAR, nil
	case "did:btcr:":
		return CoinType_COIN_TYPE_BITCOIN, nil
	case "did:ethr:":
		return CoinType_COIN_TYPE_ETHEREUM, nil
	case "did:evmos:":
		return CoinType_COIN_TYPE_EVMOS, nil
	case "did:filr:":
		return CoinType_COIN_TYPE_FILECOIN, nil
	case "did:juno:":
		return CoinType_COIN_TYPE_JUNO, nil
	case "did:osmos:":
		return CoinType_COIN_TYPE_OSMO, nil
	case "did:sol:":
		return CoinType_COIN_TYPE_SOLANA, nil
	case "did:sonr:":
		return CoinType_COIN_TYPE_SONR, nil
	case "did:starz:":
		return CoinType_COIN_TYPE_STARGAZE, nil
	default:
		return CoinType_COIN_TYPE_UNSPECIFIED, fmt.Errorf("unsupported DID method")
	}
}

// GetCoinTypeHRP returns the HRP for a given coin type
func GetCoinTypeHRP(coinType CoinType) string {
	switch coinType {
	case CoinType_COIN_TYPE_ATOM:
		return "cosmos"
	case CoinType_COIN_TYPE_AXELAR:
		return "axelar"
	case CoinType_COIN_TYPE_BITCOIN:
		return "bc"
	case CoinType_COIN_TYPE_ETHEREUM:
		return "eth"
	case CoinType_COIN_TYPE_EVMOS:
		return "evmos"
	case CoinType_COIN_TYPE_FILECOIN:
		return "fil"
	case CoinType_COIN_TYPE_JUNO:
		return "juno"
	case CoinType_COIN_TYPE_OSMO:
		return "osmo"
	case CoinType_COIN_TYPE_SOLANA:
		return "sol"
	case CoinType_COIN_TYPE_SONR:
		return "idx"
	case CoinType_COIN_TYPE_STARGAZE:
		return "starz"
	default:
		return "0x"
	}
}

func GetAddressByPublicKey(pubKey []byte, coinType CoinType) (string, error) {
	hrp := GetCoinTypeHRP(coinType)
	pubKeyHex := hex.EncodeToString(pubKey)
	return cosmos.GetAddressByPublicKey(pubKeyHex, hrp)
}
