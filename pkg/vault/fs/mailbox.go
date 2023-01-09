package fs

import "errors"

func (vfs *vaultFsImpl) ListMessages() ([][]byte, error) {
	return nil, errors.New("Method not implemented")
}
func (vfs *vaultFsImpl) SendMessage(to []byte, message []byte) error {
	return errors.New("Method not implemented")
}
