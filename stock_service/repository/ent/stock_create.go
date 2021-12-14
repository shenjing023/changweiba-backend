// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"stock_service/repository/ent/stock"
	"stock_service/repository/ent/tradedate"
	"stock_service/repository/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StockCreate is the builder for creating a Stock entity.
type StockCreate struct {
	config
	mutation *StockMutation
	hooks    []Hook
}

// SetSymbol sets the "symbol" field.
func (sc *StockCreate) SetSymbol(s string) *StockCreate {
	sc.mutation.SetSymbol(s)
	return sc
}

// SetName sets the "name" field.
func (sc *StockCreate) SetName(s string) *StockCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetID sets the "id" field.
func (sc *StockCreate) SetID(u uint64) *StockCreate {
	sc.mutation.SetID(u)
	return sc
}

// AddTradeIDs adds the "trades" edge to the TradeDate entity by IDs.
func (sc *StockCreate) AddTradeIDs(ids ...uint64) *StockCreate {
	sc.mutation.AddTradeIDs(ids...)
	return sc
}

// AddTrades adds the "trades" edges to the TradeDate entity.
func (sc *StockCreate) AddTrades(t ...*TradeDate) *StockCreate {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTradeIDs(ids...)
}

// AddSubscriberIDs adds the "subscribers" edge to the User entity by IDs.
func (sc *StockCreate) AddSubscriberIDs(ids ...uint64) *StockCreate {
	sc.mutation.AddSubscriberIDs(ids...)
	return sc
}

// AddSubscribers adds the "subscribers" edges to the User entity.
func (sc *StockCreate) AddSubscribers(u ...*User) *StockCreate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return sc.AddSubscriberIDs(ids...)
}

// Mutation returns the StockMutation object of the builder.
func (sc *StockCreate) Mutation() *StockMutation {
	return sc.mutation
}

// Save creates the Stock in the database.
func (sc *StockCreate) Save(ctx context.Context) (*Stock, error) {
	var (
		err  error
		node *Stock
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StockCreate) SaveX(ctx context.Context) *Stock {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StockCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StockCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StockCreate) check() error {
	if _, ok := sc.mutation.Symbol(); !ok {
		return &ValidationError{Name: "symbol", err: errors.New(`ent: missing required field "symbol"`)}
	}
	if v, ok := sc.mutation.Symbol(); ok {
		if err := stock.SymbolValidator(v); err != nil {
			return &ValidationError{Name: "symbol", err: fmt.Errorf(`ent: validator failed for field "symbol": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := stock.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if v, ok := sc.mutation.ID(); ok {
		if err := stock.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "id": %w`, err)}
		}
	}
	return nil
}

func (sc *StockCreate) sqlSave(ctx context.Context) (*Stock, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (sc *StockCreate) createSpec() (*Stock, *sqlgraph.CreateSpec) {
	var (
		_node = &Stock{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: stock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: stock.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Symbol(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldSymbol,
		})
		_node.Symbol = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldName,
		})
		_node.Name = value
	}
	if nodes := sc.mutation.TradesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stock.TradesTable,
			Columns: []string{stock.TradesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: tradedate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.SubscribersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   stock.SubscribersTable,
			Columns: stock.SubscribersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
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

// StockCreateBulk is the builder for creating many Stock entities in bulk.
type StockCreateBulk struct {
	config
	builders []*StockCreate
}

// Save creates the Stock entities in the database.
func (scb *StockCreateBulk) Save(ctx context.Context) ([]*Stock, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Stock, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StockMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StockCreateBulk) SaveX(ctx context.Context) []*Stock {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StockCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StockCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
