package types

import (
	"encoding/json"
	"fmt"
	"math/big"

	sonrcrypto "github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/types/crypto"
	"github.com/sonrhq/kryptology/pkg/core/curves"
	"github.com/sonrhq/kryptology/pkg/core/protocol"
	dklsv1 "github.com/sonrhq/kryptology/pkg/tecdsa/dkls/v1"

	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/sonrhq/kryptology/pkg/tecdsa/dkls/v1/dkg"
	"golang.org/x/crypto/sha3"
)

var kDefaultCurve = curves.K256()
var kDefaultVersion uint = protocol.Version1

type Keyshare struct {
	// The `dkgResultMessage` field is a pointer to a `protocol.Message` object that contains the DKG result message.
	Output *protocol.Message `json:"output"`

	// The `role` field is a `KeyshareRole` object that indicates the role of the keyshare.
	Role KeyshareRole `json:"role"`
}

func NewAliceKeyshare(dkgResultMsg *protocol.Message) *Keyshare {
	return &Keyshare{
		Output: dkgResultMsg,
		Role:   KeyshareRolePublic,
	}
}

func NewBobKeyshare(dkgResultMsg *protocol.Message) *Keyshare {
	return &Keyshare{
		Output: dkgResultMsg,
		Role:   KeyshareRoleUser,
	}
}

func (ks *Keyshare) FormatAddress(ct sonrcrypto.CoinType) (string, error) {
	spk, err := ks.PubKey()
	if err != nil {
		return "", fmt.Errorf("error getting alice public key: %v", err)
	}
	return ct.FormatAddress(crypto.NewSecp256k1PubKey(spk)), nil
}

func (ks *Keyshare) FormatDID(ct sonrcrypto.CoinType) (string, error) {
	spk, err := ks.PubKey()
	if err != nil {
		return "", fmt.Errorf("error getting alice public key: %v", err)
	}
	did, _ := ct.FormatDID(crypto.NewSecp256k1PubKey(spk))
	return did, nil
}

func (ks *Keyshare) GetAliceDKGResult() (*dkg.AliceOutput, error) {
	return dklsv1.DecodeAliceDkgResult(ks.Output)
}

func (ks *Keyshare) GetBobDKGResult() (*dkg.BobOutput, error) {
	return dklsv1.DecodeBobDkgResult(ks.Output)
}

func (ks *Keyshare) MarshalPublic() ([]byte, error) {
	dkgResult, err := ks.GetAliceDKGResult()
	if err != nil {
		return nil, fmt.Errorf("error getting alice dkg result: %v", err)
	}
	msg, err := dklsv1.EncodeAliceDkgOutput(dkgResult, protocol.Version1)
	if err != nil {
		return nil, fmt.Errorf("error encoding alice dkg result: %v", err)
	}
	return json.Marshal(msg)
}

func (ks *Keyshare) MarshalPrivate() ([]byte, error) {
	dkgResult, err := ks.GetBobDKGResult()
	if err != nil {
		return nil, fmt.Errorf("error getting bob dkg result: %v", err)
	}
	msg, err := dklsv1.EncodeBobDkgOutput(dkgResult, protocol.Version1)
	if err != nil {
		return nil, fmt.Errorf("error encoding bob dkg result: %v", err)
	}
	return json.Marshal(msg)
}

func (ks *Keyshare) UnmarshalPrivate(bz []byte) error {
	var msg protocol.Message
	if err := json.Unmarshal(bz, &msg); err != nil {
		return fmt.Errorf("error unmarshalling keyshare: %v", err)
	}
	if _, err := ks.GetBobDKGResult(); err != nil {
		return fmt.Errorf("error getting bob dkg result: %v", err)
	}
	ks.Output = &msg
	return nil
}

func (ks *Keyshare) UnmarshalPublic(bz []byte) error {
	var msg protocol.Message
	if err := json.Unmarshal(bz, &msg); err != nil {
		return fmt.Errorf("error unmarshalling keyshare: %v", err)
	}
	ks.Output = &msg
	if _, err := ks.GetAliceDKGResult(); err != nil {
		return fmt.Errorf("error getting alice dkg result: %v", err)
	}
	return nil
}

// PubKey returns the public key of the keyshare as a secp256k1.PubKey
func (ks *Keyshare) PubKey() (*secp256k1.PubKey, error) {
	buildSecp256k1 := func(bz []byte) *secp256k1.PubKey {
		return &secp256k1.PubKey{Key: bz}
	}
	if ks.Role.isAlice() {
		dkgResult, err := ks.GetAliceDKGResult()
		if err != nil {
			return nil, fmt.Errorf("error getting alice dkg result: %v", err)
		}
		return buildSecp256k1(dkgResult.PublicKey.ToAffineCompressed()), nil
	} else {
		dkgResult, err := ks.GetBobDKGResult()
		if err != nil {
			return nil, fmt.Errorf("error getting bob dkg result: %v", err)
		}
		return buildSecp256k1(dkgResult.PublicKey.ToAffineCompressed()), nil
	}
}

// PublicPoint returns the public key of the keyshare as a *curves.EcPoint
func (ks *Keyshare) PublicPoint() (*curves.EcPoint, error) {
	buildEcPoint := func(bz []byte) (*curves.EcPoint, error) {
		x := new(big.Int).SetBytes(bz[1:33])
		y := new(big.Int).SetBytes(bz[33:])
		ecCurve, err := kDefaultCurve.ToEllipticCurve()
		if err != nil {
			return nil, fmt.Errorf("error converting curve: %v", err)
		}
		return &curves.EcPoint{X: x, Y: y, Curve: ecCurve}, nil
	}

	if ks.Role.isAlice() {
		dkgResult, err := ks.GetAliceDKGResult()
		if err != nil {
			return nil, fmt.Errorf("error getting alice dkg result: %v", err)
		}
		return buildEcPoint(dkgResult.PublicKey.ToAffineUncompressed())
	} else {
		dkgResult, err := ks.GetBobDKGResult()
		if err != nil {
			return nil, fmt.Errorf("error getting bob dkg result: %v", err)
		}
		return buildEcPoint(dkgResult.PublicKey.ToAffineUncompressed())
	}
}

// Verify returns true if the signature is valid for the keyshare
func (ks *Keyshare) Verify(msg []byte, sigBz []byte) (bool, error) {
	sig, err := DeserializeECDSASecp256k1Signature(sigBz)
	if err != nil {
		return false, fmt.Errorf("error deserializing signature: %v", err)
	}
	hash := sha3.New256()
	_, err = hash.Write(msg)
	if err != nil {
		return false, fmt.Errorf("error hashing message: %v", err)
	}
	digest := hash.Sum(nil)
	publicKey, err := ks.PublicPoint()
	if err != nil {
		return false, fmt.Errorf("error getting public key: %v", err)
	}
	return curves.VerifyEcdsa(publicKey, digest[:], sig), nil
}
