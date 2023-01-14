package fs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	files "github.com/ipfs/go-ipfs-files"
)

var (
	// Path Manipulation Errors
	ErrDuplicateFilePathOption    = errors.New("Duplicate file path option")
	ErrPrefixSuffixSetWithReplace = errors.New("Prefix or Suffix set with Replace.")
	ErrSeparatorLength            = errors.New("Separator length must be 1.")
	ErrNoFileNameSet              = errors.New("File name was not set by options.")

	// Device ID Errors
	ErrEmptyDeviceID = errors.New("Device ID cannot be empty")
	ErrMissingEnvVar = errors.New("Cannot set EnvVariable with empty value")

	// Directory errors
	ErrDirectoryInvalid = errors.New("Directory Type is invalid")
	ErrDirectoryUnset   = errors.New("Directory path has not been set")
	ErrDirectoryJoin    = errors.New("Failed to join directory path")
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

// ListFiles returns a list of files in a folder.
func (f Folder) ListFiles() ([]string, error) {
	var files []string
	err := filepath.Walk(f.Path(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// ListFolders returns a list of folders in a folder.
func (f Folder) ListFolders() ([]Folder, error) {
	var folders []Folder
	err := filepath.Walk(f.Path(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			folders = append(folders, Folder(path))
		}
		return nil
	})
	return folders, err
}

// MkdirAll creates a directory and all its parents.
func (f Folder) MkdirAll() error {
	return os.MkdirAll(f.Path(), 0755)
}

// FileNode returns the ipfs file node.
func (f Folder) FileNode() (files.Node, error) {
	st, err := os.Stat(f.Path())
	if err != nil {
		return nil, err
	}
	return files.NewSerialFile(f.Path(), false, st)
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

// Name returns the name of the folder.
func (f Folder) Name() string {
	return filepath.Base(f.Path())
}

// ReadFile reads a file.
func (f Folder) ReadFile(fileName string) ([]byte, error) {
	return os.ReadFile(filepath.Join(f.Path(), fileName))
}

// Stat returns the FileInfo structure describing file.
func (f Folder) Stat(fileName string) (os.FileInfo, error) {
	return os.Stat(filepath.Join(f.Path(), fileName))
}

// WriteFile writes a file.
func (f Folder) WriteFile(fileName string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(filepath.Join(f.Path(), fileName)), 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(f.Path(), fileName), data, 0644)
}
