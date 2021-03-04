// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/credential"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/predicate"
)

// CredentialUpdate is the builder for updating Credential entities.
type CredentialUpdate struct {
	config
	hooks    []Hook
	mutation *CredentialMutation
}

// Where adds a new predicate for the CredentialUpdate builder.
func (cu *CredentialUpdate) Where(ps ...predicate.Credential) *CredentialUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CredentialUpdate) SetName(s string) *CredentialUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetType sets the "type" field.
func (cu *CredentialUpdate) SetType(s string) *CredentialUpdate {
	cu.mutation.SetType(s)
	return cu
}

// SetUsername sets the "username" field.
func (cu *CredentialUpdate) SetUsername(s string) *CredentialUpdate {
	cu.mutation.SetUsername(s)
	return cu
}

// SetPassword sets the "password" field.
func (cu *CredentialUpdate) SetPassword(s string) *CredentialUpdate {
	cu.mutation.SetPassword(s)
	return cu
}

// SetPrivateKey sets the "private_key" field.
func (cu *CredentialUpdate) SetPrivateKey(s string) *CredentialUpdate {
	cu.mutation.SetPrivateKey(s)
	return cu
}

// SetPassphrase sets the "passphrase" field.
func (cu *CredentialUpdate) SetPassphrase(s string) *CredentialUpdate {
	cu.mutation.SetPassphrase(s)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CredentialUpdate) SetCreatedAt(t time.Time) *CredentialUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CredentialUpdate) SetNillableCreatedAt(t *time.Time) *CredentialUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CredentialUpdate) SetUpdatedAt(t time.Time) *CredentialUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// Mutation returns the CredentialMutation object of the builder.
func (cu *CredentialUpdate) Mutation() *CredentialMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CredentialUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CredentialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CredentialUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CredentialUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CredentialUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CredentialUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := credential.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CredentialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   credential.Table,
			Columns: credential.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: credential.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldName,
		})
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldType,
		})
	}
	if value, ok := cu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldUsername,
		})
	}
	if value, ok := cu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldPassword,
		})
	}
	if value, ok := cu.mutation.PrivateKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldPrivateKey,
		})
	}
	if value, ok := cu.mutation.Passphrase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldPassphrase,
		})
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: credential.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: credential.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{credential.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CredentialUpdateOne is the builder for updating a single Credential entity.
type CredentialUpdateOne struct {
	config
	hooks    []Hook
	mutation *CredentialMutation
}

// SetName sets the "name" field.
func (cuo *CredentialUpdateOne) SetName(s string) *CredentialUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetType sets the "type" field.
func (cuo *CredentialUpdateOne) SetType(s string) *CredentialUpdateOne {
	cuo.mutation.SetType(s)
	return cuo
}

// SetUsername sets the "username" field.
func (cuo *CredentialUpdateOne) SetUsername(s string) *CredentialUpdateOne {
	cuo.mutation.SetUsername(s)
	return cuo
}

// SetPassword sets the "password" field.
func (cuo *CredentialUpdateOne) SetPassword(s string) *CredentialUpdateOne {
	cuo.mutation.SetPassword(s)
	return cuo
}

// SetPrivateKey sets the "private_key" field.
func (cuo *CredentialUpdateOne) SetPrivateKey(s string) *CredentialUpdateOne {
	cuo.mutation.SetPrivateKey(s)
	return cuo
}

// SetPassphrase sets the "passphrase" field.
func (cuo *CredentialUpdateOne) SetPassphrase(s string) *CredentialUpdateOne {
	cuo.mutation.SetPassphrase(s)
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CredentialUpdateOne) SetCreatedAt(t time.Time) *CredentialUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CredentialUpdateOne) SetNillableCreatedAt(t *time.Time) *CredentialUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CredentialUpdateOne) SetUpdatedAt(t time.Time) *CredentialUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// Mutation returns the CredentialMutation object of the builder.
func (cuo *CredentialUpdateOne) Mutation() *CredentialMutation {
	return cuo.mutation
}

// Save executes the query and returns the updated Credential entity.
func (cuo *CredentialUpdateOne) Save(ctx context.Context) (*Credential, error) {
	var (
		err  error
		node *Credential
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CredentialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CredentialUpdateOne) SaveX(ctx context.Context) *Credential {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CredentialUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CredentialUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CredentialUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := credential.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CredentialUpdateOne) sqlSave(ctx context.Context) (_node *Credential, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   credential.Table,
			Columns: credential.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: credential.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Credential.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldName,
		})
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldType,
		})
	}
	if value, ok := cuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldUsername,
		})
	}
	if value, ok := cuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldPassword,
		})
	}
	if value, ok := cuo.mutation.PrivateKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldPrivateKey,
		})
	}
	if value, ok := cuo.mutation.Passphrase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldPassphrase,
		})
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: credential.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: credential.FieldUpdatedAt,
		})
	}
	_node = &Credential{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{credential.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
