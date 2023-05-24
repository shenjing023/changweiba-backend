// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"cw_post_service/repository/ent/comment"
	"cw_post_service/repository/ent/post"
	"cw_post_service/repository/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *PostUpdate) SetUserID(u uint64) *PostUpdate {
	pu.mutation.ResetUserID()
	pu.mutation.SetUserID(u)
	return pu
}

// AddUserID adds u to the "user_id" field.
func (pu *PostUpdate) AddUserID(u int64) *PostUpdate {
	pu.mutation.AddUserID(u)
	return pu
}

// SetTitle sets the "title" field.
func (pu *PostUpdate) SetTitle(s string) *PostUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetContent sets the "content" field.
func (pu *PostUpdate) SetContent(s string) *PostUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetStatus sets the "status" field.
func (pu *PostUpdate) SetStatus(i int8) *PostUpdate {
	pu.mutation.ResetStatus()
	pu.mutation.SetStatus(i)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PostUpdate) SetNillableStatus(i *int8) *PostUpdate {
	if i != nil {
		pu.SetStatus(*i)
	}
	return pu
}

// AddStatus adds i to the "status" field.
func (pu *PostUpdate) AddStatus(i int8) *PostUpdate {
	pu.mutation.AddStatus(i)
	return pu
}

// SetReplyNum sets the "reply_num" field.
func (pu *PostUpdate) SetReplyNum(i int64) *PostUpdate {
	pu.mutation.ResetReplyNum()
	pu.mutation.SetReplyNum(i)
	return pu
}

// SetNillableReplyNum sets the "reply_num" field if the given value is not nil.
func (pu *PostUpdate) SetNillableReplyNum(i *int64) *PostUpdate {
	if i != nil {
		pu.SetReplyNum(*i)
	}
	return pu
}

// AddReplyNum adds i to the "reply_num" field.
func (pu *PostUpdate) AddReplyNum(i int64) *PostUpdate {
	pu.mutation.AddReplyNum(i)
	return pu
}

// SetUpdateAt sets the "update_at" field.
func (pu *PostUpdate) SetUpdateAt(i int64) *PostUpdate {
	pu.mutation.ResetUpdateAt()
	pu.mutation.SetUpdateAt(i)
	return pu
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (pu *PostUpdate) SetNillableUpdateAt(i *int64) *PostUpdate {
	if i != nil {
		pu.SetUpdateAt(*i)
	}
	return pu
}

// AddUpdateAt adds i to the "update_at" field.
func (pu *PostUpdate) AddUpdateAt(i int64) *PostUpdate {
	pu.mutation.AddUpdateAt(i)
	return pu
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (pu *PostUpdate) AddCommentIDs(ids ...uint64) *PostUpdate {
	pu.mutation.AddCommentIDs(ids...)
	return pu
}

// AddComments adds the "comments" edges to the Comment entity.
func (pu *PostUpdate) AddComments(c ...*Comment) *PostUpdate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddCommentIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// ClearComments clears all "comments" edges to the Comment entity.
func (pu *PostUpdate) ClearComments() *PostUpdate {
	pu.mutation.ClearComments()
	return pu
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (pu *PostUpdate) RemoveCommentIDs(ids ...uint64) *PostUpdate {
	pu.mutation.RemoveCommentIDs(ids...)
	return pu
}

// RemoveComments removes "comments" edges to Comment entities.
func (pu *PostUpdate) RemoveComments(c ...*Comment) *PostUpdate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCommentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PostMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PostUpdate) check() error {
	if v, ok := pu.mutation.UserID(); ok {
		if err := post.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Post.user_id": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Post.title": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Content(); ok {
		if err := post.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Post.content": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := post.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Post.status": %w`, err)}
		}
	}
	if v, ok := pu.mutation.ReplyNum(); ok {
		if err := post.ReplyNumValidator(v); err != nil {
			return &ValidationError{Name: "reply_num", err: fmt.Errorf(`ent: validator failed for field "Post.reply_num": %w`, err)}
		}
	}
	if v, ok := pu.mutation.UpdateAt(); ok {
		if err := post.UpdateAtValidator(v); err != nil {
			return &ValidationError{Name: "update_at", err: fmt.Errorf(`ent: validator failed for field "Post.update_at": %w`, err)}
		}
	}
	return nil
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeUint64))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UserID(); ok {
		_spec.SetField(post.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := pu.mutation.AddedUserID(); ok {
		_spec.AddField(post.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(post.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := pu.mutation.AddedStatus(); ok {
		_spec.AddField(post.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := pu.mutation.ReplyNum(); ok {
		_spec.SetField(post.FieldReplyNum, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedReplyNum(); ok {
		_spec.AddField(post.FieldReplyNum, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.UpdateAt(); ok {
		_spec.SetField(post.FieldUpdateAt, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedUpdateAt(); ok {
		_spec.AddField(post.FieldUpdateAt, field.TypeInt64, value)
	}
	if pu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !pu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

// SetUserID sets the "user_id" field.
func (puo *PostUpdateOne) SetUserID(u uint64) *PostUpdateOne {
	puo.mutation.ResetUserID()
	puo.mutation.SetUserID(u)
	return puo
}

// AddUserID adds u to the "user_id" field.
func (puo *PostUpdateOne) AddUserID(u int64) *PostUpdateOne {
	puo.mutation.AddUserID(u)
	return puo
}

// SetTitle sets the "title" field.
func (puo *PostUpdateOne) SetTitle(s string) *PostUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetContent sets the "content" field.
func (puo *PostUpdateOne) SetContent(s string) *PostUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetStatus sets the "status" field.
func (puo *PostUpdateOne) SetStatus(i int8) *PostUpdateOne {
	puo.mutation.ResetStatus()
	puo.mutation.SetStatus(i)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableStatus(i *int8) *PostUpdateOne {
	if i != nil {
		puo.SetStatus(*i)
	}
	return puo
}

// AddStatus adds i to the "status" field.
func (puo *PostUpdateOne) AddStatus(i int8) *PostUpdateOne {
	puo.mutation.AddStatus(i)
	return puo
}

// SetReplyNum sets the "reply_num" field.
func (puo *PostUpdateOne) SetReplyNum(i int64) *PostUpdateOne {
	puo.mutation.ResetReplyNum()
	puo.mutation.SetReplyNum(i)
	return puo
}

// SetNillableReplyNum sets the "reply_num" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableReplyNum(i *int64) *PostUpdateOne {
	if i != nil {
		puo.SetReplyNum(*i)
	}
	return puo
}

// AddReplyNum adds i to the "reply_num" field.
func (puo *PostUpdateOne) AddReplyNum(i int64) *PostUpdateOne {
	puo.mutation.AddReplyNum(i)
	return puo
}

// SetUpdateAt sets the "update_at" field.
func (puo *PostUpdateOne) SetUpdateAt(i int64) *PostUpdateOne {
	puo.mutation.ResetUpdateAt()
	puo.mutation.SetUpdateAt(i)
	return puo
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableUpdateAt(i *int64) *PostUpdateOne {
	if i != nil {
		puo.SetUpdateAt(*i)
	}
	return puo
}

// AddUpdateAt adds i to the "update_at" field.
func (puo *PostUpdateOne) AddUpdateAt(i int64) *PostUpdateOne {
	puo.mutation.AddUpdateAt(i)
	return puo
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (puo *PostUpdateOne) AddCommentIDs(ids ...uint64) *PostUpdateOne {
	puo.mutation.AddCommentIDs(ids...)
	return puo
}

// AddComments adds the "comments" edges to the Comment entity.
func (puo *PostUpdateOne) AddComments(c ...*Comment) *PostUpdateOne {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddCommentIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// ClearComments clears all "comments" edges to the Comment entity.
func (puo *PostUpdateOne) ClearComments() *PostUpdateOne {
	puo.mutation.ClearComments()
	return puo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (puo *PostUpdateOne) RemoveCommentIDs(ids ...uint64) *PostUpdateOne {
	puo.mutation.RemoveCommentIDs(ids...)
	return puo
}

// RemoveComments removes "comments" edges to Comment entities.
func (puo *PostUpdateOne) RemoveComments(c ...*Comment) *PostUpdateOne {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCommentIDs(ids...)
}

// Where appends a list predicates to the PostUpdate builder.
func (puo *PostUpdateOne) Where(ps ...predicate.Post) *PostUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	return withHooks[*Post, PostMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PostUpdateOne) check() error {
	if v, ok := puo.mutation.UserID(); ok {
		if err := post.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Post.user_id": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Post.title": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Content(); ok {
		if err := post.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Post.content": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := post.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Post.status": %w`, err)}
		}
	}
	if v, ok := puo.mutation.ReplyNum(); ok {
		if err := post.ReplyNumValidator(v); err != nil {
			return &ValidationError{Name: "reply_num", err: fmt.Errorf(`ent: validator failed for field "Post.reply_num": %w`, err)}
		}
	}
	if v, ok := puo.mutation.UpdateAt(); ok {
		if err := post.UpdateAtValidator(v); err != nil {
			return &ValidationError{Name: "update_at", err: fmt.Errorf(`ent: validator failed for field "Post.update_at": %w`, err)}
		}
	}
	return nil
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeUint64))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UserID(); ok {
		_spec.SetField(post.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := puo.mutation.AddedUserID(); ok {
		_spec.AddField(post.FieldUserID, field.TypeUint64, value)
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(post.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := puo.mutation.AddedStatus(); ok {
		_spec.AddField(post.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := puo.mutation.ReplyNum(); ok {
		_spec.SetField(post.FieldReplyNum, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedReplyNum(); ok {
		_spec.AddField(post.FieldReplyNum, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.UpdateAt(); ok {
		_spec.SetField(post.FieldUpdateAt, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedUpdateAt(); ok {
		_spec.AddField(post.FieldUpdateAt, field.TypeInt64, value)
	}
	if puo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !puo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
