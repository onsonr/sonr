package ipfs

import "github.com/ipfs/boxo/files"

type file struct {
	files.File
	name string
}

func (f *file) Name() string {
	return f.name
}

func NewFile(name string, data []byte) File {
	return &file{File: files.NewBytesFile(data), name: name}
}
