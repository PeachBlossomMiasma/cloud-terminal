// Code generated by entc, DO NOT EDIT.

package usergroup

const (
	// Label holds the string label denoting the usergroup type in the database.
	Label = "user_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"

	// Table holds the table name of the usergroup in the database.
	Table = "user_groups"
)

// Columns holds all SQL columns for usergroup fields.
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