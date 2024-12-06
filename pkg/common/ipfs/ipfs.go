package ipfs

import "github.com/ipfs/boxo/files"

type Client interface {
	Add(data []byte) (string, error)
	AddFile(file File) (string, error)
	AddFolder(folder Folder) (string, error)
	Get(cid string) ([]byte, error)
	IsPublished(ipns string) (bool, error)
	Exists(cid string) (bool, error)
	Pin(cid string) error
	Unpin(cid string) error
	Publish(cid string, name string) (string, error)
	Ls(cid string) ([]string, error)
}

type File interface {
	files.File
	Name() string
}

func convertFilesToMap(vs []File) map[string]files.Node {
	m := make(map[string]files.Node)
	for _, f := range vs {
		m[f.Name()] = f
	}
	return m
}
