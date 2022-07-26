package protocol

import (
	"github.com/ipfs/go-cid"
)

type IPFS interface {
	GetData(cid string) ([]byte, error)
	PutData(data []byte) (*cid.Cid, error)
	RemoveFile(cidstr string) error
}
