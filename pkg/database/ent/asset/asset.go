// Code generated by entc, DO NOT EDIT.

package asset

import (
	"time"
)

const (
	// Label holds the string label denoting the asset type in the database.
	Label = "asset"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldProtocol holds the string denoting the protocol field in the database.
	FieldProtocol = "protocol"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// FieldAccountType holds the string denoting the account_type field in the database.
	FieldAccountType = "account_type"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldCredentialID holds the string denoting the credential_id field in the database.
	FieldCredentialID = "credential_id"
	// FieldPrivateKey holds the string denoting the private_key field in the database.
	FieldPrivateKey = "private_key"
	// FieldPassphrase holds the string denoting the passphrase field in the database.
	FieldPassphrase = "passphrase"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"

	// EdgeSessions holds the string denoting the sessions edge name in mutations.
	EdgeSessions = "sessions"

	// Table holds the table name of the asset in the database.
	Table = "assets"
	// SessionsTable is the table the holds the sessions relation/edge.
	SessionsTable = "assets"
	// SessionsInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionsInverseTable = "sessions"
	// SessionsColumn is the table column denoting the sessions relation/edge.
	SessionsColumn = "session_assets"
)

// Columns holds all SQL columns for asset fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldIP,
	FieldProtocol,
	FieldPort,
	FieldAccountType,
	FieldUsername,
	FieldPassword,
	FieldCredentialID,
	FieldPrivateKey,
	FieldPassphrase,
	FieldDescription,
	FieldActive,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTags,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Asset type.
var ForeignKeys = []string{
	"session_assets",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
