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

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where adds a new predicate for the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.predicates = append(su.mutation.predicates, ps...)
	return su
}

// SetProtocol sets the "protocol" field.
func (su *SessionUpdate) SetProtocol(s string) *SessionUpdate {
	su.mutation.SetProtocol(s)
	return su
}

// SetIP sets the "ip" field.
func (su *SessionUpdate) SetIP(s string) *SessionUpdate {
	su.mutation.SetIP(s)
	return su
}

// SetPort sets the "port" field.
func (su *SessionUpdate) SetPort(i int) *SessionUpdate {
	su.mutation.ResetPort()
	su.mutation.SetPort(i)
	return su
}

// AddPort adds i to the "port" field.
func (su *SessionUpdate) AddPort(i int) *SessionUpdate {
	su.mutation.AddPort(i)
	return su
}

// SetConnectionID sets the "connection_id" field.
func (su *SessionUpdate) SetConnectionID(s string) *SessionUpdate {
	su.mutation.SetConnectionID(s)
	return su
}

// SetAssetID sets the "asset_id" field.
func (su *SessionUpdate) SetAssetID(s string) *SessionUpdate {
	su.mutation.SetAssetID(s)
	return su
}

// SetUsername sets the "username" field.
func (su *SessionUpdate) SetUsername(s string) *SessionUpdate {
	su.mutation.SetUsername(s)
	return su
}

// SetPassword sets the "password" field.
func (su *SessionUpdate) SetPassword(s string) *SessionUpdate {
	su.mutation.SetPassword(s)
	return su
}

// SetCreator sets the "creator" field.
func (su *SessionUpdate) SetCreator(s string) *SessionUpdate {
	su.mutation.SetCreator(s)
	return su
}

// SetClientIP sets the "client_ip" field.
func (su *SessionUpdate) SetClientIP(s string) *SessionUpdate {
	su.mutation.SetClientIP(s)
	return su
}

// SetWidth sets the "width" field.
func (su *SessionUpdate) SetWidth(i int) *SessionUpdate {
	su.mutation.ResetWidth()
	su.mutation.SetWidth(i)
	return su
}

// AddWidth adds i to the "width" field.
func (su *SessionUpdate) AddWidth(i int) *SessionUpdate {
	su.mutation.AddWidth(i)
	return su
}

// SetHeight sets the "height" field.
func (su *SessionUpdate) SetHeight(i int) *SessionUpdate {
	su.mutation.ResetHeight()
	su.mutation.SetHeight(i)
	return su
}

// AddHeight adds i to the "height" field.
func (su *SessionUpdate) AddHeight(i int) *SessionUpdate {
	su.mutation.AddHeight(i)
	return su
}

// SetStatus sets the "status" field.
func (su *SessionUpdate) SetStatus(s string) *SessionUpdate {
	su.mutation.SetStatus(s)
	return su
}

// SetRecording sets the "recording" field.
func (su *SessionUpdate) SetRecording(s string) *SessionUpdate {
	su.mutation.SetRecording(s)
	return su
}

// SetPrivateKey sets the "private_key" field.
func (su *SessionUpdate) SetPrivateKey(s string) *SessionUpdate {
	su.mutation.SetPrivateKey(s)
	return su
}

// SetPassphrase sets the "passphrase" field.
func (su *SessionUpdate) SetPassphrase(s string) *SessionUpdate {
	su.mutation.SetPassphrase(s)
	return su
}

// SetCode sets the "code" field.
func (su *SessionUpdate) SetCode(i int) *SessionUpdate {
	su.mutation.ResetCode()
	su.mutation.SetCode(i)
	return su
}

// AddCode adds i to the "code" field.
func (su *SessionUpdate) AddCode(i int) *SessionUpdate {
	su.mutation.AddCode(i)
	return su
}

// SetMessage sets the "message" field.
func (su *SessionUpdate) SetMessage(s string) *SessionUpdate {
	su.mutation.SetMessage(s)
	return su
}

// SetConnected sets the "connected" field.
func (su *SessionUpdate) SetConnected(t time.Time) *SessionUpdate {
	su.mutation.SetConnected(t)
	return su
}

// SetNillableConnected sets the "connected" field if the given value is not nil.
func (su *SessionUpdate) SetNillableConnected(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetConnected(*t)
	}
	return su
}

// SetDisconnected sets the "disconnected" field.
func (su *SessionUpdate) SetDisconnected(t time.Time) *SessionUpdate {
	su.mutation.SetDisconnected(t)
	return su
}

// AddAssetIDs adds the "assets" edge to the Asset entity by IDs.
func (su *SessionUpdate) AddAssetIDs(ids ...string) *SessionUpdate {
	su.mutation.AddAssetIDs(ids...)
	return su
}

// AddAssets adds the "assets" edges to the Asset entity.
func (su *SessionUpdate) AddAssets(a ...*Asset) *SessionUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.AddAssetIDs(ids...)
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// ClearAssets clears all "assets" edges to the Asset entity.
func (su *SessionUpdate) ClearAssets() *SessionUpdate {
	su.mutation.ClearAssets()
	return su
}

// RemoveAssetIDs removes the "assets" edge to Asset entities by IDs.
func (su *SessionUpdate) RemoveAssetIDs(ids ...string) *SessionUpdate {
	su.mutation.RemoveAssetIDs(ids...)
	return su
}

// RemoveAssets removes "assets" edges to Asset entities.
func (su *SessionUpdate) RemoveAssets(a ...*Asset) *SessionUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.RemoveAssetIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SessionUpdate) defaults() {
	if _, ok := su.mutation.Disconnected(); !ok {
		v := session.UpdateDefaultDisconnected()
		su.mutation.SetDisconnected(v)
	}
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: session.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Protocol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldProtocol,
		})
	}
	if value, ok := su.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldIP,
		})
	}
	if value, ok := su.mutation.Port(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldPort,
		})
	}
	if value, ok := su.mutation.AddedPort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldPort,
		})
	}
	if value, ok := su.mutation.ConnectionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldConnectionID,
		})
	}
	if value, ok := su.mutation.AssetID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldAssetID,
		})
	}
	if value, ok := su.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldUsername,
		})
	}
	if value, ok := su.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldPassword,
		})
	}
	if value, ok := su.mutation.Creator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldCreator,
		})
	}
	if value, ok := su.mutation.ClientIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldClientIP,
		})
	}
	if value, ok := su.mutation.Width(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldWidth,
		})
	}
	if value, ok := su.mutation.AddedWidth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldWidth,
		})
	}
	if value, ok := su.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldHeight,
		})
	}
	if value, ok := su.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldHeight,
		})
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldStatus,
		})
	}
	if value, ok := su.mutation.Recording(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldRecording,
		})
	}
	if value, ok := su.mutation.PrivateKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldPrivateKey,
		})
	}
	if value, ok := su.mutation.Passphrase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldPassphrase,
		})
	}
	if value, ok := su.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldCode,
		})
	}
	if value, ok := su.mutation.AddedCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldCode,
		})
	}
	if value, ok := su.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldMessage,
		})
	}
	if value, ok := su.mutation.Connected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldConnected,
		})
	}
	if value, ok := su.mutation.Disconnected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldDisconnected,
		})
	}
	if su.mutation.AssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   session.AssetsTable,
			Columns: []string{session.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedAssetsIDs(); len(nodes) > 0 && !su.mutation.AssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   session.AssetsTable,
			Columns: []string{session.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   session.AssetsTable,
			Columns: []string{session.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// SetProtocol sets the "protocol" field.
func (suo *SessionUpdateOne) SetProtocol(s string) *SessionUpdateOne {
	suo.mutation.SetProtocol(s)
	return suo
}

// SetIP sets the "ip" field.
func (suo *SessionUpdateOne) SetIP(s string) *SessionUpdateOne {
	suo.mutation.SetIP(s)
	return suo
}

// SetPort sets the "port" field.
func (suo *SessionUpdateOne) SetPort(i int) *SessionUpdateOne {
	suo.mutation.ResetPort()
	suo.mutation.SetPort(i)
	return suo
}

// AddPort adds i to the "port" field.
func (suo *SessionUpdateOne) AddPort(i int) *SessionUpdateOne {
	suo.mutation.AddPort(i)
	return suo
}

// SetConnectionID sets the "connection_id" field.
func (suo *SessionUpdateOne) SetConnectionID(s string) *SessionUpdateOne {
	suo.mutation.SetConnectionID(s)
	return suo
}

// SetAssetID sets the "asset_id" field.
func (suo *SessionUpdateOne) SetAssetID(s string) *SessionUpdateOne {
	suo.mutation.SetAssetID(s)
	return suo
}

// SetUsername sets the "username" field.
func (suo *SessionUpdateOne) SetUsername(s string) *SessionUpdateOne {
	suo.mutation.SetUsername(s)
	return suo
}

// SetPassword sets the "password" field.
func (suo *SessionUpdateOne) SetPassword(s string) *SessionUpdateOne {
	suo.mutation.SetPassword(s)
	return suo
}

// SetCreator sets the "creator" field.
func (suo *SessionUpdateOne) SetCreator(s string) *SessionUpdateOne {
	suo.mutation.SetCreator(s)
	return suo
}

// SetClientIP sets the "client_ip" field.
func (suo *SessionUpdateOne) SetClientIP(s string) *SessionUpdateOne {
	suo.mutation.SetClientIP(s)
	return suo
}

// SetWidth sets the "width" field.
func (suo *SessionUpdateOne) SetWidth(i int) *SessionUpdateOne {
	suo.mutation.ResetWidth()
	suo.mutation.SetWidth(i)
	return suo
}

// AddWidth adds i to the "width" field.
func (suo *SessionUpdateOne) AddWidth(i int) *SessionUpdateOne {
	suo.mutation.AddWidth(i)
	return suo
}

// SetHeight sets the "height" field.
func (suo *SessionUpdateOne) SetHeight(i int) *SessionUpdateOne {
	suo.mutation.ResetHeight()
	suo.mutation.SetHeight(i)
	return suo
}

// AddHeight adds i to the "height" field.
func (suo *SessionUpdateOne) AddHeight(i int) *SessionUpdateOne {
	suo.mutation.AddHeight(i)
	return suo
}

// SetStatus sets the "status" field.
func (suo *SessionUpdateOne) SetStatus(s string) *SessionUpdateOne {
	suo.mutation.SetStatus(s)
	return suo
}

// SetRecording sets the "recording" field.
func (suo *SessionUpdateOne) SetRecording(s string) *SessionUpdateOne {
	suo.mutation.SetRecording(s)
	return suo
}

// SetPrivateKey sets the "private_key" field.
func (suo *SessionUpdateOne) SetPrivateKey(s string) *SessionUpdateOne {
	suo.mutation.SetPrivateKey(s)
	return suo
}

// SetPassphrase sets the "passphrase" field.
func (suo *SessionUpdateOne) SetPassphrase(s string) *SessionUpdateOne {
	suo.mutation.SetPassphrase(s)
	return suo
}

// SetCode sets the "code" field.
func (suo *SessionUpdateOne) SetCode(i int) *SessionUpdateOne {
	suo.mutation.ResetCode()
	suo.mutation.SetCode(i)
	return suo
}

// AddCode adds i to the "code" field.
func (suo *SessionUpdateOne) AddCode(i int) *SessionUpdateOne {
	suo.mutation.AddCode(i)
	return suo
}

// SetMessage sets the "message" field.
func (suo *SessionUpdateOne) SetMessage(s string) *SessionUpdateOne {
	suo.mutation.SetMessage(s)
	return suo
}

// SetConnected sets the "connected" field.
func (suo *SessionUpdateOne) SetConnected(t time.Time) *SessionUpdateOne {
	suo.mutation.SetConnected(t)
	return suo
}

// SetNillableConnected sets the "connected" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableConnected(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetConnected(*t)
	}
	return suo
}

// SetDisconnected sets the "disconnected" field.
func (suo *SessionUpdateOne) SetDisconnected(t time.Time) *SessionUpdateOne {
	suo.mutation.SetDisconnected(t)
	return suo
}

// AddAssetIDs adds the "assets" edge to the Asset entity by IDs.
func (suo *SessionUpdateOne) AddAssetIDs(ids ...string) *SessionUpdateOne {
	suo.mutation.AddAssetIDs(ids...)
	return suo
}

// AddAssets adds the "assets" edges to the Asset entity.
func (suo *SessionUpdateOne) AddAssets(a ...*Asset) *SessionUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.AddAssetIDs(ids...)
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// ClearAssets clears all "assets" edges to the Asset entity.
func (suo *SessionUpdateOne) ClearAssets() *SessionUpdateOne {
	suo.mutation.ClearAssets()
	return suo
}

// RemoveAssetIDs removes the "assets" edge to Asset entities by IDs.
func (suo *SessionUpdateOne) RemoveAssetIDs(ids ...string) *SessionUpdateOne {
	suo.mutation.RemoveAssetIDs(ids...)
	return suo
}

// RemoveAssets removes "assets" edges to Asset entities.
func (suo *SessionUpdateOne) RemoveAssets(a ...*Asset) *SessionUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.RemoveAssetIDs(ids...)
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	var (
		err  error
		node *Session
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SessionUpdateOne) defaults() {
	if _, ok := suo.mutation.Disconnected(); !ok {
		v := session.UpdateDefaultDisconnected()
		suo.mutation.SetDisconnected(v)
	}
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: session.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Session.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Protocol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldProtocol,
		})
	}
	if value, ok := suo.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldIP,
		})
	}
	if value, ok := suo.mutation.Port(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldPort,
		})
	}
	if value, ok := suo.mutation.AddedPort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldPort,
		})
	}
	if value, ok := suo.mutation.ConnectionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldConnectionID,
		})
	}
	if value, ok := suo.mutation.AssetID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldAssetID,
		})
	}
	if value, ok := suo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldUsername,
		})
	}
	if value, ok := suo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldPassword,
		})
	}
	if value, ok := suo.mutation.Creator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldCreator,
		})
	}
	if value, ok := suo.mutation.ClientIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldClientIP,
		})
	}
	if value, ok := suo.mutation.Width(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldWidth,
		})
	}
	if value, ok := suo.mutation.AddedWidth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldWidth,
		})
	}
	if value, ok := suo.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldHeight,
		})
	}
	if value, ok := suo.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldHeight,
		})
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldStatus,
		})
	}
	if value, ok := suo.mutation.Recording(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldRecording,
		})
	}
	if value, ok := suo.mutation.PrivateKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldPrivateKey,
		})
	}
	if value, ok := suo.mutation.Passphrase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldPassphrase,
		})
	}
	if value, ok := suo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldCode,
		})
	}
	if value, ok := suo.mutation.AddedCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldCode,
		})
	}
	if value, ok := suo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldMessage,
		})
	}
	if value, ok := suo.mutation.Connected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldConnected,
		})
	}
	if value, ok := suo.mutation.Disconnected(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldDisconnected,
		})
	}
	if suo.mutation.AssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   session.AssetsTable,
			Columns: []string{session.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedAssetsIDs(); len(nodes) > 0 && !suo.mutation.AssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   session.AssetsTable,
			Columns: []string{session.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   session.AssetsTable,
			Columns: []string{session.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
