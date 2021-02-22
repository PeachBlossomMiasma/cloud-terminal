// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/user"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/usergroup"
)

// UserGroupCreate is the builder for creating a UserGroup entity.
type UserGroupCreate struct {
	config
	mutation *UserGroupMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ugc *UserGroupCreate) SetName(s string) *UserGroupCreate {
	ugc.mutation.SetName(s)
	return ugc
}

// SetCreatedAt sets the "created_at" field.
func (ugc *UserGroupCreate) SetCreatedAt(t time.Time) *UserGroupCreate {
	ugc.mutation.SetCreatedAt(t)
	return ugc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ugc *UserGroupCreate) SetNillableCreatedAt(t *time.Time) *UserGroupCreate {
	if t != nil {
		ugc.SetCreatedAt(*t)
	}
	return ugc
}

// SetUpdatedAt sets the "updated_at" field.
func (ugc *UserGroupCreate) SetUpdatedAt(t time.Time) *UserGroupCreate {
	ugc.mutation.SetUpdatedAt(t)
	return ugc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ugc *UserGroupCreate) SetNillableUpdatedAt(t *time.Time) *UserGroupCreate {
	if t != nil {
		ugc.SetUpdatedAt(*t)
	}
	return ugc
}

// SetID sets the "id" field.
func (ugc *UserGroupCreate) SetID(s string) *UserGroupCreate {
	ugc.mutation.SetID(s)
	return ugc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (ugc *UserGroupCreate) AddUserIDs(ids ...string) *UserGroupCreate {
	ugc.mutation.AddUserIDs(ids...)
	return ugc
}

// AddUsers adds the "users" edges to the User entity.
func (ugc *UserGroupCreate) AddUsers(u ...*User) *UserGroupCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ugc.AddUserIDs(ids...)
}

// Mutation returns the UserGroupMutation object of the builder.
func (ugc *UserGroupCreate) Mutation() *UserGroupMutation {
	return ugc.mutation
}

// Save creates the UserGroup in the database.
func (ugc *UserGroupCreate) Save(ctx context.Context) (*UserGroup, error) {
	var (
		err  error
		node *UserGroup
	)
	ugc.defaults()
	if len(ugc.hooks) == 0 {
		if err = ugc.check(); err != nil {
			return nil, err
		}
		node, err = ugc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ugc.check(); err != nil {
				return nil, err
			}
			ugc.mutation = mutation
			node, err = ugc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ugc.hooks) - 1; i >= 0; i-- {
			mut = ugc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ugc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ugc *UserGroupCreate) SaveX(ctx context.Context) *UserGroup {
	v, err := ugc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ugc *UserGroupCreate) defaults() {
	if _, ok := ugc.mutation.CreatedAt(); !ok {
		v := usergroup.DefaultCreatedAt()
		ugc.mutation.SetCreatedAt(v)
	}
	if _, ok := ugc.mutation.UpdatedAt(); !ok {
		v := usergroup.DefaultUpdatedAt()
		ugc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ugc *UserGroupCreate) check() error {
	if _, ok := ugc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := ugc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := ugc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (ugc *UserGroupCreate) sqlSave(ctx context.Context) (*UserGroup, error) {
	_node, _spec := ugc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ugc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (ugc *UserGroupCreate) createSpec() (*UserGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &UserGroup{config: ugc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usergroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: usergroup.FieldID,
			},
		}
	)
	if id, ok := ugc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ugc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usergroup.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ugc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usergroup.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ugc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usergroup.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ugc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.UsersTable,
			Columns: usergroup.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserGroupCreateBulk is the builder for creating many UserGroup entities in bulk.
type UserGroupCreateBulk struct {
	config
	builders []*UserGroupCreate
}

// Save creates the UserGroup entities in the database.
func (ugcb *UserGroupCreateBulk) Save(ctx context.Context) ([]*UserGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ugcb.builders))
	nodes := make([]*UserGroup, len(ugcb.builders))
	mutators := make([]Mutator, len(ugcb.builders))
	for i := range ugcb.builders {
		func(i int, root context.Context) {
			builder := ugcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserGroupMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ugcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ugcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ugcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ugcb *UserGroupCreateBulk) SaveX(ctx context.Context) []*UserGroup {
	v, err := ugcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
