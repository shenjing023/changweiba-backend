// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// StockColumns holds the columns for the "stock" table.
	StockColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "symbol", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "varchar(10)"}},
		{Name: "name", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "varchar(10)"}},
	}
	// StockTable holds the schema information for the "stock" table.
	StockTable = &schema.Table{
		Name:       "stock",
		Columns:    StockColumns,
		PrimaryKey: []*schema.Column{StockColumns[0]},
	}
	// TradeDateColumns holds the columns for the "trade_date" table.
	TradeDateColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "t_date", Type: field.TypeString, SchemaType: map[string]string{"mysql": "date"}},
		{Name: "close", Type: field.TypeFloat64},
		{Name: "volume", Type: field.TypeFloat64},
		{Name: "create_at", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "update_at", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "xueqiu_comment_count", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "stock_id", Type: field.TypeUint64, Nullable: true},
	}
	// TradeDateTable holds the schema information for the "trade_date" table.
	TradeDateTable = &schema.Table{
		Name:       "trade_date",
		Columns:    TradeDateColumns,
		PrimaryKey: []*schema.Column{TradeDateColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "trade_date_stock_trades",
				Columns:    []*schema.Column{TradeDateColumns[7]},
				RefColumns: []*schema.Column{StockColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "tradedate_stock_id",
				Unique:  false,
				Columns: []*schema.Column{TradeDateColumns[7]},
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "nick_name", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "varchar(20)"}},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// StockSubscribersColumns holds the columns for the "stock_subscribers" table.
	StockSubscribersColumns = []*schema.Column{
		{Name: "stock_id", Type: field.TypeUint64},
		{Name: "user_id", Type: field.TypeUint64},
	}
	// StockSubscribersTable holds the schema information for the "stock_subscribers" table.
	StockSubscribersTable = &schema.Table{
		Name:       "stock_subscribers",
		Columns:    StockSubscribersColumns,
		PrimaryKey: []*schema.Column{StockSubscribersColumns[0], StockSubscribersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "stock_subscribers_stock_id",
				Columns:    []*schema.Column{StockSubscribersColumns[0]},
				RefColumns: []*schema.Column{StockColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "stock_subscribers_user_id",
				Columns:    []*schema.Column{StockSubscribersColumns[1]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		StockTable,
		TradeDateTable,
		UserTable,
		StockSubscribersTable,
	}
)

func init() {
	StockTable.Annotation = &entsql.Annotation{
		Table: "stock",
	}
	TradeDateTable.ForeignKeys[0].RefTable = StockTable
	TradeDateTable.Annotation = &entsql.Annotation{
		Table: "trade_date",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
	StockSubscribersTable.ForeignKeys[0].RefTable = StockTable
	StockSubscribersTable.ForeignKeys[1].RefTable = UserTable
}
