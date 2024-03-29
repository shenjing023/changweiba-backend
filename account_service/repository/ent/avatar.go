// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cw_account_service/repository/ent/avatar"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Avatar is the model entity for the Avatar schema.
type Avatar struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	// 头像url
	URL string `json:"url,omitempty"`
	// Status holds the value of the "status" field.
	// 状态，0：正常，1：不可用
	Status int8 `json:"status,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Avatar) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case avatar.FieldID, avatar.FieldStatus:
			values[i] = new(sql.NullInt64)
		case avatar.FieldURL:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Avatar", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Avatar fields.
func (a *Avatar) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case avatar.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case avatar.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				a.URL = value.String
			}
		case avatar.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				a.Status = int8(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Avatar.
// Note that you need to call Avatar.Unwrap() before calling this method if this Avatar
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Avatar) Update() *AvatarUpdateOne {
	return (&AvatarClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Avatar entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Avatar) Unwrap() *Avatar {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Avatar is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Avatar) String() string {
	var builder strings.Builder
	builder.WriteString("Avatar(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", url=")
	builder.WriteString(a.URL)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", a.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Avatars is a parsable slice of Avatar.
type Avatars []*Avatar

func (a Avatars) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
