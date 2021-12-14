package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
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
		field.Uint64("id").Positive().Unique(),

		field.Uint64("user_id").Positive().Comment("The user that posted the message."),

		field.Uint64("post_id").Positive().Comment("The post that the message belongs to.").Optional(),

		field.String("content").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).NotEmpty().Comment("The content of the message."),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint unsigned", // Override MySQL.
		}).NonNegative().Default(0).Comment("状态,是否被封，0：正常，大于0被封"),

		field.Uint64("floor").Positive().Comment("第几楼"),

		field.Int64("create_at").SchemaType(map[string]string{
			dialect.MySQL: "int UNSIGNED", // Override MySQL.
		}).NonNegative().Default(0).Comment("创建时间").Immutable(),
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
		index.Fields("user_id"),
	}
}
