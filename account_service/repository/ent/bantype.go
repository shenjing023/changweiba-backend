// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cw_account_service/repository/ent/bantype"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// BanType is the model entity for the BanType schema.
type BanType struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Content holds the value of the "content" field.
	// 具体ban的内容
	Content string `json:"content,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BanType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case bantype.FieldID:
			values[i] = new(sql.NullInt64)
		case bantype.FieldContent:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BanType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BanType fields.
func (bt *BanType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bantype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bt.ID = uint64(value.Int64)
		case bantype.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				bt.Content = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this BanType.
// Note that you need to call BanType.Unwrap() before calling this method if this BanType
// was returned from a transaction, and the transaction was committed or rolled back.
func (bt *BanType) Update() *BanTypeUpdateOne {
	return (&BanTypeClient{config: bt.config}).UpdateOne(bt)
}

// Unwrap unwraps the BanType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bt *BanType) Unwrap() *BanType {
	tx, ok := bt.config.driver.(*txDriver)
	if !ok {
		panic("ent: BanType is not a transactional entity")
	}
	bt.config.driver = tx.drv
	return bt
}

// String implements the fmt.Stringer.
func (bt *BanType) String() string {
	var builder strings.Builder
	builder.WriteString("BanType(")
	builder.WriteString(fmt.Sprintf("id=%v", bt.ID))
	builder.WriteString(", content=")
	builder.WriteString(bt.Content)
	builder.WriteByte(')')
	return builder.String()
}

// BanTypes is a parsable slice of BanType.
type BanTypes []*BanType

func (bt BanTypes) config(cfg config) {
	for _i := range bt {
		bt[_i].config = cfg
	}
}
