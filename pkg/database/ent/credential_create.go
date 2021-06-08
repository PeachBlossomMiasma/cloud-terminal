// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/credential"
)

// CredentialCreate is the builder for creating a Credential entity.
type CredentialCreate struct {
	config
	mutation *CredentialMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CredentialCreate) SetName(s string) *CredentialCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetType sets the "type" field.
func (cc *CredentialCreate) SetType(s string) *CredentialCreate {
	cc.mutation.SetType(s)
	return cc
}

// SetUsername sets the "username" field.
func (cc *CredentialCreate) SetUsername(s string) *CredentialCreate {
	cc.mutation.SetUsername(s)
	return cc
}

// SetPassword sets the "password" field.
func (cc *CredentialCreate) SetPassword(s string) *CredentialCreate {
	cc.mutation.SetPassword(s)
	return cc
}

// SetPrivateKey sets the "private_key" field.
func (cc *CredentialCreate) SetPrivateKey(s string) *CredentialCreate {
	cc.mutation.SetPrivateKey(s)
	return cc
}

// SetPassphrase sets the "passphrase" field.
func (cc *CredentialCreate) SetPassphrase(s string) *CredentialCreate {
	cc.mutation.SetPassphrase(s)
	return cc
}

// SetID sets the "id" field.
func (cc *CredentialCreate) SetID(s string) *CredentialCreate {
	cc.mutation.SetID(s)
	return cc
}

// Mutation returns the CredentialMutation object of the builder.
func (cc *CredentialCreate) Mutation() *CredentialMutation {
	return cc.mutation
}

// Save creates the Credential in the database.
func (cc *CredentialCreate) Save(ctx context.Context) (*Credential, error) {
	var (
		err  error
		node *Credential
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CredentialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CredentialCreate) SaveX(ctx context.Context) *Credential {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (cc *CredentialCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	if _, ok := cc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New("ent: missing required field \"username\"")}
	}
	if _, ok := cc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if _, ok := cc.mutation.PrivateKey(); !ok {
		return &ValidationError{Name: "private_key", err: errors.New("ent: missing required field \"private_key\"")}
	}
	if _, ok := cc.mutation.Passphrase(); !ok {
		return &ValidationError{Name: "passphrase", err: errors.New("ent: missing required field \"passphrase\"")}
	}
	return nil
}

func (cc *CredentialCreate) sqlSave(ctx context.Context) (*Credential, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (cc *CredentialCreate) createSpec() (*Credential, *sqlgraph.CreateSpec) {
	var (
		_node = &Credential{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: credential.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: credential.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldType,
		})
		_node.Type = value
	}
	if value, ok := cc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := cc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := cc.mutation.PrivateKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldPrivateKey,
		})
		_node.PrivateKey = value
	}
	if value, ok := cc.mutation.Passphrase(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldPassphrase,
		})
		_node.Passphrase = value
	}
	return _node, _spec
}

// CredentialCreateBulk is the builder for creating many Credential entities in bulk.
type CredentialCreateBulk struct {
	config
	builders []*CredentialCreate
}

// Save creates the Credential entities in the database.
func (ccb *CredentialCreateBulk) Save(ctx context.Context) ([]*Credential, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Credential, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CredentialMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CredentialCreateBulk) SaveX(ctx context.Context) []*Credential {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
