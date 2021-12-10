// Code generated by entc, DO NOT EDIT.

package ent

import (
	"stock_service/repository/ent/schema"
	"stock_service/repository/ent/stock"
	"stock_service/repository/ent/tradedate"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	stockFields := schema.Stock{}.Fields()
	_ = stockFields
	// stockDescSymbol is the schema descriptor for symbol field.
	stockDescSymbol := stockFields[0].Descriptor()
	// stock.SymbolValidator is a validator for the "symbol" field. It is called by the builders before save.
	stock.SymbolValidator = stockDescSymbol.Validators[0].(func(string) error)
	// stockDescName is the schema descriptor for name field.
	stockDescName := stockFields[1].Descriptor()
	// stock.NameValidator is a validator for the "name" field. It is called by the builders before save.
	stock.NameValidator = stockDescName.Validators[0].(func(string) error)
	tradedateFields := schema.TradeDate{}.Fields()
	_ = tradedateFields
	// tradedateDescStockID is the schema descriptor for stock_id field.
	tradedateDescStockID := tradedateFields[0].Descriptor()
	// tradedate.StockIDValidator is a validator for the "stock_id" field. It is called by the builders before save.
	tradedate.StockIDValidator = tradedateDescStockID.Validators[0].(func(int) error)
	// tradedateDescTDate is the schema descriptor for t_date field.
	tradedateDescTDate := tradedateFields[1].Descriptor()
	// tradedate.TDateValidator is a validator for the "t_date" field. It is called by the builders before save.
	tradedate.TDateValidator = tradedateDescTDate.Validators[0].(func(string) error)
	// tradedateDescCreateAt is the schema descriptor for create_at field.
	tradedateDescCreateAt := tradedateFields[3].Descriptor()
	// tradedate.DefaultCreateAt holds the default value on creation for the create_at field.
	tradedate.DefaultCreateAt = tradedateDescCreateAt.Default.(int64)
	// tradedate.CreateAtValidator is a validator for the "create_at" field. It is called by the builders before save.
	tradedate.CreateAtValidator = tradedateDescCreateAt.Validators[0].(func(int64) error)
	// tradedateDescUpdateAt is the schema descriptor for update_at field.
	tradedateDescUpdateAt := tradedateFields[4].Descriptor()
	// tradedate.DefaultUpdateAt holds the default value on creation for the update_at field.
	tradedate.DefaultUpdateAt = tradedateDescUpdateAt.Default.(int64)
	// tradedate.UpdateAtValidator is a validator for the "update_at" field. It is called by the builders before save.
	tradedate.UpdateAtValidator = tradedateDescUpdateAt.Validators[0].(func(int64) error)
	// tradedateDescXueqiuCommentCount is the schema descriptor for xueqiu_comment_count field.
	tradedateDescXueqiuCommentCount := tradedateFields[5].Descriptor()
	// tradedate.DefaultXueqiuCommentCount holds the default value on creation for the xueqiu_comment_count field.
	tradedate.DefaultXueqiuCommentCount = tradedateDescXueqiuCommentCount.Default.(int64)
	// tradedate.XueqiuCommentCountValidator is a validator for the "xueqiu_comment_count" field. It is called by the builders before save.
	tradedate.XueqiuCommentCountValidator = tradedateDescXueqiuCommentCount.Validators[0].(func(int64) error)
}
