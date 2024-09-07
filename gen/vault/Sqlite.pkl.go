// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

type Sqlite interface {
	DB
}

var _ Sqlite = (*SqliteImpl)(nil)

type SqliteImpl struct {
	Filename string `pkl:"filename"`
}

func (rcv *SqliteImpl) GetFilename() string {
	return rcv.Filename
}
