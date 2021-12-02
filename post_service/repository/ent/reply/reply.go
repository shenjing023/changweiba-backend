// Code generated by entc, DO NOT EDIT.

package reply

const (
	// Label holds the string label denoting the reply type in the database.
	Label = "reply"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCommentID holds the string denoting the comment_id field in the database.
	FieldCommentID = "comment_id"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
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
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// Table holds the table name of the reply in the database.
	Table = "reply"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "reply"
	// OwnerInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	OwnerInverseTable = "comment"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "comment_id"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "reply"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "reply"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
)

// Columns holds all SQL columns for reply fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldCommentID,
	FieldParentID,
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
	UserIDValidator func(int64) error
	// CommentIDValidator is a validator for the "comment_id" field. It is called by the builders before save.
	CommentIDValidator func(int64) error
	// ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	ParentIDValidator func(int64) error
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(int8) error
	// DefaultFloor holds the default value on creation for the "floor" field.
	DefaultFloor int64
	// FloorValidator is a validator for the "floor" field. It is called by the builders before save.
	FloorValidator func(int64) error
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt int64
	// CreateAtValidator is a validator for the "create_at" field. It is called by the builders before save.
	CreateAtValidator func(int64) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)
