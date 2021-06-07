// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/asset"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/predicate"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/session"
)

// AssetQuery is the builder for querying Asset entities.
type AssetQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Asset
	// eager-loading edges.
	withSessions *SessionQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AssetQuery builder.
func (aq *AssetQuery) Where(ps ...predicate.Asset) *AssetQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AssetQuery) Limit(limit int) *AssetQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AssetQuery) Offset(offset int) *AssetQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AssetQuery) Unique(unique bool) *AssetQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AssetQuery) Order(o ...OrderFunc) *AssetQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QuerySessions chains the current query on the "sessions" edge.
func (aq *AssetQuery) QuerySessions() *SessionQuery {
	query := &SessionQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asset.SessionsTable, asset.SessionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Asset entity from the query.
// Returns a *NotFoundError when no Asset was found.
func (aq *AssetQuery) First(ctx context.Context) (*Asset, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{asset.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AssetQuery) FirstX(ctx context.Context) *Asset {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Asset ID from the query.
// Returns a *NotFoundError when no Asset ID was found.
func (aq *AssetQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{asset.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AssetQuery) FirstIDX(ctx context.Context) string {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Asset entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Asset entity is not found.
// Returns a *NotFoundError when no Asset entities are found.
func (aq *AssetQuery) Only(ctx context.Context) (*Asset, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{asset.Label}
	default:
		return nil, &NotSingularError{asset.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AssetQuery) OnlyX(ctx context.Context) *Asset {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Asset ID in the query.
// Returns a *NotSingularError when exactly one Asset ID is not found.
// Returns a *NotFoundError when no entities are found.
func (aq *AssetQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{asset.Label}
	default:
		err = &NotSingularError{asset.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AssetQuery) OnlyIDX(ctx context.Context) string {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Assets.
func (aq *AssetQuery) All(ctx context.Context) ([]*Asset, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AssetQuery) AllX(ctx context.Context) []*Asset {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Asset IDs.
func (aq *AssetQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := aq.Select(asset.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AssetQuery) IDsX(ctx context.Context) []string {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AssetQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AssetQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AssetQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AssetQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AssetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AssetQuery) Clone() *AssetQuery {
	if aq == nil {
		return nil
	}
	return &AssetQuery{
		config:       aq.config,
		limit:        aq.limit,
		offset:       aq.offset,
		order:        append([]OrderFunc{}, aq.order...),
		predicates:   append([]predicate.Asset{}, aq.predicates...),
		withSessions: aq.withSessions.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithSessions tells the query-builder to eager-load the nodes that are connected to
// the "sessions" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithSessions(opts ...func(*SessionQuery)) *AssetQuery {
	query := &SessionQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withSessions = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Asset.Query().
//		GroupBy(asset.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aq *AssetQuery) GroupBy(field string, fields ...string) *AssetGroupBy {
	group := &AssetGroupBy{config: aq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Asset.Query().
//		Select(asset.FieldName).
//		Scan(ctx, &v)
//
func (aq *AssetQuery) Select(field string, fields ...string) *AssetSelect {
	aq.fields = append([]string{field}, fields...)
	return &AssetSelect{AssetQuery: aq}
}

func (aq *AssetQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !asset.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AssetQuery) sqlAll(ctx context.Context) ([]*Asset, error) {
	var (
		nodes       = []*Asset{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withSessions != nil,
		}
	)
	if aq.withSessions != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, asset.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Asset{config: aq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aq.withSessions; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*Asset)
		for i := range nodes {
			if nodes[i].session_assets == nil {
				continue
			}
			fk := *nodes[i].session_assets
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(session.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "session_assets" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Sessions = n
			}
		}
	}

	return nodes, nil
}

func (aq *AssetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AssetQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aq *AssetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asset.Table,
			Columns: asset.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: asset.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asset.FieldID)
		for i := range fields {
			if fields[i] != asset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AssetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(asset.Table)
	selector := builder.Select(t1.Columns(asset.Columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(asset.Columns...)...)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AssetGroupBy is the group-by builder for Asset entities.
type AssetGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AssetGroupBy) Aggregate(fns ...AggregateFunc) *AssetGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AssetGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (agb *AssetGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := agb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *AssetGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AssetGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (agb *AssetGroupBy) StringsX(ctx context.Context) []string {
	v, err := agb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *AssetGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = agb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asset.Label}
	default:
		err = fmt.Errorf("ent: AssetGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (agb *AssetGroupBy) StringX(ctx context.Context) string {
	v, err := agb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *AssetGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AssetGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (agb *AssetGroupBy) IntsX(ctx context.Context) []int {
	v, err := agb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *AssetGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = agb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asset.Label}
	default:
		err = fmt.Errorf("ent: AssetGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (agb *AssetGroupBy) IntX(ctx context.Context) int {
	v, err := agb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *AssetGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AssetGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (agb *AssetGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := agb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *AssetGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = agb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asset.Label}
	default:
		err = fmt.Errorf("ent: AssetGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (agb *AssetGroupBy) Float64X(ctx context.Context) float64 {
	v, err := agb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *AssetGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AssetGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (agb *AssetGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := agb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *AssetGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = agb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asset.Label}
	default:
		err = fmt.Errorf("ent: AssetGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (agb *AssetGroupBy) BoolX(ctx context.Context) bool {
	v, err := agb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (agb *AssetGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agb.fields {
		if !asset.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AssetGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql
	columns := make([]string, 0, len(agb.fields)+len(agb.fns))
	columns = append(columns, agb.fields...)
	for _, fn := range agb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(agb.fields...)
}

// AssetSelect is the builder for selecting fields of Asset entities.
type AssetSelect struct {
	*AssetQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *AssetSelect) Scan(ctx context.Context, v interface{}) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AssetQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (as *AssetSelect) ScanX(ctx context.Context, v interface{}) {
	if err := as.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (as *AssetSelect) Strings(ctx context.Context) ([]string, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AssetSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (as *AssetSelect) StringsX(ctx context.Context) []string {
	v, err := as.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (as *AssetSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = as.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asset.Label}
	default:
		err = fmt.Errorf("ent: AssetSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (as *AssetSelect) StringX(ctx context.Context) string {
	v, err := as.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (as *AssetSelect) Ints(ctx context.Context) ([]int, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AssetSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (as *AssetSelect) IntsX(ctx context.Context) []int {
	v, err := as.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (as *AssetSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = as.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asset.Label}
	default:
		err = fmt.Errorf("ent: AssetSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (as *AssetSelect) IntX(ctx context.Context) int {
	v, err := as.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (as *AssetSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AssetSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (as *AssetSelect) Float64sX(ctx context.Context) []float64 {
	v, err := as.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (as *AssetSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = as.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asset.Label}
	default:
		err = fmt.Errorf("ent: AssetSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (as *AssetSelect) Float64X(ctx context.Context) float64 {
	v, err := as.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (as *AssetSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AssetSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (as *AssetSelect) BoolsX(ctx context.Context) []bool {
	v, err := as.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (as *AssetSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = as.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asset.Label}
	default:
		err = fmt.Errorf("ent: AssetSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (as *AssetSelect) BoolX(ctx context.Context) bool {
	v, err := as.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (as *AssetSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sqlQuery().Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (as *AssetSelect) sqlQuery() sql.Querier {
	selector := as.sql
	selector.Select(selector.Columns(as.fields...)...)
	return selector
}
