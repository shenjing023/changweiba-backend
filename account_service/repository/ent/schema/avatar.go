package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Avatar holds the schema definition for the Avatar entity.
type Avatar struct {
	ent.Schema
}

// Fields of the Avatar.
func (Avatar) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive().Unique(),

		field.String("url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).NotEmpty().Comment("头像url"),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint unsigned", // Override MySQL.
		}).NonNegative().Default(0).Comment("状态，0：正常，1：不可用"),
	}
}

// Edges of the Avatar.
func (Avatar) Edges() []ent.Edge {
	return nil
}

func (Avatar) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "avatar"},
	}
}
