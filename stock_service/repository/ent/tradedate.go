// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"stock_service/repository/ent/stock"
	"stock_service/repository/ent/tradedate"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TradeDate is the model entity for the TradeDate schema.
type TradeDate struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Stock ID
	StockID uint64 `json:"stock_id,omitempty"`
	// 交易日期
	TDate string `json:"t_date,omitempty"`
	// 收盘价
	Close float64 `json:"close,omitempty"`
	// 成交量
	Volume float64 `json:"volume,omitempty"`
	// 创建时间
	CreateAt int64 `json:"create_at,omitempty"`
	// 最后更新时间
	UpdateAt int64 `json:"update_at,omitempty"`
	// 雪球评论数
	XueqiuCommentCount int64 `json:"xueqiu_comment_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TradeDateQuery when eager-loading is set.
	Edges        TradeDateEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TradeDateEdges holds the relations/edges for other nodes in the graph.
type TradeDateEdges struct {
	// Stock holds the value of the stock edge.
	Stock *Stock `json:"stock,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// StockOrErr returns the Stock value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TradeDateEdges) StockOrErr() (*Stock, error) {
	if e.loadedTypes[0] {
		if e.Stock == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: stock.Label}
		}
		return e.Stock, nil
	}
	return nil, &NotLoadedError{edge: "stock"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TradeDate) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tradedate.FieldClose, tradedate.FieldVolume:
			values[i] = new(sql.NullFloat64)
		case tradedate.FieldID, tradedate.FieldStockID, tradedate.FieldCreateAt, tradedate.FieldUpdateAt, tradedate.FieldXueqiuCommentCount:
			values[i] = new(sql.NullInt64)
		case tradedate.FieldTDate:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TradeDate fields.
func (td *TradeDate) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tradedate.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			td.ID = uint64(value.Int64)
		case tradedate.FieldStockID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stock_id", values[i])
			} else if value.Valid {
				td.StockID = uint64(value.Int64)
			}
		case tradedate.FieldTDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field t_date", values[i])
			} else if value.Valid {
				td.TDate = value.String
			}
		case tradedate.FieldClose:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field close", values[i])
			} else if value.Valid {
				td.Close = value.Float64
			}
		case tradedate.FieldVolume:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field volume", values[i])
			} else if value.Valid {
				td.Volume = value.Float64
			}
		case tradedate.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				td.CreateAt = value.Int64
			}
		case tradedate.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				td.UpdateAt = value.Int64
			}
		case tradedate.FieldXueqiuCommentCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field xueqiu_comment_count", values[i])
			} else if value.Valid {
				td.XueqiuCommentCount = value.Int64
			}
		default:
			td.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TradeDate.
// This includes values selected through modifiers, order, etc.
func (td *TradeDate) Value(name string) (ent.Value, error) {
	return td.selectValues.Get(name)
}

// QueryStock queries the "stock" edge of the TradeDate entity.
func (td *TradeDate) QueryStock() *StockQuery {
	return NewTradeDateClient(td.config).QueryStock(td)
}

// Update returns a builder for updating this TradeDate.
// Note that you need to call TradeDate.Unwrap() before calling this method if this TradeDate
// was returned from a transaction, and the transaction was committed or rolled back.
func (td *TradeDate) Update() *TradeDateUpdateOne {
	return NewTradeDateClient(td.config).UpdateOne(td)
}

// Unwrap unwraps the TradeDate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (td *TradeDate) Unwrap() *TradeDate {
	_tx, ok := td.config.driver.(*txDriver)
	if !ok {
		panic("ent: TradeDate is not a transactional entity")
	}
	td.config.driver = _tx.drv
	return td
}

// String implements the fmt.Stringer.
func (td *TradeDate) String() string {
	var builder strings.Builder
	builder.WriteString("TradeDate(")
	builder.WriteString(fmt.Sprintf("id=%v, ", td.ID))
	builder.WriteString("stock_id=")
	builder.WriteString(fmt.Sprintf("%v", td.StockID))
	builder.WriteString(", ")
	builder.WriteString("t_date=")
	builder.WriteString(td.TDate)
	builder.WriteString(", ")
	builder.WriteString("close=")
	builder.WriteString(fmt.Sprintf("%v", td.Close))
	builder.WriteString(", ")
	builder.WriteString("volume=")
	builder.WriteString(fmt.Sprintf("%v", td.Volume))
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(fmt.Sprintf("%v", td.CreateAt))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(fmt.Sprintf("%v", td.UpdateAt))
	builder.WriteString(", ")
	builder.WriteString("xueqiu_comment_count=")
	builder.WriteString(fmt.Sprintf("%v", td.XueqiuCommentCount))
	builder.WriteByte(')')
	return builder.String()
}

// TradeDates is a parsable slice of TradeDate.
type TradeDates []*TradeDate
