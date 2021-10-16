package device

import (
	"path/filepath"
	"strings"
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

	// Properties
	fileName  string
	baseName  string
	extension string
	items     []string
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
	}

	// Set File Name
	if strings.Contains(".", name) {
		defOpts.baseName = strings.Split(name, ".")[0]
		defOpts.extension = strings.Split(name, ".")[1]
	} else {
		defOpts.baseName = name
		defOpts.extension = ""
	}
	return defOpts
}

// Merge merges the file path options.
func (fpo *filePathOptions) Apply() string {
	// Check for Replace
	if fpo.HasReplace() {
		return filepath.Join(fpo.Directory, fpo.Replace)
	}

	// Check for Prefix
	items := make([]string, 0)
	if fpo.HasPrefix() {
		items = append(items, fpo.Prefix)
	}

	items = append(items, fpo.baseName)

	// Check for suffix
	if fpo.HasSuffix() {
		items = append(items, fpo.Suffix)
	}

	// Add extension
	return filepath.Join(fpo.Directory, addExtension(fpo.extension, strings.Join(items, fpo.Separator)))
}

func (fpo *filePathOptions) HasExtension() bool {
	return fpo.extension != ""
}

func (fpo *filePathOptions) HasPrefix() bool {
	return fpo.Prefix != ""
}

func (fpo *filePathOptions) HasSuffix() bool {
	return fpo.Suffix != ""
}

func (fpo *filePathOptions) HasReplace() bool {
	return fpo.Replace != ""
}

func addExtension(ext, path string) string {
	if ext == "" {
		return path
	}
	if strings.Contains(ext, ".") {
		return path + ext
	} else {
		return path + "." + ext
	}
}
