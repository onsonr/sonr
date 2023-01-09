package fs

import "errors"

func (vfs *vaultFsImpl) Add(body []byte, name string) error {
	return errors.New("Method not implemented")
}
func (vfs *vaultFsImpl) Get(name string) ([]byte, error) {
	return nil, errors.New("Method not implemented")
}
