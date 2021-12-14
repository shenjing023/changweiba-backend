// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"stock_service/repository/ent/stock"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Stock is the model entity for the Stock schema.
type Stock struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Symbol holds the value of the "symbol" field.
	// 股票代码
	Symbol string `json:"symbol,omitempty"`
	// Name holds the value of the "name" field.
	// 股票名称
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StockQuery when eager-loading is set.
	Edges StockEdges `json:"edges"`
}

// StockEdges holds the relations/edges for other nodes in the graph.
type StockEdges struct {
	// Trades holds the value of the trades edge.
	Trades []*TradeDate `json:"trades,omitempty"`
	// Subscribers holds the value of the subscribers edge.
	Subscribers []*User `json:"subscribers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TradesOrErr returns the Trades value or an error if the edge
// was not loaded in eager-loading.
func (e StockEdges) TradesOrErr() ([]*TradeDate, error) {
	if e.loadedTypes[0] {
		return e.Trades, nil
	}
	return nil, &NotLoadedError{edge: "trades"}
}

// SubscribersOrErr returns the Subscribers value or an error if the edge
// was not loaded in eager-loading.
func (e StockEdges) SubscribersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Subscribers, nil
	}
	return nil, &NotLoadedError{edge: "subscribers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Stock) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case stock.FieldID:
			values[i] = new(sql.NullInt64)
		case stock.FieldSymbol, stock.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Stock", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Stock fields.
func (s *Stock) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case stock.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint64(value.Int64)
		case stock.FieldSymbol:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field symbol", values[i])
			} else if value.Valid {
				s.Symbol = value.String
			}
		case stock.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		}
	}
	return nil
}

// QueryTrades queries the "trades" edge of the Stock entity.
func (s *Stock) QueryTrades() *TradeDateQuery {
	return (&StockClient{config: s.config}).QueryTrades(s)
}

// QuerySubscribers queries the "subscribers" edge of the Stock entity.
func (s *Stock) QuerySubscribers() *UserQuery {
	return (&StockClient{config: s.config}).QuerySubscribers(s)
}

// Update returns a builder for updating this Stock.
// Note that you need to call Stock.Unwrap() before calling this method if this Stock
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Stock) Update() *StockUpdateOne {
	return (&StockClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Stock entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Stock) Unwrap() *Stock {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Stock is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Stock) String() string {
	var builder strings.Builder
	builder.WriteString("Stock(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", symbol=")
	builder.WriteString(s.Symbol)
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Stocks is a parsable slice of Stock.
type Stocks []*Stock

func (s Stocks) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
