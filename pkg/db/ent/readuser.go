// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/notification/pkg/db/ent/readuser"
	"github.com/google/uuid"
)

// ReadUser is the model entity for the ReadUser schema.
type ReadUser struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// AnnouncementID holds the value of the "announcement_id" field.
	AnnouncementID uuid.UUID `json:"announcement_id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ReadUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case readuser.FieldCreateAt, readuser.FieldUpdateAt, readuser.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case readuser.FieldID, readuser.FieldAppID, readuser.FieldUserID, readuser.FieldAnnouncementID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ReadUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ReadUser fields.
func (ru *ReadUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case readuser.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ru.ID = *value
			}
		case readuser.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ru.AppID = *value
			}
		case readuser.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ru.UserID = *value
			}
		case readuser.FieldAnnouncementID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field announcement_id", values[i])
			} else if value != nil {
				ru.AnnouncementID = *value
			}
		case readuser.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ru.CreateAt = uint32(value.Int64)
			}
		case readuser.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ru.UpdateAt = uint32(value.Int64)
			}
		case readuser.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				ru.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ReadUser.
// Note that you need to call ReadUser.Unwrap() before calling this method if this ReadUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (ru *ReadUser) Update() *ReadUserUpdateOne {
	return (&ReadUserClient{config: ru.config}).UpdateOne(ru)
}

// Unwrap unwraps the ReadUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ru *ReadUser) Unwrap() *ReadUser {
	_tx, ok := ru.config.driver.(*txDriver)
	if !ok {
		panic("ent: ReadUser is not a transactional entity")
	}
	ru.config.driver = _tx.drv
	return ru
}

// String implements the fmt.Stringer.
func (ru *ReadUser) String() string {
	var builder strings.Builder
	builder.WriteString("ReadUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ru.ID))
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ru.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ru.UserID))
	builder.WriteString(", ")
	builder.WriteString("announcement_id=")
	builder.WriteString(fmt.Sprintf("%v", ru.AnnouncementID))
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(fmt.Sprintf("%v", ru.CreateAt))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(fmt.Sprintf("%v", ru.UpdateAt))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(fmt.Sprintf("%v", ru.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// ReadUsers is a parsable slice of ReadUser.
type ReadUsers []*ReadUser

func (ru ReadUsers) config(cfg config) {
	for _i := range ru {
		ru[_i].config = cfg
	}
}
