// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// TotpSecret holds the value of the "totpSecret" field.
	TotpSecret string `json:"totpSecret,omitempty"`
	// Online holds the value of the "online" field.
	Online bool `json:"online,omitempty"`
	// Enable holds the value of the "enable" field.
	Enable bool `json:"enable,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges              UserEdges `json:"edges"`
	verification_users *string
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserGroups holds the value of the user_groups edge.
	UserGroups []*UserGroup
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserGroupsOrErr returns the UserGroups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserGroupsOrErr() ([]*UserGroup, error) {
	if e.loadedTypes[0] {
		return e.UserGroups, nil
	}
	return nil, &NotLoadedError{edge: "user_groups"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldOnline, user.FieldEnable:
			values[i] = &sql.NullBool{}
		case user.FieldID, user.FieldUsername, user.FieldPassword, user.FieldNickname, user.FieldTotpSecret, user.FieldType:
			values[i] = &sql.NullString{}
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = &sql.NullTime{}
		case user.ForeignKeys[0]: // verification_users
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				u.Nickname = value.String
			}
		case user.FieldTotpSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field totpSecret", values[i])
			} else if value.Valid {
				u.TotpSecret = value.String
			}
		case user.FieldOnline:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field online", values[i])
			} else if value.Valid {
				u.Online = value.Bool
			}
		case user.FieldEnable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enable", values[i])
			} else if value.Valid {
				u.Enable = value.Bool
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				u.Type = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field verification_users", values[i])
			} else if value.Valid {
				u.verification_users = new(string)
				*u.verification_users = value.String
			}
		}
	}
	return nil
}

// QueryUserGroups queries the "user_groups" edge of the User entity.
func (u *User) QueryUserGroups() *UserGroupQuery {
	return (&UserClient{config: u.config}).QueryUserGroups(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteString(", nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", totpSecret=")
	builder.WriteString(u.TotpSecret)
	builder.WriteString(", online=")
	builder.WriteString(fmt.Sprintf("%v", u.Online))
	builder.WriteString(", enable=")
	builder.WriteString(fmt.Sprintf("%v", u.Enable))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", type=")
	builder.WriteString(u.Type)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
