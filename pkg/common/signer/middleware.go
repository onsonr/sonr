package signer

import (
	"fmt"

	"github.com/onsonr/sonr/crypto/mpc"
)

func (cc *SignerContext) NewKeyset() (mpc.Keyset, error) {
	ks, err := mpc.NewKeyset()
	if err != nil {
		return nil, err
	}
	cc.keyset = ks
	cc.hasKeyset = true
	return ks, nil
}

func (cc *SignerContext) GetKeyset() (mpc.Keyset, error) {
	if !cc.hasKeyset {
		return nil, fmt.Errorf("keyset not found")
	}
	if cc.keyset == nil {
		return nil, fmt.Errorf("keyset is nil")
	}
	return cc.keyset, nil
}

func (cc *SignerContext) NewSource() (mpc.KeyshareSource, error) {
	if !cc.hasKeyset {
		return nil, fmt.Errorf("keyset not found")
	}
	if cc.keyset == nil {
		return nil, fmt.Errorf("keyset is nil")
	}
	src, err := mpc.NewSource(cc.keyset)
	if err != nil {
		return nil, err
	}
	cc.signer = src
	cc.hasSigner = true
	return src, nil
}

func (cc *SignerContext) GetSource() (mpc.KeyshareSource, error) {
	if !cc.hasSigner {
		return nil, fmt.Errorf("signer not found")
	}
	if cc.signer == nil {
		return nil, fmt.Errorf("signer is nil")
	}
	return cc.signer, nil
}
