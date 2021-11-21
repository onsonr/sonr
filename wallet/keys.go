package wallet

import (
	"errors"
	"fmt"

	"github.com/hyperledger/aries-framework-go/pkg/kms"
	"github.com/hyperledger/aries-framework-go/pkg/wallet"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/device"
	"github.com/sonr-io/core/internal/did"
	"google.golang.org/protobuf/proto"

	cpb "github.com/libp2p/go-libp2p-core/crypto/pb"
)

// NewKeyPair creates a new keypair
func NewKeyPair(kt kms.KeyType, label string) (*wallet.KeyPair, error) {
	kp, err := Instance.CreateKeyPair(kt)
	if err != nil {
		return nil, err
	}
	err = Info.AddInfo(kp.KeyID, kt, TokenType_SNR, label)
	if err != nil {
		return nil, err
	}
	return kp, nil
}

// createDefaultKeys creates the default keys
func createDefaultKeys(sname string) error {
	if device.Wallet.Exists(WALLET_INFO_FILE_NAME) {
		// Load Metadata
		infoBuf, err := device.Wallet.ReadFile(WALLET_INFO_FILE_NAME)
		if err != nil {
			logger.Errorf("Failed to read wallet info file: %s", err.Error())
			return err
		}
		winf := &WalletInfo{}
		err = proto.Unmarshal(infoBuf, winf)
		if err != nil {
			logger.Errorf("Failed to unmarshal wallet info file: %s", err.Error())
			return err
		}
		Info = winf
		return nil
	}

	// Create Sonr Device key
	id, pubbuf, err := Provider.KMS().CreateAndExportPubKeyBytes(kms.ED25519Type)
	if err != nil {
		return err
	}

	pubkey, err := crypto.UnmarshalEd25519PublicKey(pubbuf)
	if err != nil {
		return err
	}

	didDoc, err := did.NewDoc(pubkey, sname)
	if err != nil {
		return err
	}

	didJson, err := didDoc.MarshalJSON()
	if err != nil {
		return err
	}

	logger.Infof("Created DID: %s", string(didJson))

	// Create new wallet info
	Info = &WalletInfo{
		SName: sname,
		Keys:  make(map[string]*KeypairInfo),
	}

	err = Info.AddInfo(id, kms.ED25519Type, TokenType_SNR, "device")
	if err != nil {
		return err
	}
	return nil
}

func (info *WalletInfo) AddInfo(id string, kt kms.KeyType, tt TokenType, label string) error {
	i := &KeypairInfo{
		Id:        id,
		KeyType:   fmt.Sprint(kt),
		TokenType: tt,
		Label:     label,
	}
	info.Keys[id] = i
	buf, err := proto.Marshal(info)
	if err != nil {
		return err
	}
	return device.Wallet.WriteFile(WALLET_INFO_FILE_NAME, buf)
}

func (info *WalletInfo) FindInfoByLabel(label string) []*KeypairInfo {
	var kpis []*KeypairInfo
	for _, v := range info.Keys {
		if v.Label == label {
			kpis = append(kpis, v)
		}
	}
	return kpis
}

func (info *WalletInfo) FindInfoByToken(tt TokenType) []*KeypairInfo {
	var kpis []*KeypairInfo
	for _, v := range info.Keys {
		if v.TokenType == tt {
			kpis = append(kpis, v)
		}
	}
	return kpis
}

func (i *KeypairInfo) LoadPrivKH() (interface{}, error) {
	privFace, err := Provider.KMS().Get(i.Id)
	if err != nil {
		return nil, err
	}
	return privFace, nil
}

func (i *KeypairInfo) LoadPubKey() (crypto.PubKey, error) {
	pubBuf, err := Provider.KMS().ExportPubKeyBytes(i.Id)
	if err != nil {
		return nil, err
	}
	return crypto.UnmarshalEd25519PublicKey(pubBuf)
}

// DevicePubKey returns the device public key
func DevicePubKey() (crypto.PubKey, error) {
	return Info.FindInfoByLabel("device")[0].LoadPubKey()
}

// DevicePrivKH returns the device private key
func DevicePrivKH() (interface{}, error) {
	return Info.FindInfoByLabel("device")[0].LoadPrivKH()
}

// DevicePrivKey returns the device private key
func DevicePrivKey() (crypto.PrivKey, error) {
	return createPrivKey(Info.FindInfoByLabel("device")[0])
}

// sonrPrivKey implements crypto.PrivKey
type sonrPrivKey struct {
	crypto.PrivKey
	kh interface{}
	i  *KeypairInfo
}

// createPrivKey creates a private key from a keypair info
func createPrivKey(i *KeypairInfo) (crypto.PrivKey, error) {
	privFace, err := i.LoadPrivKH()
	if err != nil {
		return nil, err
	}
	return &sonrPrivKey{
		kh: privFace,
		i:  i,
	}, nil
}

// Sign implements crypto.PrivKey Sign
func (k *sonrPrivKey) Sign(msg []byte) ([]byte, error) {
	return Provider.Crypto().Sign(msg, k.kh)
}

// GetPublic implements crypto.PrivKey GetPublic
func (k *sonrPrivKey) GetPublic() crypto.PubKey {
	pub, _ := DevicePubKey()
	return pub
}

// Equals implements crypto.PrivKey Equals
func (k *sonrPrivKey) Equals(k2 crypto.Key) bool {
	return k.GetPublic().Equals(k2)
}

// Bytes implements crypto.PrivKey Bytes
func (k *sonrPrivKey) Raw() ([]byte, error) {
	return nil, errors.New("not implemented")
}

// Type implements crypto.PrivKey Type
func (k *sonrPrivKey) Type() cpb.KeyType {
	return cpb.KeyType_Ed25519
}
