// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"stock_service/repository/ent/predicate"
	"stock_service/repository/ent/tradedate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TradeDateDelete is the builder for deleting a TradeDate entity.
type TradeDateDelete struct {
	config
	hooks    []Hook
	mutation *TradeDateMutation
}

// Where appends a list predicates to the TradeDateDelete builder.
func (tdd *TradeDateDelete) Where(ps ...predicate.TradeDate) *TradeDateDelete {
	tdd.mutation.Where(ps...)
	return tdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tdd *TradeDateDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tdd.hooks) == 0 {
		affected, err = tdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TradeDateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tdd.mutation = mutation
			affected, err = tdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tdd.hooks) - 1; i >= 0; i-- {
			if tdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdd *TradeDateDelete) ExecX(ctx context.Context) int {
	n, err := tdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tdd *TradeDateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tradedate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tradedate.FieldID,
			},
		},
	}
	if ps := tdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tdd.driver, _spec)
}

// TradeDateDeleteOne is the builder for deleting a single TradeDate entity.
type TradeDateDeleteOne struct {
	tdd *TradeDateDelete
}

// Exec executes the deletion query.
func (tddo *TradeDateDeleteOne) Exec(ctx context.Context) error {
	n, err := tddo.tdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tradedate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tddo *TradeDateDeleteOne) ExecX(ctx context.Context) {
	tddo.tdd.ExecX(ctx)
}
