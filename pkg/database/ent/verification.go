// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/verification"
)

// Verification is the model entity for the Verification schema.
type Verification struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ClientIP holds the value of the "client_ip" field.
	ClientIP string `json:"client_ip,omitempty"`
	// ClientUserAgent holds the value of the "clientUserAgent" field.
	ClientUserAgent string `json:"clientUserAgent,omitempty"`
	// LoginTime holds the value of the "login_time" field.
	LoginTime time.Time `json:"login_time,omitempty"`
	// LogoutTime holds the value of the "logout_time" field.
	LogoutTime time.Time `json:"logout_time,omitempty"`
	// Remember holds the value of the "remember" field.
	Remember bool `json:"remember,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VerificationQuery when eager-loading is set.
	Edges VerificationEdges `json:"edges"`
}

// VerificationEdges holds the relations/edges for other nodes in the graph.
type VerificationEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e VerificationEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Verification) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case verification.FieldRemember:
			values[i] = new(sql.NullBool)
		case verification.FieldID, verification.FieldClientIP, verification.FieldClientUserAgent:
			values[i] = new(sql.NullString)
		case verification.FieldCreatedAt, verification.FieldUpdatedAt, verification.FieldLoginTime, verification.FieldLogoutTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Verification", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Verification fields.
func (v *Verification) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case verification.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				v.ID = value.String
			}
		case verification.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		case verification.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				v.UpdatedAt = value.Time
			}
		case verification.FieldClientIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_ip", values[i])
			} else if value.Valid {
				v.ClientIP = value.String
			}
		case verification.FieldClientUserAgent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field clientUserAgent", values[i])
			} else if value.Valid {
				v.ClientUserAgent = value.String
			}
		case verification.FieldLoginTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field login_time", values[i])
			} else if value.Valid {
				v.LoginTime = value.Time
			}
		case verification.FieldLogoutTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field logout_time", values[i])
			} else if value.Valid {
				v.LogoutTime = value.Time
			}
		case verification.FieldRemember:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field remember", values[i])
			} else if value.Valid {
				v.Remember = value.Bool
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Verification entity.
func (v *Verification) QueryUsers() *UserQuery {
	return (&VerificationClient{config: v.config}).QueryUsers(v)
}

// Update returns a builder for updating this Verification.
// Note that you need to call Verification.Unwrap() before calling this method if this Verification
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Verification) Update() *VerificationUpdateOne {
	return (&VerificationClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the Verification entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Verification) Unwrap() *Verification {
	tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Verification is not a transactional entity")
	}
	v.config.driver = tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Verification) String() string {
	var builder strings.Builder
	builder.WriteString("Verification(")
	builder.WriteString(fmt.Sprintf("id=%v", v.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(v.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", client_ip=")
	builder.WriteString(v.ClientIP)
	builder.WriteString(", clientUserAgent=")
	builder.WriteString(v.ClientUserAgent)
	builder.WriteString(", login_time=")
	builder.WriteString(v.LoginTime.Format(time.ANSIC))
	builder.WriteString(", logout_time=")
	builder.WriteString(v.LogoutTime.Format(time.ANSIC))
	builder.WriteString(", remember=")
	builder.WriteString(fmt.Sprintf("%v", v.Remember))
	builder.WriteByte(')')
	return builder.String()
}

// Verifications is a parsable slice of Verification.
type Verifications []*Verification

func (v Verifications) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}
