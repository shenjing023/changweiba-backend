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

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive().Unique(),

		field.Uint64("user_id").Positive().Comment("The user that posted the message."),

		field.String("title").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).NotEmpty().Comment("The title of the message."),

		field.String("content").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).NotEmpty().Comment("The content of the message."),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint unsigned", // Override MySQL.
		}).NonNegative().Default(0).Comment("状态,是否被封，0：正常，大于0被封"),

		field.Int64("reply_num").SchemaType(map[string]string{
			dialect.MySQL: "int UNSIGNED", // Override MySQL.
		}).NonNegative().Default(0).Comment("回复数"),

		field.Int64("create_at").SchemaType(map[string]string{
			dialect.MySQL: "int UNSIGNED", // Override MySQL.
		}).NonNegative().Default(0).Comment("创建时间").Immutable(),

		field.Int64("update_at").SchemaType(map[string]string{
			dialect.MySQL: "int UNSIGNED", // Override MySQL.
		}).NonNegative().Default(0).Comment("最后更新时间"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type),
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
		index.Fields("user_id"),
	}
}
