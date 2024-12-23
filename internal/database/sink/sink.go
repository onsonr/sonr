package sink

import (
	_ "embed"
)

//go:embed vault/schema.sql
var SchemaVaultSQL string
