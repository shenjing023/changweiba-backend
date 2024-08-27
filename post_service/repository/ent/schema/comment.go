package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{

		field.Int("author_id").Positive().Optional().Comment("The user that posted the message."),

		field.Int("post_id").Positive().Comment("The post that the message belongs to.").Optional(),

		field.String("content").MaxLen(1024).NotEmpty().Comment("The content of the message."),

		field.Enum("status").
			Values(StatusNormal, StatusBanned, StatusDeleted).
			Default(StatusNormal).Comment("状态"),

		field.Int("floor").Positive().Comment("第几楼"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Post.Type).
			Ref("comments").
			Unique().
			Field("post_id"),
		edge.To("replies", Reply.Type),
		edge.From("author", User.Type).
			Ref("comments").
			Unique().
			Field("author_id"),
	}
}

func (Comment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "comment"},
	}
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("author_id"),
		index.Fields("post_id"),
	}
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
