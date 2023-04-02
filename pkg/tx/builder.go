package tx

import (
	"fmt"
	"math/big"

	"cosmossdk.io/math"
	"github.com/sonrhq/core/internal/models"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/tx/cosmos"
	"github.com/sonrhq/core/pkg/tx/eth"
)

type SonrTxBuilder interface {
	SendTokens(to string, amount int) ([]byte, error)
}

type sonrTxBuilder struct {
	chainID string
	acc     models.Account
	ct      crypto.CoinType
}

func NewSonrTxBuilder(chainID string, acc models.Account) SonrTxBuilder {
	return &sonrTxBuilder{
		chainID: chainID,
		acc:     acc,
		ct:      acc.CoinType(),
	}
}

func (stb *sonrTxBuilder) SendTokens(to string, amount int) ([]byte, error) {
	// Ethereum transaction
	if stb.ct.IsEthereum() {
		return eth.SignEthereumTransaction(stb.acc, to, big.NewInt(int64(amount)))
	}

	// Cosmos transaction
	if stb.ct.IsCosmos() || stb.ct.IsSonr() {
		return cosmos.SignTransaction(stb.acc, to, math.NewInt(int64(amount)), stb.ct.String())
	}

	return nil, fmt.Errorf("unsupported coin type: %s", stb.ct)
}
