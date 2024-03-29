// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"stock_service/repository/ent/hot"
	"stock_service/repository/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HotDelete is the builder for deleting a Hot entity.
type HotDelete struct {
	config
	hooks    []Hook
	mutation *HotMutation
}

// Where appends a list predicates to the HotDelete builder.
func (hd *HotDelete) Where(ps ...predicate.Hot) *HotDelete {
	hd.mutation.Where(ps...)
	return hd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hd *HotDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hd.sqlExec, hd.mutation, hd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hd *HotDelete) ExecX(ctx context.Context) int {
	n, err := hd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hd *HotDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(hot.Table, sqlgraph.NewFieldSpec(hot.FieldID, field.TypeUint64))
	if ps := hd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hd.mutation.done = true
	return affected, err
}

// HotDeleteOne is the builder for deleting a single Hot entity.
type HotDeleteOne struct {
	hd *HotDelete
}

// Where appends a list predicates to the HotDelete builder.
func (hdo *HotDeleteOne) Where(ps ...predicate.Hot) *HotDeleteOne {
	hdo.hd.mutation.Where(ps...)
	return hdo
}

// Exec executes the deletion query.
func (hdo *HotDeleteOne) Exec(ctx context.Context) error {
	n, err := hdo.hd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hot.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hdo *HotDeleteOne) ExecX(ctx context.Context) {
	if err := hdo.Exec(ctx); err != nil {
		panic(err)
	}
}
