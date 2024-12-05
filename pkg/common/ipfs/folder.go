package ipfs

import "github.com/ipfs/boxo/files"

type Folder = files.Directory

func NewFolder(fs ...File) Folder {
	return files.NewMapDirectory(convertFilesToMap(fs))
}

func convertFilesToMap(vs []File) map[string]files.Node {
	m := make(map[string]files.Node)
	for _, f := range vs {
		m[f.Name()] = f
	}
	return m
}
