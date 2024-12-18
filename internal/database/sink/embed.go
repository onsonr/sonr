package sink

import (
	_ "embed"
)

//go:embed schema_hway.sql
var SchemaHwaySQL string

//go:embed schema_motr.sql
var SchemaMotrSQL string

