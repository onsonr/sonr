package fs

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ipfs/boxo/files"
)

// File represents a file in the filesystem
type File string

// NewFile creates a new File instance
func NewFile(path string) File {
	return File(path)
}

// Name returns the name of the file
func (f File) Name() string {
	return filepath.Base(string(f))
}

// Path returns the path of the file
func (f File) Path() string {
	return string(f)
}

// Read reads the contents of the file
func (f File) Read() ([]byte, error) {
	return os.ReadFile(string(f))
}

// Write writes data to the file
func (f File) Write(data []byte) error {
	return os.WriteFile(string(f), data, 0644)
}

// Append appends data to the file
func (f File) Append(data []byte) error {
	file, err := os.OpenFile(string(f), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

// Exists checks if the file exists
func (f File) Exists() bool {
	_, err := os.Stat(string(f))
	return !os.IsNotExist(err)
}

// Remove removes the file
func (f File) Remove() error {
	return os.Remove(string(f))
}

// Stat returns the FileInfo for the file
func (f File) Stat() (os.FileInfo, error) {
	return os.Stat(string(f))
}

// Open opens the file and returns an fs.File
func (f File) Open() (fs.File, error) {
	return os.Open(string(f))
}

// Node returns a files.Node representation of the file
func (f File) Node() (files.Node, error) {
	file, err := os.Open(string(f))
	if err != nil {
		return nil, err
	}
	return files.NewReaderFile(file), nil
}
