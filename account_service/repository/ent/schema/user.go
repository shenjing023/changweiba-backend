package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

const (
	UserStatusNormal = "NORMAL"
	UserStatusBanned = "BANNED"
)

const (
	UserRoleAdmin = "ADMIN"
	UserRoleUser  = "USER"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().
			MaxLen(40).Unique().Comment("昵称"),
		field.String("password").NotEmpty().
			Comment("密码").Sensitive(),
		field.String("avatar").NotEmpty().Comment("头像"),
		field.Int("score").NonNegative().
			Default(0).Comment("分数"),
		field.String("description").MaxLen(150).Optional().Comment("简介"),
		field.Enum("status").
			Values(UserStatusNormal, UserStatusBanned).
			Default(UserStatusNormal).Comment("用户状态"),
		field.Enum("role").
			Values(UserRoleUser, UserRoleAdmin).
			Default(UserRoleUser).Comment("用户角色"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
