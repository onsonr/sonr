package ipfs

import (
	"fmt"
	"io"

	"github.com/ipfs/boxo/files"
	"github.com/spf13/afero"
)

// TODO: Add a new variable called AppFs that is an instance of afero.NewMemMapFs().
var AppFs = afero.NewMemMapFs()

// VFS is an interface for interacting with a virtual file system.
type VFS interface {
	Add(path string, data []byte) error
	Get(path string) ([]byte, error)
	Rm(path string) error
	Ls(path string) ([]string, error)
	Name() string
	Node() files.Node
}

// vfs is the struct implementation of an IPFS file system
type vfs struct {
	files map[string]files.File
	name  string

	// TODO: Add afero.Fs to the vfs struct.
	fs afero.Fs
}

// NewVFS creates a new virtual file system.
func NewVFS(name string) VFS {
	fs := afero.NewMemMapFs()
	return &vfs{
		files: make(map[string]files.File, 0),
		name:  name,
		fs:    fs,
	}
}

// Name returns the name of the virtual file system.
func (v *vfs) Name() string {
	return v.name
}

// Add adds a file to the virtual file system.
func (v *vfs) Add(path string, data []byte) error {
	v.files[path] = files.NewBytesFile(data)
	return nil
}

// Get retrieves a file from the virtual file system.
func (v *vfs) Get(path string) ([]byte, error) {
	if file, ok := v.files[path]; ok {
		return io.ReadAll(file)
	}
	return nil, fmt.Errorf("file not found")
}

// Rm removes a file from the virtual file system.
func (v *vfs) Rm(path string) error {
	delete(v.files, path)
	return nil
}

// Ls lists the files in the virtual file system.
func (v *vfs) Ls(path string) ([]string, error) {
	var files []string
	for k := range v.files {
		files = append(files, k)
	}
	return files, nil
}

// Node returns the root node of the virtual file system.
func (v *vfs) Node() files.Node {
	rootDir := make(map[string]files.Node, 0)
	fileList := make([]files.DirEntry, 0)
	for k, f := range v.files {
		ent := files.FileEntry(k, f)
		fileList = append(fileList, ent)
	}
	dir := files.NewSliceDirectory(fileList)
	node := dir.Entries().Node()
	rootDir[v.name] = node
	finalDir := files.NewMapDirectory(rootDir)
	return finalDir.Entries().Node()
}
