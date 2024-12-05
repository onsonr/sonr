package ipfs

import "github.com/ipfs/boxo/files"

type Client interface {
	Add(data []byte) (string, error)
	Get(cid string) ([]byte, error)
	IsPublished(ipns string) (bool, error)
	Exists(cid string) (bool, error)
	Pin(cid string) error
	Unpin(cid string) error
	Publish(cid string, name string) (string, error)
	LsFolder(path string) ([]string, error)
}

type File interface {
	files.File
	Name() string
}
