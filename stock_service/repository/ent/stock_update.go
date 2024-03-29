// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"stock_service/repository/ent/predicate"
	"stock_service/repository/ent/stock"
	"stock_service/repository/ent/tradedate"
	"stock_service/repository/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StockUpdate is the builder for updating Stock entities.
type StockUpdate struct {
	config
	hooks    []Hook
	mutation *StockMutation
}

// Where appends a list predicates to the StockUpdate builder.
func (su *StockUpdate) Where(ps ...predicate.Stock) *StockUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetSymbol sets the "symbol" field.
func (su *StockUpdate) SetSymbol(s string) *StockUpdate {
	su.mutation.SetSymbol(s)
	return su
}

// SetName sets the "name" field.
func (su *StockUpdate) SetName(s string) *StockUpdate {
	su.mutation.SetName(s)
	return su
}

// SetBull sets the "bull" field.
func (su *StockUpdate) SetBull(i int) *StockUpdate {
	su.mutation.ResetBull()
	su.mutation.SetBull(i)
	return su
}

// SetNillableBull sets the "bull" field if the given value is not nil.
func (su *StockUpdate) SetNillableBull(i *int) *StockUpdate {
	if i != nil {
		su.SetBull(*i)
	}
	return su
}

// AddBull adds i to the "bull" field.
func (su *StockUpdate) AddBull(i int) *StockUpdate {
	su.mutation.AddBull(i)
	return su
}

// SetLastSubscribeAt sets the "last_subscribe_at" field.
func (su *StockUpdate) SetLastSubscribeAt(t time.Time) *StockUpdate {
	su.mutation.SetLastSubscribeAt(t)
	return su
}

// SetNillableLastSubscribeAt sets the "last_subscribe_at" field if the given value is not nil.
func (su *StockUpdate) SetNillableLastSubscribeAt(t *time.Time) *StockUpdate {
	if t != nil {
		su.SetLastSubscribeAt(*t)
	}
	return su
}

// SetShort sets the "short" field.
func (su *StockUpdate) SetShort(s string) *StockUpdate {
	su.mutation.SetShort(s)
	return su
}

// SetNillableShort sets the "short" field if the given value is not nil.
func (su *StockUpdate) SetNillableShort(s *string) *StockUpdate {
	if s != nil {
		su.SetShort(*s)
	}
	return su
}

// AddTradeIDs adds the "trades" edge to the TradeDate entity by IDs.
func (su *StockUpdate) AddTradeIDs(ids ...uint64) *StockUpdate {
	su.mutation.AddTradeIDs(ids...)
	return su
}

// AddTrades adds the "trades" edges to the TradeDate entity.
func (su *StockUpdate) AddTrades(t ...*TradeDate) *StockUpdate {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.AddTradeIDs(ids...)
}

// AddSubscriberIDs adds the "subscribers" edge to the User entity by IDs.
func (su *StockUpdate) AddSubscriberIDs(ids ...uint64) *StockUpdate {
	su.mutation.AddSubscriberIDs(ids...)
	return su
}

// AddSubscribers adds the "subscribers" edges to the User entity.
func (su *StockUpdate) AddSubscribers(u ...*User) *StockUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.AddSubscriberIDs(ids...)
}

// Mutation returns the StockMutation object of the builder.
func (su *StockUpdate) Mutation() *StockMutation {
	return su.mutation
}

// ClearTrades clears all "trades" edges to the TradeDate entity.
func (su *StockUpdate) ClearTrades() *StockUpdate {
	su.mutation.ClearTrades()
	return su
}

// RemoveTradeIDs removes the "trades" edge to TradeDate entities by IDs.
func (su *StockUpdate) RemoveTradeIDs(ids ...uint64) *StockUpdate {
	su.mutation.RemoveTradeIDs(ids...)
	return su
}

// RemoveTrades removes "trades" edges to TradeDate entities.
func (su *StockUpdate) RemoveTrades(t ...*TradeDate) *StockUpdate {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.RemoveTradeIDs(ids...)
}

// ClearSubscribers clears all "subscribers" edges to the User entity.
func (su *StockUpdate) ClearSubscribers() *StockUpdate {
	su.mutation.ClearSubscribers()
	return su
}

// RemoveSubscriberIDs removes the "subscribers" edge to User entities by IDs.
func (su *StockUpdate) RemoveSubscriberIDs(ids ...uint64) *StockUpdate {
	su.mutation.RemoveSubscriberIDs(ids...)
	return su
}

// RemoveSubscribers removes "subscribers" edges to User entities.
func (su *StockUpdate) RemoveSubscribers(u ...*User) *StockUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.RemoveSubscriberIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StockUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StockUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StockUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StockUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StockUpdate) check() error {
	if v, ok := su.mutation.Symbol(); ok {
		if err := stock.SymbolValidator(v); err != nil {
			return &ValidationError{Name: "symbol", err: fmt.Errorf(`ent: validator failed for field "Stock.symbol": %w`, err)}
		}
	}
	if v, ok := su.mutation.Name(); ok {
		if err := stock.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Stock.name": %w`, err)}
		}
	}
	if v, ok := su.mutation.Short(); ok {
		if err := stock.ShortValidator(v); err != nil {
			return &ValidationError{Name: "short", err: fmt.Errorf(`ent: validator failed for field "Stock.short": %w`, err)}
		}
	}
	return nil
}

func (su *StockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(stock.Table, stock.Columns, sqlgraph.NewFieldSpec(stock.FieldID, field.TypeUint64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Symbol(); ok {
		_spec.SetField(stock.FieldSymbol, field.TypeString, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(stock.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Bull(); ok {
		_spec.SetField(stock.FieldBull, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedBull(); ok {
		_spec.AddField(stock.FieldBull, field.TypeInt, value)
	}
	if value, ok := su.mutation.LastSubscribeAt(); ok {
		_spec.SetField(stock.FieldLastSubscribeAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.Short(); ok {
		_spec.SetField(stock.FieldShort, field.TypeString, value)
	}
	if su.mutation.TradesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stock.TradesTable,
			Columns: []string{stock.TradesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tradedate.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedTradesIDs(); len(nodes) > 0 && !su.mutation.TradesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stock.TradesTable,
			Columns: []string{stock.TradesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tradedate.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TradesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stock.TradesTable,
			Columns: []string{stock.TradesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tradedate.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.SubscribersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   stock.SubscribersTable,
			Columns: stock.SubscribersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedSubscribersIDs(); len(nodes) > 0 && !su.mutation.SubscribersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   stock.SubscribersTable,
			Columns: stock.SubscribersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SubscribersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   stock.SubscribersTable,
			Columns: stock.SubscribersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{stock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StockUpdateOne is the builder for updating a single Stock entity.
type StockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StockMutation
}

// SetSymbol sets the "symbol" field.
func (suo *StockUpdateOne) SetSymbol(s string) *StockUpdateOne {
	suo.mutation.SetSymbol(s)
	return suo
}

// SetName sets the "name" field.
func (suo *StockUpdateOne) SetName(s string) *StockUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetBull sets the "bull" field.
func (suo *StockUpdateOne) SetBull(i int) *StockUpdateOne {
	suo.mutation.ResetBull()
	suo.mutation.SetBull(i)
	return suo
}

// SetNillableBull sets the "bull" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableBull(i *int) *StockUpdateOne {
	if i != nil {
		suo.SetBull(*i)
	}
	return suo
}

// AddBull adds i to the "bull" field.
func (suo *StockUpdateOne) AddBull(i int) *StockUpdateOne {
	suo.mutation.AddBull(i)
	return suo
}

// SetLastSubscribeAt sets the "last_subscribe_at" field.
func (suo *StockUpdateOne) SetLastSubscribeAt(t time.Time) *StockUpdateOne {
	suo.mutation.SetLastSubscribeAt(t)
	return suo
}

// SetNillableLastSubscribeAt sets the "last_subscribe_at" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableLastSubscribeAt(t *time.Time) *StockUpdateOne {
	if t != nil {
		suo.SetLastSubscribeAt(*t)
	}
	return suo
}

// SetShort sets the "short" field.
func (suo *StockUpdateOne) SetShort(s string) *StockUpdateOne {
	suo.mutation.SetShort(s)
	return suo
}

// SetNillableShort sets the "short" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableShort(s *string) *StockUpdateOne {
	if s != nil {
		suo.SetShort(*s)
	}
	return suo
}

// AddTradeIDs adds the "trades" edge to the TradeDate entity by IDs.
func (suo *StockUpdateOne) AddTradeIDs(ids ...uint64) *StockUpdateOne {
	suo.mutation.AddTradeIDs(ids...)
	return suo
}

// AddTrades adds the "trades" edges to the TradeDate entity.
func (suo *StockUpdateOne) AddTrades(t ...*TradeDate) *StockUpdateOne {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.AddTradeIDs(ids...)
}

// AddSubscriberIDs adds the "subscribers" edge to the User entity by IDs.
func (suo *StockUpdateOne) AddSubscriberIDs(ids ...uint64) *StockUpdateOne {
	suo.mutation.AddSubscriberIDs(ids...)
	return suo
}

// AddSubscribers adds the "subscribers" edges to the User entity.
func (suo *StockUpdateOne) AddSubscribers(u ...*User) *StockUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.AddSubscriberIDs(ids...)
}

// Mutation returns the StockMutation object of the builder.
func (suo *StockUpdateOne) Mutation() *StockMutation {
	return suo.mutation
}

// ClearTrades clears all "trades" edges to the TradeDate entity.
func (suo *StockUpdateOne) ClearTrades() *StockUpdateOne {
	suo.mutation.ClearTrades()
	return suo
}

// RemoveTradeIDs removes the "trades" edge to TradeDate entities by IDs.
func (suo *StockUpdateOne) RemoveTradeIDs(ids ...uint64) *StockUpdateOne {
	suo.mutation.RemoveTradeIDs(ids...)
	return suo
}

// RemoveTrades removes "trades" edges to TradeDate entities.
func (suo *StockUpdateOne) RemoveTrades(t ...*TradeDate) *StockUpdateOne {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.RemoveTradeIDs(ids...)
}

// ClearSubscribers clears all "subscribers" edges to the User entity.
func (suo *StockUpdateOne) ClearSubscribers() *StockUpdateOne {
	suo.mutation.ClearSubscribers()
	return suo
}

// RemoveSubscriberIDs removes the "subscribers" edge to User entities by IDs.
func (suo *StockUpdateOne) RemoveSubscriberIDs(ids ...uint64) *StockUpdateOne {
	suo.mutation.RemoveSubscriberIDs(ids...)
	return suo
}

// RemoveSubscribers removes "subscribers" edges to User entities.
func (suo *StockUpdateOne) RemoveSubscribers(u ...*User) *StockUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.RemoveSubscriberIDs(ids...)
}

// Where appends a list predicates to the StockUpdate builder.
func (suo *StockUpdateOne) Where(ps ...predicate.Stock) *StockUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StockUpdateOne) Select(field string, fields ...string) *StockUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Stock entity.
func (suo *StockUpdateOne) Save(ctx context.Context) (*Stock, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StockUpdateOne) SaveX(ctx context.Context) *Stock {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StockUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StockUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StockUpdateOne) check() error {
	if v, ok := suo.mutation.Symbol(); ok {
		if err := stock.SymbolValidator(v); err != nil {
			return &ValidationError{Name: "symbol", err: fmt.Errorf(`ent: validator failed for field "Stock.symbol": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Name(); ok {
		if err := stock.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Stock.name": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Short(); ok {
		if err := stock.ShortValidator(v); err != nil {
			return &ValidationError{Name: "short", err: fmt.Errorf(`ent: validator failed for field "Stock.short": %w`, err)}
		}
	}
	return nil
}

func (suo *StockUpdateOne) sqlSave(ctx context.Context) (_node *Stock, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(stock.Table, stock.Columns, sqlgraph.NewFieldSpec(stock.FieldID, field.TypeUint64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Stock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, stock.FieldID)
		for _, f := range fields {
			if !stock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != stock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Symbol(); ok {
		_spec.SetField(stock.FieldSymbol, field.TypeString, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(stock.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Bull(); ok {
		_spec.SetField(stock.FieldBull, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedBull(); ok {
		_spec.AddField(stock.FieldBull, field.TypeInt, value)
	}
	if value, ok := suo.mutation.LastSubscribeAt(); ok {
		_spec.SetField(stock.FieldLastSubscribeAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Short(); ok {
		_spec.SetField(stock.FieldShort, field.TypeString, value)
	}
	if suo.mutation.TradesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stock.TradesTable,
			Columns: []string{stock.TradesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tradedate.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedTradesIDs(); len(nodes) > 0 && !suo.mutation.TradesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stock.TradesTable,
			Columns: []string{stock.TradesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tradedate.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TradesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stock.TradesTable,
			Columns: []string{stock.TradesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tradedate.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.SubscribersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   stock.SubscribersTable,
			Columns: stock.SubscribersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedSubscribersIDs(); len(nodes) > 0 && !suo.mutation.SubscribersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   stock.SubscribersTable,
			Columns: stock.SubscribersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SubscribersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   stock.SubscribersTable,
			Columns: stock.SubscribersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Stock{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{stock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
