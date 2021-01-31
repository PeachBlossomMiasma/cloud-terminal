// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/asset"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/session"
)

// AssetCreate is the builder for creating a Asset entity.
type AssetCreate struct {
	config
	mutation *AssetMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AssetCreate) SetName(s string) *AssetCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetIP sets the "ip" field.
func (ac *AssetCreate) SetIP(s string) *AssetCreate {
	ac.mutation.SetIP(s)
	return ac
}

// SetProtocol sets the "protocol" field.
func (ac *AssetCreate) SetProtocol(s string) *AssetCreate {
	ac.mutation.SetProtocol(s)
	return ac
}

// SetPort sets the "port" field.
func (ac *AssetCreate) SetPort(i int) *AssetCreate {
	ac.mutation.SetPort(i)
	return ac
}

// SetAccountType sets the "accountType" field.
func (ac *AssetCreate) SetAccountType(s string) *AssetCreate {
	ac.mutation.SetAccountType(s)
	return ac
}

// SetUsername sets the "username" field.
func (ac *AssetCreate) SetUsername(s string) *AssetCreate {
	ac.mutation.SetUsername(s)
	return ac
}

// SetPassword sets the "password" field.
func (ac *AssetCreate) SetPassword(s string) *AssetCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetCredentialId sets the "credentialId" field.
func (ac *AssetCreate) SetCredentialId(s string) *AssetCreate {
	ac.mutation.SetCredentialId(s)
	return ac
}

// SetPrivateKey sets the "privateKey" field.
func (ac *AssetCreate) SetPrivateKey(s string) *AssetCreate {
	ac.mutation.SetPrivateKey(s)
	return ac
}

// SetPassphrase sets the "passphrase" field.
func (ac *AssetCreate) SetPassphrase(s string) *AssetCreate {
	ac.mutation.SetPassphrase(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *AssetCreate) SetDescription(s string) *AssetCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetActive sets the "active" field.
func (ac *AssetCreate) SetActive(b bool) *AssetCreate {
	ac.mutation.SetActive(b)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AssetCreate) SetCreatedAt(t time.Time) *AssetCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AssetCreate) SetNillableCreatedAt(t *time.Time) *AssetCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AssetCreate) SetUpdatedAt(t time.Time) *AssetCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AssetCreate) SetNillableUpdatedAt(t *time.Time) *AssetCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetTags sets the "tags" field.
func (ac *AssetCreate) SetTags(s string) *AssetCreate {
	ac.mutation.SetTags(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AssetCreate) SetID(s string) *AssetCreate {
	ac.mutation.SetID(s)
	return ac
}

// SetSessionsID sets the "sessions" edge to the Session entity by ID.
func (ac *AssetCreate) SetSessionsID(id string) *AssetCreate {
	ac.mutation.SetSessionsID(id)
	return ac
}

// SetNillableSessionsID sets the "sessions" edge to the Session entity by ID if the given value is not nil.
func (ac *AssetCreate) SetNillableSessionsID(id *string) *AssetCreate {
	if id != nil {
		ac = ac.SetSessionsID(*id)
	}
	return ac
}

// SetSessions sets the "sessions" edge to the Session entity.
func (ac *AssetCreate) SetSessions(s *Session) *AssetCreate {
	return ac.SetSessionsID(s.ID)
}

// Mutation returns the AssetMutation object of the builder.
func (ac *AssetCreate) Mutation() *AssetMutation {
	return ac.mutation
}

// Save creates the Asset in the database.
func (ac *AssetCreate) Save(ctx context.Context) (*Asset, error) {
	var (
		err  error
		node *Asset
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AssetCreate) SaveX(ctx context.Context) *Asset {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ac *AssetCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := asset.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := asset.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AssetCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := ac.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New("ent: missing required field \"ip\"")}
	}
	if _, ok := ac.mutation.Protocol(); !ok {
		return &ValidationError{Name: "protocol", err: errors.New("ent: missing required field \"protocol\"")}
	}
	if _, ok := ac.mutation.Port(); !ok {
		return &ValidationError{Name: "port", err: errors.New("ent: missing required field \"port\"")}
	}
	if _, ok := ac.mutation.AccountType(); !ok {
		return &ValidationError{Name: "accountType", err: errors.New("ent: missing required field \"accountType\"")}
	}
	if _, ok := ac.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New("ent: missing required field \"username\"")}
	}
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if _, ok := ac.mutation.CredentialId(); !ok {
		return &ValidationError{Name: "credentialId", err: errors.New("ent: missing required field \"credentialId\"")}
	}
	if _, ok := ac.mutation.PrivateKey(); !ok {
		return &ValidationError{Name: "privateKey", err: errors.New("ent: missing required field \"privateKey\"")}
	}
	if _, ok := ac.mutation.Passphrase(); !ok {
		return &ValidationError{Name: "passphrase", err: errors.New("ent: missing required field \"passphrase\"")}
	}
	if _, ok := ac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if _, ok := ac.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New("ent: missing required field \"active\"")}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := ac.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New("ent: missing required field \"tags\"")}
	}
	return nil
}

func (ac *AssetCreate) sqlSave(ctx context.Context) (*Asset, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (ac *AssetCreate) createSpec() (*Asset, *sqlgraph.CreateSpec) {
	var (
		_node = &Asset{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: asset.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: asset.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.IP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldIP,
		})
		_node.IP = value
	}
	if value, ok := ac.mutation.Protocol(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldProtocol,
		})
		_node.Protocol = value
	}
	if value, ok := ac.mutation.Port(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: asset.FieldPort,
		})
		_node.Port = value
	}
	if value, ok := ac.mutation.AccountType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldAccountType,
		})
		_node.AccountType = value
	}
	if value, ok := ac.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := ac.mutation.CredentialId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldCredentialId,
		})
		_node.CredentialId = value
	}
	if value, ok := ac.mutation.PrivateKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldPrivateKey,
		})
		_node.PrivateKey = value
	}
	if value, ok := ac.mutation.Passphrase(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldPassphrase,
		})
		_node.Passphrase = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := ac.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: asset.FieldActive,
		})
		_node.Active = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asset.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asset.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Tags(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldTags,
		})
		_node.Tags = value
	}
	if nodes := ac.mutation.SessionsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AssetCreateBulk is the builder for creating many Asset entities in bulk.
type AssetCreateBulk struct {
	config
	builders []*AssetCreate
}

// Save creates the Asset entities in the database.
func (acb *AssetCreateBulk) Save(ctx context.Context) ([]*Asset, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Asset, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AssetMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AssetCreateBulk) SaveX(ctx context.Context) []*Asset {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}