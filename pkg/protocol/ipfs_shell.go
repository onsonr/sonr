package protocol

import (
	"github.com/ipfs/go-cid"
	shell "github.com/ipfs/go-ipfs-api"
)

// IPFSShell implements a protocol.IPFS featuring:
// 	- IPFS initialization and error fallback
//	- In-memory cache mechanism
type IPFSShell struct {
	*shell.Shell
}

func NewIPFSShell(url string) *IPFSShell {
	return &IPFSShell{shell.NewShell(url)}
}

func (I IPFSShell) GetData(cid string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (I IPFSShell) PutData(data []byte) (*cid.Cid, error) {
	//TODO implement me
	panic("implement me")
}

func (I IPFSShell) PinFile(cidstr string) error {
	//TODO implement me
	panic("implement me")
}

func (I IPFSShell) RemoveFile(cidstr string) error {
	//TODO implement me
	panic("implement me")
}
