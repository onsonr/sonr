package types

import "github.com/di-dao/sonr/crypto"

// Controller is the interface for the controller
type ControllerI interface {
	Set(key, value string) ([]byte, error)
	PublicKey() crypto.PublicKey
	Refresh() error
	Sign(msg []byte) ([]byte, error)
	Remove(key, value string) error
	Check(key string, w []byte) bool
}
