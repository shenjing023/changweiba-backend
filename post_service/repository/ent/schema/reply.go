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

// Reply holds the schema definition for the Reply entity.
type Reply struct {
	ent.Schema
}

// Fields of the Reply.
func (Reply) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive().Unique(),

		field.Uint64("user_id").Positive().Comment("The user that posted the message."),

		field.Uint64("comment_id").Positive().Comment("The comment that this reply is for.").Optional(),

		field.Uint64("parent_id").Positive().Comment("回复哪个回复的id").Optional(),

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
		index.Fields("user_id"),
	}
}
