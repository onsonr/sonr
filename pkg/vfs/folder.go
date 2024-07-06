package vfs

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ipfs/boxo/files"
)

// Folder represents a folder in the filesystem
type Folder string

// NewFolder creates a new Folder instance
func NewFolder(path string) Folder {
	return Folder(path)
}

// Name returns the name of the folder
func (f Folder) Name() string {
	return filepath.Base(string(f))
}

// Path returns the path of the folder
func (f Folder) Path() string {
	return string(f)
}

// Create creates the folder if it doesn't exist
func (f Folder) Create() error {
	return os.MkdirAll(string(f), os.ModePerm)
}

// Exists checks if the folder exists
func (f Folder) Exists() bool {
	info, err := os.Stat(string(f))
	return err == nil && info.IsDir()
}

// Ls lists the contents of the folder
func (f Folder) Ls() ([]fs.DirEntry, error) {
	return os.ReadDir(string(f))
}

// WriteFile writes data to a file in the folder
func (f Folder) WriteFile(name string, data []byte, perm os.FileMode) (File, error) {
	err := os.WriteFile(filepath.Join(string(f), name), data, perm)
	return File(filepath.Join(string(f), name)), err
}

// ReadFile reads the contents of a file in the folder
func (f Folder) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(string(f), name))
}

// DeleteFile deletes a file from the folder
func (f Folder) DeleteFile(name string) error {
	return os.Remove(filepath.Join(string(f), name))
}

// Remove removes the folder and its contents
func (f Folder) Remove() error {
	return os.RemoveAll(string(f))
}

// Join joins the folder path with the given elements
func (f Folder) Join(elem ...string) Folder {
	return Folder(filepath.Join(append([]string{string(f)}, elem...)...))
}

// IsDir checks if the folder is a directory
func (f Folder) IsDir() bool {
	info, err := os.Stat(string(f))
	return err == nil && info.IsDir()
}

// Touch creates an empty file if it doesn't exist, or updates its access and modification times if it does
func (f Folder) Touch(name string) (File, error) {
	var err error
	filePath := filepath.Join(string(f), name)
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return File(filePath), err
		}
		if err := file.Close(); err != nil {
			return File(filePath), err
		}
		return File(filePath), nil
	}
	return File(filePath), err
}

// Node returns a files.Node representation of the folder
func (f Folder) Node() (files.Node, error) {
	entries, err := f.Ls()
	if err != nil {
		return nil, err
	}
	dirEntries := make([]files.DirEntry, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			subFolder := f.Join(entry.Name())
			subNode, err := subFolder.Node()
			if err != nil {
				return nil, err
			}
			dirEntries = append(dirEntries, files.FileEntry(entry.Name(), subNode))
		} else {
			file := File(filepath.Join(string(f), entry.Name()))
			fileNode, err := file.Node()
			if err != nil {
				return nil, err
			}
			dirEntries = append(dirEntries, files.FileEntry(entry.Name(), fileNode))
		}
	}
	return files.NewSliceDirectory(dirEntries), nil
}
