// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AnnouncementsColumns holds the columns for the "announcements" table.
	AnnouncementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AnnouncementsTable holds the schema information for the "announcements" table.
	AnnouncementsTable = &schema.Table{
		Name:       "announcements",
		Columns:    AnnouncementsColumns,
		PrimaryKey: []*schema.Column{AnnouncementsColumns[0]},
	}
	// MailBoxesColumns holds the columns for the "mail_boxes" table.
	MailBoxesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "from_user_id", Type: field.TypeUUID},
		{Name: "to_user_id", Type: field.TypeUUID},
		{Name: "alread_read", Type: field.TypeBool},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// MailBoxesTable holds the schema information for the "mail_boxes" table.
	MailBoxesTable = &schema.Table{
		Name:       "mail_boxes",
		Columns:    MailBoxesColumns,
		PrimaryKey: []*schema.Column{MailBoxesColumns[0]},
	}
	// NotificationsColumns holds the columns for the "notifications" table.
	NotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "already_read", Type: field.TypeBool},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// NotificationsTable holds the schema information for the "notifications" table.
	NotificationsTable = &schema.Table{
		Name:       "notifications",
		Columns:    NotificationsColumns,
		PrimaryKey: []*schema.Column{NotificationsColumns[0]},
	}
	// ReadUsersColumns holds the columns for the "read_users" table.
	ReadUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "announcement_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// ReadUsersTable holds the schema information for the "read_users" table.
	ReadUsersTable = &schema.Table{
		Name:       "read_users",
		Columns:    ReadUsersColumns,
		PrimaryKey: []*schema.Column{ReadUsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AnnouncementsTable,
		MailBoxesTable,
		NotificationsTable,
		ReadUsersTable,
	}
)

func init() {
}
