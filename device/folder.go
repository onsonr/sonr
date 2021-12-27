package device

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// IsFile returns true if the given path is a file
func IsFile(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

// Exists checks if a file exists.
func Exists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

type Folder string

// Path returns the path of the folder.
func (f Folder) Path() string {
	return string(f)
}

// Create creates a file.
func (f Folder) Create(fileName string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(filepath.Join(f.Path(), fileName)), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(filepath.Join(f.Path(), fileName))
}

// CreateFolder creates a folder.
func (f Folder) CreateFolder(dirName string) (Folder, error) {
	path := filepath.Join(f.Path(), dirName)
	return Folder(path), os.MkdirAll(path, 0755)
}

// Delete removes a file or a folder.
func (f Folder) Delete(fileName string) error {
	return os.Remove(filepath.Join(f.Path(), fileName))
}

// Exists checks if a file exists.
func (f Folder) Exists(fileName string) bool {
	_, err := os.Stat(filepath.Join(f.Path(), fileName))
	return !os.IsNotExist(err)
}

// MkdirAll creates a directory and all its parents.
func (f Folder) MkdirAll() error {
	return os.MkdirAll(f.Path(), 0755)
}

// GenPath generates a path from a folder and a file name.
func (f Folder) GenPath(path string, opts ...FilePathOption) (string, error) {
	// Increment fileName to avoid overwriting
	num := 0
	name := filepath.Base(path)
	for f.Exists(name) {
		num++
		strPtr := strings.Split(filepath.Base(path), ".")
		if len(strPtr) == 1 {
			name = fmt.Sprintf("%s(%d)", filepath.Base(path), num)
		} else {
			name = fmt.Sprintf("%s(%d).%s", strPtr[0], num, strPtr[1])
		}
	}

	// Initialize options list
	fpoList := make([]*filePathOptions, len(opts))
	for _, opt := range opts {
		fpoList = append(fpoList, opt.Apply())
	}

	// Merge options
	fpo := &filePathOptions{}
	err := fpo.Merge(name, fpoList...)
	if err != nil {
		return "", err
	}

	// Build path
	return fpo.Apply(f.Path())
}

// JoinPath joins a folder and a file name.
func (f Folder) JoinPath(ps ...string) string {
	return filepath.Join(f.Path(), filepath.Join(ps...))
}

// ReadFile reads a file.
func (f Folder) ReadFile(fileName string) ([]byte, error) {
	return os.ReadFile(filepath.Join(f.Path(), fileName))
}

// WriteFile writes a file.
func (f Folder) WriteFile(fileName string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(filepath.Join(f.Path(), fileName)), 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(f.Path(), fileName), data, 0644)
}
