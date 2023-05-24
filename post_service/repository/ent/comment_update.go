// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"cw_post_service/repository/ent/comment"
	"cw_post_service/repository/ent/post"
	"cw_post_service/repository/ent/predicate"
	"cw_post_service/repository/ent/reply"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CommentUpdate) SetUserID(u uint64) *CommentUpdate {
	cu.mutation.ResetUserID()
	cu.mutation.SetUserID(u)
	return cu
}

// AddUserID adds u to the "user_id" field.
func (cu *CommentUpdate) AddUserID(u int64) *CommentUpdate {
	cu.mutation.AddUserID(u)
	return cu
}

// SetPostID sets the "post_id" field.
func (cu *CommentUpdate) SetPostID(u uint64) *CommentUpdate {
	cu.mutation.SetPostID(u)
	return cu
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillablePostID(u *uint64) *CommentUpdate {
	if u != nil {
		cu.SetPostID(*u)
	}
	return cu
}

// ClearPostID clears the value of the "post_id" field.
func (cu *CommentUpdate) ClearPostID() *CommentUpdate {
	cu.mutation.ClearPostID()
	return cu
}

// SetContent sets the "content" field.
func (cu *CommentUpdate) SetContent(s string) *CommentUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetStatus sets the "status" field.
func (cu *CommentUpdate) SetStatus(i int8) *CommentUpdate {
	cu.mutation.ResetStatus()
	cu.mutation.SetStatus(i)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableStatus(i *int8) *CommentUpdate {
	if i != nil {
		cu.SetStatus(*i)
	}
	return cu
}

// AddStatus adds i to the "status" field.
func (cu *CommentUpdate) AddStatus(i int8) *CommentUpdate {
	cu.mutation.AddStatus(i)
	return cu
}

// SetFloor sets the "floor" field.
func (cu *CommentUpdate) SetFloor(u uint64) *CommentUpdate {
	cu.mutation.ResetFloor()
	cu.mutation.SetFloor(u)
	return cu
}

// AddFloor adds u to the "floor" field.
func (cu *CommentUpdate) AddFloor(u int64) *CommentUpdate {
	cu.mutation.AddFloor(u)
	return cu
}

// SetOwnerID sets the "owner" edge to the Post entity by ID.
func (cu *CommentUpdate) SetOwnerID(id uint64) *CommentUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetNillableOwnerID sets the "owner" edge to the Post entity by ID if the given value is not nil.
func (cu *CommentUpdate) SetNillableOwnerID(id *uint64) *CommentUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the "owner" edge to the Post entity.
func (cu *CommentUpdate) SetOwner(p *Post) *CommentUpdate {
	return cu.SetOwnerID(p.ID)
}

// AddReplyIDs adds the "replies" edge to the Reply entity by IDs.
func (cu *CommentUpdate) AddReplyIDs(ids ...uint64) *CommentUpdate {
	cu.mutation.AddReplyIDs(ids...)
	return cu
}

// AddReplies adds the "replies" edges to the Reply entity.
func (cu *CommentUpdate) AddReplies(r ...*Reply) *CommentUpdate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.AddReplyIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the Post entity.
func (cu *CommentUpdate) ClearOwner() *CommentUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// ClearReplies clears all "replies" edges to the Reply entity.
func (cu *CommentUpdate) ClearReplies() *CommentUpdate {
	cu.mutation.ClearReplies()
	return cu
}

// RemoveReplyIDs removes the "replies" edge to Reply entities by IDs.
func (cu *CommentUpdate) RemoveReplyIDs(ids ...uint64) *CommentUpdate {
	cu.mutation.RemoveReplyIDs(ids...)
	return cu
}

// RemoveReplies removes "replies" edges to Reply entities.
func (cu *CommentUpdate) RemoveReplies(r ...*Reply) *CommentUpdate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.RemoveReplyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CommentMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CommentUpdate) check() error {
	if v, ok := cu.mutation.UserID(); ok {
		if err := comment.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Comment.user_id": %w`, err)}
		}
	}
	if v, ok := cu.mutation.PostID(); ok {
		if err := comment.PostIDValidator(v); err != nil {
			return &ValidationError{Name: "post_id", err: fmt.Errorf(`ent: validator failed for field "Comment.post_id": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Content(); ok {
		if err := comment.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Comment.content": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Status(); ok {
		if err := comment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Comment.status": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Floor(); ok {
		if err := comment.FloorValidator(v); err != nil {
			return &ValidationError{Name: "floor", err: fmt.Errorf(`ent: validator failed for field "Comment.floor": %w`, err)}
		}
	}
	return nil
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UserID(); ok {
		_spec.SetField(comment.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.AddedUserID(); ok {
		_spec.AddField(comment.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(comment.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := cu.mutation.AddedStatus(); ok {
		_spec.AddField(comment.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := cu.mutation.Floor(); ok {
		_spec.SetField(comment.FieldFloor, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.AddedFloor(); ok {
		_spec.AddField(comment.FieldFloor, field.TypeUint64, value)
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.OwnerTable,
			Columns: []string{comment.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.OwnerTable,
			Columns: []string{comment.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.RepliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedRepliesIDs(); len(nodes) > 0 && !cu.mutation.RepliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RepliesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetUserID sets the "user_id" field.
func (cuo *CommentUpdateOne) SetUserID(u uint64) *CommentUpdateOne {
	cuo.mutation.ResetUserID()
	cuo.mutation.SetUserID(u)
	return cuo
}

// AddUserID adds u to the "user_id" field.
func (cuo *CommentUpdateOne) AddUserID(u int64) *CommentUpdateOne {
	cuo.mutation.AddUserID(u)
	return cuo
}

// SetPostID sets the "post_id" field.
func (cuo *CommentUpdateOne) SetPostID(u uint64) *CommentUpdateOne {
	cuo.mutation.SetPostID(u)
	return cuo
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillablePostID(u *uint64) *CommentUpdateOne {
	if u != nil {
		cuo.SetPostID(*u)
	}
	return cuo
}

// ClearPostID clears the value of the "post_id" field.
func (cuo *CommentUpdateOne) ClearPostID() *CommentUpdateOne {
	cuo.mutation.ClearPostID()
	return cuo
}

// SetContent sets the "content" field.
func (cuo *CommentUpdateOne) SetContent(s string) *CommentUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CommentUpdateOne) SetStatus(i int8) *CommentUpdateOne {
	cuo.mutation.ResetStatus()
	cuo.mutation.SetStatus(i)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableStatus(i *int8) *CommentUpdateOne {
	if i != nil {
		cuo.SetStatus(*i)
	}
	return cuo
}

// AddStatus adds i to the "status" field.
func (cuo *CommentUpdateOne) AddStatus(i int8) *CommentUpdateOne {
	cuo.mutation.AddStatus(i)
	return cuo
}

// SetFloor sets the "floor" field.
func (cuo *CommentUpdateOne) SetFloor(u uint64) *CommentUpdateOne {
	cuo.mutation.ResetFloor()
	cuo.mutation.SetFloor(u)
	return cuo
}

// AddFloor adds u to the "floor" field.
func (cuo *CommentUpdateOne) AddFloor(u int64) *CommentUpdateOne {
	cuo.mutation.AddFloor(u)
	return cuo
}

// SetOwnerID sets the "owner" edge to the Post entity by ID.
func (cuo *CommentUpdateOne) SetOwnerID(id uint64) *CommentUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetNillableOwnerID sets the "owner" edge to the Post entity by ID if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableOwnerID(id *uint64) *CommentUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the "owner" edge to the Post entity.
func (cuo *CommentUpdateOne) SetOwner(p *Post) *CommentUpdateOne {
	return cuo.SetOwnerID(p.ID)
}

// AddReplyIDs adds the "replies" edge to the Reply entity by IDs.
func (cuo *CommentUpdateOne) AddReplyIDs(ids ...uint64) *CommentUpdateOne {
	cuo.mutation.AddReplyIDs(ids...)
	return cuo
}

// AddReplies adds the "replies" edges to the Reply entity.
func (cuo *CommentUpdateOne) AddReplies(r ...*Reply) *CommentUpdateOne {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.AddReplyIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the Post entity.
func (cuo *CommentUpdateOne) ClearOwner() *CommentUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// ClearReplies clears all "replies" edges to the Reply entity.
func (cuo *CommentUpdateOne) ClearReplies() *CommentUpdateOne {
	cuo.mutation.ClearReplies()
	return cuo
}

// RemoveReplyIDs removes the "replies" edge to Reply entities by IDs.
func (cuo *CommentUpdateOne) RemoveReplyIDs(ids ...uint64) *CommentUpdateOne {
	cuo.mutation.RemoveReplyIDs(ids...)
	return cuo
}

// RemoveReplies removes "replies" edges to Reply entities.
func (cuo *CommentUpdateOne) RemoveReplies(r ...*Reply) *CommentUpdateOne {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.RemoveReplyIDs(ids...)
}

// Where appends a list predicates to the CommentUpdate builder.
func (cuo *CommentUpdateOne) Where(ps ...predicate.Comment) *CommentUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	return withHooks[*Comment, CommentMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CommentUpdateOne) check() error {
	if v, ok := cuo.mutation.UserID(); ok {
		if err := comment.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Comment.user_id": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.PostID(); ok {
		if err := comment.PostIDValidator(v); err != nil {
			return &ValidationError{Name: "post_id", err: fmt.Errorf(`ent: validator failed for field "Comment.post_id": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Content(); ok {
		if err := comment.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Comment.content": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Status(); ok {
		if err := comment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Comment.status": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Floor(); ok {
		if err := comment.FloorValidator(v); err != nil {
			return &ValidationError{Name: "floor", err: fmt.Errorf(`ent: validator failed for field "Comment.floor": %w`, err)}
		}
	}
	return nil
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UserID(); ok {
		_spec.SetField(comment.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.AddedUserID(); ok {
		_spec.AddField(comment.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(comment.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := cuo.mutation.AddedStatus(); ok {
		_spec.AddField(comment.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := cuo.mutation.Floor(); ok {
		_spec.SetField(comment.FieldFloor, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.AddedFloor(); ok {
		_spec.AddField(comment.FieldFloor, field.TypeUint64, value)
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.OwnerTable,
			Columns: []string{comment.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.OwnerTable,
			Columns: []string{comment.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.RepliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedRepliesIDs(); len(nodes) > 0 && !cuo.mutation.RepliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RepliesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reply.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
