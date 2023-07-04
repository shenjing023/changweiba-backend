package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Stock holds the schema definition for the Stock entity.
type Hot struct {
	ent.Schema
}

// Fields of the Stock.
func (Hot) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive().Unique(),
		field.String("symbol").SchemaType(map[string]string{
			dialect.MySQL: "varchar(10)", // Override MySQL.
		}).NotEmpty().Comment("股票代码"),
		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(10)", // Override MySQL.
		}).NotEmpty().Comment("股票名称"),
		field.String("t_date").SchemaType(map[string]string{
			dialect.MySQL: "date", // Override MySQL.
		}).NotEmpty().Comment("交易日期"),
		field.Int("order").Comment("排行榜排名").Default(0),
		field.String("tag").Comment("题材标签").Default("[]"),
		field.Int("bull").Comment("持仓建议").Default(0),
		field.String("short").NotEmpty().Comment("短期趋势").Default("---"),
		field.Text("analyse").Comment("投资建议分析").Default(""),
	}
}

func (Hot) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "stock_hot"},
	}
}

func (Hot) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("t_date"),
	}
}
