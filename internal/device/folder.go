package device

import (
	"os"

	"path/filepath"
)

type FolderType int

const (
	FolderType_NONE FolderType = iota
	FolderType_HOME
	FolderType_DOCUMENTS
	FolderType_DOWNLOADS
	FolderType_TEMPORARY
	FolderType_TEMPORARY_SUBDIR
	FolderType_SUPPORT
	FolderType_SUPPORT_SUBDIR
)


type Folder struct {
	Type FolderType
	Path string
}

func buildMobileFolders(home Folder) []Folder {
	homePath := home.Path
	supportPath := filepath.Join(homePath, "Support")
	return []Folder{
		home,
		NewFolder(homePath, FolderType_DOWNLOADS),
		NewFolder(supportPath, FolderType_SUPPORT),
		NewFolder(filepath.Join(supportPath, "Textile"), FolderType_SUPPORT_SUBDIR),
		NewFolder(filepath.Join(supportPath, "Database"), FolderType_SUPPORT_SUBDIR),
	}
}

func NewFolder(path string, t FolderType) Folder {
	return Folder{
		Type: t,
		Path: path,
	}
}

func (c Folder) Name() string {
	return filepath.Base(c.Path)
}

func (c Folder) Delete(fileName string) error {
	return os.Remove(filepath.Join(c.Path, fileName))
}

func (c Folder) Create(fileName string) (*os.File, error) {
	err := c.CreateParentDir(fileName)
	if err != nil {
		return nil, err
	}
	return os.Create(filepath.Join(c.Path, fileName))
}

func (c Folder) ReadFile(fileName string) ([]byte, error) {
	return os.ReadFile(filepath.Join(c.Path, fileName))
}

// CreateParentDir creates the parent directory of fileName inside c. fileName
// is a relative path inside c, containing zero or more path separators.
func (c Folder) CreateParentDir(fileName string) error {
	return os.MkdirAll(filepath.Dir(filepath.Join(c.Path, fileName)), 0755)
}

func (c Folder) WriteFile(fileName string, data []byte) error {
	err := c.CreateParentDir(fileName)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(c.Path, fileName), data, 0644)
}

func (c Folder) MkdirAll() error {
	return os.MkdirAll(c.Path, 0755)
}

func (d Folder) Equals(other Folder) bool {
	return d.Path == other.Path && d.Type == other.Type
}

// Exists returns true if the directory exists.
func (c Folder) Exists(fileName string) bool {
	_, err := os.Stat(filepath.Join(c.Path, fileName))
	return !os.IsNotExist(err)
}

// Join returns the path for the directory joined with the path.
func (d Folder) Join(path string) (string, error) {
	// Join the directory with the path
	return filepath.Join(d.Path, path), nil
}

// Has returns true if the directory has the file or directory.
func (d Folder) Has(p string) bool {
	// Get the directory path
	path, err := d.Join(p)
	if err != nil {
		logger.Error("Failed to determine joined path", err)
		return false
	}

	// Check if the directory has the file or directory
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
