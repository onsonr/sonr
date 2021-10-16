package device

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kataras/golog"
)

// NewPath returns a new path with the given file name and specified folder.
func NewPath(dir string, path string, opts ...FilePathOption) string {
	// Initialize options list
	name := filepath.Base(path)
	options := defaultFilePathOptions(dir, name)
	for _, opt := range opts {
		opt(options)
	}

	// Build path
	return options.Apply()
}

// NewDatabasePath Returns a new path in database dir with given file name.
func NewDatabasePath(path string, opts ...FilePathOption) string {
	// Initialize options list
	name := filepath.Base(path)
	options := defaultFilePathOptions(DatabasePath, name)
	for _, opt := range opts {
		opt(options)
	}

	// Build path
	return options.Apply()
}

// NewDocsPath Returns a new path in docs dir with given file name.
func NewDocsPath(path string, opts ...FilePathOption) string {
	// Initialize options list
	name := filepath.Base(path)
	options := defaultFilePathOptions(DocsPath, name)
	for _, opt := range opts {
		opt(options)
	}

	// Build path
	return options.Apply()
}

// NewDownloadsPath Returns a new path in downloads dir with given file name.
func NewDownloadsPath(path string, opts ...FilePathOption) string {
	// Initialize options list
	name := filepath.Base(path)
	options := defaultFilePathOptions(DownloadsPath, name)
	for _, opt := range opts {
		opt(options)
	}

	// Build path
	return options.Apply()
}

// NewTempPath Returns a new path in temp dir with given file name.
func NewTempPath(path string, opts ...FilePathOption) string {
	// Initialize options list
	name := filepath.Base(path)
	options := defaultFilePathOptions(TempPath, name)
	for _, opt := range opts {
		opt(options)
	}

	// Build path
	return options.Apply()
}

// NewSupportPath Returns a new path in support dir with given file name.
func NewSupportPath(path string, opts ...FilePathOption) string {
	// Initialize options list
	name := filepath.Base(path)
	options := defaultFilePathOptions(SupportPath, name)
	for _, opt := range opts {
		opt(options)
	}

	// Build path
	return options.Apply()
}

// NodeOption is a function that modifies the node options.
type FilePathOption func(*filePathOptions)

// CreateDirIfNotExist creates the directory if it does not exist.
func CreateDirIfNotExist() FilePathOption {
	return func(o *filePathOptions) {
		o.CreateDir = true
	}
}

// WithSuffix sets the suffix for the file name.
func WithSuffix(path string) FilePathOption {
	return func(o *filePathOptions) {
		o.Suffix = path
	}
}

// WithPrefix sets the prefix for the file path.
func WithPrefix(v string) FilePathOption {
	return func(o *filePathOptions) {
		o.Prefix = v
	}
}

// WithReplace sets the replace string for the file path.
func WithReplace(v string) FilePathOption {
	return func(o *filePathOptions) {
		o.Replace = v
	}
}

// WithSeparator sets the separator for the file path.
func WithSeparator(v string) FilePathOption {
	return func(o *filePathOptions) {
		o.Separator = v
	}
}

// filePathOptions is a struct for holding file path options.
type filePathOptions struct {
	// Options
	Suffix    string // Add Suffix to file name
	Prefix    string // Add Prefix to file name
	Replace   string // Replace filename with this string
	Separator string // Default is "-"
	Directory string // Directory for File
	CreateDir bool   // Create Directory if not exist

	// Properties
	fileName string
	baseName string
	ext      string
}

// defaultFilePathOptions returns the default file path options.
func defaultFilePathOptions(dir string, name string) *filePathOptions {
	// Initialize options list
	defOpts := &filePathOptions{
		Suffix:    "",
		Prefix:    "",
		Replace:   "",
		Separator: "-",
		Directory: dir,
		CreateDir: false,
	}

	// Set File Name
	if strings.Contains(".", name) {
		defOpts.baseName = strings.Split(name, ".")[0]
		defOpts.ext = strings.Split(name, ".")[1]
	} else {
		defOpts.baseName = name
		defOpts.ext = ""
	}
	return defOpts
}

// Merge merges the file path options.
func (fpo *filePathOptions) Apply() string {
	// Check for Replace
	if fpo.Replace != "" {
		return filepath.Join(fpo.Directory, fpo.Replace)
	}

	// Check for Prefix
	rawPath := strings.Join([]string{fpo.Prefix, fpo.baseName, fpo.Suffix}, fpo.Separator)
	path := filepath.Join(fpo.Directory, (rawPath + "." + fpo.ext))

	// Check for Create Dir
	if fpo.CreateDir {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			logger.Error("Failed to Create Dir as FilePathOption", golog.Fields{"error": err})
		} else {
			logger.Info("Created new Directory for Path", golog.Fields{"path": path, "root": fpo.Directory, "parent": filepath.Base(fpo.Directory)})
		}
	}
	return path
}
