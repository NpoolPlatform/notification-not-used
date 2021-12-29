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
	"github.com/NpoolPlatform/notification/pkg/db/ent/mailbox"
	"github.com/NpoolPlatform/notification/pkg/db/ent/predicate"
)

// MailBoxQuery is the builder for querying MailBox entities.
type MailBoxQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MailBox
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MailBoxQuery builder.
func (mbq *MailBoxQuery) Where(ps ...predicate.MailBox) *MailBoxQuery {
	mbq.predicates = append(mbq.predicates, ps...)
	return mbq
}

// Limit adds a limit step to the query.
func (mbq *MailBoxQuery) Limit(limit int) *MailBoxQuery {
	mbq.limit = &limit
	return mbq
}

// Offset adds an offset step to the query.
func (mbq *MailBoxQuery) Offset(offset int) *MailBoxQuery {
	mbq.offset = &offset
	return mbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mbq *MailBoxQuery) Unique(unique bool) *MailBoxQuery {
	mbq.unique = &unique
	return mbq
}

// Order adds an order step to the query.
func (mbq *MailBoxQuery) Order(o ...OrderFunc) *MailBoxQuery {
	mbq.order = append(mbq.order, o...)
	return mbq
}

// First returns the first MailBox entity from the query.
// Returns a *NotFoundError when no MailBox was found.
func (mbq *MailBoxQuery) First(ctx context.Context) (*MailBox, error) {
	nodes, err := mbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mailbox.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mbq *MailBoxQuery) FirstX(ctx context.Context) *MailBox {
	node, err := mbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MailBox ID from the query.
// Returns a *NotFoundError when no MailBox ID was found.
func (mbq *MailBoxQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mailbox.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mbq *MailBoxQuery) FirstIDX(ctx context.Context) int {
	id, err := mbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MailBox entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one MailBox entity is not found.
// Returns a *NotFoundError when no MailBox entities are found.
func (mbq *MailBoxQuery) Only(ctx context.Context) (*MailBox, error) {
	nodes, err := mbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mailbox.Label}
	default:
		return nil, &NotSingularError{mailbox.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mbq *MailBoxQuery) OnlyX(ctx context.Context) *MailBox {
	node, err := mbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MailBox ID in the query.
// Returns a *NotSingularError when exactly one MailBox ID is not found.
// Returns a *NotFoundError when no entities are found.
func (mbq *MailBoxQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mailbox.Label}
	default:
		err = &NotSingularError{mailbox.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mbq *MailBoxQuery) OnlyIDX(ctx context.Context) int {
	id, err := mbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MailBoxes.
func (mbq *MailBoxQuery) All(ctx context.Context) ([]*MailBox, error) {
	if err := mbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mbq *MailBoxQuery) AllX(ctx context.Context) []*MailBox {
	nodes, err := mbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MailBox IDs.
func (mbq *MailBoxQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mbq.Select(mailbox.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mbq *MailBoxQuery) IDsX(ctx context.Context) []int {
	ids, err := mbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mbq *MailBoxQuery) Count(ctx context.Context) (int, error) {
	if err := mbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mbq *MailBoxQuery) CountX(ctx context.Context) int {
	count, err := mbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mbq *MailBoxQuery) Exist(ctx context.Context) (bool, error) {
	if err := mbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mbq *MailBoxQuery) ExistX(ctx context.Context) bool {
	exist, err := mbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MailBoxQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mbq *MailBoxQuery) Clone() *MailBoxQuery {
	if mbq == nil {
		return nil
	}
	return &MailBoxQuery{
		config:     mbq.config,
		limit:      mbq.limit,
		offset:     mbq.offset,
		order:      append([]OrderFunc{}, mbq.order...),
		predicates: append([]predicate.MailBox{}, mbq.predicates...),
		// clone intermediate query.
		sql:  mbq.sql.Clone(),
		path: mbq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (mbq *MailBoxQuery) GroupBy(field string, fields ...string) *MailBoxGroupBy {
	group := &MailBoxGroupBy{config: mbq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mbq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (mbq *MailBoxQuery) Select(fields ...string) *MailBoxSelect {
	mbq.fields = append(mbq.fields, fields...)
	return &MailBoxSelect{MailBoxQuery: mbq}
}

func (mbq *MailBoxQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mbq.fields {
		if !mailbox.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mbq.path != nil {
		prev, err := mbq.path(ctx)
		if err != nil {
			return err
		}
		mbq.sql = prev
	}
	return nil
}

func (mbq *MailBoxQuery) sqlAll(ctx context.Context) ([]*MailBox, error) {
	var (
		nodes = []*MailBox{}
		_spec = mbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MailBox{config: mbq.config}
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
	if err := sqlgraph.QueryNodes(ctx, mbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mbq *MailBoxQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mbq.querySpec()
	return sqlgraph.CountNodes(ctx, mbq.driver, _spec)
}

func (mbq *MailBoxQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mbq *MailBoxQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mailbox.Table,
			Columns: mailbox.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mailbox.FieldID,
			},
		},
		From:   mbq.sql,
		Unique: true,
	}
	if unique := mbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mailbox.FieldID)
		for i := range fields {
			if fields[i] != mailbox.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mbq *MailBoxQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mbq.driver.Dialect())
	t1 := builder.Table(mailbox.Table)
	columns := mbq.fields
	if len(columns) == 0 {
		columns = mailbox.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mbq.sql != nil {
		selector = mbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range mbq.predicates {
		p(selector)
	}
	for _, p := range mbq.order {
		p(selector)
	}
	if offset := mbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MailBoxGroupBy is the group-by builder for MailBox entities.
type MailBoxGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mbgb *MailBoxGroupBy) Aggregate(fns ...AggregateFunc) *MailBoxGroupBy {
	mbgb.fns = append(mbgb.fns, fns...)
	return mbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mbgb *MailBoxGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mbgb.path(ctx)
	if err != nil {
		return err
	}
	mbgb.sql = query
	return mbgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mbgb *MailBoxGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mbgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mbgb *MailBoxGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mbgb.fields) > 1 {
		return nil, errors.New("ent: MailBoxGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mbgb *MailBoxGroupBy) StringsX(ctx context.Context) []string {
	v, err := mbgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mbgb *MailBoxGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mbgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailbox.Label}
	default:
		err = fmt.Errorf("ent: MailBoxGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mbgb *MailBoxGroupBy) StringX(ctx context.Context) string {
	v, err := mbgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mbgb *MailBoxGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mbgb.fields) > 1 {
		return nil, errors.New("ent: MailBoxGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mbgb *MailBoxGroupBy) IntsX(ctx context.Context) []int {
	v, err := mbgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mbgb *MailBoxGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mbgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailbox.Label}
	default:
		err = fmt.Errorf("ent: MailBoxGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mbgb *MailBoxGroupBy) IntX(ctx context.Context) int {
	v, err := mbgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mbgb *MailBoxGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mbgb.fields) > 1 {
		return nil, errors.New("ent: MailBoxGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mbgb *MailBoxGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mbgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mbgb *MailBoxGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mbgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailbox.Label}
	default:
		err = fmt.Errorf("ent: MailBoxGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mbgb *MailBoxGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mbgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mbgb *MailBoxGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mbgb.fields) > 1 {
		return nil, errors.New("ent: MailBoxGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mbgb *MailBoxGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mbgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mbgb *MailBoxGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mbgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailbox.Label}
	default:
		err = fmt.Errorf("ent: MailBoxGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mbgb *MailBoxGroupBy) BoolX(ctx context.Context) bool {
	v, err := mbgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mbgb *MailBoxGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mbgb.fields {
		if !mailbox.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mbgb *MailBoxGroupBy) sqlQuery() *sql.Selector {
	selector := mbgb.sql.Select()
	aggregation := make([]string, 0, len(mbgb.fns))
	for _, fn := range mbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mbgb.fields)+len(mbgb.fns))
		for _, f := range mbgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mbgb.fields...)...)
}

// MailBoxSelect is the builder for selecting fields of MailBox entities.
type MailBoxSelect struct {
	*MailBoxQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mbs *MailBoxSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mbs.prepareQuery(ctx); err != nil {
		return err
	}
	mbs.sql = mbs.MailBoxQuery.sqlQuery(ctx)
	return mbs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mbs *MailBoxSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mbs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mbs *MailBoxSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mbs.fields) > 1 {
		return nil, errors.New("ent: MailBoxSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mbs *MailBoxSelect) StringsX(ctx context.Context) []string {
	v, err := mbs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mbs *MailBoxSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mbs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailbox.Label}
	default:
		err = fmt.Errorf("ent: MailBoxSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mbs *MailBoxSelect) StringX(ctx context.Context) string {
	v, err := mbs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mbs *MailBoxSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mbs.fields) > 1 {
		return nil, errors.New("ent: MailBoxSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mbs *MailBoxSelect) IntsX(ctx context.Context) []int {
	v, err := mbs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mbs *MailBoxSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mbs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailbox.Label}
	default:
		err = fmt.Errorf("ent: MailBoxSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mbs *MailBoxSelect) IntX(ctx context.Context) int {
	v, err := mbs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mbs *MailBoxSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mbs.fields) > 1 {
		return nil, errors.New("ent: MailBoxSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mbs *MailBoxSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mbs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mbs *MailBoxSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mbs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailbox.Label}
	default:
		err = fmt.Errorf("ent: MailBoxSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mbs *MailBoxSelect) Float64X(ctx context.Context) float64 {
	v, err := mbs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mbs *MailBoxSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mbs.fields) > 1 {
		return nil, errors.New("ent: MailBoxSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mbs *MailBoxSelect) BoolsX(ctx context.Context) []bool {
	v, err := mbs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mbs *MailBoxSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mbs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mailbox.Label}
	default:
		err = fmt.Errorf("ent: MailBoxSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mbs *MailBoxSelect) BoolX(ctx context.Context) bool {
	v, err := mbs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mbs *MailBoxSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mbs.sql.Query()
	if err := mbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
