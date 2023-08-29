package types

// SFSColumn is the interface for a column.
type SFSColumn interface {
	// Name returns the column name.
	Name() (string)

	// Kind returns the column data type.
	Kind() (string)
}

// SFSRow is the interface for a row.
type SFSRow interface {
	// Columns returns the row columns.
	Columns() ([]SFSColumn, error)

	// Index returns the row index.
	Index() (int)

	// Get returns the value at the given column.
	Get(column SFSColumn) (interface{}, error)

	// Set sets the value at the given column.
	Set(column SFSColumn, value interface{}) error

	// Values returns the row values.
	Values() ([]interface{}, error)
}

// SFSTable is the interface for a table.
type SFSTable interface {
	// Name returns the table name.
	Name() (string)

	// Columns returns the table columns.
	Columns() ([]SFSColumn, error)

	// Rows returns the table rows.
	Rows() ([]SFSRow, error)

	// AddColumn adds a column to the table.
	AddColumn(column SFSColumn) error

	// HasColumn checks if a column is in the table.
	HasColumn(column SFSColumn) (bool, error)

	// RemoveColumn removes a column from the table.
	RemoveColumn(column SFSColumn) error

	// InsertRows inserts the rows into the table.
	InsertRows(rows []SFSRow) error

	// SelectRows executes the query on the table.
	SelectRows(query SFSQuery) error

	// DeleteRows deletes the rows from the table.
	DeleteRows(rows []SFSRow) error
}
