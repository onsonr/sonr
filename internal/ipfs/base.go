package ipfs

import (
	"fmt"
	"io"
	"time"
)

var (
	ErrNotInitialized = fmt.Errorf("IPFS client not initialized")
	ErrNotMounted     = fmt.Errorf("IPFS client not mounted")
	ErrCIDNotFound    = fmt.Errorf("CID not found")
	ErrIPNSNotFound   = fmt.Errorf("IPNS not found")
	ErrInternal       = fmt.Errorf("internal error")
)

type FileSystem interface {
	// NewFile initializes a File on the specified volume at path 'absFilePath'.
	//
	//   * Accepts volume and an absolute file path.
	//   * Upon success, a vfs.File, representing the file's new path (location path + file relative path), will be returned.
	//   * On error, nil is returned for the file.
	//   * Note that not all file systems will have a "volume" and will therefore be "":
	//       file:///path/to/file has a volume of "" and name /path/to/file
	//     whereas
	//       s3://mybucket/path/to/file has a volume of "mybucket and name /path/to/file
	//     results in /tmp/dir1/newerdir/file.txt for the final vfs.File path.
	//   * The file may or may not already exist.
	NewFile(volume string, absFilePath string) (File, error)

	// Name returns the name of the FileSystem ie: Amazon S3, os, Google Cloud Storage, etc.
	Name() string
}

type File interface {
	io.Closer
	io.Reader
	io.Seeker
	io.Writer
	fmt.Stringer

	// Exists returns boolean if the file exists on the file system.  Returns an error, if any.
	Exists() (bool, error)

	// CopyToFile will copy the current file to the provided file instance.
	//
	//   * In the case of an error, nil is returned for the file.
	//   * CopyToLocation should use native functions when possible within the same scheme.
	//   * If the file already exists, the contents will be overwritten with the current file's contents.
	//   * CopyToFile will Close both the source and target Files which therefore can't be appended to without first
	//     calling Seek() to move the cursor to the end of the file.
	CopyToFile(file File) error

	// MoveToFile will move the current file to the provided file instance.
	//
	//   * If the file already exists, the contents will be overwritten with the current file's contents.
	//   * The current instance of the file will be removed.
	//   * MoveToFile will Close both the source and target Files which therefore can't be appended to without first
	//     calling Seek() to move the cursor to the end of the file.
	MoveToFile(file File) error

	// Delete unlinks the File on the file system.
	Delete() error

	// LastModified returns the timestamp the file was last modified (as *time.Time).
	LastModified() (*time.Time, error)

	// Size returns the size of the file in bytes.
	Size() (uint64, error)

	// Path returns absolute path, including filename, ie /some/path/to/file.txt
	//
	// If the directory portion of a file is desired, call
	//   someFile.Location().Path()
	Path() string

	// Name returns the base name of the file path.
	//
	// For file:///some/path/to/file.txt, it would return file.txt
	Name() string

	// Touch creates a zero-length file on the vfs.File if no File exists.  Update File's last modified timestamp.
	// Returns error if unable to touch File.
	Touch() error

	// URI returns the fully qualified absolute URI for the File.  IE, s3://bucket/some/path/to/file.txt
	URI() string
}
