package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
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
			MaxLen(20).Unique().Comment("昵称"),
		field.String("avatar").NotEmpty().Comment("头像"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("comments", Comment.Type),
		edge.To("likes", Post.Type),
		edge.To("replies", Reply.Type),
	}
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
