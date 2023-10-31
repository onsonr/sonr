package didcontroller

import (
	"context"
	"fmt"

	"github.com/highlight/highlight/sdk/highlight-go"

	"github.com/sonrhq/sonr/internal/crypto"
	"github.com/sonrhq/sonr/pkg/didcommon"
	"github.com/sonrhq/sonr/pkg/didwallets"
	identitytypes "github.com/sonrhq/sonr/x/identity/types"
)

// Account returns the controller account
func (c *SonrController) Account() *identitytypes.ControllerAccount {
	return c.account
}

// GetPrimaryWallet returns the primary wallet sonr account
func (c *SonrController) GetPrimaryWallet() *didwallets.SonrAccount {
	return c.primary
}

// CreateWallet creates a new wallet for the given coin type
func (c *SonrController) CreateWallet(ct crypto.CoinType) (*crypto.AccountData, error) {
	ctx := context.Background()
	secKey, err := c.Authenticator.DIDSecretKey(c.email)
	if err != nil {
		highlight.RecordError(ctx, err)
		return nil, err
	}
	switch ct {
	case crypto.BTCCoinType:
		bacc, err := didwallets.NewBitcoinAccount(secKey)
		if err != nil {
			highlight.RecordError(ctx, err)
			return nil, err
		}
		did := didcommon.NewUrl(bacc.Method(), didcommon.Identifier(bacc.Address())).String()
		c.ID.AppendKeyList("wallets", did)
		c.account.Wallets = append(c.account.Wallets, did)
		return bacc.Info(), nil
	case crypto.ETHCoinType:
		eacc, err := didwallets.NewEthereumAccount(secKey)
		if err != nil {
			highlight.RecordError(ctx, err)
			return nil, err
		}
		did := didcommon.NewUrl(eacc.Method(), didcommon.Identifier(eacc.Address())).String()
		c.ID.AppendKeyList("wallets", did)
		c.account.Wallets = append(c.account.Wallets, did)
		return eacc.Info(), nil
	default:
		return nil, fmt.Errorf("unsupported coin type: %s", ct)
	}
}

// GetWallet returns the wallet info for the given DID
func (c *SonrController) GetWallet(did string) (*crypto.AccountData, error) {
	wallet, err := c.useWallet(did)
	if err != nil {
		return nil, err
	}
	return wallet.Info(), nil
}

// ListWallets returns the list of wallets for the controller
func (c *SonrController) ListWallets() []string {
	return c.ID.GetKeyList("wallets")
}

// SignWithWallet signs the given message with the wallet
func (c *SonrController) SignWithWallet(did string, msg []byte) ([]byte, error) {
	wallet, err := c.useWallet(did)
	if err != nil {
		return nil, err
	}
	return wallet.Sign(msg)
}

// VerifyWithWallet verifies the given signature with the wallet
func (c *SonrController) VerifyWithWallet(did string, msg []byte, sig []byte) (bool, error) {
	wallet, err := c.useWallet(did)
	if err != nil {
		return false, err
	}
	return wallet.Verify(msg, sig)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                 Helper Methods                                 ||
// ! ||--------------------------------------------------------------------------------||

// useWallet returns the wallet account for the given DID
func (c *SonrController) useWallet(did string) (didcommon.WalletAccount, error) {
	ctx := context.Background()
	secKey, err := c.Authenticator.DIDSecretKey(c.email)
	if err != nil {
		highlight.RecordError(ctx, err)
		return nil, err
	}
	m, _, err := didcommon.Parse(did)
	if err != nil {
		return nil, err
	}
	switch m {
	case didwallets.BitcoinMethod:
		bacc, err := didwallets.ResolveBitcoinAccount(did, secKey)
		if err != nil {
			highlight.RecordError(ctx, err)
			return nil, err
		}
		return bacc, nil
	case didwallets.EthereumMethod:
		eacc, err := didwallets.ResolveEthereumAccount(did, secKey)
		if err != nil {
			highlight.RecordError(ctx, err)
			return nil, err
		}
		return eacc, nil
	default:
		return nil, fmt.Errorf("unsupported method: %s", m)
	}
}
