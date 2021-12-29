// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notification/pkg/db/ent/mailbox"
)

// MailBox is the model entity for the MailBox schema.
type MailBox struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MailBox) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mailbox.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MailBox", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MailBox fields.
func (mb *MailBox) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mailbox.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mb.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this MailBox.
// Note that you need to call MailBox.Unwrap() before calling this method if this MailBox
// was returned from a transaction, and the transaction was committed or rolled back.
func (mb *MailBox) Update() *MailBoxUpdateOne {
	return (&MailBoxClient{config: mb.config}).UpdateOne(mb)
}

// Unwrap unwraps the MailBox entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mb *MailBox) Unwrap() *MailBox {
	tx, ok := mb.config.driver.(*txDriver)
	if !ok {
		panic("ent: MailBox is not a transactional entity")
	}
	mb.config.driver = tx.drv
	return mb
}

// String implements the fmt.Stringer.
func (mb *MailBox) String() string {
	var builder strings.Builder
	builder.WriteString("MailBox(")
	builder.WriteString(fmt.Sprintf("id=%v", mb.ID))
	builder.WriteByte(')')
	return builder.String()
}

// MailBoxes is a parsable slice of MailBox.
type MailBoxes []*MailBox

func (mb MailBoxes) config(cfg config) {
	for _i := range mb {
		mb[_i].config = cfg
	}
}
