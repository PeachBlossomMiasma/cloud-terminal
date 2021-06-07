// Code generated by entc, DO NOT EDIT.

package loginlog

import (
	"time"
)

const (
	// Label holds the string label denoting the loginlog type in the database.
	Label = "login_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldClientIP holds the string denoting the client_ip field in the database.
	FieldClientIP = "client_ip"
	// FieldClentUsetAgent holds the string denoting the clent_uset_agent field in the database.
	FieldClentUsetAgent = "clent_uset_agent"
	// FieldLoginTime holds the string denoting the login_time field in the database.
	FieldLoginTime = "login_time"
	// FieldLogoutTime holds the string denoting the logout_time field in the database.
	FieldLogoutTime = "logout_time"
	// FieldRemember holds the string denoting the remember field in the database.
	FieldRemember = "remember"
	// Table holds the table name of the loginlog in the database.
	Table = "loginlogs"
)

// Columns holds all SQL columns for loginlog fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldClientIP,
	FieldClentUsetAgent,
	FieldLoginTime,
	FieldLogoutTime,
	FieldRemember,
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

var (
	// DefaultLoginTime holds the default value on creation for the "login_time" field.
	DefaultLoginTime func() time.Time
	// DefaultLogoutTime holds the default value on creation for the "logout_time" field.
	DefaultLogoutTime func() time.Time
	// UpdateDefaultLogoutTime holds the default value on update for the "logout_time" field.
	UpdateDefaultLogoutTime func() time.Time
)
