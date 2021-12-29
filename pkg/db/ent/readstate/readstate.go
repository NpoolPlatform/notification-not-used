// Code generated by entc, DO NOT EDIT.

package readstate

const (
	// Label holds the string label denoting the readstate type in the database.
	Label = "read_state"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the readstate in the database.
	Table = "read_states"
)

// Columns holds all SQL columns for readstate fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
