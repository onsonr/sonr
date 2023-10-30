package types

import (
	"encoding/json"
	"fmt"

	"github.com/sonr-io/kryptology/pkg/core/protocol"
	dklsv1 "github.com/sonr-io/kryptology/pkg/tecdsa/dkls/v1"

	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"sonr.io/core/internal/crypto"
	"golang.org/x/crypto/sha3"
)

type EncKeyshareSet struct {
	Public    *Keyshare `json:"public"`
	Encrypted []byte    `json:"user"`
}

// FormatDID returns the DID of the account based on the coin type
func (kss *EncKeyshareSet) FormatDID(ct crypto.CoinType) string {
	did, err := kss.Public.FormatDID(ct)
	if err != nil {
		panic(err)
	}
	return did
}

// The `PublicKey()` function is a method of the `KeyshareSet` type. It returns the public key corresponding to Alice's keyshare in the keyshare set. It does this by calling the `PubKey()` method of the `Keyshare` object corresponding to Alice's keyshare. If the keyshare set is not
// valid or if there is an error in retrieving the public key, it returns an error.
func (kss *EncKeyshareSet) PublicKey() *secp256k1.PubKey {
	pub, err := kss.Public.PubKey()
	if err != nil {
		panic(err)
	}
	return pub
}

// GetAccountData returns the proto representation of the account
func (kss *EncKeyshareSet) GetAccountData(ct crypto.CoinType) *crypto.AccountData {
	dat, err := crypto.NewDefaultAccountData(ct, crypto.NewSecp256k1PubKey(kss.PublicKey()))
	if err != nil {
		panic(err)
	}
	return dat
}

// Marshal returns the JSON encoding of the EncKeyshareSet.
func (kss *EncKeyshareSet) Marshal() ([]byte, error) {
	return json.Marshal(kss)
}

// Unmarshal parses the JSON-encoded data and stores the result in the EncKeyshareSet.
func (kss *EncKeyshareSet) Unmarshal(bz []byte) error {
	return json.Unmarshal(bz, kss)
}

func (kss *EncKeyshareSet) DecryptUserKeyshare(key crypto.EncryptionKey) (KeyshareSet, error) {
	alice := kss.Public
	if alice == nil {
		return EmptyKeyshareSet(), fmt.Errorf("alice keyshare is nil")
	}
	bz, err := key.Decrypt(kss.Encrypted)
	if err != nil {
		return EmptyKeyshareSet(), fmt.Errorf("error decrypting keyshare: %v", err)
	}
	// Deserialize keyshare
	msg := &protocol.Message{}
	if err := json.Unmarshal(bz, msg); err != nil {
		return EmptyKeyshareSet(), fmt.Errorf("error unmarshalling keyshare: %v", err)
	}
	bob := NewBobKeyshare(msg)
	return KeyshareSet{
		alice,
		bob,
	}, nil
}

type KeyshareSet [2]*Keyshare

// The function returns an empty KeyshareSet.
func EmptyKeyshareSet() KeyshareSet {
	return KeyshareSet{nil, nil}
}

// The function creates a new KeyshareSet object using the provided Alice and Bob DKG result messages.
func NewKSS(a *Keyshare, b *Keyshare) KeyshareSet {
	return KeyshareSet{
		a,
		b,
	}
}

// The function creates a new KeyshareSet object using the provided Alice and Bob DKG result messages.
func NewKeyshareSet(aliceDkgResultMsg *protocol.Message, bobDkgResultMsg *protocol.Message) KeyshareSet {
	return KeyshareSet{
		NewAliceKeyshare(aliceDkgResultMsg),
		NewBobKeyshare(bobDkgResultMsg),
	}
}

// The `Alice()` function is a method of the `KeyshareSet` type. It returns the `Keyshare` object corresponding to Alice's keyshare in the keyshare set.
func (kss KeyshareSet) Alice() *Keyshare {
	a := kss[0]
	if a == nil {
		panic("alice keyshare is nil")
	}
	return a
}

// The `Bob()` function is a method of the `KeyshareSet` type. It returns the `Keyshare` object corresponding to Bob's keyshare in the keyshare set.
func (kss KeyshareSet) Bob() *Keyshare {
	b := kss[1]
	if b == nil {
		panic("bob keyshare is nil")
	}
	return b
}

// The `DKGAtIndex` function is a method of the `KeyshareSet` type. It takes an integer `i` as input and returns the DKG (Distributed Key Generation) result message at the specified index.
func (kss KeyshareSet) DKGAtIndex(i int) *protocol.Message {
	if i == 0 {
		return kss.Alice().Output
	} else if i == 1 {
		return kss.Bob().Output
	} else {
		fmt.Println("DKGAtIndex(): invalid index")
		return nil
	}
}

// FormatAddress returns the address of the account based on the coin type
func (kss KeyshareSet) FormatAddress(ct crypto.CoinType) string {
	ad, err := kss.Alice().FormatAddress(ct)
	if err != nil {
		panic(err)
	}
	return ad
}

// FormatDID returns the DID of the account based on the coin type
func (kss KeyshareSet) FormatDID(ct crypto.CoinType) string {
	did, err := kss.Alice().FormatDID(ct)
	if err != nil {
		panic(err)
	}
	return did
}

// GetAccountData returns the proto representation of the account
func (wa KeyshareSet) GetAccountData(ct crypto.CoinType) *crypto.AccountData {
	dat, err := crypto.NewDefaultAccountData(ct, crypto.NewSecp256k1PubKey(wa.PublicKey()))
	if err != nil {
		panic(err)
	}
	return dat
}

// The `IsValid()` function is a method of the `KeyshareSet` type. It checks if the `KeyshareSet` object is valid by performing the following checks:
func (kss KeyshareSet) IsValid() error {
	if len(kss) != 2 {
		return fmt.Errorf("keyshare set must have exactly 2 keyshares")
	}
	alice := kss[0]
	if alice == nil {
		return fmt.Errorf("alice keyshare is nil")
	}
	bob := kss[1]
	if bob == nil {
		return fmt.Errorf("bob keyshare is nil")
	}
	return nil
}

// The `PublicKey()` function is a method of the `KeyshareSet` type. It returns the public key corresponding to Alice's keyshare in the keyshare set. It does this by calling the `PubKey()` method of the `Keyshare` object corresponding to Alice's keyshare. If the keyshare set is not
// valid or if there is an error in retrieving the public key, it returns an error.
func (kss KeyshareSet) PublicKey() *secp256k1.PubKey {
	pub, err := kss.Alice().PubKey()
	if err != nil {
		panic(err)
	}
	return pub
}

// The `Sign` function is a method of the `KeyshareSet` type. It takes a byte slice `msg` as input and returns a byte slice and an error.
func (kss KeyshareSet) Sign(msg []byte) ([]byte, error) {
	if err := kss.IsValid(); err != nil {
		return nil, fmt.Errorf("error validating keyshare set: %v", err)
	}
	aliceSign, err := dklsv1.NewAliceSign(kDefaultCurve, sha3.New256(), msg, kss.Alice().Output, protocol.Version1)
	if err != nil {
		return nil, fmt.Errorf("error creating Alice sign: %v", err)
	}
	bobSign, err := dklsv1.NewBobSign(kDefaultCurve, sha3.New256(), msg, kss.Bob().Output, protocol.Version1)
	if err != nil {
		return nil, fmt.Errorf("error creating Bob sign: %v", err)
	}

	aErr, bErr := RunIteratedProtocol(aliceSign, bobSign)
	if aErr != protocol.ErrProtocolFinished || bErr != protocol.ErrProtocolFinished {
		return nil, fmt.Errorf("error running protocol: aErr=%v, bErr=%v", aErr, bErr)
	}

	resultMessage, err := bobSign.Result(protocol.Version1)
	if err != nil {
		return nil, fmt.Errorf("error getting result: %v", err)
	}

	result, err := dklsv1.DecodeSignature(resultMessage)
	if err != nil {
		return nil, fmt.Errorf("error decoding signature: %v", err)
	}
	sigBytes, err := SerializeECDSASecp256k1Signature(result)
	if err != nil {
		return nil, fmt.Errorf("error serializing signature: %v", err)
	}
	return sigBytes, nil
}

// The `Verify` function is a method of the `KeyshareSet` type. It takes a byte slice `msg` and a byte slice `sigBz` as input and returns a boolean value and an error.
func (kss KeyshareSet) Verify(msg []byte, sigBz []byte) (bool, error) {
	if err := kss.IsValid(); err != nil {
		return false, fmt.Errorf("error validating keyshare set: %v", err)
	}
	return kss[0].Verify(msg, sigBz)
}

func (kss KeyshareSet) EncryptUserKeyshare(c crypto.EncryptionKey) (*EncKeyshareSet, error) {
	if err := kss.IsValid(); err != nil {
		return nil, fmt.Errorf("error validating keyshare set: %v", err)
	}
	bz, err := kss.Bob().MarshalPrivate()
	if err != nil {
		return nil, fmt.Errorf("error marshaling bob keyshare: %v", err)
	}
	enc, err := c.Encrypt(bz)
	if err != nil {
		return nil, fmt.Errorf("error encrypting keyshare: %v", err)
	}
	return &EncKeyshareSet{
		Public:    kss.Alice(),
		Encrypted: enc,
	}, nil
}
