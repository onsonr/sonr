package network

import (
	"sync"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/vault/internal/mpc"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// OfflineWallet is a slice of WalletShare
type OfflineWallet []common.WalletShare

// Address returns the address of the Wallet.
func (ws OfflineWallet) Address() string {
	return ws[0].Address()
}

// Bip32Derive creates a new WalletShare that is derived from the given path.
func (ws OfflineWallet) Bip32Derive(i uint32) (common.WalletShare, error) {
	return ws[0].Bip32Derive(i)
}

// EncryptKey
func (ws OfflineWallet) EncryptKey() ([]byte, error) {
	return ws.Sign("vault", []byte(ws.Address()))
}

// Find returns the WalletShare with the given ID.
func (ws OfflineWallet) Find(id party.ID) common.WalletShare {
	for _, w := range ws {
		if w.SelfID() == id {
			return w
		}
	}
	return nil
}

// Creating a map of party.ID to cmp.Config.
func (ws OfflineWallet) GetConfigMap() map[party.ID]*cmp.Config {
	configMap := make(map[party.ID]*cmp.Config)
	for _, w := range ws {
		configMap[w.SelfID()] = w.CMPConfig()
	}
	return configMap
}

// Network creates a new offline network from a list of wallet shares.
func (ws OfflineWallet) Network() common.Network {
	parties := make(party.IDSlice, 0, len(ws))
	for _, s := range ws {
		parties = append(parties, s.SelfID())
	}
	return NewOfflineNetwork(parties)
}

func (ws OfflineWallet) PublicKey() (*secp256k1.PubKey, error) {
	return ws[0].PublicKey()
}

// Refreshing the wallet shares.
func (ws OfflineWallet) Refresh(current party.ID) (common.Wallet, error) {
	var mtx sync.Mutex
	net := ws.Network()
	configs := ws.GetConfigMap()

	var wg sync.WaitGroup
	for _, id := range net.Ls() {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			conf, err := mpc.CmpRefresh(configs[id], net, &wg, pl)
			if err != nil {
				return
			}

			mtx.Lock()
			configs[conf.ID] = conf
			mtx.Unlock()
		}(id)
	}
	wg.Wait()
	shares := make([]common.WalletShare, 0)
	for _, conf := range configs {
		shares = append(shares, mpc.NewWalletShare("snr", conf))
	}
	return OfflineWallet(shares), nil
}

// Signing a message with the current party.ID.
func (ws OfflineWallet) Sign(current party.ID, m []byte) ([]byte, error) {
	net := ws.Network()
	configs := ws.GetConfigMap()
	doneChan := make(chan *ecdsa.Signature)
	var wg sync.WaitGroup
	for _, id := range net.Ls() {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			sig, err := mpc.CmpSign(configs[id], m, net.Ls(), net, &wg, pl)
			if err != nil {
				return
			}
			if id == current {
				doneChan <- sig
			}
		}(id)
	}
	wg.Wait()
	return mpc.SerializeSignature(<-doneChan)
}

// SignTx constructs a Tx with all required info for a wallet to sign.
func (ws OfflineWallet) SignTx(current party.ID, msgs ...sdk.Msg) ([]byte, error) {
	txb, err := buildTx(msgs...)
	if err != nil {
		return nil, err
	}

	ai, err := getAuthInfoSingle(ws.Find(current), 2)
	if err != nil {
		return nil, err
	}

	sigDocBz, err := getSignDocBytes(ai, txb)
	if err != nil {
		return nil, err
	}

	sig, err := ws.Sign(current, sigDocBz)
	if err != nil {
		return nil, err
	}
	return createRawTxBytes(txb, sig, ai)
}

// Verify a signature with all the WalletShares in the slice.
func (ws OfflineWallet) Verify(msg, sig []byte) bool {
	for _, w := range ws {
		if w.Verify(msg, sig) {
			return true
		}
	}
	return false
}

//
// Helper functions
//

// buildTx builds a transaction from the given inputs.
func buildTx(msgs ...sdk.Msg) (*txtypes.TxBody, error) {
	// func BuildTx(w *crypto.MPCWallet, msgs ...sdk.Msg) (*txtypes.TxBody, error) {
	// Create Any for each message
	anyMsgs := make([]*codectypes.Any, len(msgs))
	for i, m := range msgs {
		msg, err := codectypes.NewAnyWithValue(m)
		if err != nil {
			return nil, err
		}
		anyMsgs[i] = msg
	}

	// Create TXRaw and Marshal
	txBody := txtypes.TxBody{
		Messages: anyMsgs,
	}
	return &txBody, nil
}

// createRawTxBytes is a helper function to create a raw raw transaction and Marshal it to bytes
func createRawTxBytes(txBody *txtypes.TxBody, sig []byte, authInfo *txtypes.AuthInfo) ([]byte, error) {
	// Serialize the tx body
	txBytes, err := txBody.Marshal()
	if err != nil {
		return nil, err
	}

	// Serialize the authInfo
	authInfoBytes, err := authInfo.Marshal()
	if err != nil {
		return nil, err
	}

	// Create a signature list and append the signature
	sigList := make([][]byte, 1)
	sigList[0] = sig

	// Create Raw TX
	txRaw := &txtypes.TxRaw{
		BodyBytes:     txBytes,
		AuthInfoBytes: authInfoBytes,
		Signatures:    sigList,
	}

	// Marshal the txRaw
	return txRaw.Marshal()
}

// getAuthInfoSingle returns the authentication information for the given message.
func getAuthInfoSingle(w common.WalletShare, gas int) (*txtypes.AuthInfo, error) {
	// Get PublicKey
	pubKey, err := w.PublicKey()
	if err != nil {
		return nil, err
	}

	// Build signerInfo parameters
	anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
	if err != nil {
		return nil, err
	}

	// Create AuthInfo
	authInfo := txtypes.AuthInfo{
		SignerInfos: []*txtypes.SignerInfo{
			{
				PublicKey: anyPubKey,
				ModeInfo: &txtypes.ModeInfo{
					Sum: &txtypes.ModeInfo_Single_{
						Single: &txtypes.ModeInfo_Single{
							Mode: 1,
						},
					},
				},
				Sequence: 1,
			},
		},
		Fee: &txtypes.Fee{
			Amount:   sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(int64(gas)))),
			GasLimit: uint64(300000),
			Granter:  w.Address(),
		},
	}
	return &authInfo, nil
}

// It takes a transaction body and auth info, serializes them, and then creates a SignDoc object that
// contains the serialized transaction body and auth info, and the chain ID
func getSignDocBytes(authInfo *txtypes.AuthInfo, txBody *txtypes.TxBody) ([]byte, error) {
	// Serialize the transaction body.
	txBodyBz, err := txBody.Marshal()
	if err != nil {
		return nil, err
	}

	// Serialize the auth info.
	authInfoBz, err := authInfo.Marshal()
	if err != nil {
		return nil, err
	}

	// Create SignDoc
	signDoc := &txtypes.SignDoc{
		BodyBytes:     txBodyBz,
		AuthInfoBytes: authInfoBz,
		ChainId:       "sonr",
	}
	return signDoc.Marshal()
}
