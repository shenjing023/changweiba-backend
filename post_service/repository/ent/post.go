// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cw_post_service/repository/ent/post"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Post is the model entity for the Post schema.
type Post struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	// The user that posted the message.
	UserID int64 `json:"user_id,omitempty"`
	// Topic holds the value of the "topic" field.
	// The topic of the message.
	Topic string `json:"topic,omitempty"`
	// Status holds the value of the "status" field.
	// 状态,是否被封，0：正常，大于0被封
	Status int8 `json:"status,omitempty"`
	// ReplyNum holds the value of the "reply_num" field.
	// 回复数
	ReplyNum int64 `json:"reply_num,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	// 创建时间
	CreateAt int64 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	// 最后更新时间
	UpdateAt int64 `json:"update_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostQuery when eager-loading is set.
	Edges PostEdges `json:"edges"`
}

// PostEdges holds the relations/edges for other nodes in the graph.
type PostEdges struct {
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[0] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Post) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldID, post.FieldUserID, post.FieldStatus, post.FieldReplyNum, post.FieldCreateAt, post.FieldUpdateAt:
			values[i] = new(sql.NullInt64)
		case post.FieldTopic:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Post", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Post fields.
func (po *Post) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int64(value.Int64)
		case post.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				po.UserID = value.Int64
			}
		case post.FieldTopic:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field topic", values[i])
			} else if value.Valid {
				po.Topic = value.String
			}
		case post.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				po.Status = int8(value.Int64)
			}
		case post.FieldReplyNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reply_num", values[i])
			} else if value.Valid {
				po.ReplyNum = value.Int64
			}
		case post.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				po.CreateAt = value.Int64
			}
		case post.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				po.UpdateAt = value.Int64
			}
		}
	}
	return nil
}

// QueryComments queries the "comments" edge of the Post entity.
func (po *Post) QueryComments() *CommentQuery {
	return (&PostClient{config: po.config}).QueryComments(po)
}

// Update returns a builder for updating this Post.
// Note that you need to call Post.Unwrap() before calling this method if this Post
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Post) Update() *PostUpdateOne {
	return (&PostClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Post entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Post) Unwrap() *Post {
	tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Post is not a transactional entity")
	}
	po.config.driver = tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", po.UserID))
	builder.WriteString(", topic=")
	builder.WriteString(po.Topic)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", po.Status))
	builder.WriteString(", reply_num=")
	builder.WriteString(fmt.Sprintf("%v", po.ReplyNum))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", po.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", po.UpdateAt))
	builder.WriteByte(')')
	return builder.String()
}

// Posts is a parsable slice of Post.
type Posts []*Post

func (po Posts) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
