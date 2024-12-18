package sqlite

import (
	_ "embed"
)

//go:embed schema.sql
var SchemaMotrSQL string
