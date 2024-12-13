package mpc

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/bech32"
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
