// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/asset"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/predicate"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/session"
)

// AssetUpdate is the builder for updating Asset entities.
type AssetUpdate struct {
	config
	hooks    []Hook
	mutation *AssetMutation
}

// Where adds a new predicate for the AssetUpdate builder.
func (au *AssetUpdate) Where(ps ...predicate.Asset) *AssetUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetName sets the "name" field.
func (au *AssetUpdate) SetName(s string) *AssetUpdate {
	au.mutation.SetName(s)
	return au
}

// SetIP sets the "ip" field.
func (au *AssetUpdate) SetIP(s string) *AssetUpdate {
	au.mutation.SetIP(s)
	return au
}

// SetProtocol sets the "protocol" field.
func (au *AssetUpdate) SetProtocol(s string) *AssetUpdate {
	au.mutation.SetProtocol(s)
	return au
}

// SetPort sets the "port" field.
func (au *AssetUpdate) SetPort(i int) *AssetUpdate {
	au.mutation.ResetPort()
	au.mutation.SetPort(i)
	return au
}

// AddPort adds i to the "port" field.
func (au *AssetUpdate) AddPort(i int) *AssetUpdate {
	au.mutation.AddPort(i)
	return au
}

// SetAccountType sets the "account_type" field.
func (au *AssetUpdate) SetAccountType(s string) *AssetUpdate {
	au.mutation.SetAccountType(s)
	return au
}

// SetUsername sets the "username" field.
func (au *AssetUpdate) SetUsername(s string) *AssetUpdate {
	au.mutation.SetUsername(s)
	return au
}

// SetPassword sets the "password" field.
func (au *AssetUpdate) SetPassword(s string) *AssetUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetCredentialID sets the "credential_id" field.
func (au *AssetUpdate) SetCredentialID(s string) *AssetUpdate {
	au.mutation.SetCredentialID(s)
	return au
}

// SetPrivateKey sets the "private_key" field.
func (au *AssetUpdate) SetPrivateKey(s string) *AssetUpdate {
	au.mutation.SetPrivateKey(s)
	return au
}

// SetPassphrase sets the "passphrase" field.
func (au *AssetUpdate) SetPassphrase(s string) *AssetUpdate {
	au.mutation.SetPassphrase(s)
	return au
}

// SetDescription sets the "description" field.
func (au *AssetUpdate) SetDescription(s string) *AssetUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetActive sets the "active" field.
func (au *AssetUpdate) SetActive(b bool) *AssetUpdate {
	au.mutation.SetActive(b)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AssetUpdate) SetCreatedAt(t time.Time) *AssetUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AssetUpdate) SetNillableCreatedAt(t *time.Time) *AssetUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AssetUpdate) SetUpdatedAt(t time.Time) *AssetUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetTags sets the "tags" field.
func (au *AssetUpdate) SetTags(s string) *AssetUpdate {
	au.mutation.SetTags(s)
	return au
}

// SetSessionsID sets the "sessions" edge to the Session entity by ID.
func (au *AssetUpdate) SetSessionsID(id string) *AssetUpdate {
	au.mutation.SetSessionsID(id)
	return au
}

// SetNillableSessionsID sets the "sessions" edge to the Session entity by ID if the given value is not nil.
func (au *AssetUpdate) SetNillableSessionsID(id *string) *AssetUpdate {
	if id != nil {
		au = au.SetSessionsID(*id)
	}
	return au
}

// SetSessions sets the "sessions" edge to the Session entity.
func (au *AssetUpdate) SetSessions(s *Session) *AssetUpdate {
	return au.SetSessionsID(s.ID)
}

// Mutation returns the AssetMutation object of the builder.
func (au *AssetUpdate) Mutation() *AssetMutation {
	return au.mutation
}

// ClearSessions clears the "sessions" edge to the Session entity.
func (au *AssetUpdate) ClearSessions() *AssetUpdate {
	au.mutation.ClearSessions()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AssetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AssetUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AssetUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AssetUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AssetUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := asset.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

func (au *AssetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asset.Table,
			Columns: asset.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: asset.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldName,
		})
	}
	if value, ok := au.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldIP,
		})
	}
	if value, ok := au.mutation.Protocol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldProtocol,
		})
	}
	if value, ok := au.mutation.Port(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: asset.FieldPort,
		})
	}
	if value, ok := au.mutation.AddedPort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: asset.FieldPort,
		})
	}
	if value, ok := au.mutation.AccountType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldAccountType,
		})
	}
	if value, ok := au.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldUsername,
		})
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldPassword,
		})
	}
	if value, ok := au.mutation.CredentialID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldCredentialID,
		})
	}
	if value, ok := au.mutation.PrivateKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldPrivateKey,
		})
	}
	if value, ok := au.mutation.Passphrase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldPassphrase,
		})
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldDescription,
		})
	}
	if value, ok := au.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: asset.FieldActive,
		})
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asset.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asset.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldTags,
		})
	}
	if au.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.SessionsTable,
			Columns: []string{asset.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: session.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.SessionsTable,
			Columns: []string{asset.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asset.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AssetUpdateOne is the builder for updating a single Asset entity.
type AssetUpdateOne struct {
	config
	hooks    []Hook
	mutation *AssetMutation
}

// SetName sets the "name" field.
func (auo *AssetUpdateOne) SetName(s string) *AssetUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetIP sets the "ip" field.
func (auo *AssetUpdateOne) SetIP(s string) *AssetUpdateOne {
	auo.mutation.SetIP(s)
	return auo
}

// SetProtocol sets the "protocol" field.
func (auo *AssetUpdateOne) SetProtocol(s string) *AssetUpdateOne {
	auo.mutation.SetProtocol(s)
	return auo
}

// SetPort sets the "port" field.
func (auo *AssetUpdateOne) SetPort(i int) *AssetUpdateOne {
	auo.mutation.ResetPort()
	auo.mutation.SetPort(i)
	return auo
}

// AddPort adds i to the "port" field.
func (auo *AssetUpdateOne) AddPort(i int) *AssetUpdateOne {
	auo.mutation.AddPort(i)
	return auo
}

// SetAccountType sets the "account_type" field.
func (auo *AssetUpdateOne) SetAccountType(s string) *AssetUpdateOne {
	auo.mutation.SetAccountType(s)
	return auo
}

// SetUsername sets the "username" field.
func (auo *AssetUpdateOne) SetUsername(s string) *AssetUpdateOne {
	auo.mutation.SetUsername(s)
	return auo
}

// SetPassword sets the "password" field.
func (auo *AssetUpdateOne) SetPassword(s string) *AssetUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetCredentialID sets the "credential_id" field.
func (auo *AssetUpdateOne) SetCredentialID(s string) *AssetUpdateOne {
	auo.mutation.SetCredentialID(s)
	return auo
}

// SetPrivateKey sets the "private_key" field.
func (auo *AssetUpdateOne) SetPrivateKey(s string) *AssetUpdateOne {
	auo.mutation.SetPrivateKey(s)
	return auo
}

// SetPassphrase sets the "passphrase" field.
func (auo *AssetUpdateOne) SetPassphrase(s string) *AssetUpdateOne {
	auo.mutation.SetPassphrase(s)
	return auo
}

// SetDescription sets the "description" field.
func (auo *AssetUpdateOne) SetDescription(s string) *AssetUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetActive sets the "active" field.
func (auo *AssetUpdateOne) SetActive(b bool) *AssetUpdateOne {
	auo.mutation.SetActive(b)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AssetUpdateOne) SetCreatedAt(t time.Time) *AssetUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AssetUpdateOne) SetNillableCreatedAt(t *time.Time) *AssetUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AssetUpdateOne) SetUpdatedAt(t time.Time) *AssetUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetTags sets the "tags" field.
func (auo *AssetUpdateOne) SetTags(s string) *AssetUpdateOne {
	auo.mutation.SetTags(s)
	return auo
}

// SetSessionsID sets the "sessions" edge to the Session entity by ID.
func (auo *AssetUpdateOne) SetSessionsID(id string) *AssetUpdateOne {
	auo.mutation.SetSessionsID(id)
	return auo
}

// SetNillableSessionsID sets the "sessions" edge to the Session entity by ID if the given value is not nil.
func (auo *AssetUpdateOne) SetNillableSessionsID(id *string) *AssetUpdateOne {
	if id != nil {
		auo = auo.SetSessionsID(*id)
	}
	return auo
}

// SetSessions sets the "sessions" edge to the Session entity.
func (auo *AssetUpdateOne) SetSessions(s *Session) *AssetUpdateOne {
	return auo.SetSessionsID(s.ID)
}

// Mutation returns the AssetMutation object of the builder.
func (auo *AssetUpdateOne) Mutation() *AssetMutation {
	return auo.mutation
}

// ClearSessions clears the "sessions" edge to the Session entity.
func (auo *AssetUpdateOne) ClearSessions() *AssetUpdateOne {
	auo.mutation.ClearSessions()
	return auo
}

// Save executes the query and returns the updated Asset entity.
func (auo *AssetUpdateOne) Save(ctx context.Context) (*Asset, error) {
	var (
		err  error
		node *Asset
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AssetUpdateOne) SaveX(ctx context.Context) *Asset {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AssetUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AssetUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AssetUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := asset.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

func (auo *AssetUpdateOne) sqlSave(ctx context.Context) (_node *Asset, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asset.Table,
			Columns: asset.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: asset.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Asset.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldName,
		})
	}
	if value, ok := auo.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldIP,
		})
	}
	if value, ok := auo.mutation.Protocol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldProtocol,
		})
	}
	if value, ok := auo.mutation.Port(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: asset.FieldPort,
		})
	}
	if value, ok := auo.mutation.AddedPort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: asset.FieldPort,
		})
	}
	if value, ok := auo.mutation.AccountType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldAccountType,
		})
	}
	if value, ok := auo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldUsername,
		})
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldPassword,
		})
	}
	if value, ok := auo.mutation.CredentialID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldCredentialID,
		})
	}
	if value, ok := auo.mutation.PrivateKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldPrivateKey,
		})
	}
	if value, ok := auo.mutation.Passphrase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldPassphrase,
		})
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldDescription,
		})
	}
	if value, ok := auo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: asset.FieldActive,
		})
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asset.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asset.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldTags,
		})
	}
	if auo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.SessionsTable,
			Columns: []string{asset.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: session.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.SessionsTable,
			Columns: []string{asset.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Asset{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asset.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
