//go:build js && wasm
// +build js,wasm

package idb

import (
	"context"
	"encoding/json"
	"errors"
	"syscall/js"

	"github.com/hack-pad/go-indexeddb/idb"
)

// Model is an interface that must be implemented by types used with Table
type Model interface {
	Table() string
}

// Table is a generic wrapper around IDB for easier database operations on a specific table
type Table[T Model] struct {
	db      *idb.Database
	dbName  string
	keyPath string
}

// NewTable creates a new Table instance
func NewTable[T Model](dbName string, version uint, keyPath string) (*Table[T], error) {
	ctx := context.Background()
	factory := idb.Global()

	var model T
	tableName := model.Table()

	openRequest, err := factory.Open(ctx, dbName, version, func(db *idb.Database, oldVersion, newVersion uint) error {
		_, err := db.CreateObjectStore(tableName, idb.ObjectStoreOptions{
			KeyPath: js.ValueOf(keyPath),
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	db, err := openRequest.Await(ctx)
	if err != nil {
		return nil, err
	}
	return &Table[T]{
		db:      db,
		dbName:  dbName,
		keyPath: keyPath,
	}, nil
}

// Insert adds a new record to the table
func (t *Table[T]) Insert(data T) error {
	tx, err := t.db.Transaction(idb.TransactionReadWrite, data.Table())
	if err != nil {
		return err
	}
	defer tx.Commit()

	objectStore, err := tx.ObjectStore(data.Table())
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = objectStore.Add(js.ValueOf(string(jsonData)))
	return err
}

// Query retrieves a record from the table based on a key
func (t *Table[T]) Query(key interface{}) (T, error) {
	var result T

	tx, err := t.db.Transaction(idb.TransactionReadOnly, result.Table())
	if err != nil {
		return result, err
	}
	defer tx.Commit()

	objectStore, err := tx.ObjectStore(result.Table())
	if err != nil {
		return result, err
	}

	request, err := objectStore.Get(js.ValueOf(key))
	if err != nil {
		return result, err
	}

	value, err := request.Await(context.Background())
	if err != nil {
		return result, err
	}

	if value.IsUndefined() || value.IsNull() {
		return result, errors.New("record not found")
	}

	err = json.Unmarshal([]byte(value.String()), &result)
	return result, err
}

// Delete removes a record from the table based on a key
func (t *Table[T]) Delete(key interface{}) error {
	var model T
	tx, err := t.db.Transaction(idb.TransactionReadWrite, model.Table())
	if err != nil {
		return err
	}
	defer tx.Commit()

	objectStore, err := tx.ObjectStore(model.Table())
	if err != nil {
		return err
	}

	_, err = objectStore.Delete(js.ValueOf(key))
	return err
}

// Close closes the database connection
func (t *Table[T]) Close() error {
	return t.db.Close()
}
