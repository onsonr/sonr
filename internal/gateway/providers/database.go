package providers

import (
	"database/sql"

	"github.com/onsonr/sonr/internal/gateway/repository"
)

type DatabaseProvider struct {
	*repository.Queries
}

func NewDatabaseService(conn *sql.Conn) DatabaseProvider {
	return DatabaseProvider{
		Queries: repository.New(conn),
	}
}
