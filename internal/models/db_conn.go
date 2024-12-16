package models

import (
	"context"
	"database/sql"

	"github.com/onsonr/sonr/internal/models/sink"
	config "github.com/onsonr/sonr/pkg/config/hway"
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
