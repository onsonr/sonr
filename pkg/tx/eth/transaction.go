package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
"github.com/sonrhq/core/internal/protocol/packages/controller"
)

// EthereumTransaction represents the Ethereum transaction data
type EthereumTransaction struct {
	Nonce    uint64
	To       string
	Value    *big.Int
	GasLimit uint64
	GasPrice *big.Int
	Data     []byte
}

// SignEthereumTransaction signs an Ethereum transaction using the wallet account abstraction
func SignEthereumTransaction(wa controller.Account, to string, amount *big.Int) ([]byte, error) {
	// Set default gas limit and gas price
	defaultGasLimit := uint64(21000)
	defaultGasPrice := big.NewInt(20000000000) // 20 Gwei

	// Create EthereumTransaction
	etx := &EthereumTransaction{
		Nonce:    0,
		To:       to,
		Value:    amount,
		GasLimit: defaultGasLimit,
		GasPrice: defaultGasPrice,
		Data:     []byte{},
	}

	// Serialize the Ethereum transaction data
	txData := types.NewTransaction(etx.Nonce, common.HexToAddress(etx.To), etx.Value, etx.GasLimit, etx.GasPrice, etx.Data)
	encodedTx, err := rlp.EncodeToBytes(txData)
	if err != nil {
		return nil, err
	}

	// Sign the serialized transaction data using the existing Sign method
	signedTx, err := wa.Sign(encodedTx)
	if err != nil {
		return nil, err
	}

	// Add the Ethereum-specific ECDSA recovery id (v) to the signature
	recId := byte(27)
	if crypto.VerifySignature(wa.PubKey().Bytes(), encodedTx, signedTx[:64]) {
		recId = byte(28)
	}

	signedTx[64] = recId - byte(27)
	return signedTx, nil
}
