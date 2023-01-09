package fs

import "errors"

func (vfs *vaultFsImpl) SignData(data []byte) ([]byte, []byte, error) {
	return nil, nil, errors.New("Method unimplemented")
}
func (vfs *vaultFsImpl) StoreShare(share []byte, partyId string) error {
	return errors.New("Method unimplemented")
}
func (vfs *vaultFsImpl) VerifyData(data []byte, signature []byte) bool {
	return false
}

// // assembleWalletFromShares takes a WalletShareConfig and CID to return a Offline Wallet
// func (v *VaultService) assembleWalletFromShares(cid string, current *common.WalletShareConfig) (party.ID, common.Wallet, error) {
// 	// Initialize provided share
// 	shares := make([]*common.WalletShareConfig, 0)
// 	shares = append(shares, current)

// 	// Fetch Vault share from IPFS
// 	oldbz, err := v.highway.Get(cid)
// 	if err != nil {
// 		return "", nil, err
// 	}

// 	// Unmarshal share
// 	share := &common.WalletShareConfig{}
// 	err = share.Unmarshal(oldbz)
// 	if err != nil {
// 		return "", nil, err
// 	}

// 	// Load wallet
// 	wallet, err := network.LoadOfflineWallet(shares)
// 	if err != nil {
// 		return "", nil, err
// 	}
// 	return party.ID(current.SelfId), wallet, nil
// }
