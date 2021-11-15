// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNickName holds the string denoting the nick_name field in the database.
	FieldNickName = "nick_name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// Table holds the table name of the user in the database.
	Table = "user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldNickName,
	FieldPassword,
	FieldAvatar,
	FieldStatus,
	FieldScore,
	FieldRole,
	FieldCreateAt,
	FieldUpdateAt,
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
	// NickNameValidator is a validator for the "nick_name" field. It is called by the builders before save.
	NickNameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	AvatarValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(int8) error
	// DefaultScore holds the default value on creation for the "score" field.
	DefaultScore int
	// ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	ScoreValidator func(int) error
	// DefaultRole holds the default value on creation for the "role" field.
	DefaultRole int8
	// RoleValidator is a validator for the "role" field. It is called by the builders before save.
	RoleValidator func(int8) error
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt int64
	// CreateAtValidator is a validator for the "create_at" field. It is called by the builders before save.
	CreateAtValidator func(int64) error
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt int64
	// UpdateAtValidator is a validator for the "update_at" field. It is called by the builders before save.
	UpdateAtValidator func(int64) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)
