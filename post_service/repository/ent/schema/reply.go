package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Reply holds the schema definition for the Reply entity.
type Reply struct {
	ent.Schema
}

// Fields of the Reply.
func (Reply) Fields() []ent.Field {
	return []ent.Field{

		field.Int("author_id").Positive().Optional().Comment("The user that posted the message."),

		field.Int("comment_id").Positive().Comment("The comment that this reply is for.").Optional(),

		field.Int("parent_id").Positive().Comment("回复哪个回复的id").Optional(),

		field.String("content").MaxLen(50).NotEmpty().Comment("The content of the message."),

		field.Enum("status").
			Values(StatusNormal, StatusBanned, StatusDeleted).
			Default(StatusNormal).Comment("状态"),
	}
}

// Edges of the Reply.
func (Reply) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Comment.Type).
			Ref("replies").
			Unique().
			Field("comment_id"),
		edge.To("children", Reply.Type).
			From("parent").
			Unique().
			Field("parent_id"),
		edge.From("author", User.Type).
			Ref("replies").
			Field("author_id").
			Unique(),
	}
}

func (Reply) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "reply"},
	}
}

func (Reply) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("author_id"),
		index.Fields("comment_id"),
	}
}

func (Reply) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
