package fs

import (
	"path/filepath"
	"strings"
)

type filePathOptType int

// Option Type Enum
const (
	filePathOptionTypeSuffix    filePathOptType = iota // Suffix
	filePathOptionTypePrefix                           // Prefix
	filePathOptionTypeReplace                          // Replace
	filePathOptionTypeSeparator                        // Separator
)

// FilePathOption is a function option for FilePath.
type FilePathOption interface {
	Apply() *filePathOptions
}

// filePathOpt is a struct that implements FilePathOption.
type filePathOpt struct {
	FilePathOption
	value string
	filePathOptType
}

// Apply returns the value set for the specified option.
func (fio *filePathOpt) Apply() *filePathOptions {
	if fio.filePathOptType == filePathOptionTypeSuffix {
		return &filePathOptions{Suffix: fio.value}
	} else if fio.filePathOptType == filePathOptionTypePrefix {
		return &filePathOptions{Prefix: fio.value}
	} else if fio.filePathOptType == filePathOptionTypeReplace {
		return &filePathOptions{Replace: fio.value}
	} else if fio.filePathOptType == filePathOptionTypeSeparator {
		return &filePathOptions{Separator: fio.value}
	}
	return nil
}

// FilePathOptFunc is a function that returns a FilePathOption.
type FilePathOptFunc func(path string) FilePathOption

// WithSuffix sets the suffix for the file path.
var WithSuffix FilePathOptFunc = func(path string) FilePathOption {
	return &filePathOpt{
		filePathOptType: filePathOptionTypeSuffix,
		value:           path,
	}
}

// WithPrefix sets the prefix for the file path.
var WithPrefix FilePathOptFunc = func(path string) FilePathOption {
	return &filePathOpt{
		filePathOptType: filePathOptionTypePrefix,
		value:           path,
	}
}

// WithReplace sets the replace string for the file path.
var WithReplace FilePathOptFunc = func(path string) FilePathOption {
	return &filePathOpt{
		filePathOptType: filePathOptionTypeReplace,
		value:           path,
	}
}

// WithSeparator sets the separator for the file path.
var WithSeparator FilePathOptFunc = func(path string) FilePathOption {
	return &filePathOpt{
		filePathOptType: filePathOptionTypeSeparator,
		value:           path,
	}
}

// filePathOptions is a struct for holding file path options.
type filePathOptions struct {
	// Options
	Suffix    string // Add Suffix to file name
	Prefix    string // Add Prefix to file name
	Replace   string // Replace filename with this string
	Separator string // Default is "-"

	// Properties
	fileName  string
	baseName  string
	extension string

	// Internal
	suffixSet    bool
	prefixSet    bool
	replaceSet   bool
	separatorSet bool
}

// Merge merges the file path options.
func (fpo *filePathOptions) Merge(name string, optsList ...*filePathOptions) error {
	// Split name into base name and extension
	if strings.Contains(".", name) {
		// Set extension
		ptr := strings.Split(name, ".")
		fpo.extension = ptr[len(ptr)-1]

		// Set base name
		if len(ptr) > 2 {
			fpo.baseName = strings.Join(ptr[:len(ptr)-2], "_")
		} else {
			fpo.baseName = ptr[0]
		}
	} else {
		fpo.baseName = name
		fpo.extension = ""
	}

	// Set Checkers
	fpo.suffixSet = false
	fpo.prefixSet = false
	fpo.replaceSet = false
	fpo.separatorSet = false

	// Merge options
	for _, opts := range optsList {
		// Check if suffix is set
		if opts.Suffix != "" {
			if !fpo.suffixSet {
				fpo.Suffix = opts.Suffix
				fpo.suffixSet = true
			} else {
				return ErrDuplicateFilePathOption
			}
		}

		// Check if prefix is set
		if opts.Prefix != "" {
			if !fpo.prefixSet {
				fpo.Prefix = opts.Prefix
				fpo.prefixSet = true
			} else {
				return ErrDuplicateFilePathOption
			}
		}

		// Check if replace is set
		if opts.Replace != "" {
			if !fpo.replaceSet {
				if fpo.Prefix == "" && fpo.Suffix == "" {
					fpo.Replace = opts.Replace
					fpo.replaceSet = true
				} else {
					return ErrPrefixSuffixSetWithReplace
				}
			} else {
				return ErrDuplicateFilePathOption
			}
		}

		// Check separator
		if opts.Separator != "" {
			if !fpo.separatorSet {
				if len(opts.Separator) == 1 {
					fpo.Separator = opts.Separator
					fpo.separatorSet = true
				} else {
					return ErrSeparatorLength
				}
			} else {
				return ErrDuplicateFilePathOption
			}
		}
	}
	return nil
}

// Apply method applies the file path options to the given path.
func (fpo *filePathOptions) Apply(dir string) (string, error) {
	// Set Default Separator if not set
	if !fpo.separatorSet {
		fpo.Separator = "-"
	}

	// Check for Replace
	if fpo.replaceSet {
		// Check if prefix or suffix is set
		if fpo.suffixSet || fpo.prefixSet {
			return "", ErrPrefixSuffixSetWithReplace
		}
		// Set Filename to replace
		fpo.fileName = fpo.Replace
	} else {
		// Check for prefix
		if fpo.prefixSet {
			fpo.fileName = fpo.Prefix + fpo.Separator + fpo.baseName
		} else {
			fpo.fileName = fpo.baseName
		}

		// Check for suffix
		if fpo.suffixSet {
			fpo.fileName = fpo.fileName + fpo.Separator + fpo.Suffix
		}
	}

	// Add extension
	if fpo.extension != "" {
		fpo.fileName = fpo.fileName + "." + fpo.extension
	}

	// Check if file name is set
	if fpo.fileName != "" {
		path := filepath.Join(dir, fpo.fileName)
		return path, nil
	} else {
		return "", ErrNoFileNameSet
	}
}
