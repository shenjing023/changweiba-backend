// Code generated by ent, DO NOT EDIT.

package tradedate

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tradedate type in the database.
	Label = "trade_date"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStockID holds the string denoting the stock_id field in the database.
	FieldStockID = "stock_id"
	// FieldTDate holds the string denoting the t_date field in the database.
	FieldTDate = "t_date"
	// FieldClose holds the string denoting the close field in the database.
	FieldClose = "close"
	// FieldVolume holds the string denoting the volume field in the database.
	FieldVolume = "volume"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldXueqiuCommentCount holds the string denoting the xueqiu_comment_count field in the database.
	FieldXueqiuCommentCount = "xueqiu_comment_count"
	// EdgeStock holds the string denoting the stock edge name in mutations.
	EdgeStock = "stock"
	// Table holds the table name of the tradedate in the database.
	Table = "trade_date"
	// StockTable is the table that holds the stock relation/edge.
	StockTable = "trade_date"
	// StockInverseTable is the table name for the Stock entity.
	// It exists in this package in order to avoid circular dependency with the "stock" package.
	StockInverseTable = "stock"
	// StockColumn is the table column denoting the stock relation/edge.
	StockColumn = "stock_id"
)

// Columns holds all SQL columns for tradedate fields.
var Columns = []string{
	FieldID,
	FieldStockID,
	FieldTDate,
	FieldClose,
	FieldVolume,
	FieldCreateAt,
	FieldUpdateAt,
	FieldXueqiuCommentCount,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// StockIDValidator is a validator for the "stock_id" field. It is called by the builders before save.
	StockIDValidator func(uint64) error
	// TDateValidator is a validator for the "t_date" field. It is called by the builders before save.
	TDateValidator func(string) error
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt int64
	// CreateAtValidator is a validator for the "create_at" field. It is called by the builders before save.
	CreateAtValidator func(int64) error
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt int64
	// UpdateAtValidator is a validator for the "update_at" field. It is called by the builders before save.
	UpdateAtValidator func(int64) error
	// DefaultXueqiuCommentCount holds the default value on creation for the "xueqiu_comment_count" field.
	DefaultXueqiuCommentCount int64
	// XueqiuCommentCountValidator is a validator for the "xueqiu_comment_count" field. It is called by the builders before save.
	XueqiuCommentCountValidator func(int64) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint64) error
)

// OrderOption defines the ordering options for the TradeDate queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStockID orders the results by the stock_id field.
func ByStockID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStockID, opts...).ToFunc()
}

// ByTDate orders the results by the t_date field.
func ByTDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTDate, opts...).ToFunc()
}

// ByClose orders the results by the close field.
func ByClose(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClose, opts...).ToFunc()
}

// ByVolume orders the results by the volume field.
func ByVolume(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVolume, opts...).ToFunc()
}

// ByCreateAt orders the results by the create_at field.
func ByCreateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateAt, opts...).ToFunc()
}

// ByUpdateAt orders the results by the update_at field.
func ByUpdateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateAt, opts...).ToFunc()
}

// ByXueqiuCommentCount orders the results by the xueqiu_comment_count field.
func ByXueqiuCommentCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldXueqiuCommentCount, opts...).ToFunc()
}

// ByStockField orders the results by stock field.
func ByStockField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStockStep(), sql.OrderByField(field, opts...))
	}
}
func newStockStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StockInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, StockTable, StockColumn),
	)
}
