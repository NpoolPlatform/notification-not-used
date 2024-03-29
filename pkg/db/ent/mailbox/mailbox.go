// Code generated by ent, DO NOT EDIT.

package mailbox

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the mailbox type in the database.
	Label = "mail_box"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldFromUserID holds the string denoting the from_user_id field in the database.
	FieldFromUserID = "from_user_id"
	// FieldToUserID holds the string denoting the to_user_id field in the database.
	FieldToUserID = "to_user_id"
	// FieldAlreadyRead holds the string denoting the already_read field in the database.
	FieldAlreadyRead = "already_read"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the mailbox in the database.
	Table = "mail_boxes"
)

// Columns holds all SQL columns for mailbox fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldFromUserID,
	FieldToUserID,
	FieldAlreadyRead,
	FieldTitle,
	FieldContent,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
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
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
