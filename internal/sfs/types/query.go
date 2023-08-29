package types

// SFSQuery is the interface for a query.
type SFSQuery interface {
	// Table returns the table name.
	Table() (string)

	// Columns returns the table columns.
	Columns() ([]SFSColumn, error)

	// Where returns the where clause.
	Where(k string, v interface{}) (error)

	// Or returns the or clause.
	Or(k string, v interface{}) (error)

	// Xor returns the xor clause.
	Xor(k string, v interface{}) (error)

	// And returns the and clause.
	And(k string, v interface{}) (error)

	// Not returns the not clause.
	Not(k string, v interface{}) (error)

	// InsertInto returns the insert into clause.
	InsertInto(cols []SFSColumn, vs []interface{}) (error)

	// String returns the constructed query string.
	String() (string, error)
}
