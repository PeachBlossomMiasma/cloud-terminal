// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccessSecuritysColumns holds the columns for the "accessSecuritys" table.
	AccessSecuritysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "rule", Type: field.TypeString},
		{Name: "ip", Type: field.TypeString},
		{Name: "source", Type: field.TypeString},
		{Name: "priority", Type: field.TypeInt64},
	}
	// AccessSecuritysTable holds the schema information for the "accessSecuritys" table.
	AccessSecuritysTable = &schema.Table{
		Name:        "accessSecuritys",
		Columns:     AccessSecuritysColumns,
		PrimaryKey:  []*schema.Column{AccessSecuritysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
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
		{Name: "tags", Type: field.TypeString},
		{Name: "access_security_assets", Type: field.TypeString, Nullable: true},
		{Name: "session_assets", Type: field.TypeString, Nullable: true},
		{Name: "user_assets", Type: field.TypeString, Nullable: true},
	}
	// AssetsTable holds the schema information for the "assets" table.
	AssetsTable = &schema.Table{
		Name:       "assets",
		Columns:    AssetsColumns,
		PrimaryKey: []*schema.Column{AssetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "assets_accessSecuritys_assets",
				Columns:    []*schema.Column{AssetsColumns[14]},
				RefColumns: []*schema.Column{AccessSecuritysColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "assets_sessions_assets",
				Columns:    []*schema.Column{AssetsColumns[15]},
				RefColumns: []*schema.Column{SessionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "assets_users_assets",
				Columns:    []*schema.Column{AssetsColumns[16]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommandsColumns holds the columns for the "commands" table.
	CommandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "content", Type: field.TypeJSON},
	}
	// CommandsTable holds the schema information for the "commands" table.
	CommandsTable = &schema.Table{
		Name:        "commands",
		Columns:     CommandsColumns,
		PrimaryKey:  []*schema.Column{CommandsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
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
	}
	// CredentialsTable holds the schema information for the "credentials" table.
	CredentialsTable = &schema.Table{
		Name:        "credentials",
		Columns:     CredentialsColumns,
		PrimaryKey:  []*schema.Column{CredentialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:        "groups",
		Columns:     GroupsColumns,
		PrimaryKey:  []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// JobsColumns holds the columns for the "jobs" table.
	JobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "cronjobid", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "func", Type: field.TypeString},
		{Name: "cron", Type: field.TypeString},
		{Name: "mode", Type: field.TypeString},
		{Name: "resource_ids", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeString},
	}
	// JobsTable holds the schema information for the "jobs" table.
	JobsTable = &schema.Table{
		Name:        "jobs",
		Columns:     JobsColumns,
		PrimaryKey:  []*schema.Column{JobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// LoginlogsColumns holds the columns for the "loginlogs" table.
	LoginlogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "user_id", Type: field.TypeString},
		{Name: "client_ip", Type: field.TypeString},
		{Name: "clent_uset_agent", Type: field.TypeString},
		{Name: "login_time", Type: field.TypeTime},
		{Name: "logout_time", Type: field.TypeTime},
		{Name: "remember", Type: field.TypeBool},
	}
	// LoginlogsTable holds the schema information for the "loginlogs" table.
	LoginlogsTable = &schema.Table{
		Name:        "loginlogs",
		Columns:     LoginlogsColumns,
		PrimaryKey:  []*schema.Column{LoginlogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
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
		{Name: "mode", Type: field.TypeString},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:        "sessions",
		Columns:     SessionsColumns,
		PrimaryKey:  []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TestsColumns holds the columns for the "tests" table.
	TestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
	}
	// TestsTable holds the schema information for the "tests" table.
	TestsTable = &schema.Table{
		Name:        "tests",
		Columns:     TestsColumns,
		PrimaryKey:  []*schema.Column{TestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString},
		{Name: "totp_secret", Type: field.TypeString},
		{Name: "online", Type: field.TypeBool},
		{Name: "enable", Type: field.TypeBool},
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
				Symbol:     "users_verifications_users",
				Columns:    []*schema.Column{UsersColumns[11]},
				RefColumns: []*schema.Column{VerificationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
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
	}
	// GroupUsersColumns holds the columns for the "group_users" table.
	GroupUsersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// GroupUsersTable holds the schema information for the "group_users" table.
	GroupUsersTable = &schema.Table{
		Name:       "group_users",
		Columns:    GroupUsersColumns,
		PrimaryKey: []*schema.Column{GroupUsersColumns[0], GroupUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_users_group_id",
				Columns:    []*schema.Column{GroupUsersColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_users_user_id",
				Columns:    []*schema.Column{GroupUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccessSecuritysTable,
		AssetsTable,
		CommandsTable,
		CredentialsTable,
		GroupsTable,
		JobsTable,
		LoginlogsTable,
		PropertiesTable,
		ResourceSharersTable,
		SessionsTable,
		TestsTable,
		UsersTable,
		VerificationsTable,
		GroupUsersTable,
	}
)

func init() {
	AccessSecuritysTable.Annotation = &entsql.Annotation{
		Table: "accessSecuritys",
	}
	AssetsTable.ForeignKeys[0].RefTable = AccessSecuritysTable
	AssetsTable.ForeignKeys[1].RefTable = SessionsTable
	AssetsTable.ForeignKeys[2].RefTable = UsersTable
	AssetsTable.Annotation = &entsql.Annotation{
		Table: "assets",
	}
	CommandsTable.Annotation = &entsql.Annotation{
		Table: "commands",
	}
	CredentialsTable.Annotation = &entsql.Annotation{
		Table: "credentials",
	}
	GroupsTable.Annotation = &entsql.Annotation{
		Table: "groups",
	}
	JobsTable.Annotation = &entsql.Annotation{
		Table: "jobs",
	}
	LoginlogsTable.Annotation = &entsql.Annotation{
		Table: "loginlogs",
	}
	PropertiesTable.Annotation = &entsql.Annotation{
		Table: "properties",
	}
	ResourceSharersTable.Annotation = &entsql.Annotation{
		Table: "resourceSharers",
	}
	SessionsTable.Annotation = &entsql.Annotation{
		Table: "sessions",
	}
	UsersTable.ForeignKeys[0].RefTable = VerificationsTable
	UsersTable.Annotation = &entsql.Annotation{
		Table: "users",
	}
	VerificationsTable.Annotation = &entsql.Annotation{
		Table: "verifications",
	}
	GroupUsersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersTable.ForeignKeys[1].RefTable = UsersTable
}
