package mpc

import (
	genericecdsa "crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/onsonr/sonr/crypto/core/curves"
	"github.com/onsonr/sonr/crypto/core/protocol"
	"golang.org/x/crypto/sha3"
)

type (
	ExportedKeyset = []byte
)

type Keyset interface {
	Address() string
	Val() *ValKeyshare
	ValJSON() string
	User() *UserKeyshare
	UserJSON() string
}

type keyset struct {
	val  *ValKeyshare
	user *UserKeyshare
	addr string
}

func (k keyset) Address() string {
	return k.addr
}

func (k keyset) Val() *ValKeyshare {
	return k.val
}

func (k keyset) User() *UserKeyshare {
	return k.user
}

func (k keyset) ValJSON() string {
	return k.val.String()
}

func (k keyset) UserJSON() string {
	return k.user.String()
}

func ComputeIssuerDID(pk []byte) (string, string, error) {
	addr, err := ComputeSonrAddr(pk)
	if err != nil {
		return "", "", err
	}
	return fmt.Sprintf("did:sonr:%s", addr), addr, nil
}

func ComputeSonrAddr(pk []byte) (string, error) {
	sonrAddr, err := bech32.ConvertAndEncode("idx", pk)
	if err != nil {
		return "", err
	}
	return sonrAddr, nil
}

// For DKG bob starts first. For refresh and sign, Alice starts first.
func RunProtocol(firstParty protocol.Iterator, secondParty protocol.Iterator) (error, error) {
	var (
		message *protocol.Message
		aErr    error
		bErr    error
	)

	for aErr != protocol.ErrProtocolFinished || bErr != protocol.ErrProtocolFinished {
		// Crank each protocol forward one iteration
		message, bErr = firstParty.Next(message)
		if bErr != nil && bErr != protocol.ErrProtocolFinished {
			return nil, bErr
		}

		message, aErr = secondParty.Next(message)
		if aErr != nil && aErr != protocol.ErrProtocolFinished {
			return aErr, nil
		}
	}
	return aErr, bErr
}

// ComputeEcPoint builds an elliptic curve point from a compressed byte slice
func ComputeEcPoint(pubKey []byte) (*curves.EcPoint, error) {
	crv := curves.K256()
	x := new(big.Int).SetBytes(pubKey[1:33])
	y := new(big.Int).SetBytes(pubKey[33:])
	ecCurve, err := crv.ToEllipticCurve()
	if err != nil {
		return nil, fmt.Errorf("error converting curve: %v", err)
	}
	return &curves.EcPoint{X: x, Y: y, Curve: ecCurve}, nil
}

func ComputeEcdsaPublicKey(pubKey []byte) (*genericecdsa.PublicKey, error) {
	pk, err := ComputeEcPoint(pubKey)
	if err != nil {
		return nil, err
	}
	return &genericecdsa.PublicKey{
		Curve: pk.Curve,
		X:     pk.X,
		Y:     pk.Y,
	}, nil
}

// VerifySignature verifies the signature of a message
func VerifySignature(pk []byte, msg []byte, sig []byte) (bool, error) {
	pp, err := ComputeEcPoint(pk)
	if err != nil {
		return false, err
	}
	sigEd, err := DeserializeSignature(sig)
	if err != nil {
		return false, err
	}
	hash := sha3.New256()
	_, err = hash.Write(msg)
	if err != nil {
		return false, err
	}
	digest := hash.Sum(nil)
	return curves.VerifyEcdsa(pp, digest[:], sigEd), nil
}

func checkIteratedErrors(aErr, bErr error) error {
	if aErr == protocol.ErrProtocolFinished && bErr == protocol.ErrProtocolFinished {
		return nil
	}
	if aErr != protocol.ErrProtocolFinished {
		return aErr
	}
	if bErr != protocol.ErrProtocolFinished {
		return bErr
	}
	return nil
}
