package ipfs

import (
	"context"
	"fmt"
	"time"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
)

type IPFSFile struct {
	node   files.Node
	path   string
	name   string
	client *rpc.HttpApi
}

func (f *IPFSFile) Close() error {
	return nil // IPFS nodes don't need to be closed
}

func (f *IPFSFile) Read(p []byte) (n int, err error) {
	if file, ok := f.node.(files.File); ok {
		return file.Read(p)
	}
	return 0, fmt.Errorf("not a file")
}

func (f *IPFSFile) Seek(offset int64, whence int) (int64, error) {
	if file, ok := f.node.(files.File); ok {
		return file.Seek(offset, whence)
	}
	return 0, fmt.Errorf("not a file")
}

func (f *IPFSFile) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("write operation not supported for IPFS files")
}

func (f *IPFSFile) String() string {
	return f.path
}

func (f *IPFSFile) Exists() (bool, error) {
	// In IPFS, if we have the node, it exists
	return true, nil
}

func (f *IPFSFile) CopyToFile(file File) error {
	// Implementation depends on how you want to handle copying between IPFS and other file types
	return fmt.Errorf("CopyToFile not implemented for IPFS files")
}

func (f *IPFSFile) MoveToFile(file File) error {
	// Moving files in IPFS doesn't make sense in the traditional way
	return fmt.Errorf("MoveToFile not applicable for IPFS files")
}

func (f *IPFSFile) Delete() error {
	// Deleting in IPFS is not straightforward, might need to implement unpinning
	return fmt.Errorf("Delete operation not supported for IPFS files")
}

func (f *IPFSFile) LastModified() (*time.Time, error) {
	// IPFS doesn't have a concept of last modified time
	return nil, fmt.Errorf("LastModified not applicable for IPFS files")
}

func (f *IPFSFile) Size() (uint64, error) {
	if file, ok := f.node.(files.File); ok {
		s, _ := file.Size()
		return uint64(s), nil
	}
	return 0, fmt.Errorf("not a file")
}

func (f *IPFSFile) Path() string {
	return f.path
}

func (f *IPFSFile) Name() string {
	return f.name
}

func (f *IPFSFile) Touch() error {
	return fmt.Errorf("Touch operation not supported for IPFS files")
}

func (f *IPFSFile) URI() string {
	return fmt.Sprintf("ipfs://%s", f.path)
}

type IPFSFileSystem struct {
	client *rpc.HttpApi
}

func (fs *IPFSFileSystem) NewFile(volume string, absFilePath string) (File, error) {
	p, err := path.NewPath(absFilePath)
	if err != nil {
		return nil, err
	}
	node, err := fs.client.Unixfs().Get(context.Background(), p)
	if err != nil {
		return nil, err
	}

	return &IPFSFile{
		node:   node,
		path:   absFilePath,
		name:   p.String(),
		client: fs.client,
	}, nil
}

func (fs *IPFSFileSystem) Name() string {
	return "IPFS"
}
