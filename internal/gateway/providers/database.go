package providers

import (
	"database/sql"

	"github.com/onsonr/sonr/internal/gateway/repository"
)

type DatabaseProvider struct {
	*repository.Queries
}

func NewDatabaseService(conn *sql.DB) DatabaseProvider {
	return DatabaseProvider{
		Queries: repository.New(conn),
	}
}
