package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Int("author_id").Positive().Optional().Comment("The user that posted the message."),

		field.String("title").NotEmpty().MaxLen(100).Comment("The title of the message."),

		field.String("content").NotEmpty().MaxLen(2048).Comment("The content of the message."),

		field.Enum("status").
			Values(StatusNormal, StatusBanned, StatusDeleted).
			Default(StatusNormal).Comment("状态"),

		field.Int("reply_num").NonNegative().Default(0).Comment("回复数"),

		field.Int8("pin").NonNegative().Default(0).Comment("是否置顶，0：否，1是"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type),
		edge.From("author", User.Type).
			Ref("posts").Field("author_id").Unique(),
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "post"},
	}
}

func (Post) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("author_id"),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
