// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"cw_post_service/repository/ent/comment"
	"cw_post_service/repository/ent/post"
	"cw_post_service/repository/ent/predicate"
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
func (pu *PostUpdate) SetUserID(i int64) *PostUpdate {
	pu.mutation.ResetUserID()
	pu.mutation.SetUserID(i)
	return pu
}

// AddUserID adds i to the "user_id" field.
func (pu *PostUpdate) AddUserID(i int64) *PostUpdate {
	pu.mutation.AddUserID(i)
	return pu
}

// SetTopic sets the "topic" field.
func (pu *PostUpdate) SetTopic(s string) *PostUpdate {
	pu.mutation.SetTopic(s)
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

// SetCreateAt sets the "create_at" field.
func (pu *PostUpdate) SetCreateAt(i int64) *PostUpdate {
	pu.mutation.ResetCreateAt()
	pu.mutation.SetCreateAt(i)
	return pu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (pu *PostUpdate) SetNillableCreateAt(i *int64) *PostUpdate {
	if i != nil {
		pu.SetCreateAt(*i)
	}
	return pu
}

// AddCreateAt adds i to the "create_at" field.
func (pu *PostUpdate) AddCreateAt(i int64) *PostUpdate {
	pu.mutation.AddCreateAt(i)
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
func (pu *PostUpdate) AddCommentIDs(ids ...int64) *PostUpdate {
	pu.mutation.AddCommentIDs(ids...)
	return pu
}

// AddComments adds the "comments" edges to the Comment entity.
func (pu *PostUpdate) AddComments(c ...*Comment) *PostUpdate {
	ids := make([]int64, len(c))
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
func (pu *PostUpdate) RemoveCommentIDs(ids ...int64) *PostUpdate {
	pu.mutation.RemoveCommentIDs(ids...)
	return pu
}

// RemoveComments removes "comments" edges to Comment entities.
func (pu *PostUpdate) RemoveComments(c ...*Comment) *PostUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCommentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
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
			return &ValidationError{Name: "user_id", err: fmt.Errorf("ent: validator failed for field \"user_id\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Topic(); ok {
		if err := post.TopicValidator(v); err != nil {
			return &ValidationError{Name: "topic", err: fmt.Errorf("ent: validator failed for field \"topic\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := post.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := pu.mutation.ReplyNum(); ok {
		if err := post.ReplyNumValidator(v); err != nil {
			return &ValidationError{Name: "reply_num", err: fmt.Errorf("ent: validator failed for field \"reply_num\": %w", err)}
		}
	}
	if v, ok := pu.mutation.CreateAt(); ok {
		if err := post.CreateAtValidator(v); err != nil {
			return &ValidationError{Name: "create_at", err: fmt.Errorf("ent: validator failed for field \"create_at\": %w", err)}
		}
	}
	if v, ok := pu.mutation.UpdateAt(); ok {
		if err := post.UpdateAtValidator(v); err != nil {
			return &ValidationError{Name: "update_at", err: fmt.Errorf("ent: validator failed for field \"update_at\": %w", err)}
		}
	}
	return nil
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: post.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUserID,
		})
	}
	if value, ok := pu.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUserID,
		})
	}
	if value, ok := pu.mutation.Topic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldTopic,
		})
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: post.FieldStatus,
		})
	}
	if value, ok := pu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: post.FieldStatus,
		})
	}
	if value, ok := pu.mutation.ReplyNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldReplyNum,
		})
	}
	if value, ok := pu.mutation.AddedReplyNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldReplyNum,
		})
	}
	if value, ok := pu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldCreateAt,
		})
	}
	if value, ok := pu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldCreateAt,
		})
	}
	if value, ok := pu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUpdateAt,
		})
	}
	if value, ok := pu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUpdateAt,
		})
	}
	if pu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: comment.FieldID,
				},
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
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: comment.FieldID,
				},
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
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: comment.FieldID,
				},
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
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
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
func (puo *PostUpdateOne) SetUserID(i int64) *PostUpdateOne {
	puo.mutation.ResetUserID()
	puo.mutation.SetUserID(i)
	return puo
}

// AddUserID adds i to the "user_id" field.
func (puo *PostUpdateOne) AddUserID(i int64) *PostUpdateOne {
	puo.mutation.AddUserID(i)
	return puo
}

// SetTopic sets the "topic" field.
func (puo *PostUpdateOne) SetTopic(s string) *PostUpdateOne {
	puo.mutation.SetTopic(s)
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

// SetCreateAt sets the "create_at" field.
func (puo *PostUpdateOne) SetCreateAt(i int64) *PostUpdateOne {
	puo.mutation.ResetCreateAt()
	puo.mutation.SetCreateAt(i)
	return puo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableCreateAt(i *int64) *PostUpdateOne {
	if i != nil {
		puo.SetCreateAt(*i)
	}
	return puo
}

// AddCreateAt adds i to the "create_at" field.
func (puo *PostUpdateOne) AddCreateAt(i int64) *PostUpdateOne {
	puo.mutation.AddCreateAt(i)
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
func (puo *PostUpdateOne) AddCommentIDs(ids ...int64) *PostUpdateOne {
	puo.mutation.AddCommentIDs(ids...)
	return puo
}

// AddComments adds the "comments" edges to the Comment entity.
func (puo *PostUpdateOne) AddComments(c ...*Comment) *PostUpdateOne {
	ids := make([]int64, len(c))
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
func (puo *PostUpdateOne) RemoveCommentIDs(ids ...int64) *PostUpdateOne {
	puo.mutation.RemoveCommentIDs(ids...)
	return puo
}

// RemoveComments removes "comments" edges to Comment entities.
func (puo *PostUpdateOne) RemoveComments(c ...*Comment) *PostUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCommentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	var (
		err  error
		node *Post
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
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
			return &ValidationError{Name: "user_id", err: fmt.Errorf("ent: validator failed for field \"user_id\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Topic(); ok {
		if err := post.TopicValidator(v); err != nil {
			return &ValidationError{Name: "topic", err: fmt.Errorf("ent: validator failed for field \"topic\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := post.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := puo.mutation.ReplyNum(); ok {
		if err := post.ReplyNumValidator(v); err != nil {
			return &ValidationError{Name: "reply_num", err: fmt.Errorf("ent: validator failed for field \"reply_num\": %w", err)}
		}
	}
	if v, ok := puo.mutation.CreateAt(); ok {
		if err := post.CreateAtValidator(v); err != nil {
			return &ValidationError{Name: "create_at", err: fmt.Errorf("ent: validator failed for field \"create_at\": %w", err)}
		}
	}
	if v, ok := puo.mutation.UpdateAt(); ok {
		if err := post.UpdateAtValidator(v); err != nil {
			return &ValidationError{Name: "update_at", err: fmt.Errorf("ent: validator failed for field \"update_at\": %w", err)}
		}
	}
	return nil
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: post.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Post.ID for update")}
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
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUserID,
		})
	}
	if value, ok := puo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUserID,
		})
	}
	if value, ok := puo.mutation.Topic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldTopic,
		})
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: post.FieldStatus,
		})
	}
	if value, ok := puo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: post.FieldStatus,
		})
	}
	if value, ok := puo.mutation.ReplyNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldReplyNum,
		})
	}
	if value, ok := puo.mutation.AddedReplyNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldReplyNum,
		})
	}
	if value, ok := puo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldCreateAt,
		})
	}
	if value, ok := puo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldCreateAt,
		})
	}
	if value, ok := puo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUpdateAt,
		})
	}
	if value, ok := puo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUpdateAt,
		})
	}
	if puo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: comment.FieldID,
				},
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
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: comment.FieldID,
				},
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
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: comment.FieldID,
				},
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
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
