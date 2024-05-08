package types

// IController is the interface for the controller
type IController interface {
	Set(key, value string) ([]byte, error)
	PublicKey() *PublicKey
	Refresh() error
	Sign(msg []byte) ([]byte, error)
	Remove(key, value string) error
	Check(key string, w []byte) bool
}