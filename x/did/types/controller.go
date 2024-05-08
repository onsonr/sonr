package types

// Controller is the interface for the controller
type ControllerI interface {
	Set(key, value string) ([]byte, error)
	PublicKey() *PublicKey
	Refresh() error
	Sign(msg []byte) ([]byte, error)
	Remove(key, value string) error
	Check(key string, w []byte) bool
}
