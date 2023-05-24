// Code generated by ent, DO NOT EDIT.

package comment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldPostID holds the string denoting the post_id field in the database.
	FieldPostID = "post_id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldFloor holds the string denoting the floor field in the database.
	FieldFloor = "floor"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeReplies holds the string denoting the replies edge name in mutations.
	EdgeReplies = "replies"
	// Table holds the table name of the comment in the database.
	Table = "comment"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "comment"
	// OwnerInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	OwnerInverseTable = "post"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "post_id"
	// RepliesTable is the table that holds the replies relation/edge.
	RepliesTable = "reply"
	// RepliesInverseTable is the table name for the Reply entity.
	// It exists in this package in order to avoid circular dependency with the "reply" package.
	RepliesInverseTable = "reply"
	// RepliesColumn is the table column denoting the replies relation/edge.
	RepliesColumn = "comment_id"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldPostID,
	FieldContent,
	FieldStatus,
	FieldFloor,
	FieldCreateAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(uint64) error
	// PostIDValidator is a validator for the "post_id" field. It is called by the builders before save.
	PostIDValidator func(uint64) error
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(int8) error
	// FloorValidator is a validator for the "floor" field. It is called by the builders before save.
	FloorValidator func(uint64) error
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt int64
	// CreateAtValidator is a validator for the "create_at" field. It is called by the builders before save.
	CreateAtValidator func(int64) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint64) error
)

// OrderOption defines the ordering options for the Comment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByPostID orders the results by the post_id field.
func ByPostID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostID, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByFloor orders the results by the floor field.
func ByFloor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFloor, opts...).ToFunc()
}

// ByCreateAt orders the results by the create_at field.
func ByCreateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateAt, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByRepliesCount orders the results by replies count.
func ByRepliesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRepliesStep(), opts...)
	}
}

// ByReplies orders the results by replies terms.
func ByReplies(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRepliesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newRepliesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RepliesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RepliesTable, RepliesColumn),
	)
}
