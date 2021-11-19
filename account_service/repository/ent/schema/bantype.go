package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// BanType holds the schema definition for the BanType entity.
type BanType struct {
	ent.Schema
}

// Fields of the BanType.
func (BanType) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "int UNSIGNED", // Override MySQL.
		}).NonNegative().Unique(),

		field.String("content").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).NotEmpty().Comment("具体ban的内容"),
	}
}

// Edges of the BanType.
func (BanType) Edges() []ent.Edge {
	return nil
}

func (BanType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "ban_type"},
	}
}
