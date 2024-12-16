package database

import (
	"context"
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	config "github.com/onsonr/sonr/internal/config/hway"
	"github.com/onsonr/sonr/internal/database/sink"
)

// NewDB initializes and returns a configured database connection
func NewDB(env config.Hway) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// create tables
	if _, err := db.ExecContext(context.Background(), sink.SchemaSQL); err != nil {
		return nil, err
	}
	return db, nil
}
