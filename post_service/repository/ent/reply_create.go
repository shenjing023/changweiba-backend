// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"cw_post_service/repository/ent/comment"
	"cw_post_service/repository/ent/reply"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReplyCreate is the builder for creating a Reply entity.
type ReplyCreate struct {
	config
	mutation *ReplyMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (rc *ReplyCreate) SetUserID(u uint64) *ReplyCreate {
	rc.mutation.SetUserID(u)
	return rc
}

// SetCommentID sets the "comment_id" field.
func (rc *ReplyCreate) SetCommentID(u uint64) *ReplyCreate {
	rc.mutation.SetCommentID(u)
	return rc
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (rc *ReplyCreate) SetNillableCommentID(u *uint64) *ReplyCreate {
	if u != nil {
		rc.SetCommentID(*u)
	}
	return rc
}

// SetParentID sets the "parent_id" field.
func (rc *ReplyCreate) SetParentID(u uint64) *ReplyCreate {
	rc.mutation.SetParentID(u)
	return rc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (rc *ReplyCreate) SetNillableParentID(u *uint64) *ReplyCreate {
	if u != nil {
		rc.SetParentID(*u)
	}
	return rc
}

// SetContent sets the "content" field.
func (rc *ReplyCreate) SetContent(s string) *ReplyCreate {
	rc.mutation.SetContent(s)
	return rc
}

// SetStatus sets the "status" field.
func (rc *ReplyCreate) SetStatus(i int8) *ReplyCreate {
	rc.mutation.SetStatus(i)
	return rc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rc *ReplyCreate) SetNillableStatus(i *int8) *ReplyCreate {
	if i != nil {
		rc.SetStatus(*i)
	}
	return rc
}

// SetFloor sets the "floor" field.
func (rc *ReplyCreate) SetFloor(u uint64) *ReplyCreate {
	rc.mutation.SetFloor(u)
	return rc
}

// SetCreateAt sets the "create_at" field.
func (rc *ReplyCreate) SetCreateAt(i int64) *ReplyCreate {
	rc.mutation.SetCreateAt(i)
	return rc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (rc *ReplyCreate) SetNillableCreateAt(i *int64) *ReplyCreate {
	if i != nil {
		rc.SetCreateAt(*i)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *ReplyCreate) SetID(u uint64) *ReplyCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetOwnerID sets the "owner" edge to the Comment entity by ID.
func (rc *ReplyCreate) SetOwnerID(id uint64) *ReplyCreate {
	rc.mutation.SetOwnerID(id)
	return rc
}

// SetNillableOwnerID sets the "owner" edge to the Comment entity by ID if the given value is not nil.
func (rc *ReplyCreate) SetNillableOwnerID(id *uint64) *ReplyCreate {
	if id != nil {
		rc = rc.SetOwnerID(*id)
	}
	return rc
}

// SetOwner sets the "owner" edge to the Comment entity.
func (rc *ReplyCreate) SetOwner(c *Comment) *ReplyCreate {
	return rc.SetOwnerID(c.ID)
}

// SetParent sets the "parent" edge to the Reply entity.
func (rc *ReplyCreate) SetParent(r *Reply) *ReplyCreate {
	return rc.SetParentID(r.ID)
}

// AddChildIDs adds the "children" edge to the Reply entity by IDs.
func (rc *ReplyCreate) AddChildIDs(ids ...uint64) *ReplyCreate {
	rc.mutation.AddChildIDs(ids...)
	return rc
}

// AddChildren adds the "children" edges to the Reply entity.
func (rc *ReplyCreate) AddChildren(r ...*Reply) *ReplyCreate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddChildIDs(ids...)
}

// Mutation returns the ReplyMutation object of the builder.
func (rc *ReplyCreate) Mutation() *ReplyMutation {
	return rc.mutation
}

// Save creates the Reply in the database.
func (rc *ReplyCreate) Save(ctx context.Context) (*Reply, error) {
	rc.defaults()
	return withHooks[*Reply, ReplyMutation](ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReplyCreate) SaveX(ctx context.Context) *Reply {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReplyCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReplyCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReplyCreate) defaults() {
	if _, ok := rc.mutation.Status(); !ok {
		v := reply.DefaultStatus
		rc.mutation.SetStatus(v)
	}
	if _, ok := rc.mutation.CreateAt(); !ok {
		v := reply.DefaultCreateAt
		rc.mutation.SetCreateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReplyCreate) check() error {
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Reply.user_id"`)}
	}
	if v, ok := rc.mutation.UserID(); ok {
		if err := reply.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Reply.user_id": %w`, err)}
		}
	}
	if v, ok := rc.mutation.CommentID(); ok {
		if err := reply.CommentIDValidator(v); err != nil {
			return &ValidationError{Name: "comment_id", err: fmt.Errorf(`ent: validator failed for field "Reply.comment_id": %w`, err)}
		}
	}
	if v, ok := rc.mutation.ParentID(); ok {
		if err := reply.ParentIDValidator(v); err != nil {
			return &ValidationError{Name: "parent_id", err: fmt.Errorf(`ent: validator failed for field "Reply.parent_id": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Reply.content"`)}
	}
	if v, ok := rc.mutation.Content(); ok {
		if err := reply.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Reply.content": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Reply.status"`)}
	}
	if v, ok := rc.mutation.Status(); ok {
		if err := reply.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Reply.status": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Floor(); !ok {
		return &ValidationError{Name: "floor", err: errors.New(`ent: missing required field "Reply.floor"`)}
	}
	if v, ok := rc.mutation.Floor(); ok {
		if err := reply.FloorValidator(v); err != nil {
			return &ValidationError{Name: "floor", err: fmt.Errorf(`ent: validator failed for field "Reply.floor": %w`, err)}
		}
	}
	if _, ok := rc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "Reply.create_at"`)}
	}
	if v, ok := rc.mutation.CreateAt(); ok {
		if err := reply.CreateAtValidator(v); err != nil {
			return &ValidationError{Name: "create_at", err: fmt.Errorf(`ent: validator failed for field "Reply.create_at": %w`, err)}
		}
	}
	if v, ok := rc.mutation.ID(); ok {
		if err := reply.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Reply.id": %w`, err)}
		}
	}
	return nil
}

func (rc *ReplyCreate) sqlSave(ctx context.Context) (*Reply, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReplyCreate) createSpec() (*Reply, *sqlgraph.CreateSpec) {
	var (
		_node = &Reply{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(reply.Table, sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUint64))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.UserID(); ok {
		_spec.SetField(reply.FieldUserID, field.TypeUint64, value)
		_node.UserID = value
	}
	if value, ok := rc.mutation.Content(); ok {
		_spec.SetField(reply.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.SetField(reply.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := rc.mutation.Floor(); ok {
		_spec.SetField(reply.FieldFloor, field.TypeUint64, value)
		_node.Floor = value
	}
	if value, ok := rc.mutation.CreateAt(); ok {
		_spec.SetField(reply.FieldCreateAt, field.TypeInt64, value)
		_node.CreateAt = value
	}
	if nodes := rc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reply.OwnerTable,
			Columns: []string{reply.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CommentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reply.ParentTable,
			Columns: []string{reply.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   reply.ChildrenTable,
			Columns: []string{reply.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReplyCreateBulk is the builder for creating many Reply entities in bulk.
type ReplyCreateBulk struct {
	config
	builders []*ReplyCreate
}

// Save creates the Reply entities in the database.
func (rcb *ReplyCreateBulk) Save(ctx context.Context) ([]*Reply, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Reply, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReplyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReplyCreateBulk) SaveX(ctx context.Context) []*Reply {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReplyCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReplyCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
