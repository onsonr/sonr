package ipfs

import "github.com/ipfs/boxo/files"

type Client interface {
	Add(data []byte) (string, error)
	AddFile(file File) (string, error)
	AddFolder(folder Folder) (string, error)
	Exists(cid string) (bool, error)
	Get(cid string) ([]byte, error)
	IsPinned(ipns string) (bool, error)
	Ls(cid string) ([]string, error)
	Pin(cid string, name string) error
	Unpin(cid string) error
}

type File interface {
	files.File
	Name() string
}

func NewFileMap(vs []File) map[string]files.Node {
	m := make(map[string]files.Node)
	for _, f := range vs {
		m[f.Name()] = f
	}
	return m
}
