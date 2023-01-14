package fs

import (
	"fmt"
	"strconv"

	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/common/crypto"
	"github.com/sonr-hq/sonr/pkg/vault/internal/mpc"
	"github.com/sonr-hq/sonr/pkg/vault/internal/network"
)

// Storing the share of the wallet.
// Index at -1 is for root share, otherwise it is for a derived share.
func (c *VaultConfig) StoreShare(share []byte, partyId string, index int) error {
	boxer, err := c.newBoxer()
	if err != nil {
		return err
	}
	enc, err := boxer.Seal(share)
	if err != nil {
		return err
	}
	if index == -1 {
		err := c.authDir.WriteFile(partyId, enc)
		if err != nil {
			return err
		}
	} else {
		newDir, err := c.authDir.CreateFolder(strconv.Itoa(index))
		if err != nil {
			return err
		}
		err = newDir.WriteFile(partyId, enc)
		if err != nil {
			return err
		}
	}
	return nil
}

// StoreOfflineWallet stores the offline wallet in the vault.
func (c *VaultConfig) StoreOfflineWallet(wallet network.OfflineWallet) error {
	for _, share := range wallet.List() {
		bz, err := share.Marshal()
		if err != nil {
			return err
		}
		err = c.StoreShare(bz, string(share.SelfID()), share.Index())
		if err != nil {
			return err
		}
	}
	return nil
}

// Getting the share of the wallet.
// Index at -1 is for root share, otherwise it is for a derived share.
func (c *VaultConfig) GetShare(partyId string, index int) ([]byte, error) {
	folderList, err := c.authDir.ListFolders()
	if err != nil {
		return nil, err
	}

	boxer, err := c.newBoxer()
	if err != nil {
		return nil, err
	}
	if index == -1 {
		if c.authDir.Exists(partyId) {
			enc, err := c.authDir.ReadFile(partyId)
			if err != nil {
				return nil, err
			}
			dec, err := boxer.Open(enc)
			if err != nil {
				return nil, err
			}
			return dec, nil
		}
	} else {
		for _, folder := range folderList {
			if folder.Name() == strconv.Itoa(index) {
				if folder.Exists(partyId) {
					enc, err := folder.ReadFile(partyId)
					if err != nil {
						return nil, err
					}
					dec, err := boxer.Open(enc)
					if err != nil {
						return nil, err
					}
					return dec, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("share not found for party " + partyId + " and index " + strconv.Itoa(index))
}

// GetOfflineWallet gets the offline wallet from the vault.
func (c *VaultConfig) GetOfflineWallet() (network.OfflineWallet, error) {
	var shares []crypto.WalletShare
	fileList, err := c.authDir.ListFiles()
	if err != nil {
		return nil, err
	}
	for _, file := range fileList {
		bz, err := c.GetShare(file, -1)
		if err != nil {
			return nil, err
		}
		share, err := common.NewWalletShare(bz)
		err = share.Unmarshal(bz)
		if err != nil {
			return nil, err
		}

		walletShare, err := mpc.LoadWalletShare(share)
		if err != nil {
			return nil, err
		}
		shares = append(shares, walletShare)
	}
	return network.OfflineWallet(shares), nil
}
