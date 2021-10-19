package fs

import (
	"os"
	"path/filepath"
)

type Folder string

func (f Folder) Path() string {
	return string(f)
}

func (f Folder) Create(fileName string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(filepath.Join(f.Path(), fileName)), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(filepath.Join(f.Path(), fileName))
}

func (f Folder) CreateFolder(dirName string) (Folder, error) {
	path := filepath.Join(f.Path(), dirName)
	return Folder(path), os.MkdirAll(path, 0755)
}

func (f Folder) Delete(fileName string) error {
	return os.Remove(filepath.Join(f.Path(), fileName))
}

func (f Folder) Exists(fileName string) bool {
	_, err := os.Stat(filepath.Join(f.Path(), fileName))
	return !os.IsNotExist(err)
}

func (f Folder) MkdirAll() error {
	return os.MkdirAll(f.Path(), 0755)
}

func (f Folder) GenPath(path string, opts ...FilePathOption) (string, error) {
	// Initialize options list
	name := filepath.Base(path)
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

func (f Folder) ReadFile(fileName string) ([]byte, error) {
	return os.ReadFile(filepath.Join(f.Path(), fileName))
}

func (f Folder) WriteFile(fileName string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(filepath.Join(f.Path(), fileName)), 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(f.Path(), fileName), data, 0644)
}
