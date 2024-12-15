package models

import "github.com/onsonr/sonr/crypto/mpc"

type StagedEnclave struct {
	Origin string `json:"origin"`
	Handle string `json:"handle"`
	Nonce  []byte `json:"nonce"`
	mpc.Enclave
}
