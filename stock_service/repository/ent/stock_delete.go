// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"stock_service/repository/ent/predicate"
	"stock_service/repository/ent/stock"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StockDelete is the builder for deleting a Stock entity.
type StockDelete struct {
	config
	hooks    []Hook
	mutation *StockMutation
}

// Where appends a list predicates to the StockDelete builder.
func (sd *StockDelete) Where(ps ...predicate.Stock) *StockDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *StockDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *StockDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *StockDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(stock.Table, sqlgraph.NewFieldSpec(stock.FieldID, field.TypeUint64))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// StockDeleteOne is the builder for deleting a single Stock entity.
type StockDeleteOne struct {
	sd *StockDelete
}

// Where appends a list predicates to the StockDelete builder.
func (sdo *StockDeleteOne) Where(ps ...predicate.Stock) *StockDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *StockDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{stock.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *StockDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
