package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Stock holds the schema definition for the Stock entity.
type Stock struct {
	ent.Schema
}

// Fields of the Stock.
func (Stock) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive().Unique(),
		field.String("symbol").SchemaType(map[string]string{
			dialect.MySQL: "varchar(10)", // Override MySQL.
		}).NotEmpty().Unique().Comment("股票代码"),
		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(10)", // Override MySQL.
		}).NotEmpty().Unique().Comment("股票名称"),
		field.Int("bull").Comment("持仓建议").Default(0),
		field.Time("last_subscribe_at").
			Default(time.Now).Comment("最后被订阅的时间"),
		field.String("short").NotEmpty().Comment("短期趋势").Default("---"),
	}
}

// Edges of the Stock.
func (Stock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("trades", TradeDate.Type),
		edge.To("subscribers", User.Type),
	}
}

func (Stock) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "stock"},
	}
}
