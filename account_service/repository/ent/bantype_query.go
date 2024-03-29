// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"cw_account_service/repository/ent/bantype"
	"cw_account_service/repository/ent/predicate"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BanTypeQuery is the builder for querying BanType entities.
type BanTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BanType
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BanTypeQuery builder.
func (btq *BanTypeQuery) Where(ps ...predicate.BanType) *BanTypeQuery {
	btq.predicates = append(btq.predicates, ps...)
	return btq
}

// Limit adds a limit step to the query.
func (btq *BanTypeQuery) Limit(limit int) *BanTypeQuery {
	btq.limit = &limit
	return btq
}

// Offset adds an offset step to the query.
func (btq *BanTypeQuery) Offset(offset int) *BanTypeQuery {
	btq.offset = &offset
	return btq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (btq *BanTypeQuery) Unique(unique bool) *BanTypeQuery {
	btq.unique = &unique
	return btq
}

// Order adds an order step to the query.
func (btq *BanTypeQuery) Order(o ...OrderFunc) *BanTypeQuery {
	btq.order = append(btq.order, o...)
	return btq
}

// First returns the first BanType entity from the query.
// Returns a *NotFoundError when no BanType was found.
func (btq *BanTypeQuery) First(ctx context.Context) (*BanType, error) {
	nodes, err := btq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bantype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (btq *BanTypeQuery) FirstX(ctx context.Context) *BanType {
	node, err := btq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BanType ID from the query.
// Returns a *NotFoundError when no BanType ID was found.
func (btq *BanTypeQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = btq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bantype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (btq *BanTypeQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := btq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BanType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one BanType entity is not found.
// Returns a *NotFoundError when no BanType entities are found.
func (btq *BanTypeQuery) Only(ctx context.Context) (*BanType, error) {
	nodes, err := btq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bantype.Label}
	default:
		return nil, &NotSingularError{bantype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (btq *BanTypeQuery) OnlyX(ctx context.Context) *BanType {
	node, err := btq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BanType ID in the query.
// Returns a *NotSingularError when exactly one BanType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (btq *BanTypeQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = btq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bantype.Label}
	default:
		err = &NotSingularError{bantype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (btq *BanTypeQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := btq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BanTypes.
func (btq *BanTypeQuery) All(ctx context.Context) ([]*BanType, error) {
	if err := btq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return btq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (btq *BanTypeQuery) AllX(ctx context.Context) []*BanType {
	nodes, err := btq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BanType IDs.
func (btq *BanTypeQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := btq.Select(bantype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (btq *BanTypeQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := btq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (btq *BanTypeQuery) Count(ctx context.Context) (int, error) {
	if err := btq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return btq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (btq *BanTypeQuery) CountX(ctx context.Context) int {
	count, err := btq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (btq *BanTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := btq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return btq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (btq *BanTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := btq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BanTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (btq *BanTypeQuery) Clone() *BanTypeQuery {
	if btq == nil {
		return nil
	}
	return &BanTypeQuery{
		config:     btq.config,
		limit:      btq.limit,
		offset:     btq.offset,
		order:      append([]OrderFunc{}, btq.order...),
		predicates: append([]predicate.BanType{}, btq.predicates...),
		// clone intermediate query.
		sql:  btq.sql.Clone(),
		path: btq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Content string `json:"content,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BanType.Query().
//		GroupBy(bantype.FieldContent).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (btq *BanTypeQuery) GroupBy(field string, fields ...string) *BanTypeGroupBy {
	group := &BanTypeGroupBy{config: btq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := btq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return btq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Content string `json:"content,omitempty"`
//	}
//
//	client.BanType.Query().
//		Select(bantype.FieldContent).
//		Scan(ctx, &v)
//
func (btq *BanTypeQuery) Select(fields ...string) *BanTypeSelect {
	btq.fields = append(btq.fields, fields...)
	return &BanTypeSelect{BanTypeQuery: btq}
}

func (btq *BanTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range btq.fields {
		if !bantype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if btq.path != nil {
		prev, err := btq.path(ctx)
		if err != nil {
			return err
		}
		btq.sql = prev
	}
	return nil
}

func (btq *BanTypeQuery) sqlAll(ctx context.Context) ([]*BanType, error) {
	var (
		nodes = []*BanType{}
		_spec = btq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &BanType{config: btq.config}
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
	if err := sqlgraph.QueryNodes(ctx, btq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (btq *BanTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := btq.querySpec()
	return sqlgraph.CountNodes(ctx, btq.driver, _spec)
}

func (btq *BanTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := btq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (btq *BanTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bantype.Table,
			Columns: bantype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: bantype.FieldID,
			},
		},
		From:   btq.sql,
		Unique: true,
	}
	if unique := btq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := btq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bantype.FieldID)
		for i := range fields {
			if fields[i] != bantype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := btq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := btq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := btq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := btq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (btq *BanTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(btq.driver.Dialect())
	t1 := builder.Table(bantype.Table)
	columns := btq.fields
	if len(columns) == 0 {
		columns = bantype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if btq.sql != nil {
		selector = btq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range btq.predicates {
		p(selector)
	}
	for _, p := range btq.order {
		p(selector)
	}
	if offset := btq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := btq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BanTypeGroupBy is the group-by builder for BanType entities.
type BanTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (btgb *BanTypeGroupBy) Aggregate(fns ...AggregateFunc) *BanTypeGroupBy {
	btgb.fns = append(btgb.fns, fns...)
	return btgb
}

// Scan applies the group-by query and scans the result into the given value.
func (btgb *BanTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := btgb.path(ctx)
	if err != nil {
		return err
	}
	btgb.sql = query
	return btgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (btgb *BanTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := btgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (btgb *BanTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(btgb.fields) > 1 {
		return nil, errors.New("ent: BanTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := btgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (btgb *BanTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := btgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (btgb *BanTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = btgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bantype.Label}
	default:
		err = fmt.Errorf("ent: BanTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (btgb *BanTypeGroupBy) StringX(ctx context.Context) string {
	v, err := btgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (btgb *BanTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(btgb.fields) > 1 {
		return nil, errors.New("ent: BanTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := btgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (btgb *BanTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := btgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (btgb *BanTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = btgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bantype.Label}
	default:
		err = fmt.Errorf("ent: BanTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (btgb *BanTypeGroupBy) IntX(ctx context.Context) int {
	v, err := btgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (btgb *BanTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(btgb.fields) > 1 {
		return nil, errors.New("ent: BanTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := btgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (btgb *BanTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := btgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (btgb *BanTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = btgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bantype.Label}
	default:
		err = fmt.Errorf("ent: BanTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (btgb *BanTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := btgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (btgb *BanTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(btgb.fields) > 1 {
		return nil, errors.New("ent: BanTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := btgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (btgb *BanTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := btgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (btgb *BanTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = btgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bantype.Label}
	default:
		err = fmt.Errorf("ent: BanTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (btgb *BanTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := btgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (btgb *BanTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range btgb.fields {
		if !bantype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := btgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := btgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (btgb *BanTypeGroupBy) sqlQuery() *sql.Selector {
	selector := btgb.sql.Select()
	aggregation := make([]string, 0, len(btgb.fns))
	for _, fn := range btgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(btgb.fields)+len(btgb.fns))
		for _, f := range btgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(btgb.fields...)...)
}

// BanTypeSelect is the builder for selecting fields of BanType entities.
type BanTypeSelect struct {
	*BanTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bts *BanTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bts.prepareQuery(ctx); err != nil {
		return err
	}
	bts.sql = bts.BanTypeQuery.sqlQuery(ctx)
	return bts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bts *BanTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (bts *BanTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bts.fields) > 1 {
		return nil, errors.New("ent: BanTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bts *BanTypeSelect) StringsX(ctx context.Context) []string {
	v, err := bts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (bts *BanTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bantype.Label}
	default:
		err = fmt.Errorf("ent: BanTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bts *BanTypeSelect) StringX(ctx context.Context) string {
	v, err := bts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (bts *BanTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bts.fields) > 1 {
		return nil, errors.New("ent: BanTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bts *BanTypeSelect) IntsX(ctx context.Context) []int {
	v, err := bts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (bts *BanTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bantype.Label}
	default:
		err = fmt.Errorf("ent: BanTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bts *BanTypeSelect) IntX(ctx context.Context) int {
	v, err := bts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (bts *BanTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bts.fields) > 1 {
		return nil, errors.New("ent: BanTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bts *BanTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (bts *BanTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bantype.Label}
	default:
		err = fmt.Errorf("ent: BanTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bts *BanTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := bts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (bts *BanTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bts.fields) > 1 {
		return nil, errors.New("ent: BanTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bts *BanTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := bts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (bts *BanTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bantype.Label}
	default:
		err = fmt.Errorf("ent: BanTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bts *BanTypeSelect) BoolX(ctx context.Context) bool {
	v, err := bts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bts *BanTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bts.sql.Query()
	if err := bts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
