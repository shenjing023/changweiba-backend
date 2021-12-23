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

// TradeDate holds the schema definition for the TradeDate entity.
type TradeDate struct {
	ent.Schema
}

// Fields of the TradeDate.
func (TradeDate) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive().Unique(),
		field.Uint64("stock_id").Positive().Comment("Stock ID").Optional(),
		field.String("t_date").SchemaType(map[string]string{
			dialect.MySQL: "date", // Override MySQL.
		}).NotEmpty().Comment("交易日期"),
		field.Float("close").Comment("收盘价"),
		field.Float("volume").Comment("成交量"),
		field.Int64("create_at").SchemaType(map[string]string{
			dialect.MySQL: "int UNSIGNED", // Override MySQL.
		}).NonNegative().Default(0).Comment("创建时间").Immutable(),

		field.Int64("update_at").SchemaType(map[string]string{
			dialect.MySQL: "int UNSIGNED", // Override MySQL.
		}).NonNegative().Default(0).Comment("最后更新时间"),

		field.Int64("xueqiu_comment_count").SchemaType(map[string]string{
			dialect.MySQL: "int UNSIGNED", // Override MySQL.
		}).NonNegative().Default(0).Comment("雪球评论数"),
	}
}

// Edges of the TradeDate.
func (TradeDate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("stock", Stock.Type).Ref("trades").Unique().Field("stock_id"),
	}
}

func (TradeDate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "trade_date"},
	}
}

func (TradeDate) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("stock_id"),
		index.Fields("stock_id", "t_date").Unique(),
	}
}
