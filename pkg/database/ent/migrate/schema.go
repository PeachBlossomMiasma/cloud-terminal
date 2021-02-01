// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AssetsColumns holds the columns for the "assets" table.
	AssetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "ip", Type: field.TypeString},
		{Name: "protocol", Type: field.TypeString},
		{Name: "port", Type: field.TypeInt},
		{Name: "account_type", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "credential_id", Type: field.TypeString},
		{Name: "private_key", Type: field.TypeString},
		{Name: "passphrase", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "active", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "tags", Type: field.TypeString},
		{Name: "session_assets", Type: field.TypeString, Nullable: true},
	}
	// AssetsTable holds the schema information for the "assets" table.
	AssetsTable = &schema.Table{
		Name:       "assets",
		Columns:    AssetsColumns,
		PrimaryKey: []*schema.Column{AssetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "assets_sessions_assets",
				Columns: []*schema.Column{AssetsColumns[16]},

				RefColumns: []*schema.Column{SessionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "assets"},
	}
	// CommandsColumns holds the columns for the "commands" table.
	CommandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "content", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CommandsTable holds the schema information for the "commands" table.
	CommandsTable = &schema.Table{
		Name:        "commands",
		Columns:     CommandsColumns,
		PrimaryKey:  []*schema.Column{CommandsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "commands"},
	}
	// CredentialsColumns holds the columns for the "credentials" table.
	CredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "private_key", Type: field.TypeString},
		{Name: "passphrase", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CredentialsTable holds the schema information for the "credentials" table.
	CredentialsTable = &schema.Table{
		Name:        "credentials",
		Columns:     CredentialsColumns,
		PrimaryKey:  []*schema.Column{CredentialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "credentials"},
	}
	// PropertiesColumns holds the columns for the "properties" table.
	PropertiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "value", Type: field.TypeString},
	}
	// PropertiesTable holds the schema information for the "properties" table.
	PropertiesTable = &schema.Table{
		Name:        "properties",
		Columns:     PropertiesColumns,
		PrimaryKey:  []*schema.Column{PropertiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "properties"},
	}
	// ResourceSharersColumns holds the columns for the "resourceSharers" table.
	ResourceSharersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "resource_id", Type: field.TypeString},
		{Name: "resource_type", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "user_group_id", Type: field.TypeString},
	}
	// ResourceSharersTable holds the schema information for the "resourceSharers" table.
	ResourceSharersTable = &schema.Table{
		Name:        "resourceSharers",
		Columns:     ResourceSharersColumns,
		PrimaryKey:  []*schema.Column{ResourceSharersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "resourceSharers"},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "protocol", Type: field.TypeString},
		{Name: "ip", Type: field.TypeString},
		{Name: "port", Type: field.TypeInt},
		{Name: "connection_id", Type: field.TypeString},
		{Name: "asset_id", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "creator", Type: field.TypeString},
		{Name: "client_ip", Type: field.TypeString},
		{Name: "width", Type: field.TypeInt},
		{Name: "height", Type: field.TypeInt},
		{Name: "status", Type: field.TypeString},
		{Name: "recording", Type: field.TypeString},
		{Name: "private_key", Type: field.TypeString},
		{Name: "passphrase", Type: field.TypeString},
		{Name: "code", Type: field.TypeInt},
		{Name: "message", Type: field.TypeString},
		{Name: "connected", Type: field.TypeTime},
		{Name: "disconnected", Type: field.TypeTime},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:        "sessions",
		Columns:     SessionsColumns,
		PrimaryKey:  []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "sessions"},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString},
		{Name: "totp_secret", Type: field.TypeString},
		{Name: "online", Type: field.TypeBool},
		{Name: "enable", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString},
		{Name: "verification_users", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_verifications_users",
				Columns: []*schema.Column{UsersColumns[10]},

				RefColumns: []*schema.Column{VerificationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "users"},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:        "user_groups",
		Columns:     UserGroupsColumns,
		PrimaryKey:  []*schema.Column{UserGroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "user_groups"},
	}
	// VerificationsColumns holds the columns for the "verifications" table.
	VerificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "client_ip", Type: field.TypeString},
		{Name: "client_user_agent", Type: field.TypeString},
		{Name: "login_time", Type: field.TypeTime},
		{Name: "logout_time", Type: field.TypeTime},
		{Name: "remember", Type: field.TypeBool},
	}
	// VerificationsTable holds the schema information for the "verifications" table.
	VerificationsTable = &schema.Table{
		Name:        "verifications",
		Columns:     VerificationsColumns,
		PrimaryKey:  []*schema.Column{VerificationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "verifications"},
	}
	// UserGroupUsersColumns holds the columns for the "user_group_users" table.
	UserGroupUsersColumns = []*schema.Column{
		{Name: "user_group_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// UserGroupUsersTable holds the schema information for the "user_group_users" table.
	UserGroupUsersTable = &schema.Table{
		Name:       "user_group_users",
		Columns:    UserGroupUsersColumns,
		PrimaryKey: []*schema.Column{UserGroupUsersColumns[0], UserGroupUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_group_users_user_group_id",
				Columns: []*schema.Column{UserGroupUsersColumns[0]},

				RefColumns: []*schema.Column{UserGroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "user_group_users_user_id",
				Columns: []*schema.Column{UserGroupUsersColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AssetsTable,
		CommandsTable,
		CredentialsTable,
		PropertiesTable,
		ResourceSharersTable,
		SessionsTable,
		UsersTable,
		UserGroupsTable,
		VerificationsTable,
		UserGroupUsersTable,
	}
)

func init() {
	AssetsTable.ForeignKeys[0].RefTable = SessionsTable
	UsersTable.ForeignKeys[0].RefTable = VerificationsTable
	UserGroupUsersTable.ForeignKeys[0].RefTable = UserGroupsTable
	UserGroupUsersTable.ForeignKeys[1].RefTable = UsersTable
}
