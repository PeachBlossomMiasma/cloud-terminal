// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/command"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/predicate"
)

// CommandQuery is the builder for querying Command entities.
type CommandQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Command
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommandQuery builder.
func (cq *CommandQuery) Where(ps ...predicate.Command) *CommandQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *CommandQuery) Limit(limit int) *CommandQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *CommandQuery) Offset(offset int) *CommandQuery {
	cq.offset = &offset
	return cq
}

// Order adds an order step to the query.
func (cq *CommandQuery) Order(o ...OrderFunc) *CommandQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// First returns the first Command entity from the query.
// Returns a *NotFoundError when no Command was found.
func (cq *CommandQuery) First(ctx context.Context) (*Command, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{command.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CommandQuery) FirstX(ctx context.Context) *Command {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Command ID from the query.
// Returns a *NotFoundError when no Command ID was found.
func (cq *CommandQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{command.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CommandQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Command entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Command entity is not found.
// Returns a *NotFoundError when no Command entities are found.
func (cq *CommandQuery) Only(ctx context.Context) (*Command, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{command.Label}
	default:
		return nil, &NotSingularError{command.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CommandQuery) OnlyX(ctx context.Context) *Command {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Command ID in the query.
// Returns a *NotSingularError when exactly one Command ID is not found.
// Returns a *NotFoundError when no entities are found.
func (cq *CommandQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{command.Label}
	default:
		err = &NotSingularError{command.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CommandQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Commands.
func (cq *CommandQuery) All(ctx context.Context) ([]*Command, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *CommandQuery) AllX(ctx context.Context) []*Command {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Command IDs.
func (cq *CommandQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cq.Select(command.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CommandQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CommandQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CommandQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CommandQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CommandQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommandQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CommandQuery) Clone() *CommandQuery {
	if cq == nil {
		return nil
	}
	return &CommandQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]OrderFunc{}, cq.order...),
		predicates: append([]predicate.Command{}, cq.predicates...),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (cq *CommandQuery) GroupBy(field string, fields ...string) *CommandGroupBy {
	group := &CommandGroupBy{config: cq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (cq *CommandQuery) Select(field string, fields ...string) *CommandSelect {
	cq.fields = append([]string{field}, fields...)
	return &CommandSelect{CommandQuery: cq}
}

func (cq *CommandQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !command.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CommandQuery) sqlAll(ctx context.Context) ([]*Command, error) {
	var (
		nodes = []*Command{}
		_spec = cq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Command{config: cq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cq *CommandQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CommandQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (cq *CommandQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   command.Table,
			Columns: command.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: command.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, command.FieldID)
		for i := range fields {
			if fields[i] != command.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, command.ValidColumn)
			}
		}
	}
	return _spec
}

func (cq *CommandQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(command.Table)
	selector := builder.Select(t1.Columns(command.Columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(command.Columns...)...)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector, command.ValidColumn)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommandGroupBy is the group-by builder for Command entities.
type CommandGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CommandGroupBy) Aggregate(fns ...AggregateFunc) *CommandGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cgb *CommandGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cgb *CommandGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cgb *CommandGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CommandGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cgb *CommandGroupBy) StringsX(ctx context.Context) []string {
	v, err := cgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cgb *CommandGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{command.Label}
	default:
		err = fmt.Errorf("ent: CommandGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cgb *CommandGroupBy) StringX(ctx context.Context) string {
	v, err := cgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cgb *CommandGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CommandGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cgb *CommandGroupBy) IntsX(ctx context.Context) []int {
	v, err := cgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cgb *CommandGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{command.Label}
	default:
		err = fmt.Errorf("ent: CommandGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cgb *CommandGroupBy) IntX(ctx context.Context) int {
	v, err := cgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cgb *CommandGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CommandGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cgb *CommandGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cgb *CommandGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{command.Label}
	default:
		err = fmt.Errorf("ent: CommandGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cgb *CommandGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cgb *CommandGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CommandGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cgb *CommandGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cgb *CommandGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{command.Label}
	default:
		err = fmt.Errorf("ent: CommandGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cgb *CommandGroupBy) BoolX(ctx context.Context) bool {
	v, err := cgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cgb *CommandGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cgb.fields {
		if !command.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *CommandGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql
	columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
	columns = append(columns, cgb.fields...)
	for _, fn := range cgb.fns {
		columns = append(columns, fn(selector, command.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(cgb.fields...)
}

// CommandSelect is the builder for selecting fields of Command entities.
type CommandSelect struct {
	*CommandQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CommandSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.sql = cs.CommandQuery.sqlQuery()
	return cs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cs *CommandSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cs *CommandSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CommandSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cs *CommandSelect) StringsX(ctx context.Context) []string {
	v, err := cs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cs *CommandSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{command.Label}
	default:
		err = fmt.Errorf("ent: CommandSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cs *CommandSelect) StringX(ctx context.Context) string {
	v, err := cs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cs *CommandSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CommandSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cs *CommandSelect) IntsX(ctx context.Context) []int {
	v, err := cs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cs *CommandSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{command.Label}
	default:
		err = fmt.Errorf("ent: CommandSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cs *CommandSelect) IntX(ctx context.Context) int {
	v, err := cs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cs *CommandSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CommandSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cs *CommandSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cs *CommandSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{command.Label}
	default:
		err = fmt.Errorf("ent: CommandSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cs *CommandSelect) Float64X(ctx context.Context) float64 {
	v, err := cs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cs *CommandSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CommandSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cs *CommandSelect) BoolsX(ctx context.Context) []bool {
	v, err := cs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cs *CommandSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{command.Label}
	default:
		err = fmt.Errorf("ent: CommandSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cs *CommandSelect) BoolX(ctx context.Context) bool {
	v, err := cs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cs *CommandSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cs.sqlQuery().Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cs *CommandSelect) sqlQuery() sql.Querier {
	selector := cs.sql
	selector.Select(selector.Columns(cs.fields...)...)
	return selector
}
