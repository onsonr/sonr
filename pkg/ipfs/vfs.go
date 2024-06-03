package ipfs

import (
	"fmt"
	"io"

	"github.com/ipfs/boxo/files"
)

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
}

// NewVFS creates a new virtual file system.
func NewFileSystem(name string) VFS {
	return &vfs{
		files: make(map[string]files.File, 0),
		name:  name,
	}
}

// Load creates a new virtual file system from a given files.Node.
func Load(name string, node files.Node) (VFS, error) {
	entry := files.FileEntry(name, node)
	rootDir := files.DirFromEntry(entry)
	vfs := &vfs{
		files: make(map[string]files.File, 0),
		name:  name,
	}

	err := loadDirectory(rootDir, vfs)
	if err != nil {
		return nil, err
	}

	return vfs, nil
}

// loadDirectory recursively loads the files and directories from a given directory node.
func loadDirectory(dir files.Directory, vfs *vfs) error {
	it := dir.Entries()
	for it.Next() {
		name, node := it.Name(), it.Node()
		switch node := node.(type) {
		case files.File:
			data, err := io.ReadAll(node)
			if err != nil {
				return err
			}
			vfs.files[name] = files.NewBytesFile(data)

		case files.Directory:
			subDir := files.DirFromEntry(files.FileEntry(name, node))
			err := loadDirectory(subDir, vfs)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported node type")
		}
	}
	return nil
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
