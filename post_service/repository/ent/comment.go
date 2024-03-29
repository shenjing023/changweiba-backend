// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cw_post_service/repository/ent/comment"
	"cw_post_service/repository/ent/post"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// The user that posted the message.
	UserID uint64 `json:"user_id,omitempty"`
	// The post that the message belongs to.
	PostID uint64 `json:"post_id,omitempty"`
	// The content of the message.
	Content string `json:"content,omitempty"`
	// 状态,是否被封，0：正常，大于0被封
	Status int8 `json:"status,omitempty"`
	// 第几楼
	Floor uint64 `json:"floor,omitempty"`
	// 创建时间
	CreateAt int64 `json:"create_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentQuery when eager-loading is set.
	Edges        CommentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CommentEdges holds the relations/edges for other nodes in the graph.
type CommentEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Post `json:"owner,omitempty"`
	// Replies holds the value of the replies edge.
	Replies []*Reply `json:"replies,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) OwnerOrErr() (*Post, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: post.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// RepliesOrErr returns the Replies value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) RepliesOrErr() ([]*Reply, error) {
	if e.loadedTypes[1] {
		return e.Replies, nil
	}
	return nil, &NotLoadedError{edge: "replies"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldID, comment.FieldUserID, comment.FieldPostID, comment.FieldStatus, comment.FieldFloor, comment.FieldCreateAt:
			values[i] = new(sql.NullInt64)
		case comment.FieldContent:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint64(value.Int64)
		case comment.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				c.UserID = uint64(value.Int64)
			}
		case comment.FieldPostID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field post_id", values[i])
			} else if value.Valid {
				c.PostID = uint64(value.Int64)
			}
		case comment.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				c.Content = value.String
			}
		case comment.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = int8(value.Int64)
			}
		case comment.FieldFloor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field floor", values[i])
			} else if value.Valid {
				c.Floor = uint64(value.Int64)
			}
		case comment.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				c.CreateAt = value.Int64
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Comment.
// This includes values selected through modifiers, order, etc.
func (c *Comment) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Comment entity.
func (c *Comment) QueryOwner() *PostQuery {
	return NewCommentClient(c.config).QueryOwner(c)
}

// QueryReplies queries the "replies" edge of the Comment entity.
func (c *Comment) QueryReplies() *ReplyQuery {
	return NewCommentClient(c.config).QueryReplies(c)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return NewCommentClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", c.UserID))
	builder.WriteString(", ")
	builder.WriteString("post_id=")
	builder.WriteString(fmt.Sprintf("%v", c.PostID))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(c.Content)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("floor=")
	builder.WriteString(fmt.Sprintf("%v", c.Floor))
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(fmt.Sprintf("%v", c.CreateAt))
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment
