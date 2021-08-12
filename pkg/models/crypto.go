package models

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/textileio/go-threads/core/thread"
)

// ** ─── KeyPair MANAGEMENT ────────────────────────────────────────────────────────

// Method Creates New Key Pair
func (d *Device) copyKeyPair(kp *KeyPair) (*KeyPair, *SonrError) {
	if d.HasKeys(kp.GetType()) {
		// Fetch Private, Public Keys
		privKey, pubKey := kp.PrivPubKeys()

		// Marshal Data
		privBuf, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return nil, NewError(err, ErrorEvent_MARSHAL)
		}

		// Marshal Data
		pubBuf, err := crypto.MarshalPublicKey(pubKey)
		if err != nil {
			return nil, NewError(err, ErrorEvent_MARSHAL)
		}

		// Write Private Key to File
		path, werr := d.WriteKey(privBuf, kp.GetType())
		if werr != nil {
			return nil, NewError(err, ErrorEvent_USER_SAVE)
		}

		// Set Keys
		return &KeyPair{
			Type:      kp.GetType(),
			Signature: kp.GetSignature(),
			Public: &KeyPair_Public{
				Base64: crypto.ConfigEncodeKey(pubBuf),
				Buffer: pubBuf,
			},
			Private: &KeyPair_Private{
				Path:   path,
				Buffer: privBuf,
			},
		}, nil
	} else {
		return nil, NewError(errors.New("Key Pair Already exists for this type at path. To replace delete keys first."), ErrorEvent_USER_CREATE)
	}
}

// Method to Load all keys in Device
func (d *Device) loadKeyChain() (*KeyChain, *SonrError) {
	// Load AccountKeys
	accountKeys, err := d.loadKeyPair(KeyPair_ACCOUNT)
	if err != nil {
		return nil, err
	}

	// Load Device Keys
	deviceKeys, err := d.loadKeyPair(KeyPair_DEVICE)
	if err != nil {
		return nil, err
	}

	// Load Group Keys
	groupKeys, err := d.loadKeyPair(KeyPair_GROUP)
	if err != nil {
		return nil, err
	}

	// Set Device Key Values
	d.LinkPubKey = deviceKeys.GetPublic()
	d.AccountKeys = accountKeys
	d.GroupKeys = groupKeys

	return &KeyChain{
		Account: accountKeys,
		Device:  deviceKeys,
		Group:   groupKeys,
	}, nil
}

// Method Loads Existing Key Pair
func (d *Device) loadKeyPair(t KeyPair_Type) (*KeyPair, *SonrError) {
	if d.HasKeys(t) {
		// Get PrivKey File
		privBuf, serr := d.ReadKey(t)
		if serr != nil {
			return nil, serr
		}

		// Get Private Key from Buffer
		privKey, err := crypto.UnmarshalPrivateKey(privBuf)
		if err != nil {
			return nil, NewError(err, ErrorEvent_KEY_INVALID)
		}

		// Get Public Key from Private and Marshal
		pubKey := privKey.GetPublic()
		pubBuf, err := crypto.MarshalPublicKey(pubKey)
		if err != nil {
			return nil, NewError(err, ErrorEvent_KEY_SET)
		}

		// Create Account AccountKeys
		keys := &KeyPair{
			Type:      t,
			Signature: Signature_Ed25519,
			Public: &KeyPair_Public{
				Base64: crypto.ConfigEncodeKey(pubBuf),
				Buffer: pubBuf,
			},
			Private: &KeyPair_Private{
				Path:   d.WorkingKeyPath(t),
				Buffer: privBuf,
			},
		}

		// Set Key Pair
		return keys, nil
	}
	return nil, NewError(errors.New("Keys dont exist need migration."), ErrorEvent_KEY_SET)
}

// Method to Create all keys in Device
func (d *Device) newKeyChain() (*KeyChain, *SonrError) {
	// Load AccountKeys
	accountKeys, err := d.newKeyPair(KeyPair_ACCOUNT)
	if err != nil {
		return nil, err
	}

	// Load Device Keys
	deviceKeys, err := d.newKeyPair(KeyPair_DEVICE)
	if err != nil {
		return nil, err
	}

	// Load Group Keys
	groupKeys, err := d.newKeyPair(KeyPair_GROUP)
	if err != nil {
		return nil, err
	}

	// Set Device Key Values
	d.LinkPubKey = deviceKeys.GetPublic()
	d.AccountKeys = accountKeys
	d.GroupKeys = groupKeys

	// Return Key Chain
	return &KeyChain{
		Account: accountKeys,
		Device:  deviceKeys,
		Group:   groupKeys,
	}, nil
}

// Method Creates New Key Pair
func (d *Device) newKeyPair(t KeyPair_Type) (*KeyPair, *SonrError) {
	// Create New Key
	privKey, pubKey, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, NewError(err, ErrorEvent_HOST_KEY)
	}

	// Marshal Data
	privBuf, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		return nil, NewError(err, ErrorEvent_MARSHAL)
	}

	// Marshal Data
	pubBuf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		return nil, NewError(err, ErrorEvent_MARSHAL)
	}

	// Write Private Key to File
	path, werr := d.WriteKey(privBuf, t)
	if werr != nil {
		return nil, NewError(err, ErrorEvent_USER_SAVE)
	}

	// Set Keys
	return &KeyPair{
		Type:      t,
		Signature: Signature_Ed25519,
		Public: &KeyPair_Public{
			Base64: crypto.ConfigEncodeKey(pubBuf),
			Buffer: pubBuf,
		},
		Private: &KeyPair_Private{
			Path:   path,
			Buffer: privBuf,
		},
	}, nil
}

// Method to Create all keys in Device
func (d *Device) tempKeyChain() (*KeyChain, *SonrError) {
	// Load AccountKeys
	tempKeys, err := d.newKeyPair(KeyPair_TEMPORARY)
	if err != nil {
		return nil, err
	}

	// Set Device Key Values
	d.AccountKeys = tempKeys

	// Return Key Chain
	return &KeyChain{
		Account: tempKeys,
	}, nil
}

// Method Deletes Existing Keys and Creates New Pair
func (d *Device) deleteKeyPair(t KeyPair_Type) *SonrError {
	// Delete Key Pair
	err := os.Remove(d.WorkingKeyPath(t))
	if err != nil {
		LogInfo("ERROR: " + err.Error())
		return NewError(err, ErrorEvent_USER_FS)
	}

	// Create New Key
	return nil
}

// Method Returns PeerID from Public Key
func (kp *KeyPair) ID() (peer.ID, *SonrError) {
	id, err := peer.IDFromPublicKey(kp.PubKey())
	if err != nil {
		return "", NewError(err, ErrorEvent_KEY_ID)
	}
	return id, nil
}

// Method Returns Private Key and Public Key
func (kp *KeyPair) PrivPubKeys() (crypto.PrivKey, crypto.PubKey) {
	// Get Key from Buffer
	return kp.PrivKey(), kp.PubKey()
}

// Method Returns Private Key
func (kp *KeyPair) PrivKey() crypto.PrivKey {
	// Get Key from Buffer
	key, err := crypto.UnmarshalPrivateKey(kp.GetPrivate().GetBuffer())
	if err != nil {
		return nil
	}
	return key
}

// Method Returns Private Key
func (kp *KeyPair) PrivBuffer() []byte {
	return kp.GetPrivate().GetBuffer()
}

// Method Returns Public Key
func (kp *KeyPair) PubKey() crypto.PubKey {
	// Get Key from Buffer
	privKey, err := crypto.UnmarshalPrivateKey(kp.GetPrivate().GetBuffer())
	if err != nil {
		return nil
	}
	return privKey.GetPublic()
}

// Method Returns Public Key as Base64 String
func (kp *KeyPair) PubKeyBase64() string {
	return kp.GetPublic().GetBase64()
}

// Method Signs given data and returns response
func (kp *KeyPair) Sign(value string) string {
	h := hmac.New(sha256.New, kp.PrivBuffer())
	h.Write([]byte(value))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

// Method verifies 'sig' is the signed hash of 'data'
func (kp *KeyPair) Verify(data []byte, sig []byte) (bool, error) {
	// Check for Public Key
	if pubKey := kp.PubKey(); pubKey != nil {
		result, err := pubKey.Verify(data, sig)
		if err != nil {
			return false, err
		}
		return result, nil
	}
	// Return Error
	return false, errors.New("Public Key Doesnt Exist")
}

// Method verifies pubkey is from this device
func (kp *KeyPair) VerifyPubKey(pubKey crypto.PubKey) bool {
	return kp.PrivKey().GetPublic().Equals(pubKey)
}

// Method Checks if Device has Given Key Type
func (d *Device) HasKeys(t KeyPair_Type) bool {
	if _, err := os.Stat(d.WorkingKeyPath(t)); os.IsNotExist(err) {
		return false
	}
	return true
}

// Method Checks if Device has Account Keys
func (d *Device) HasAccountKeys() bool {
	if _, err := os.Stat(d.WorkingKeyPath(KeyPair_ACCOUNT)); os.IsNotExist(err) {
		return false
	}
	return true
}

// Method Checks if Device has Device-Link Keys
func (d *Device) HasDeviceKeys() bool {
	if _, err := os.Stat(d.WorkingKeyPath(KeyPair_DEVICE)); os.IsNotExist(err) {
		return false
	}
	return true
}

// Method Checks if Device has Group Keys
func (d *Device) HasGroupKeys() bool {
	if _, err := os.Stat(d.WorkingKeyPath(KeyPair_GROUP)); os.IsNotExist(err) {
		return false
	}
	return true
}

// Method Checks if Device has Temporary Keys
func (d *Device) HasTemporaryKeys() bool {
	if _, err := os.Stat(d.WorkingKeyPath(KeyPair_TEMPORARY)); os.IsNotExist(err) {
		return false
	}
	return true
}

// Returns Short ID for this Device
func (d *Device) ShortID() string {
	// Check for Keys
	if d.HasTemporaryKeys() {
		// Write Device ID as New sha256 String
		h := hmac.New(sha256.New, d.AccountKeys.PrivBuffer())
		h.Write([]byte(d.GetId()))
		hexCode := hex.EncodeToString(h.Sum(nil))

		// Fetch Length of ID
		nLen := 0
		for i := 0; i < len(hexCode); i++ {
			if b := hexCode[i]; '0' <= b && b <= '9' {
				nLen++
			}
		}

		// Iterate Over Coded String
		var n = make([]int, 0, nLen)
		for i := 0; i < len(hexCode); i++ {
			if b := hexCode[i]; '0' <= b && b <= '9' {
				n = append(n, int(b)-'0')
			}
		}

		// Convert int array into string
		result := ""
		for _, v := range n[:6] {
			if v < 10 {
				result = result + fmt.Sprintf("%d", v)
			}
		}

		// Return Short ID
		return result
	} else {
		LogError(errors.New("Device does not have a Key Pair"))
		return ""
	}
}

// Method returns Thread Identity for Device
func (d *Device) ThreadIdentity() thread.Identity {
	return thread.NewLibp2pIdentity(d.AccountKeys.PrivKey())
}

// Returns FileName of a Given KeyPair Type
func (t KeyPair_Type) FileName() string {
	switch t {
	case KeyPair_ACCOUNT:
		return ".account_private_key"
	case KeyPair_GROUP:
		return ".group_private_key"
	case KeyPair_DEVICE:
		return ".link_private_key"
	case KeyPair_TEMPORARY:
		return ".temporary_private_key"
	}
	return ""
}
