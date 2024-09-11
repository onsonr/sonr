package builder

import (
	"github.com/onsonr/crypto"
	"github.com/onsonr/sonr/x/did/types"
)

type Signer interface {
	Sign(msg []byte) ([]byte, error)
	Verify(msg []byte, sig []byte) error
	PublicKey() []byte
}

type signer struct {
	user *types.Keyshare
	val  *types.Keyshare
}

func (k signer) Sign(msg []byte) ([]byte, error) {
	valSignFunc, err := crypto.GetSignFunc(k.val, msg)
	if err != nil {
		return nil, err
	}
	usrSignFunc, err := crypto.GetSignFunc(k.user, msg)
	if err != nil {
		return nil, err
	}
	sig, err := crypto.RunMPCSign(valSignFunc, usrSignFunc)
	if err != nil {
		return nil, err
	}
	return crypto.SerializeMPCSignature(sig)
}

func (k signer) Verify(msg []byte, sig []byte) error {
	sigMpc, err := crypto.DeserializeMPCSignature(sig)
	if err != nil {
		return err
	}
	pk, err := crypto.GetECDSAPublicKey(k.val)
	if err != nil {
		return err
	}
	ok := crypto.VerifyMPCSignature(sigMpc, msg, pk)
	if !ok {
		return types.ErrInvalidSignature
	}
	return nil
}

func (k signer) PublicKey() []byte {
	if k.user != nil {
		return k.user.PublicKey
	}
	if k.val != nil {
		return k.val.PublicKey
	}
	return nil
}
